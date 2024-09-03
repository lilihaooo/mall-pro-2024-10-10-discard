package product

import (
	"admin-gin/api/v1"
	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

func (*TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("tag")
	tagApi := v1.ApiGroupApp.ProductApiGroup.TagApi
	{
		customerRouter.GET("/list", tagApi.TagList) // 创建客户
	}
}
