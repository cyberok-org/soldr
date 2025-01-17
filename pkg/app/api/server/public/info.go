package public

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"soldr/pkg/app/api/logger"
	"soldr/pkg/app/api/models"
	"soldr/pkg/app/api/server/response"
	"soldr/pkg/version"
)

type info struct {
	Type     string            `json:"type"`
	Service  *models.Service   `json:"service"`
	Develop  bool              `json:"develop"`
	User     models.User       `json:"user"`
	Role     models.Role       `json:"role"`
	Tenant   models.Tenant     `json:"tenant"`
	Privs    []string          `json:"privileges"`
	Services []*models.Service `json:"services"`
}

// Info is function to return settings and current information about system and config
// @Summary Retrieve current user and system settings
// @Tags Public
// @Produce json
// @Param refresh_cookie query boolean false "boolean arg to refresh current cookie, use explicit false"
// @Success 200 {object} response.successResp{data=info} "info received successful"
// @Failure 403 {object} response.errorResp "getting info not permitted"
// @Failure 404 {object} response.errorResp "user not found"
// @Failure 500 {object} response.errorResp "internal error on getting information about system and config"
// @Router /info [get]
func (s *AuthService) Info(c *gin.Context) {
	var (
		err   error
		expt  int64
		gtmt  int64
		ok    bool
		privs []string
		resp  info
	)

	nowt := time.Now().Unix()
	session := sessions.Default(c)
	uid := session.Get("uid")
	exp := session.Get("exp")
	gtm := session.Get("gtm")
	svc := session.Get("svc")
	resp.Develop = version.IsDevelop == "true"

	if privs := session.Get("prm"); privs != nil {
		if resp.Privs, ok = privs.([]string); !ok || resp.Privs == nil {
			resp.Privs = make([]string, 0)
		}
	}

	if exp != nil && gtm != nil {
		expt, _ = exp.(int64)
		gtmt, _ = gtm.(int64)
	}

	if uid == nil || expt == 0 || gtmt == 0 || nowt > expt {
		resp.Type = "guest"
		resp.Privs = make([]string, 0)
	} else {
		resp.Type = "user"
		err = s.db.Take(&resp.User, "id = ?", uid).
			Related(&resp.Role).
			Related(&resp.Tenant).Error
		if err != nil {
			response.Error(c, response.ErrInfoUserNotFound, err)
			return
		} else if err = resp.User.Valid(); err != nil {
			logger.FromContext(c).WithError(err).Errorf("error validating user data '%s'", resp.User.Hash)
			response.Error(c, response.ErrInfoInvalidUserData, err)
			return
		}

		if err = s.db.Table("privileges").Where("role_id = ?", resp.User.RoleID).Pluck("name", &privs).Error; err != nil {
			logger.FromContext(c).WithError(err).Errorf("error getting user privileges list '%s'", resp.User.Hash)
			response.Error(c, response.ErrInfoInvalidUserData, err)
			return
		}

		if err = s.db.Find(&resp.Services, "tenant_id = ?", resp.User.TenantID).Error; err != nil {
			logger.FromContext(c).WithError(err).Errorf("error getting user services list '%s'", resp.User.Hash)
			response.Error(c, response.ErrInfoInvalidServiceData, err)
			return
		}

		for _, service := range resp.Services {
			service.Info = nil
			if svc == nil {
				continue
			}
			if svcHash, ok := svc.(string); ok && service.Hash == svcHash {
				resp.Service = service
			}
		}

		// check 5 minutes timeout to refresh current token
		var fiveMins int64 = 5 * 60
		if nowt >= gtmt+fiveMins && c.Query("refresh_cookie") != "false" {
			if err = s.refreshCookie(c, &resp, privs); err != nil {
				logger.FromContext(c).WithError(err).Errorf("failed to refresh token")
				// raise error when there is elapsing last five minutes
				if nowt >= gtmt+int64(s.cfg.SessionTimeout)-fiveMins {
					response.Error(c, response.ErrInternal, err)
					return
				}
			}
		}
	}

	// raise error when there is elapsing last five minutes
	// and user hasn't permissions in the session auth cookie
	if resp.Type != "guest" && resp.Privs == nil {
		logger.FromContext(c).
			WithFields(logrus.Fields{
				"uid": resp.User.ID,
				"rid": resp.User.RoleID,
				"tid": resp.User.TenantID,
			}).
			Errorf("failed to get user privileges for '%s' '%s'", resp.User.Mail, resp.User.Name)
		response.Error(c, response.ErrInternal, err)
		return
	}

	response.Success(c, http.StatusOK, resp)
}
