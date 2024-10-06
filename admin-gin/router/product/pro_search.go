package product

import (
	"admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type SearchRouter struct{}

func (*SearchRouter) InitSearchRouter(Router *gin.RouterGroup) {
	var apiGroup = "search"
	searchRouter := Router.Group(apiGroup)
	searchApi := v1.ApiGroupApp.ProductApiGroup.SearchApi
	{
		common.RegisterRouteWithComment(searchRouter, "GET", "hotSearch/list", "获取热搜列表", apiGroup, searchApi.HotList)
		common.RegisterRouteWithComment(searchRouter, "GET", "suggestion/list", "获取推荐词条列表", apiGroup, searchApi.SuggestionList)

	}
}
