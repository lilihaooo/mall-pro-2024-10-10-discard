package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	apiGroup := "base"
	baseRouter := Router.Group(apiGroup)
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		common.RegisterRouteWithComment(baseRouter, "POST", "login", "登录", apiGroup, baseApi.Login)
		common.RegisterRouteWithComment(baseRouter, "POST", "captcha", "验证码", apiGroup, baseApi.Captcha)
	}
	return baseRouter
}
