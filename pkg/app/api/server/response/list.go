// Code generated by errgen. DO NOT EDIT.

package response

// agents

var ErrAgentsNotFound = NewHttpError(404, "Agents.NotFound", "agent not found")
var ErrAgentsInvalidData = NewHttpError(500, "Agents.InvalidData", "invalid agent data")
var ErrGetAgentsInvalidRequest = NewHttpError(400, "Agents.GetAgents.InvalidRequest", "invalid agent request data")
var ErrGetAgentsInvalidQuery = NewHttpError(500, "Agents.GetAgents.InvalidQuery", "invalid agent query")
var ErrPatchAgentsInvalidAction = NewHttpError(400, "Agents.PatchAgents.InvalidAction", "invalid agents action")
var ErrPatchAgentsInvalidQuery = NewHttpError(500, "Agents.PatchAgents.InvalidQuery", "invalid agent query")
var ErrPatchAgentsUpdateTasksFail = NewHttpError(500, "Agents.PatchAgents.UpdateTasksFail", "failed to update tasks by filter")
var ErrPatchAgentsUpdateAgentsFail = NewHttpError(500, "Agents.PatchAgents.UpdateAgentsFail", "failed to update agents by filter")
var ErrPatchAgentsDeleteAgentsFail = NewHttpError(500, "Agents.PatchAgents.DeleteAgentsFail", "failed to delete agents by filter")
var ErrPatchAgentsMoveFail = NewHttpError(500, "Agents.PatchAgents.MoveFail", "move fail")
var ErrGetAgentDetailsNotFound = NewHttpError(404, "Agents.GetAgent.DetailsNotFound", "internal error on retrieving agent details")
var ErrGetAgentGroupNotFound = NewHttpError(404, "Agents.GetAgent.GroupNotFound", "group not found")
var ErrGetAgentInvalidGroupData = NewHttpError(500, "Agents.GetAgent.InvalidGroupData", "invalid group data")
var ErrGetAgentGroupModulesNotFound = NewHttpError(404, "Agents.GetAgent.GroupModulesNotFound", "group modules not found")
var ErrGetAgentInvalidAgentModuleData = NewHttpError(500, "Agents.GetAgent.InvalidAgentModuleData", "invalid agent module data")
var ErrGetAgentsInvalidAgentModuleData = NewHttpError(500, "Agents.GetAgents.InvalidAgentModuleData", "invalid agent module data")
var ErrGetAgentsGetSystemModulesFail = NewHttpError(500, "Agents.GetAgents.GetSystemModulesFail", "internal error on retrieving system modules list")
var ErrGetAgentPoliciesNotFound = NewHttpError(404, "Agents.GetAgent.PoliciesNotFound", "group policies not found")
var ErrPatchAgentValidationError = NewHttpError(400, "Agents.PatchAgent.ValidationError", "failed to validate agent")
var ErrPatchAgentTaskUpdateFail = NewHttpError(500, "Agents.PatchAgent.TaskUpdateFail", "failed to update tasks by agent")
var ErrCreateAgentValidationError = NewHttpError(400, "Agents.CreateAgent.ValidationError", "failed to valid agent info")
var ErrCreateAgentCreateError = NewHttpError(500, "Agents.CreateAgent.CreateError", "failed to create agent to db")

// auth

var ErrAuthInvalidLoginRequest = NewHttpError(400, "Auth.InvalidLoginRequest", "invalid login data")
var ErrAuthInvalidAuthorizeQuery = NewHttpError(400, "Auth.InvalidAuthorizeQuery", "invalid authorize query")
var ErrAuthInvalidLoginCallbackRequest = NewHttpError(400, "Auth.InvalidLoginCallbackRequest", "invalid login callback data")
var ErrAuthInvalidAuthorizationState = NewHttpError(400, "Auth.InvalidAuthorizationState", "invalid authorization state data")
var ErrAuthInvalidSwitchServiceHash = NewHttpError(400, "Auth.InvalidSwitchServiceHash", "invalid switch service hash input data")
var ErrAuthInvalidAuthorizationNonce = NewHttpError(400, "Auth.InvalidAuthorizationNonce", "invalid authorization nonce data")
var ErrAuthInvalidCredentials = NewHttpError(401, "Auth.InvalidCredentials", "invalid login or password")
var ErrAuthInvalidUserData = NewHttpError(500, "Auth.InvalidUserData", "invalid user data")
var ErrAuthInactiveUser = NewHttpError(403, "Auth.InactiveUser", "user is inactive")
var ErrAuthExchangeTokenFail = NewHttpError(403, "Auth.ExchangeTokenFail", "error on exchanging token")
var ErrAuthTokenExpired = NewHttpError(403, "Auth.TokenExpired", "token is expired")
var ErrAuthVerificationTokenFail = NewHttpError(403, "Auth.VerificationTokenFail", "error on verifying token")
var ErrAuthInvalidServiceData = NewHttpError(500, "Auth.InvalidServiceData", "invalid service data")
var ErrAuthInvalidTenantData = NewHttpError(500, "Auth.InvalidTenantData", "invalid tenant data")

