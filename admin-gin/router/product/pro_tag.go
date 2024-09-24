package product

import (
	"admin-gin/api/v1"
	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

func (*TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	tagRouter := Router.Group("tag")
	tagApi := v1.ApiGroupApp.ProductApiGroup.TagApi
	{
		tagRouter.GET("/list", tagApi.TagList) // 创建客户
	}
}
