package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/middleware"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	apiGroup := "authority"
	authorityRouter := Router.Group(apiGroup).Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group(apiGroup)
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		common.RegisterRouteWithComment(authorityRouter, "POST", "createAuthority", "创建角色", apiGroup, authorityApi.CreateAuthority)
		common.RegisterRouteWithComment(authorityRouter, "POST", "deleteAuthority", "删除角色", apiGroup, authorityApi.DeleteAuthority)
		common.RegisterRouteWithComment(authorityRouter, "PUT", "updateAuthority", "更新角色", apiGroup, authorityApi.UpdateAuthority)
		common.RegisterRouteWithComment(authorityRouter, "POST", "copyAuthority", "拷贝角色", apiGroup, authorityApi.CopyAuthority)
		common.RegisterRouteWithComment(authorityRouter, "POST", "setDataAuthority", "设置角色资源权限", apiGroup, authorityApi.SetDataAuthority)

	}
	{
		common.RegisterRouteWithComment(authorityRouterWithoutRecord, "POST", "getAuthorityList", "获取角色列表", apiGroup, authorityApi.GetAuthorityList)

	}
}