// binaries

var ErrAgentBinariesInvalidRequest = NewHttpError(400, "Binaries.AgentBinaries.InvalidRequest", "invalid agent binaries request data")
var ErrAgentBinariesInvalidData = NewHttpError(500, "Binaries.AgentBinaries.InvalidData", "invalid agent binaries data")
var ErrAgentBinaryFileInvalidOS = NewHttpError(400, "Binaries.AgentBinaryFile.InvalidOS", "failed to valid agent os")
var ErrAgentBinaryFileInvalidArch = NewHttpError(400, "Binaries.AgentBinaryFile.InvalidArch", "failed to valid agent arch")
var ErrAgentBinaryFileNotFound = NewHttpError(404, "Binaries.AgentBinaryFile.NotFound", "agent binary record not found")
var ErrAgentBinaryFileCorrupted = NewHttpError(404, "Binaries.AgentBinaryFile.Corrupted", "agent binary file was corrupted")

// events

var ErrEventsInvalidRequest = NewHttpError(400, "Events.InvalidRequest", "invalid event request data")
var ErrEventsInvalidQuery = NewHttpError(500, "Events.InvalidQuery", "invalid event query")
var ErrEventsInvalidData = NewHttpError(500, "Events.InvalidData", "invalid event data")

// general

var ErrInternal = NewHttpError(500, "Internal", "internal server error")
var ErrInternalDBNotFound = NewHttpError(500, "Internal.DBNotFound", "db not found")
var ErrInternalServiceNotFound = NewHttpError(500, "Internal.ServiceNotFound", "service not found")
var ErrInternalDBEncryptorNotFound = NewHttpError(500, "Internal.DBEncryptorNotFound", "DBEncryptor not found")
var ErrNotPermitted = NewHttpError(403, "NotPermitted", "action not permitted")
var ErrAuthRequired = NewHttpError(403, "AuthRequired", "auth required")
var ErrLocalUserRequired = NewHttpError(403, "LocalUserRequired", "local user required")
var ErrPrivilegesRequired = NewHttpError(403, "PrivilegesRequired", "some privileges required")
var ErrAdminRequired = NewHttpError(403, "AdminRequired", "admin required")
var ErrSuperRequired = NewHttpError(403, "SuperRequired", "super admin required")

// group-policy

var ErrGroupPolicyPoliciesNotFound = NewHttpError(404, "GroupPolicy.PoliciesNotFound", "group policies not found")
var ErrGroupPolicyGroupsNotFound = NewHttpError(404, "GroupPolicy.GroupsNotFound", "policy groups not found")
var ErrGroupPolicyMergeModulesFail = NewHttpError(400, "GroupPolicy.MergeModulesFail", "failed to merge modules")
var ErrGroupPolicyDuplicateModules = NewHttpError(403, "GroupPolicy.DuplicateModules", "found duplicate modules")
var ErrGroupPolicyLinkFail = NewHttpError(500, "GroupPolicy.LinkFail", "failed to create link")
var ErrGroupPolicyLinkExists = NewHttpError(400, "GroupPolicy.LinkExists", "link already exists")
var ErrGroupPolicyRemoveLink = NewHttpError(500, "GroupPolicy.RemoveLink", "failed to remove link")
var ErrGroupPolicyLinkNotFound = NewHttpError(404, "GroupPolicy.LinkNotFound", "link not found")
var ErrGroupPolicyUnkownAction = NewHttpError(400, "GroupPolicy.UnkownAction", "action not found or unknown")

// groups

var ErrGroupsInvalidRequest = NewHttpError(400, "Groups.InvalidRequest", "invalid group request data")
var ErrGroupsInvalidQuery = NewHttpError(500, "Groups.InvalidQuery", "invalid group query")
var ErrGroupsInvalidData = NewHttpError(500, "Groups.InvalidData", "invalid group data")
var ErrGroupsNotFound = NewHttpError(404, "Groups.NotFound", "group not found")
var ErrGroupsValidationFail = NewHttpError(400, "Groups.ValidationFail", "failed to validate group")
var ErrGetGroupDetailsNotFound = NewHttpError(404, "Groups.GetGroup.DetailsNotFound", "internal error on retrieving group details")
var ErrGetGroupModulesNotFound = NewHttpError(404, "Groups.GetGroup.ModulesNotFound", "modules not found while getting group")
var ErrGetGroupsDetailsNotFound = NewHttpError(404, "Groups.GetGroups.DetailsNotFound", "internal error on retrieving groups details")
var ErrGetGroupsModulesNotFound = NewHttpError(404, "Groups.GetGroups.ModulesNotFound", "modules not found while getting groups")
var ErrGetGroupsInvalidModuleData = NewHttpError(500, "Groups.GetGroups.InvalidModuleData", "invalid group module data")
var ErrGetGroupsGetSystemModulesFail = NewHttpError(500, "Groups.GetGroups.GetSystemModulesFail", "internal error on retrieving system modules list")
var ErrGetGroupsPoliciesNotFound = NewHttpError(404, "Groups.GetGroups.PoliciesNotFound", "group policies not found")
var ErrCreateGroupSourceNotFound = NewHttpError(404, "Groups.CreateGroup.SourceNotFound", "source group not found")
var ErrCreateGroupCreateFail = NewHttpError(500, "Groups.CreateGroup.CreateFail", "failed to create group to db")
var ErrCreateGroupGetPolicies = NewHttpError(500, "Groups.CreateGroup.GetPolicies", "failed to get group policies")
var ErrCreateGroupCreatePolicies = NewHttpError(500, "Groups.CreateGroup.CreatePolicies", "failed to create group policies to db")

