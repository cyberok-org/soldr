package private

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"soldr/pkg/app/api/client"
	"soldr/pkg/app/api/logger"
	"soldr/pkg/app/api/models"
	"soldr/pkg/app/api/server/response"
	"soldr/pkg/app/api/storage"
	"soldr/pkg/app/api/useraction"
)

type groupDetails struct {
	Hash          string                   `json:"hash"`
	Agents        int                      `json:"agents"`
	ActiveModules int                      `json:"active_modules"`
	JoinedModules string                   `json:"joined_modules"`
	Consistency   bool                     `json:"consistency"`
	Dependencies  []models.GroupDependency `json:"dependencies"`
	Policies      []models.Policy          `json:"policies,omitempty"`
	Modules       []models.ModuleAShort    `json:"modules,omitempty"`
}

type groups struct {
	Groups  []models.Group `json:"groups"`
	Details []groupDetails `json:"details"`
	Total   uint64         `json:"total"`
}

type group struct {
	Group   models.Group `json:"group"`
	Details groupDetails `json:"details"`
}

type groupInfo struct {
	Name string   `json:"name" binding:"max=255,required_without=From"`
	Tags []string `json:"tags" binding:"omitempty"`
	From uint64   `json:"from" binding:"min=0,numeric,omitempty"`
}

type groupPolicyPatch struct {
	// Action on group policy must be one of activate, deactivate
	Action string        `form:"action" json:"action" binding:"oneof=activate deactivate,required" default:"activate" enums:"activate,deactivate"`
	Policy models.Policy `form:"policy" json:"policy" binding:"required"`
}

var groupsSQLMappers = map[string]interface{}{
	"id":           "`{{table}}`.id",
	"hash":         "`{{table}}`.hash",
	"created_date": "`{{table}}`.created_date",
	"policy_id":    "`gtp`.policy_id",
	"policy_name":  "JSON_UNQUOTE(JSON_EXTRACT(`policies`.info, '$.name.{{lang}}'))",
	"module_name":  "`modules`.name",
	"agent_count":  "count(`agents`.id)",
	"name":         "JSON_UNQUOTE(JSON_EXTRACT(`{{table}}`.info, '$.name.{{lang}}'))",
	"tags":         storage.TagsMapper,
	"data": "CONCAT(`{{table}}`.hash, ' | ', " +
		"COALESCE(JSON_EXTRACT(`{{table}}`.info, '$.name.ru'), ''), ' | ', " +
		"COALESCE(JSON_EXTRACT(`{{table}}`.info, '$.name.en'), ''), ' | ', " +
		"COALESCE(JSON_EXTRACT(`{{table}}`.info, '$.tags'), ''))",
}

const sqlGroupDetails = `
	SELECT g.hash,
		(SELECT COUNT(a.id) FROM agents a
			WHERE a.group_id = g.id AND a.deleted_at IS NULL) AS agents,
		(SELECT COUNT(m.id) FROM modules m
			LEFT JOIN policies p ON m.policy_id = p.id AND p.deleted_at IS NULL
			LEFT JOIN groups_to_policies AS gtp ON p.id = gtp.policy_id
			WHERE gtp.group_id = g.id AND m.status = 'joined' AND
				m.deleted_at IS NULL) AS active_modules,
		(SELECT GROUP_CONCAT(m.name SEPARATOR ',') FROM modules m
			LEFT JOIN policies p ON m.policy_id = p.id AND p.deleted_at IS NULL
			LEFT JOIN groups_to_policies AS gtp ON p.id = gtp.policy_id
			WHERE gtp.group_id = g.id AND m.status = 'joined' AND m.deleted_at IS NULL
			GROUP BY gtp.group_id) AS joined_modules
	FROM groups AS g`

func getGroupConsistency(modules []models.ModuleAShort) (bool, []models.GroupDependency) {
	var (
		rdeps bool = true
		gdeps []models.GroupDependency
		pdeps []models.PolicyDependency
	)
	getPolicyID := func(name string) uint64 {
		for _, mod := range modules {
			if name == mod.Info.Name {
				return mod.PolicyID
			}
		}
		return 0
	}

	rdeps, pdeps = getPolicyConsistency(modules)

	for _, pdep := range pdeps {
		gdeps = append(gdeps, models.GroupDependency{
			PolicyID:         getPolicyID(pdep.SourceModuleName),
			PolicyDependency: pdep,
		})
	}

	return rdeps, gdeps
}

