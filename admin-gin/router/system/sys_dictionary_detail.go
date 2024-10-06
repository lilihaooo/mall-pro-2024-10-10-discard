package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/middleware"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type DictionaryDetailRouter struct{}

func (s *DictionaryDetailRouter) InitSysDictionaryDetailRouter(Router *gin.RouterGroup) {

	apiGroup := "sysDictionaryDetail"
	dictionaryDetailRouter := Router.Group(apiGroup).Use(middleware.OperationRecord())
	dictionaryDetailRouterWithoutRecord := Router.Group(apiGroup)
	sysDictionaryDetailApi := v1.ApiGroupApp.SystemApiGroup.DictionaryDetailApi
	{
		common.RegisterRouteWithComment(dictionaryDetailRouter, "POST", "createSysDictionaryDetail", "新建SysDictionaryDetail", apiGroup, sysDictionaryDetailApi.CreateSysDictionaryDetail)
		common.RegisterRouteWithComment(dictionaryDetailRouter, "DELETE", "deleteSysDictionaryDetail", "删除SysDictionaryDetail", apiGroup, sysDictionaryDetailApi.DeleteSysDictionaryDetail)
		common.RegisterRouteWithComment(dictionaryDetailRouter, "PUT", "updateSysDictionaryDetail", "更新SysDictionaryDetail", apiGroup, sysDictionaryDetailApi.UpdateSysDictionaryDetail)
	}
	{
		common.RegisterRouteWithComment(dictionaryDetailRouterWithoutRecord, "GET", "findSysDictionaryDetail", "根据ID获取SysDictionaryDetail", apiGroup, sysDictionaryDetailApi.FindSysDictionaryDetail)
		common.RegisterRouteWithComment(dictionaryDetailRouterWithoutRecord, "GET", "getSysDictionaryDetailList", "获取SysDictionaryDetail列表", apiGroup, sysDictionaryDetailApi.GetSysDictionaryDetailList)

	}
}