// info

var ErrInfoUserNotFound = NewHttpError(404, "Info.UserNotFound", "user not found")
var ErrInfoInvalidUserData = NewHttpError(500, "Info.InvalidUserData", "invalid user data")
var ErrInfoInvalidServiceData = NewHttpError(500, "Info.InvalidServiceData", "invalid service data")

// modules

var ErrModulesInvalidRequest = NewHttpError(400, "Modules.InvalidRequest", "invalid agent modules request data")
var ErrModulesNotFound = NewHttpError(404, "Modules.NotFound", "module not found")
var ErrModulesSystemModuleNotFound = NewHttpError(404, "Modules.SystemModuleNotFound", "system module not found")
var ErrModulesInvalidSystemModuleData = NewHttpError(500, "Modules.InvalidSystemModuleData", "invalid system module data")
var ErrModulesInvalidData = NewHttpError(500, "Modules.InvalidData", "invalid module data")
var ErrGetAgentModulesAgentNotFound = NewHttpError(404, "Modules.GetAgentModules.AgentNotFound", "agent not found")
var ErrGetAgentModulesAgentPoliciesNotFound = NewHttpError(404, "Modules.GetAgentModules.AgentPoliciesNotFound", "agent policies not found")
var ErrGetAgentModulesInvalidAgentPoliciesData = NewHttpError(500, "Modules.GetAgentModules.InvalidAgentPoliciesData", "invalid agent policies data")
var ErrGetAgentModulesInvalidQuery = NewHttpError(500, "Modules.GetAgentModules.InvalidQuery", "invalid agent modules query")
var ErrGetAgentModulesInvalidAgentData = NewHttpError(500, "Modules.GetAgentModules.InvalidAgentData", "invalid agent module data")
var ErrGetAgentModulesDetailsNotFound = NewHttpError(404, "Modules.GetAgentModules.DetailsNotFound", "failed to retrieve agents modules details")
var ErrGetAgentModuleAgentNotFound = NewHttpError(404, "Modules.GetAgentModule.AgentNotFound", "agent not found")
var ErrGetAgentModuleAgentPoliceNotFound = NewHttpError(404, "Modules.GetAgentModule.AgentPoliceNotFound", "agent policies not found")
var ErrGetAgentModuleInvalidAgentPoliceData = NewHttpError(500, "Modules.GetAgentModule.InvalidAgentPoliceData", "invalid agent policies data")
var ErrGetAgentBModuleInvalidAgentPoliceData = NewHttpError(500, "Modules.GetAgentBModule.InvalidAgentPoliceData", "invalid agent policies data")
var ErrGetGroupModulesGroupNotFound = NewHttpError(404, "Modules.GetGroupModules.GroupNotFound", "group not found")
var ErrGetGroupModulesGroupPoliciesNotFound = NewHttpError(404, "Modules.GetGroupModules.GroupPoliciesNotFound", "group policies not found")
var ErrGetGroupModulesInvalidGroupPoliciesData = NewHttpError(500, "Modules.GetGroupModules.InvalidGroupPoliciesData", "invalid group policies data")
var ErrGetGroupModulesInvalidGroupQuery = NewHttpError(500, "Modules.GetGroupModules.InvalidGroupQuery", "invalid group modules query")
var ErrGetGroupModulesInvalidGroupData = NewHttpError(500, "Modules.GetGroupModules.InvalidGroupData", "invalid group modules data")
var ErrGetGroupModulesDetailsNotFound = NewHttpError(404, "Modules.GetGroupModules.DetailsNotFound", "failed to retrieve group modules details")
var ErrGetGroupModuleGroupNotFound = NewHttpError(404, "Modules.GetGroupModule.GroupNotFound", "group not found")
var ErrGetGroupModuleGroupPoliciesNotFound = NewHttpError(404, "Modules.GetGroupModule.GroupPoliciesNotFound", "group policies not found")
var ErrGetGroupModuleInvalidGroupPoliciesData = NewHttpError(500, "Modules.GetGroupModule.InvalidGroupPoliciesData", "invalid group policies data")
var ErrGetGroupBModuleInvalidGroupPoliciesData = NewHttpError(500, "Modules.GetGroupBModule.InvalidGroupPoliciesData", "invalid group policies data")
var ErrGetPolicyModulesPolicyNotFound = NewHttpError(404, "Modules.GetPolicyModules.PolicyNotFound", "policy not found")
var ErrGetPolicyModulesInvalidPolicyData = NewHttpError(500, "Modules.GetPolicyModules.InvalidPolicyData", "invalid policy data")
var ErrGetPolicyModulesInvalidPolicyQuery = NewHttpError(500, "Modules.GetPolicyModules.InvalidPolicyQuery", "invalid policy modules query")
var ErrGetPolicyModulesInvalidModulesQuery = NewHttpError(500, "Modules.GetPolicyModules.InvalidModulesQuery", "invalid system modules query")
var ErrGetPolicyModulesDetailsNotFound = NewHttpError(404, "Modules.GetPolicyModules.DetailsNotFound", "failed to retrieve policy modules details")
var ErrGetPolicyModulePolicyNotFound = NewHttpError(404, "Modules.GetPolicyModule.PolicyNotFound", "policy not found")
var ErrGetPolicyModuleInvalidPolicyData = NewHttpError(500, "Modules.GetPolicyModule.InvalidPolicyData", "invalid policy data")
var ErrPatchPolicyModulePolicyNotFound = NewHttpError(404, "Modules.PatchPolicyModule.PolicyNotFound", "policy not found")
var ErrPatchPolicyModuleInvalidPolicyData = NewHttpError(500, "Modules.PatchPolicyModule.InvalidPolicyData", "invalid policy data")
var ErrPatchPolicyModuleDuplicatedModule = NewHttpError(403, "Modules.PatchPolicyModule.DuplicatedModule", "found duplicate modules")
var ErrPatchPolicyModuleNewModuleInvalid = NewHttpError(400, "Modules.PatchPolicyModule.NewModuleInvalid", "failed to valid new module data")
var ErrPatchPolicyModuleAcceptFail = NewHttpError(500, "Modules.PatchPolicyModule.AcceptFail", "failed accept system changes")
var ErrPatchPolicyModuleActionNotFound = NewHttpError(404, "Modules.PatchPolicyModule.ActionNotFound", "action not found or unknown")
var ErrDeletePolicyModulePolicyNotFound = NewHttpError(404, "Modules.DeletePolicyModule.PolicyNotFound", "policy not found")
var ErrDeletePolicyModuleInvalidPolicyData = NewHttpError(500, "Modules.DeletePolicyModule.InvalidPolicyData", "invalid policy data")
var ErrGetModulesInvalidModulesQuery = NewHttpError(500, "Modules.GetModules.InvalidModulesQuery", "invalid system modules query")
var ErrCreateModuleInvalidInfo = NewHttpError(400, "Modules.CreateModule.InvalidInfo", "failed to valid module info")
var ErrCreateModuleGetCountFail = NewHttpError(404, "Modules.CreateModule.GetCountFail", "failed to get number of system module")
var ErrCreateModuleSecondSystemModule = NewHttpError(400, "Modules.CreateModule.SecondSystemModule", "failed to create second system module")
var ErrCreateModuleLoadFail = NewHttpError(500, "Modules.CreateModule.LoadFail", "failed to load module")
var ErrCreateModuleValidationFail = NewHttpError(500, "Modules.CreateModule.ValidationFail", "failed to validate module")
var ErrCreateModuleStoreS3Fail = NewHttpError(500, "Modules.CreateModule.StoreS3Fail", "failed to store module to s3")
var ErrCreateModuleStoreDBFail = NewHttpError(500, "Modules.CreateModule.StoreDBFail", "failed to store module to db")
var ErrDeleteModuleServiceNotFound = NewHttpError(404, "Modules.DeleteModule.ServiceNotFound", "services not found")
var ErrDeleteModuleDeleteFail = NewHttpError(500, "Modules.DeleteModule.DeleteFail", "failed to delete policy modules")
var ErrDeleteModuleDeleteFilesFail = NewHttpError(500, "Modules.DeleteModule.DeleteFilesFail", "failed to delete system module files")
var ErrGetModuleVersionsInvalidModulesQuery = NewHttpError(500, "Modules.GetModuleVersions.InvalidModulesQuery", "invalid system modules query")
var ErrPatchModuleVersionAcceptReleaseChangesFail = NewHttpError(500, "Modules.PatchModuleVersion.AcceptReleaseChangesFail", "failed to accept changes for released module")
var ErrPatchModuleVersionAcceptSystemChangesFail = NewHttpError(500, "Modules.PatchModuleVersion.AcceptSystemChangesFail", "failed accept system changes")
var ErrPatchModuleVersionUpdateFail = NewHttpError(500, "Modules.PatchModuleVersion.UpdateFail", "failed to update module into db")
var ErrPatchModuleVersionBuildFilesFail = NewHttpError(500, "Modules.PatchModuleVersion.BuildFilesFail", "failed to build module files")
var ErrPatchModuleVersionUpdateS3Fail = NewHttpError(500, "Modules.PatchModuleVersion.UpdateS3Fail", "failed to update module to s3")
var ErrPatchModuleVersionServiceNotFound = NewHttpError(404, "Modules.PatchModuleVersion.ServiceNotFound", "services not found")
var ErrCreateModuleVersionGetDraftNumberFail = NewHttpError(404, "Modules.CreateModuleVersion.GetDraftNumberFail", "failed to get number of system module drafts")
var ErrCreateModuleVersionSecondSystemModuleDraft = NewHttpError(400, "Modules.CreateModuleVersion.SecondSystemModuleDraft", "failed to create system module second draft")
var ErrCreateModuleVersionInvalidModuleVersionFormat = NewHttpError(500, "Modules.CreateModuleVersion.InvalidModuleVersionFormat", "invalid new system module version format")
var ErrCreateModuleVersionInvalidModuleVersion = NewHttpError(500, "Modules.CreateModuleVersion.InvalidModuleVersion", "invalid new system module version")
var ErrCreateModuleVersionBuildFilesFail = NewHttpError(500, "Modules.CreateModuleVersion.BuildFilesFail", "failed to build module files")
var ErrCreateModuleVersionValidationFail = NewHttpError(500, "Modules.CreateModuleVersion.ValidationFail", "failed to validate module")
var ErrCreateModuleVersionStoreS3Fail = NewHttpError(500, "Modules.CreateModuleVersion.StoreS3Fail", "failed to store module to s3")
var ErrCreateModuleVersionStoreDBFail = NewHttpError(500, "Modules.CreateModuleVersion.StoreDBFail", "failed to store module to db")
var ErrDeleteModuleVersionGetVersionNumberFail = NewHttpError(404, "Modules.DeleteModuleVersion.GetVersionNumberFail", "failed to get number of system module versions")
var ErrDeleteModuleVersionDeleteLastVersionFail = NewHttpError(404, "Modules.DeleteModuleVersion.DeleteLastVersionFail", "can't delete last module version")
var ErrDeleteModuleVersionDeleteFilesFail = NewHttpError(500, "Modules.DeleteModuleVersion.DeleteFilesFail", "failed to delete system module files")
var ErrGetModuleVersionUpdatesInvalidPolicyData = NewHttpError(500, "Modules.GetModuleVersionUpdates.InvalidPolicyData", "invalid policy data")
var ErrGetModuleVersionFilesListenFail = NewHttpError(500, "Modules.GetModuleVersionFiles.ListenFail", "failed to listening module files")
var ErrGetModuleVersionFileParsePathFail = NewHttpError(403, "Modules.GetModuleVersionFile.ParsePathFail", "failed to parse path to file")
var ErrGetModuleVersionFileReadFail = NewHttpError(500, "Modules.GetModuleVersionFile.ReadFail", "failed to read module file")
var ErrPatchModuleVersionFileUpdateFail = NewHttpError(403, "Modules.PatchModuleVersionFile.UpdateFail", "failed to update released module")
var ErrPatchModuleVersionFileParsePathFail = NewHttpError(403, "Modules.PatchModuleVersionFile.ParsePathFail", "failed to parse path to object")
var ErrPatchModuleVersionFileParseModuleFileFail = NewHttpError(500, "Modules.PatchModuleVersionFile.ParseModuleFileFail", "failed to parse module file content")
var ErrPatchModuleVersionFileWriteModuleFileFail = NewHttpError(500, "Modules.PatchModuleVersionFile.WriteModuleFileFail", "failed to write module file")
var ErrPatchModuleVersionFileWriteModuleObjectFail = NewHttpError(500, "Modules.PatchModuleVersionFile.WriteModuleObjectFail", "failed to write module object")
var ErrPatchModuleVersionFileParseNewpathFail = NewHttpError(403, "Modules.PatchModuleVersionFile.ParseNewpathFail", "failed to parse newpath to object")
var ErrPatchModuleVersionFileObjectNotFound = NewHttpError(500, "Modules.PatchModuleVersionFile.ObjectNotFound", "failed to find object by path")
var ErrPatchModuleVersionFilePathIdentical = NewHttpError(500, "Modules.PatchModuleVersionFile.PathIdentical", "newpath is identical to path")
var ErrPatchModuleVersionFileObjectMoveFail = NewHttpError(500, "Modules.PatchModuleVersionFile.ObjectMoveFail", "failed to move object")
var ErrPatchModuleVersionFileGetFilesFail = NewHttpError(500, "Modules.PatchModuleVersionFile.GetFilesFail", "failed to get files by path")
var ErrPatchModuleVersionFileActionNotFound = NewHttpError(404, "Modules.PatchModuleVersionFile.ActionNotFound", "action not found or unknown")
var ErrPatchModuleVersionFileSystemModuleUpdateFail = NewHttpError(500, "Modules.PatchModuleVersionFile.SystemModuleUpdateFail", "failed to update system module")
var ErrGetModuleVersionOptionMakeJsonFail = NewHttpError(500, "Modules.GetModuleVersionOption.MakeJsonFail", "failed to make json")
var ErrGetModuleVersionOptionParseJsonFail = NewHttpError(500, "Modules.GetModuleVersionOption.ParseJsonFail", "failed to parse json")
var ErrGetModuleVersionOptionNotFound = NewHttpError(404, "Modules.GetModuleVersionOption.NotFound", "option not found")
var ErrModulesInvalidParameterValue = NewHttpError(400, "Modules.InvalidParameterValue", "invalid parameter value type")
var ErrModulesFailedToEncryptSecureConfig = NewHttpError(500, "Modules.FailedToEncryptSecureConfig", "failed to encrypt secure config")
var ErrModulesFailedToDecryptSecureConfig = NewHttpError(500, "Modules.FailedToDecryptSecureConfig", "failed to decrypt secure config")
var ErrModulesDataNotEncryptedOnDBInsert = NewHttpError(500, "Modules.DataNotEncryptedOnDBInsert", "module data not encrypted on inserting to DB")
var ErrModulesFailedToCompareChanges = NewHttpError(500, "Modules.FailedToCompareChanges", "failed to compare requested module data")
var ErrModulesFailedRefreshModulesCache = NewHttpError(500, "Modules.FailedRefreshModulesCache", "failed to refresh modules cache")

