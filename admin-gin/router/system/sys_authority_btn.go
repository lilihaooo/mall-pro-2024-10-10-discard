package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type AuthorityBtnRouter struct{}

func (s *AuthorityBtnRouter) InitAuthorityBtnRouterRouter(Router *gin.RouterGroup) {
	apiGroup := "authorityBtn"
	//authorityRouter := Router.Group("authorityBtn").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authorityBtn")
	authorityBtnApi := v1.ApiGroupApp.SystemApiGroup.AuthorityBtnApi
	{
		common.RegisterRouteWithComment(authorityRouterWithoutRecord, "POST", "getAuthorityBtn", "获取角色按钮权限", apiGroup, authorityBtnApi.GetAuthorityBtn)
		common.RegisterRouteWithComment(authorityRouterWithoutRecord, "POST", "setAuthorityBtn", "设置角色按钮权限", apiGroup, authorityBtnApi.SetAuthorityBtn)
		common.RegisterRouteWithComment(authorityRouterWithoutRecord, "POST", "canRemoveAuthorityBtn", "判断是否能移除角色按钮权限", apiGroup, authorityBtnApi.CanRemoveAuthorityBtn)

	}
}
