package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup) {
	apiGroup := "autoCode"
	autoCodeRouter := Router.Group(apiGroup)
	autoCodeApi := v1.ApiGroupApp.SystemApiGroup.AutoCodeApi
	{
		common.RegisterRouteWithComment(autoCodeRouter, "GET", "getDB", "获取数据库", apiGroup, autoCodeApi.GetDB)
		common.RegisterRouteWithComment(autoCodeRouter, "GET", "getTables", "获取对应数据库的表", apiGroup, autoCodeApi.GetTables)
		common.RegisterRouteWithComment(autoCodeRouter, "GET", "getColumn", "获取指定表所有字段信息", apiGroup, autoCodeApi.GetColumn)
		common.RegisterRouteWithComment(autoCodeRouter, "POST", "preview", "获取自动创建代码预览", apiGroup, autoCodeApi.PreviewTemp)
		common.RegisterRouteWithComment(autoCodeRouter, "POST", "createTemp", "创建自动化代码", apiGroup, autoCodeApi.CreateTemp)
		common.RegisterRouteWithComment(autoCodeRouter, "POST", "createPackage", "创建package包", apiGroup, autoCodeApi.CreatePackage)
		common.RegisterRouteWithComment(autoCodeRouter, "POST", "getPackage", "获取package包", apiGroup, autoCodeApi.GetPackage)
		common.RegisterRouteWithComment(autoCodeRouter, "POST", "delPackage", "删除package包", apiGroup, autoCodeApi.DelPackage)
		common.RegisterRouteWithComment(autoCodeRouter, "POST", "createPlug", "自动插件包模板", apiGroup, autoCodeApi.AutoPlug)
		common.RegisterRouteWithComment(autoCodeRouter, "POST", "installPlugin", "自动安装插件", apiGroup, autoCodeApi.InstallPlugin)
		common.RegisterRouteWithComment(autoCodeRouter, "POST", "pubPlug", "打包插件", apiGroup, autoCodeApi.PubPlug)

	}
}