// notifications

var ErrNotificationsSubscribesEmptyList = NewHttpError(500, "Notifications.SubscribesEmptyList", "subscribes empty list after filtering")
var ErrNotificationsInvalidServiceValue = NewHttpError(500, "Notifications.InvalidServiceValue", "invalid service value")
var ErrNotificationsUpgradeWSConnFail = NewHttpError(500, "Notifications.UpgradeWSConnFail", "failed to upgrade ws connection")
var ErrNotificationsServiceNotFound = NewHttpError(500, "Notifications.ServiceNotFound", "service not found")

// options

var ErrOptionsInvalidRequestData = NewHttpError(400, "Options.InvalidRequestData", "invalid global option, list request data")
var ErrOptionsInvalidQuery = NewHttpError(500, "Options.InvalidQuery", "invalid global option, list query")
var ErrOptionsInvalidData = NewHttpError(500, "Options.InvalidData", "invalid global option, data")

// policies

var ErrPoliciesInvalidRequest = NewHttpError(400, "Policies.InvalidRequest", "invalid policy request data")
var ErrPoliciesNotFound = NewHttpError(404, "Policies.NotFound", "policy not found")
var ErrPoliciesInvalidData = NewHttpError(500, "Policies.InvalidData", "invalid policy data")
var ErrPoliciesInvalidQuery = NewHttpError(500, "Policies.InvalidQuery", "invalid policy query")
var ErrGetPoliciesSystemModulesNotFound = NewHttpError(404, "Policies.GetPolicies.SystemModulesNotFound", "internal error on retrieving system modules")
var ErrGetPoliciesDetailsNotFound = NewHttpError(404, "Policies.GetPolicies.DetailsNotFound", "internal error on retrieving policies details")
var ErrGetPoliciesModulesNotFound = NewHttpError(404, "Policies.GetPolicies.ModulesNotFound", "policy modules not found")
var ErrGetPoliciesInvalidGroupData = NewHttpError(500, "Policies.GetPolicies.InvalidGroupData", "invalid policy group data")
var ErrGetPoliciesInvalidModuleData = NewHttpError(500, "Policies.GetPolicies.InvalidModuleData", "invalid policy module data")
var ErrGetPoliciesGroupsNotFound = NewHttpError(404, "Policies.GetPolicies.GroupsNotFound", "policy groups not found")
var ErrGetPolicySystemModulesNotFound = NewHttpError(404, "Policies.GetPolicy.SystemModulesNotFound", "internal error on retrieving system modules")
var ErrGetPolicyDetailsNotFound = NewHttpError(404, "Policies.GetPolicy.DetailsNotFound", "internal error on retrieving policy details")
var ErrGetPolicyModulesNotFound = NewHttpError(404, "Policies.GetPolicy.ModulesNotFound", "policy modules not found")
var ErrGetPolicyInvalidModuleData = NewHttpError(500, "Policies.GetPolicy.InvalidModuleData", "invalid policy module data")
var ErrGetPolicyGroupsNotFound = NewHttpError(404, "Policies.GetPolicy.GroupsNotFound", "policy groups not found")
var ErrPatchPolicyGroupInvalidGroupData = NewHttpError(500, "Policies.PatchPolicyGroup.InvalidGroupData", "invalid group data")
var ErrCreatePolicySourceNotFound = NewHttpError(404, "Policies.CreatePolicy.SourceNotFound", "source policy not found")
var ErrCreatePolicyCreateFail = NewHttpError(500, "Policies.CreatePolicy.CreateFail", "failed to create policy to db")
var ErrCreatePolicyModulesNotFound = NewHttpError(500, "Policies.CreatePolicy.ModulesNotFound", "failed to get policy modules")
var ErrCreatePolicyCreateModulesFail = NewHttpError(500, "Policies.CreatePolicy.CreateModulesFail", "failed to create policy modules to db")
var ErrDeletePolicyPolicyLinkedToGroups = NewHttpError(403, "Policies.DeletePolicy.PolicyLinkedToGroups", "failed to remove linked policy to groups")
var ErrDeletePolicySystemPolicy = NewHttpError(403, "Policies.DeletePolicy.SystemPolicy", "failed to remove system policy")
var ErrDeletePolicyModulesNotFound = NewHttpError(404, "Policies.DeletePolicy.ModulesNotFound", "policy modules not found")
var ErrDeletePolicyGroupsNotFound = NewHttpError(404, "Policies.DeletePolicy.GroupsNotFound", "policy groups not found")