func makeGroupPolicyAction(act string, iDB *gorm.DB, g models.Group, p models.Policy) (*response.HttpError, error) {
	gps := models.GroupPolicies{
		Group: g,
	}
	if err := iDB.Model(gps).Association("policies").Find(&gps.Policies).Error; err != nil {
		return response.ErrGroupPolicyPoliciesNotFound, err
	}

	isPolicyActive := false
	groupToPolicy := models.GroupToPolicy{
		PolicyID: p.ID,
		GroupID:  g.ID,
	}
	pids := []uint64{0}
	for _, gp := range gps.Policies {
		pids = append(pids, gp.ID)
		if gp.ID == p.ID {
			isPolicyActive = true
		}
	}

	switch act {
	case "activate":
		var cnts []int64
		findDupsQuery := iDB.
			Table((&models.ModuleA{}).TableName()).
			Select("count(*) AS cnt").
			Where("deleted_at IS NULL").
			Where("(policy_id IN (?) OR policy_id = ?) AND status = 'joined'", pids, p.ID).
			Group("name").
			Having("cnt > 1").
			Find(&cnts)
		if err := findDupsQuery.Error; err != nil {
			return response.ErrGroupPolicyMergeModulesFail, err
		}

		if len(cnts) != 0 {
			return response.ErrGroupPolicyDuplicateModules, nil
		}

		if !isPolicyActive {
			if err := iDB.Create(&groupToPolicy).Error; err != nil {
				return response.ErrGroupPolicyLinkFail, err
			}
			return nil, nil
		} else {
			return response.ErrGroupPolicyLinkExists, nil
		}

	case "deactivate":
		if isPolicyActive {
			err := iDB.Where("policy_id = ? AND group_id = ?", groupToPolicy.PolicyID, groupToPolicy.GroupID).
				Delete(&groupToPolicy).Error
			if err != nil {
				return response.ErrGroupPolicyRemoveLink, err
			}
			return nil, nil
		} else {
			return response.ErrGroupPolicyLinkNotFound, nil
		}

	default:
		return response.ErrGroupPolicyUnkownAction, nil
	}
}

func getGroupName(db *gorm.DB, hash string) (string, error) {
	var group models.Group
	if err := db.Take(&group, "hash = ?", hash).Error; err != nil {
		return "", err
	}
	return group.Info.Name.En, nil
}

type GroupService struct {
	serverConnector  *client.AgentServerClient
	userActionWriter useraction.Writer
	modulesStorage   *storage.ModuleStorage
}

func NewGroupService(
	serverConnector *client.AgentServerClient,
	userActionWriter useraction.Writer,
	modulesStorage *storage.ModuleStorage,
) *GroupService {
	return &GroupService{
		serverConnector:  serverConnector,
		userActionWriter: userActionWriter,
		modulesStorage:   modulesStorage,
	}
}

