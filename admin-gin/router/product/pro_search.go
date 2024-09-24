package product

import (
	"admin-gin/api/v1"
	"github.com/gin-gonic/gin"
)

type SearchRouter struct{}

func (*SearchRouter) InitSearchRouter(Router *gin.RouterGroup) {
	searchRouter := Router.Group("search")
	searchApi := v1.ApiGroupApp.ProductApiGroup.SearchApi
	{
		searchRouter.GET("/hotSearch/list", searchApi.HotList)         // 创建客户
		searchRouter.GET("/suggestion/list", searchApi.SuggestionList) // 搜索词条
	}
}