// porting

var ErrPortingModuleNotFound = NewHttpError(404, "Porting.ModuleNotFound", "system module not found")
var ErrExportInvalidModuleData = NewHttpError(500, "Porting.Export.InvalidModuleData", "invalid system module data")
var ErrExportLoadFilesFail = NewHttpError(500, "Porting.Export.LoadFilesFail", "failed to load system module files")
var ErrExportBuildConfigFail = NewHttpError(500, "Porting.Export.BuildConfigFail", "failed to build system module config")
var ErrExportAddFileFail = NewHttpError(500, "Porting.Export.AddFileFail", "failed to add system module file")
var ErrExportWriteFileFail = NewHttpError(500, "Porting.Export.WriteFileFail", "failed to write system module file")
var ErrExportCloseArchiveFail = NewHttpError(500, "Porting.Export.CloseArchiveFail", "failed to close system module archive")
var ErrImportReadArchiveFail = NewHttpError(400, "Porting.Import.ReadArchiveFail", "failed to read system module archive")
var ErrImportParseArchiveFail = NewHttpError(400, "Porting.Import.ParseArchiveFail", "failed to parse system module archive")
var ErrImportParseConfigFail = NewHttpError(400, "Porting.Import.ParseConfigFail", "failed to parse system module config")
var ErrImportParseFileFail = NewHttpError(400, "Porting.Import.ParseFileFail", "failed to parse system module file info")
var ErrImportValidateConfigFail = NewHttpError(400, "Porting.Import.ValidateConfigFail", "failed to validate system module config")
var ErrImportOverrideNotPermitted = NewHttpError(403, "Porting.Import.OverrideNotPermitted", "override system module version not permitted")
var ErrImportStoreS3Fail = NewHttpError(500, "Porting.Import.StoreS3Fail", "failed to store system module to S3")
var ErrImportStoreDBFail = NewHttpError(500, "Porting.Import.StoreDBFail", "failed to store system module to DB")

