package common

import (
	"admin-gin/global"
	"admin-gin/model/system"
	"github.com/gin-gonic/gin"
)

// 自定义注册路由, 并保留路由信息
func RegisterRouteWithComment(iRoutes gin.IRoutes, method, path, description, apiGroup string, handler gin.HandlerFunc) {
	iRoutes.Handle(method, path, handler)
	allPath := "/" + apiGroup + "/" + path

	// 保存路由信息和注释
	global.Apis = append(global.Apis, system.SysApi{
		Method:      method,
		Path:        allPath,
		Description: description,
		ApiGroup:    apiGroup,
		Type:        system.CustomType,
	})
}
