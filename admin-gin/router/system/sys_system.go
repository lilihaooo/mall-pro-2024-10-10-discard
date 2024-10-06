package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/middleware"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	apiGroup := "system"
	sysRouter := Router.Group(apiGroup).Use(middleware.OperationRecord())
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		common.RegisterRouteWithComment(sysRouter, "POST", "getSystemConfig", "获取配置文件内容", apiGroup, systemApi.GetSystemConfig)
		common.RegisterRouteWithComment(sysRouter, "POST", "setSystemConfig", "设置配置文件内容", apiGroup, systemApi.SetSystemConfig)
		common.RegisterRouteWithComment(sysRouter, "POST", "getServerInfo", "获取服务器信息", apiGroup, systemApi.GetServerInfo)
		common.RegisterRouteWithComment(sysRouter, "POST", "reloadSystem", "重启服务", apiGroup, systemApi.ReloadSystem)
	}
}
