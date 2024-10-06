package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type AutoCodeHistoryRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeHistoryRouter(Router *gin.RouterGroup) {
	apiGroup := "autoCode"
	autoCodeHistoryRouter := Router.Group(apiGroup)
	autoCodeHistoryApi := v1.ApiGroupApp.SystemApiGroup.AutoCodeHistoryApi
	{
		common.RegisterRouteWithComment(autoCodeHistoryRouter, "POST", "getMeta", "根据id获取meta信息", apiGroup, autoCodeHistoryApi.First)
		common.RegisterRouteWithComment(autoCodeHistoryRouter, "POST", "rollback", "回滚", apiGroup, autoCodeHistoryApi.RollBack)
		common.RegisterRouteWithComment(autoCodeHistoryRouter, "POST", "delSysHistory", "删除回滚记录", apiGroup, autoCodeHistoryApi.Delete)
		common.RegisterRouteWithComment(autoCodeHistoryRouter, "POST", "getSysHistory", "获取回滚记录分页", apiGroup, autoCodeHistoryApi.GetList)
	}
}
