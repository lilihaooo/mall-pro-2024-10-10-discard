package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/middleware"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type DictionaryRouter struct{}

func (s *DictionaryRouter) InitSysDictionaryRouter(Router *gin.RouterGroup) {
	apiGroup := "sysDictionary"
	sysDictionaryRouter := Router.Group(apiGroup).Use(middleware.OperationRecord())
	sysDictionaryRouterWithoutRecord := Router.Group(apiGroup)
	sysDictionaryApi := v1.ApiGroupApp.SystemApiGroup.DictionaryApi
	{
		common.RegisterRouteWithComment(sysDictionaryRouter, "POST", "createSysDictionary", "新建SysDictionary", apiGroup, sysDictionaryApi.CreateSysDictionary)
		common.RegisterRouteWithComment(sysDictionaryRouter, "DELETE", "deleteSysDictionary", "删除SysDictionary", apiGroup, sysDictionaryApi.DeleteSysDictionary)
		common.RegisterRouteWithComment(sysDictionaryRouter, "PUT", "updateSysDictionary", "更新SysDictionary", apiGroup, sysDictionaryApi.UpdateSysDictionary)

	}
	{
		common.RegisterRouteWithComment(sysDictionaryRouterWithoutRecord, "GET", "findSysDictionary", "根据ID获取SysDictionary", apiGroup, sysDictionaryApi.FindSysDictionary)
		common.RegisterRouteWithComment(sysDictionaryRouterWithoutRecord, "GET", "getSysDictionaryList", "获取SysDictionary列表", apiGroup, sysDictionaryApi.GetSysDictionaryList)
	}
}
