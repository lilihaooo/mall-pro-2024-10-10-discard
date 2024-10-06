package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	apiGroup := "init"
	initRouter := Router.Group(apiGroup)
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		common.RegisterRouteWithComment(initRouter, "POST", "initdb", "初始化数据库", apiGroup, dbApi.InitDB)
		common.RegisterRouteWithComment(initRouter, "POST", "checkdb", "检测是否需要初始化数据库", apiGroup, dbApi.CheckDB)
	}
}