// GetGroups is a function to return group list view on dashboard
// @Summary Retrieve groups list by filters
// @Tags Groups
// @Produce json
// @Param request query storage.TableQuery true "query table params"
// @Success 200 {object} response.successResp{data=groups} "groups list received successful"
// @Failure 400 {object} response.errorResp "invalid query request data"
// @Failure 403 {object} response.errorResp "getting groups not permitted"
// @Failure 404 {object} response.errorResp "groups not found"
// @Failure 500 {object} response.errorResp "internal error on getting groups"
// @Router /groups/ [get]
func (s *GroupService) GetGroups(c *gin.Context) {
	var (
		gids          []uint64
		gpss          []models.GroupToPolicy
		policiesa     []models.Policy
		query         storage.TableQuery
		resp          groups
		groupedResp   response.GroupedData
		useModule     bool
		usePolicy     bool
		usePolicyName bool
	)

	if err := c.ShouldBindQuery(&query); err != nil {
		logger.FromContext(c).WithError(err).Errorf("error binding query")
		response.Error(c, response.ErrGroupsInvalidRequest, err)
		return
	}

	serviceHash := c.GetString("svc")
	if serviceHash == "" {
		logger.FromContext(c).Errorf("could not get service hash")
		response.Error(c, response.ErrInternal, nil)
		return
	}
	iDB, err := s.serverConnector.GetDB(c, serviceHash)
	if err != nil {
		logger.FromContext(c).WithError(err).Error()
		response.Error(c, response.ErrInternalDBNotFound, err)
		return
	}

	if err = query.Init("groups", groupsSQLMappers); err != nil {
		logger.FromContext(c).WithError(err).Errorf("error binding query")
		response.Error(c, response.ErrGroupsInvalidRequest, err)
		return
	}

	setUsingTables := func(sfield string) {
		if sfield == "policy_id" {
			usePolicy = true
		}
		if sfield == "module_name" {
			usePolicy = true
			useModule = true
		}
		if sfield == "policy_name" {
			usePolicy = true
			usePolicyName = true
		}
	}
	setUsingTables(query.Sort.Prop)
	setUsingTables(query.Group)
	for _, filter := range query.Filters {
		setUsingTables(filter.Field)
	}
	query.SetFilters([]func(db *gorm.DB) *gorm.DB{
		func(db *gorm.DB) *gorm.DB {
			return db.Where("groups.deleted_at IS NULL")
		},
	})
	funcs := []func(db *gorm.DB) *gorm.DB{
		func(db *gorm.DB) *gorm.DB {
			db = db.Group(`groups.id`)
			if usePolicy {
				db = db.Joins(`LEFT JOIN groups_to_policies gtp ON gtp.group_id = groups.id`)
			}
			if useModule {
				db = db.Joins(`LEFT JOIN modules ON gtp.policy_id = modules.policy_id AND modules.status = 'joined' AND modules.deleted_at IS NULL`)
			}
			if usePolicyName {
				db = db.Joins("LEFT JOIN policies ON gtp.policy_id = policies.id AND policies.deleted_at IS NULL")
			}
			if usePolicy || useModule || usePolicyName {
				db = db.Group("groups.id")
			}
			return db.Joins(`LEFT JOIN agents ON agents.group_id = groups.id AND agents.deleted_at IS NULL`)
		},
	}

	if query.Group == "" {
		if resp.Total, err = query.Query(iDB, &resp.Groups, funcs...); err != nil {
			logger.FromContext(c).WithError(err).Errorf("error finding groups")
			response.Error(c, response.ErrGroupsInvalidQuery, err)
			return
		}
	} else {
		if groupedResp.Total, err = query.QueryGrouped(iDB, &groupedResp.Grouped, funcs...); err != nil {
			logger.FromContext(c).WithError(err).Errorf("error finding grouped groups")
			response.Error(c, response.ErrGetAgentsInvalidQuery, err)
			return
		}
		response.Success(c, http.StatusOK, groupedResp)
		return
	}

	for i := 0; i < len(resp.Groups); i++ {
		gids = append(gids, resp.Groups[i].ID)
		if err = resp.Groups[i].Valid(); err != nil {
			logger.FromContext(c).WithError(err).Errorf("error validating group data '%s'", resp.Groups[i].Hash)
			response.Error(c, response.ErrGroupsInvalidData, err)
			return
		}
	}

	sqlQuery := sqlGroupDetails + ` WHERE g.id IN (?) AND g.deleted_at IS NULL`
	if err = iDB.Raw(sqlQuery, gids).Scan(&resp.Details).Error; err != nil {
		logger.FromContext(c).WithError(err).Errorf("error loading groups details")
		response.Error(c, response.ErrGetGroupsDetailsNotFound, err)
		return
	}

	if err = iDB.Find(&gpss, "group_id IN (?)", gids).Error; err != nil {
		logger.FromContext(c).WithError(err).Errorf("error finding policy to groups links")
		response.Error(c, response.ErrGroupPolicyGroupsNotFound, err)
		return
	}

	polsToGroups := make(map[uint64][]models.Policy)
	err = iDB.Model(&models.Policy{}).
		Group("policies.id").
		Joins(`LEFT JOIN groups_to_policies gtp ON gtp.policy_id = policies.id AND gtp.group_id IN (?)`, gids).
		Find(&policiesa).Error
	if err != nil {
		logger.FromContext(c).WithError(err).Errorf("error finding group policies")
		response.Error(c, response.ErrGroupPolicyPoliciesNotFound, err)
		return
	} else {
		for i := 0; i < len(policiesa); i++ {
			id := policiesa[i].ID
			name := policiesa[i].Info.Name
			if err = policiesa[i].Valid(); err != nil {
				logger.FromContext(c).WithError(err).Errorf("error validating policy data '%d' '%s'", id, name)
				response.Error(c, response.ErrGetGroupsInvalidModuleData, err)
				return
			}
			for idx := range gpss {
				if gpss[idx].PolicyID != id {
					continue
				}
				group_id := gpss[idx].GroupID
				if pols, ok := polsToGroups[group_id]; ok {
					polsToGroups[group_id] = append(pols, policiesa[i])
				} else {
					polsToGroups[group_id] = []models.Policy{policiesa[i]}
				}
			}
		}
	}

	modsToPolicies, err := s.modulesStorage.GetPoliciesModulesAShortByGroups(iDB, gids)
	if err != nil {
		logger.FromContext(c).WithError(err).Errorf("error finding group modules")
		response.Error(c, response.ErrGetGroupModulesInvalidGroupData, err)
		return
	}
	for _, group := range resp.Groups {
		var details *groupDetails
		for idx := range resp.Details {
			if resp.Details[idx].Hash == group.Hash {
				details = &resp.Details[idx]
			}
		}
		if details == nil {
			continue
		}

		if pols, ok := polsToGroups[group.ID]; ok {
			details.Policies = pols
		}

		for idx := range details.Policies {
			if mods, ok := modsToPolicies[details.Policies[idx].ID]; ok {
				details.Modules = append(details.Modules, mods...)
			}
		}

		details.Consistency, details.Dependencies = getGroupConsistency(details.Modules)
	}

	response.Success(c, http.StatusOK, resp)
}

