package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/middleware"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	apiGroup := "user"
	userRouter := Router.Group(apiGroup).Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group(apiGroup)
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		common.RegisterRouteWithComment(userRouter, "POST", "admin_register", "管理员注册账号", apiGroup, baseApi.Register)
		common.RegisterRouteWithComment(userRouter, "POST", "changePassword", "用户修改密码", apiGroup, baseApi.ChangePassword)
		common.RegisterRouteWithComment(userRouter, "POST", "setUserAuthority", "设置用户权限", apiGroup, baseApi.SetUserAuthority)
		common.RegisterRouteWithComment(userRouter, "DELETE", "deleteUser", "删除用户", apiGroup, baseApi.DeleteUser)
		common.RegisterRouteWithComment(userRouter, "PUT", "setUserInfo", "设置用户信息", apiGroup, baseApi.SetUserInfo)
		common.RegisterRouteWithComment(userRouter, "PUT", "setSelfInfo", "设置自身信息", apiGroup, baseApi.SetSelfInfo)
		common.RegisterRouteWithComment(userRouter, "POST", "setUserAuthorities", "设置用户权限组", apiGroup, baseApi.SetUserAuthorities)
		common.RegisterRouteWithComment(userRouter, "POST", "resetPassword", "重置密码", apiGroup, baseApi.ResetPassword)

		common.RegisterRouteWithComment(userRouter, "GET", "getSetting", "获取用户vue app设置", apiGroup, baseApi.GetUserSettingData)
		common.RegisterRouteWithComment(userRouter, "POST", "setSetting", "保存用户vue app设置", apiGroup, baseApi.SetUserSettingData)

	}
	{
		common.RegisterRouteWithComment(userRouterWithoutRecord, "POST", "getUserList", "分页获取用户列表", apiGroup, baseApi.GetUserList)
		common.RegisterRouteWithComment(userRouterWithoutRecord, "GET", "getUserInfo", "获取自身信息", apiGroup, baseApi.GetUserInfo)
	}
}
