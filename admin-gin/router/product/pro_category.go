package product

import (
	v1 "admin-gin/api/v1"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct{}

func (p *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("category")
	categoryApi := v1.ApiGroupApp.ProductApiGroup.CategoryApi
	{
		customerRouter.GET("", categoryApi.GetCategoryTree) // 创建客户
	}
}
