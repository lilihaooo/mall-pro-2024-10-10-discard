package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type JwtRouter struct{}

func (s *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	apiGroup := "jwt"
	jwtRouter := Router.Group(apiGroup)
	jwtApi := v1.ApiGroupApp.SystemApiGroup.JwtApi
	{
		common.RegisterRouteWithComment(jwtRouter, "POST", "jsonInBlacklist", "jwt加入黑名单", apiGroup, jwtApi.JsonInBlacklist)
	}
}