// proto

var ErrProtoInvalidRequest = NewHttpError(400, "Proto.InvalidRequest", "failed to validate auth token request")
var ErrProtoCreateTokenFail = NewHttpError(400, "Proto.CreateTokenFail", "failed to make auth token")
var ErrProtoInvalidToken = NewHttpError(400, "Proto.InvalidToken", "failed to valid auth token")

// proto_proto

var ErrProtoSockMismatch = NewHttpError(400, "Proto.SockMismatch", "mismatch socket type")
var ErrProtoInvalidAgentID = NewHttpError(400, "Proto.InvalidAgentID", "error validating agent ID")
var ErrProtoNoServiceInfo = NewHttpError(400, "Proto.NoServiceInfo", "error getting service information")
var ErrProtoUpgradeFail = NewHttpError(400, "Proto.UpgradeFail", "error upgrading to websockets")

// roles

var ErrRolesInvalidRequest = NewHttpError(400, "Roles.InvalidRequest", "invalid role request data")
var ErrRolesInvalidData = NewHttpError(500, "Roles.InvalidData", "invalid role data")

// services

var ErrServicesInvalidRequest = NewHttpError(400, "Services.InvalidRequest", "invalid service request data")
var ErrServicesInvalidData = NewHttpError(500, "Services.InvalidData", "invalid service data")
var ErrServicesNotFound = NewHttpError(404, "Services.NotFound", "service not found")

