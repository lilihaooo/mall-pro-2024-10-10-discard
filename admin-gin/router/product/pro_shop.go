package product

import (
	"admin-gin/api/v1"
	"github.com/gin-gonic/gin"
)

type ShopRouter struct{}

func (*ShopRouter) InitShopRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("shop")
	ShopApi := v1.ApiGroupApp.ProductApiGroup.ShopApi
	{
		customerRouter.GET("/list", ShopApi.ShopList) // 创建客户
	}
}
