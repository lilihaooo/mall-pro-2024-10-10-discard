package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/middleware"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	apiGroup := "menu"
	menuRouter := Router.Group(apiGroup).Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group(apiGroup)
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		common.RegisterRouteWithComment(menuRouter, "POST", "addBaseMenu", "新增菜单", apiGroup, authorityMenuApi.AddBaseMenu)
		common.RegisterRouteWithComment(menuRouter, "POST", "addMenuAuthority", "增加menu和角色关联关系", apiGroup, authorityMenuApi.AddMenuAuthority)
		common.RegisterRouteWithComment(menuRouter, "POST", "deleteBaseMenu", "删除菜单", apiGroup, authorityMenuApi.DeleteBaseMenu)
		common.RegisterRouteWithComment(menuRouter, "POST", "updateBaseMenu", "更新菜单", apiGroup, authorityMenuApi.UpdateBaseMenu)
	}
	{
		common.RegisterRouteWithComment(menuRouterWithoutRecord, "POST", "getMenu", "获取菜单树", apiGroup, authorityMenuApi.GetMenu)
		common.RegisterRouteWithComment(menuRouterWithoutRecord, "POST", "getMenuList", "分页获取基础menu列表", apiGroup, authorityMenuApi.GetMenuList)
		common.RegisterRouteWithComment(menuRouterWithoutRecord, "POST", "getBaseMenuTree", "获取用户动态路由", apiGroup, authorityMenuApi.GetBaseMenuTree)
		common.RegisterRouteWithComment(menuRouterWithoutRecord, "POST", "getMenuAuthority", "获取指定角色menu", apiGroup, authorityMenuApi.GetMenuAuthority)
		common.RegisterRouteWithComment(menuRouterWithoutRecord, "POST", "getBaseMenuById", "根据id获取菜单", apiGroup, authorityMenuApi.GetBaseMenuById)

	}
	return menuRouter
}