// tags

var ErrTagsInvalidRequest = NewHttpError(400, "Tags.InvalidRequest", "invalid tags request data")
var ErrTagsMappersNotFound = NewHttpError(400, "Tags.MappersNotFound", "failed to get tag mappers by query")
var ErrTagsInvalidQuery = NewHttpError(500, "Tags.InvalidQuery", "invalid tags query")

// tenants

var ErrTenantsInvalidRequest = NewHttpError(400, "Tenants.InvalidRequest", "invalid tenant request data")
var ErrTenantsInvalidData = NewHttpError(500, "Tenants.InvalidData", "invalid tenant data")
var ErrTenantsNotFound = NewHttpError(404, "Tenants.NotFound", "tenant not found")

// upgrades

var ErrGetAgentsUpgradesInvalidRequest = NewHttpError(400, "Upgrades.GetAgentsUpgrades.InvalidRequest", "invalid agents upgrades request data")
var ErrGetAgentsUpgradesInvalidData = NewHttpError(500, "Upgrades.GetAgentsUpgrades.InvalidData", "invalid agents upgrades data")
var ErrCreateAgentsUpgradesInvalidRequest = NewHttpError(400, "Upgrades.CreateAgentsUpgrades.InvalidRequest", "invalid agents upgrades request data")
var ErrCreateAgentsUpgradesAgentNotFound = NewHttpError(404, "Upgrades.CreateAgentsUpgrades.AgentNotFound", "agent binary record not found")
var ErrCreateAgentsUpgradesCreateTaskFail = NewHttpError(500, "Upgrades.CreateAgentsUpgrades.CreateTaskFail", "creating new upgrade tasks failed")
var ErrCreateAgentsUpgradesUpdateAgentBinariesFail = NewHttpError(400, "Upgrades.CreateAgentsUpgrades.UpdateAgentBinariesFail", "failed to update agent binaries")
var ErrGetLastAgentUpgradeLastUpgradeNotFound = NewHttpError(404, "Upgrades.GetLastAgentUpgrade.LastUpgradeNotFound", "last agent upgrade information not found")
var ErrGetLastAgentUpgradeInvalidLastData = NewHttpError(500, "Upgrades.GetLastAgentUpgrade.InvalidLastData", "invalid last agent upgrade data")
var ErrGetLastAgentUpgradeAgentNotFound = NewHttpError(404, "Upgrades.GetLastAgentUpgrade.AgentNotFound", "agent not found")
var ErrGetLastAgentUpgradeInvalidAgentData = NewHttpError(500, "Upgrades.GetLastAgentUpgrade.InvalidAgentData", "invalid agent data")
var ErrGetLastAgentUpgradeGroupNotFound = NewHttpError(404, "Upgrades.GetLastAgentUpgrade.GroupNotFound", "group not found")
var ErrGetLastAgentUpgradeInvalidGroupData = NewHttpError(500, "Upgrades.GetLastAgentUpgrade.InvalidGroupData", "invalid group data")
var ErrPatchLastAgentUpgradeInvalidAgentUpgradeInfo = NewHttpError(400, "Upgrades.PatchLastAgentUpgrade.InvalidAgentUpgradeInfo", "failed to validate last agent upgrade information")
var ErrPatchLastAgentUpgradeAgentNotFound = NewHttpError(404, "Upgrades.PatchLastAgentUpgrade.AgentNotFound", "agent not found")
var ErrPatchLastAgentUpgradeInvalidAgentData = NewHttpError(500, "Upgrades.PatchLastAgentUpgrade.InvalidAgentData", "invalid agent data")
var ErrPatchLastAgentUpgradeAgentBinaryNotFound = NewHttpError(404, "Upgrades.PatchLastAgentUpgrade.AgentBinaryNotFound", "agent binary record not found")
var ErrPatchLastAgentUpgradeUpdateAgentBinariesFail = NewHttpError(400, "Upgrades.PatchLastAgentUpgrade.UpdateAgentBinariesFail", "failed to update agent binaries")
var ErrPatchLastAgentUpgradeLastUpgradeInfoNotFound = NewHttpError(404, "Upgrades.PatchLastAgentUpgrade.LastUpgradeInfoNotFound", "last agent upgrade information not found")

