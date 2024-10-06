package product

import (
	v1 "admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct{}

func (p *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	var apiGroup = "category"
	customerRouter := Router.Group(apiGroup)
	categoryApi := v1.ApiGroupApp.ProductApiGroup.CategoryApi
	{
		common.RegisterRouteWithComment(customerRouter, "GET", "", "创建分类", apiGroup, categoryApi.GetCategoryTree)
	}
}
