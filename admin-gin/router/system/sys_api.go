package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/middleware"
	"admin-gin/router/common"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	var apiGroup = "api"
	apiRouter := Router.Group(apiGroup).Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group(apiGroup)
	apiPublicRouterWithoutRecord := RouterPub.Group(apiGroup)

	// api handles
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi

	{
		common.RegisterRouteWithComment(apiRouter, "POST", "createApi", "创建Api", apiGroup, apiRouterApi.CreateApi)
		common.RegisterRouteWithComment(apiRouter, "POST", "deleteApi", "删除Api", apiGroup, apiRouterApi.DeleteApi)
		common.RegisterRouteWithComment(apiRouter, "POST", "getApiById", "获取单条Api消息", apiGroup, apiRouterApi.GetApiById)
		common.RegisterRouteWithComment(apiRouter, "POST", "updateApi", "更新api", apiGroup, apiRouterApi.UpdateApi)
		common.RegisterRouteWithComment(apiRouter, "DELETE", "deleteApisByIds", "删除选中api", apiGroup, apiRouterApi.DeleteApisByIds)

	}
	{
		common.RegisterRouteWithComment(apiRouterWithoutRecord, "POST", "getAllApis", "获取所有api", apiGroup, apiRouterApi.GetAllApis)
		common.RegisterRouteWithComment(apiRouterWithoutRecord, "POST", "getApiList", "获取Api列表", apiGroup, apiRouterApi.GetApiList)
		//common.RegisterRouteWithComment(apiRouterWithoutRecord, "POST", "renew", "更新全部api", apiGroup, apiRouterApi.RenewApi)
	}
	{
		common.RegisterRouteWithComment(apiPublicRouterWithoutRecord, "GET", "freshCasbin", "刷新casbin权限", apiGroup, apiRouterApi.FreshCasbin)
	}
}