// GetGroup is a function to return group info and details view
// @Summary Retrieve group info by group hash
// @Tags Groups
// @Produce json
// @Param hash path string true "group hash in hex format (md5)" minlength(32) maxlength(32)
// @Success 200 {object} response.successResp{data=group} "group info received successful"
// @Failure 403 {object} response.errorResp "getting group info not permitted"
// @Failure 404 {object} response.errorResp "group not found"
// @Failure 500 {object} response.errorResp "internal error on getting group"
// @Router /groups/{hash} [get]
func (s *GroupService) GetGroup(c *gin.Context) {
	var (
		hash = c.Param("hash")
		resp group
	)

	serviceHash := c.GetString("svc")
	if serviceHash == "" {
		logger.FromContext(c).Errorf("could not get service hash")
		response.Error(c, response.ErrInternal, nil)
		return
	}
	iDB, err := s.serverConnector.GetDB(c, serviceHash)
	if err != nil {
		logger.FromContext(c).WithError(err).Error()
		response.Error(c, response.ErrInternalDBNotFound, err)
		return
	}

	if err = iDB.Take(&resp.Group, "hash = ?", hash).Error; err != nil {
		logger.FromContext(c).WithError(err).Errorf("error finding group by hash")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, response.ErrGroupsNotFound, err)
		} else {
			response.Error(c, response.ErrInternal, err)
		}
		return
	} else if err = resp.Group.Valid(); err != nil {
		logger.FromContext(c).WithError(err).Errorf("error validating group data '%s'", resp.Group.Hash)
		response.Error(c, response.ErrGroupsInvalidData, err)
		return
	}

	sqlQuery := sqlGroupDetails + ` WHERE g.hash = ? AND g.deleted_at IS NULL`
	if err = iDB.Raw(sqlQuery, hash).Scan(&resp.Details).Error; err != nil {
		logger.FromContext(c).WithError(err).Errorf("error loading details by group hash '%s'", hash)
		response.Error(c, response.ErrGetGroupDetailsNotFound, err)
		return
	}
	resp.Details.Modules, err = s.modulesStorage.GetModulesAShortByGroup(iDB, resp.Group.ID)
	if err != nil {
		logger.FromContext(c).WithError(err).Errorf("error finding group modules")
		response.Error(c, response.ErrGetGroupModulesInvalidGroupData, err)
		return
	}
	resp.Details.Consistency, resp.Details.Dependencies = getGroupConsistency(resp.Details.Modules)

	gps := models.GroupPolicies{
		Group: resp.Group,
	}
	if err = iDB.Model(gps).Association("policies").Find(&gps.Policies).Error; err != nil {
		logger.FromContext(c).WithError(err).Errorf("error finding group policies by group model")
		response.Error(c, response.ErrGetGroupsPoliciesNotFound, err)
		return
	}
	resp.Details.Policies = gps.Policies

	response.Success(c, http.StatusOK, resp)
}

