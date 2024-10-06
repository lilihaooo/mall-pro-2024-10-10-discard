package product

import (
	"admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

func (*TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	var apiGroup = "tag"
	tagRouter := Router.Group(apiGroup)
	tagApi := v1.ApiGroupApp.ProductApiGroup.TagApi
	{
		common.RegisterRouteWithComment(tagRouter, "GET", "list", "获取标签列表", apiGroup, tagApi.TagList)
	}
}
