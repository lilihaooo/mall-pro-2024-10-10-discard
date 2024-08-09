package product

import (
	v1 "admin-gin/api/v1"
	"github.com/gin-gonic/gin"
)

type GrabRouter struct{}

func (p *GrabRouter) InitGrabRouter(Router *gin.RouterGroup) {
	grabInfoRouter := Router.Group("grab")
	grabApi := v1.ApiGroupApp.ProductApiGroup.GrobApi
	{
		grabInfoRouter.GET("record", grabApi.GetGrabRecord) // 根据ID抓取数据
		grabInfoRouter.GET("log", grabApi.GetGrabLog)       // 根据ID抓取数据
	}
}
