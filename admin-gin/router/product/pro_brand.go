package product

import (
	"admin-gin/api/v1"
	"github.com/gin-gonic/gin"
)

type BrandRouter struct{}

func (*BrandRouter) InitBrandRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("brand")
	brandApi := v1.ApiGroupApp.ProductApiGroup.BrandApi
	{
		customerRouter.GET("/list", brandApi.BrandList) // 创建客户
	}
}
