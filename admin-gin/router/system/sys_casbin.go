package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/middleware"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	apiGroup := "casbin"
	casbinRouter := Router.Group(apiGroup).Use(middleware.OperationRecord())
	casbinRouterWithoutRecord := Router.Group(apiGroup)
	casbinApi := v1.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		common.RegisterRouteWithComment(casbinRouter, "POST", "updateCasbin", "更新Casbin", apiGroup, casbinApi.UpdateCasbin)
	}
	{
		common.RegisterRouteWithComment(casbinRouterWithoutRecord, "POST", "getPolicyPathByAuthorityId", "根据AuthorityId获取策略路径", apiGroup, casbinApi.GetPolicyPathByAuthorityId)
	}
}
