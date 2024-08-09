package system

import "admin-gin/service"

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	SystemApi
	CasbinApi
	AutoCodeApi
	SystemApiApi
	AuthorityApi
	DictionaryApi
	AuthorityMenuApi
	OperationRecordApi
	AutoCodeHistoryApi
	DictionaryDetailApi
	AuthorityBtnApi
	SysExportTemplateApi
}

var (
	apiService              = service.GroupApp.SystemServiceGroup.ApiService
	jwtService              = service.GroupApp.SystemServiceGroup.JwtService
	menuService             = service.GroupApp.SystemServiceGroup.MenuService
	userService             = service.GroupApp.SystemServiceGroup.UserService
	initDBService           = service.GroupApp.SystemServiceGroup.InitDBService
	casbinService           = service.GroupApp.SystemServiceGroup.CasbinService
	autoCodeService         = service.GroupApp.SystemServiceGroup.AutoCodeService
	baseMenuService         = service.GroupApp.SystemServiceGroup.BaseMenuService
	authorityService        = service.GroupApp.SystemServiceGroup.AuthorityService
	dictionaryService       = service.GroupApp.SystemServiceGroup.DictionaryService
	systemConfigService     = service.GroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService  = service.GroupApp.SystemServiceGroup.OperationRecordService
	autoCodeHistoryService  = service.GroupApp.SystemServiceGroup.AutoCodeHistoryService
	dictionaryDetailService = service.GroupApp.SystemServiceGroup.DictionaryDetailService
	authorityBtnService     = service.GroupApp.SystemServiceGroup.AuthorityBtnService
)