// users

var ErrUsersNotFound = NewHttpError(404, "Users.NotFound", "user not found")
var ErrUsersInvalidData = NewHttpError(500, "Users.InvalidData", "invalid user data")
var ErrUsersInvalidRequest = NewHttpError(400, "Users.InvalidRequest", "invalid user request data")
var ErrChangePasswordCurrentUserInvalidPassword = NewHttpError(400, "Users.ChangePasswordCurrentUser.InvalidPassword", "failed to validate user password")
var ErrChangePasswordCurrentUserInvalidCurrentPassword = NewHttpError(403, "Users.ChangePasswordCurrentUser.InvalidCurrentPassword", "invalid current password")
var ErrChangePasswordCurrentUserInvalidNewPassword = NewHttpError(400, "Users.ChangePasswordCurrentUser.InvalidNewPassword", "invalid new password form data")
var ErrGetUserModelsNotFound = NewHttpError(404, "Users.GetUser.ModelsNotFound", "user linked models not found")
var ErrCreateUserInvalidUser = NewHttpError(400, "Users.CreateUser.InvalidUser", "failed to validate user")
var ErrPatchUserModelsNotFound = NewHttpError(404, "Users.PatchUser.ModelsNotFound", "user linked models not found")
var ErrDeleteUserModelsNotFound = NewHttpError(404, "Users.DeleteUser.ModelsNotFound", "user linked models not found")

// versions

var ErrVersionsInvalidRequest = NewHttpError(400, "Versions.InvalidRequest", "invalid versions request data")
var ErrVersionsMapperNotFound = NewHttpError(400, "Versions.MapperNotFound", "failed to get version mappers by query")
var ErrVersionsInvalidQuery = NewHttpError(500, "Versions.InvalidQuery", "invalid versions query")