// PatchGroup is a function to update group public info only
// @Summary Update group info by group hash
// @Tags Groups
// @Accept json
// @Produce json
// @Param hash path string true "group hash in hex format (md5)" minlength(32) maxlength(32)
// @Param json body models.Group true "group info as JSON data"
// @Success 200 {object} response.successResp{data=models.Group} "group info updated successful"
// @Failure 400 {object} response.errorResp "invalid group info"
// @Failure 403 {object} response.errorResp "updating group info not permitted"
// @Failure 404 {object} response.errorResp "group not found"
// @Failure 500 {object} response.errorResp "internal error on updating group"
// @Router /groups/{hash} [put]
func (s *GroupService) PatchGroup(c *gin.Context) {
	var (
		count int64
		group models.Group
		hash  = c.Param("hash")
	)

	uaf := useraction.NewFields(c, "group", "group", "editing", hash, useraction.UnknownObjectDisplayName)
	defer s.userActionWriter.WriteUserAction(c, uaf)

	serviceHash := c.GetString("svc")
	if serviceHash == "" {
		logger.FromContext(c).Errorf("could not get service hash")
		response.Error(c, response.ErrInternal, nil)
		return
	}
	iDB, err := s.serverConnector.GetDB(c, serviceHash)
	if err != nil {
		logger.FromContext(c).WithError(err).Error()
		response.Error(c, response.ErrInternalDBNotFound, err)
		return
	}

	if err := c.ShouldBindJSON(&group); err != nil || group.Valid() != nil {
		if err == nil {
			err = group.Valid()
		}
		name, nameErr := getGroupName(iDB, hash)
		if nameErr == nil {
			uaf.ObjectDisplayName = name
		}
		logger.FromContext(c).WithError(err).Errorf("error binding JSON")
		response.Error(c, response.ErrGroupsValidationFail, err)
		return
	}

	uaf.ObjectDisplayName = group.Info.Name.En

	if hash != group.Hash {
		logger.FromContext(c).Errorf("mismatch group hash to requested one")
		response.Error(c, response.ErrGroupsValidationFail, nil)
		return
	}

	if err = iDB.Model(&group).Count(&count).Error; err != nil || count == 0 {
		logger.FromContext(c).Errorf("error updating group by hash '%s', group not found", hash)
		response.Error(c, response.ErrGroupsNotFound, err)
		return
	}

	public_info := []interface{}{"info", "updated_at"}
	err = iDB.Select("", public_info...).Save(&group).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.FromContext(c).Errorf("error updating group by hash '%s', group not found", hash)
		response.Error(c, response.ErrGroupsNotFound, err)
		return
	} else if err != nil {
		logger.FromContext(c).WithError(err).Errorf("error updating group by hash '%s'", hash)
		response.Error(c, response.ErrInternal, err)
		return
	}

	response.Success(c, http.StatusOK, group)
}

