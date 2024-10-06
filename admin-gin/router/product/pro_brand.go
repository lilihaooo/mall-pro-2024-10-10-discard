package product

import (
	"admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type BrandRouter struct{}

func (*BrandRouter) InitBrandRouter(Router *gin.RouterGroup) {
	var apiGroup = "brand"
	customerRouter := Router.Group(apiGroup)
	brandApi := v1.ApiGroupApp.ProductApiGroup.BrandApi
	{
		common.RegisterRouteWithComment(customerRouter, "GET", "list", "获取品牌列表", apiGroup, brandApi.BrandList)
	}
}
