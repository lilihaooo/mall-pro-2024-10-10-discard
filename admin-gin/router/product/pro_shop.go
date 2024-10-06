package product

import (
	"admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type ShopRouter struct{}

func (*ShopRouter) InitShopRouter(Router *gin.RouterGroup) {

	var apiGroup = "shop"

	customerRouter := Router.Group(apiGroup)
	ShopApi := v1.ApiGroupApp.ProductApiGroup.ShopApi
	{
		common.RegisterRouteWithComment(customerRouter, "GET", "list", "获取商店列表", apiGroup, ShopApi.ShopList)
	}
}