// PatchGroupPolicy is a function to update group policy linking
// @Summary Update or patch group policy by group hash and policy object
// @Tags Groups,Policies
// @Accept json
// @Produce json
// @Param hash path string true "group hash in hex format (md5)" minlength(32) maxlength(32)
// @Param json body groupPolicyPatch true "action on group policy as JSON data (activate, deactivate)"
// @Success 200 {object} response.successResp "group policy patched successful"
// @Failure 400 {object} response.errorResp "invalid patch request data"
// @Failure 403 {object} response.errorResp "updating group policy not permitted"
// @Failure 404 {object} response.errorResp "group or policy not found"
// @Failure 500 {object} response.errorResp "internal error on getting updating group policy"
// @Router /groups/{hash}/policies [put]
func (s *GroupService) PatchGroupPolicy(c *gin.Context) {
	var (
		form   groupPolicyPatch
		group  models.Group
		hash   = c.Param("hash")
		policy models.Policy
	)

	uaf := useraction.NewFields(c, "policy", "policy", "undefined action", "", useraction.UnknownObjectDisplayName)
	defer s.userActionWriter.WriteUserAction(c, uaf)

	if err := c.ShouldBindJSON(&form); err != nil {
		logger.FromContext(c).WithError(err).Errorf("error binding JSON")
		response.Error(c, response.ErrGroupsInvalidRequest, err)
		return
	}

	if form.Action == "activate" {
		uaf.ActionCode = "creation of the connection with the group"
	} else {
		uaf.ActionCode = "deletion of the connection with the group"
	}
	uaf.ObjectID = form.Policy.Hash
	uaf.ObjectDisplayName = form.Policy.Info.Name.En

	serviceHash := c.GetString("svc")
	if serviceHash == "" {
		logger.FromContext(c).Errorf("could not get service hash")
		response.Error(c, response.ErrInternal, nil)
		return
	}
	iDB, err := s.serverConnector.GetDB(c, serviceHash)
	if err != nil {
		logger.FromContext(c).WithError(err).Error()
		response.Error(c, response.ErrInternalDBNotFound, err)
		return
	}

	if err = iDB.Take(&group, "hash = ?", hash).Error; err != nil {
		logger.FromContext(c).WithError(err).Errorf("error finding group by hash")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, response.ErrGroupsNotFound, err)
		} else {
			response.Error(c, response.ErrInternal, err)
		}
		return
	} else if err = group.Valid(); err != nil {
		logger.FromContext(c).WithError(err).Errorf("error validating group data '%s'", group.Hash)
		response.Error(c, response.ErrGroupsInvalidData, err)
		return
	}

	if err = iDB.Take(&policy, "hash = ?", form.Policy.Hash).Error; err != nil {
		logger.FromContext(c).WithError(err).Errorf("error finding policy by hash")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, response.ErrGetGroupsPoliciesNotFound, err)
		} else {
			response.Error(c, response.ErrInternal, err)
		}
		return
	} else if err = policy.Valid(); err != nil {
		logger.FromContext(c).WithError(err).Errorf("error validating policy data '%s'", policy.Hash)
		response.Error(c, response.ErrInternal, err)
		return
	}

	httpErr, err := makeGroupPolicyAction(form.Action, iDB, group, policy)
	if httpErr != nil {
		logger.FromContext(c).WithError(err).Errorf("error patching group policy by action: %s", httpErr.Error())
		response.Error(c, httpErr, err)
	}

	response.Success(c, http.StatusOK, struct{}{})
}

// CreateGroup is a function to create new group
// @Summary Create new group in service
// @Tags Groups
// @Accept json
// @Produce json
// @Param json body groupInfo true "group info to create one"
// @Success 201 {object} response.successResp{data=models.Group} "group created successful"
// @Failure 400 {object} response.errorResp "invalid group info"
// @Failure 403 {object} response.errorResp "creating group not permitted"
// @Failure 500 {object} response.errorResp "internal error on creating group"
// @Router /groups/ [post]
func (s *GroupService) CreateGroup(c *gin.Context) {
	var (
		groupFrom models.Group
		info      groupInfo
	)

	uaf := useraction.NewFields(c, "group", "group", "creation", "", useraction.UnknownObjectDisplayName)
	defer s.userActionWriter.WriteUserAction(c, uaf)

	if err := c.ShouldBindJSON(&info); err != nil {
		logger.FromContext(c).WithError(err).Errorf("error binding JSON")
		response.Error(c, response.ErrEventsInvalidRequest, err)
		return
	}
	uaf.ObjectDisplayName = info.Name

	serviceHash := c.GetString("svc")
	if serviceHash == "" {
		logger.FromContext(c).Errorf("could not get service hash")
		response.Error(c, response.ErrInternal, nil)
		return
	}
	iDB, err := s.serverConnector.GetDB(c, serviceHash)
	if err != nil {
		logger.FromContext(c).WithError(err).Error()
		response.Error(c, response.ErrInternalDBNotFound, err)
		return
	}

	group := models.Group{
		Hash: storage.MakeGroupHash(info.Name),
		Info: models.GroupInfo{
			Name: models.GroupItemLocale{
				Ru: info.Name,
				En: info.Name,
			},
			Tags:   info.Tags,
			System: false,
		},
	}
	uaf.ObjectID = group.Hash

	if info.From != 0 {
		if err = iDB.Take(&groupFrom, "id = ?", info.From).Error; err != nil {
			logger.FromContext(c).WithError(err).Errorf("error finding source group by ID")
			response.Error(c, response.ErrCreateGroupSourceNotFound, err)
			return
		} else if err = groupFrom.Valid(); err != nil {
			logger.FromContext(c).WithError(err).Errorf("error validating group data '%s'", groupFrom.Hash)
			response.Error(c, response.ErrGetAgentInvalidGroupData, err)
			return
		}

		group = groupFrom
		group.ID = 0
		group.Info.System = false
		group.CreatedDate = time.Time{}
		group.Hash = storage.MakeGroupHash(group.Hash)
		if info.Name != "" {
			group.Info.Name = models.GroupItemLocale{
				Ru: info.Name,
				En: info.Name,
			}
		} else {
			group.Info.Name.Ru += " (копия)"
			group.Info.Name.En += " (copy)"
		}
		if len(info.Tags) != 0 {
			group.Info.Tags = info.Tags
		}
	}

	if err = iDB.Create(&group).Error; err != nil {
		logger.FromContext(c).WithError(err).Errorf("error creating group")
		response.Error(c, response.ErrCreateGroupCreateFail, err)
		return
	}

	if groupFrom.ID != 0 {
		var groupToPolicy []models.GroupToPolicy
		err = iDB.Where("group_id = ?", groupFrom.ID).Find(&groupToPolicy).Error
		if err != nil {
			logger.FromContext(c).WithError(err).Errorf("error finding group policies by group ID")
			response.Error(c, response.ErrCreateGroupGetPolicies, err)
			return
		}
		for _, gpt := range groupToPolicy {
			gpt.ID = 0
			gpt.GroupID = group.ID
			if err = iDB.Create(&gpt).Error; err != nil {
				logger.FromContext(c).WithError(err).Errorf("error creating group policies")
				response.Error(c, response.ErrCreateGroupCreatePolicies, err)
				return
			}
		}
	}

	response.Success(c, http.StatusCreated, group)
}

// DeleteGroup is a function to cascade delete group
// @Summary Delete group from instance DB
// @Tags Groups
// @Produce json
// @Param hash path string true "group hash in hex format (md5)" minlength(32) maxlength(32)
// @Success 200 {object} response.successResp "group deleted successful"
// @Failure 403 {object} response.errorResp "deleting group not permitted"
// @Failure 404 {object} response.errorResp "group not found"
// @Failure 500 {object} response.errorResp "internal error on deleting group"
// @Router /groups/{hash} [delete]
func (s *GroupService) DeleteGroup(c *gin.Context) {
	var (
		group models.Group
		hash  = c.Param("hash")
	)

	uaf := useraction.NewFields(c, "group", "group", "deletion", hash, useraction.UnknownObjectDisplayName)
	defer s.userActionWriter.WriteUserAction(c, uaf)

	serviceHash := c.GetString("svc")
	if serviceHash == "" {
		logger.FromContext(c).Errorf("could not get service hash")
		response.Error(c, response.ErrInternal, nil)
		return
	}
	iDB, err := s.serverConnector.GetDB(c, serviceHash)
	if err != nil {
		logger.FromContext(c).WithError(err).Error()
		response.Error(c, response.ErrInternalDBNotFound, err)
		return
	}

	if err = iDB.Take(&group, "hash = ?", hash).Error; err != nil {
		logger.FromContext(c).WithError(err).Errorf("error finding group by hash")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, response.ErrGroupsNotFound, err)
			return
		}
		response.Error(c, response.ErrInternal, err)
		return
	} else if err = group.Valid(); err != nil {
		logger.FromContext(c).WithError(err).Errorf("error validating group data '%s'", group.Hash)
		response.Error(c, response.ErrGroupsInvalidData, err)
		return
	}

	uaf.ObjectDisplayName = group.Info.Name.En

	if err = iDB.Delete(&group).Error; err != nil {
		logger.FromContext(c).WithError(err).Errorf("error deleting group by hash '%s'", hash)
		response.Error(c, response.ErrInternal, err)
		return
	}

	response.Success(c, http.StatusOK, struct{}{})
}
