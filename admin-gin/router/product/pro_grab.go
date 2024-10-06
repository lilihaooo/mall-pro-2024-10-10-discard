package product

import (
	v1 "admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type GrabRouter struct{}

func (p *GrabRouter) InitGrabRouter(Router *gin.RouterGroup) {
	var apiGroup = "grab"
	grabInfoRouter := Router.Group(apiGroup)
	grabApi := v1.ApiGroupApp.ProductApiGroup.GrobApi
	{
		common.RegisterRouteWithComment(grabInfoRouter, "GET", "record", "获取抓取记录", apiGroup, grabApi.GetGrabRecord)
		common.RegisterRouteWithComment(grabInfoRouter, "GET", "log", "获取抓取日志", apiGroup, grabApi.GetGrabLog)

	}
}
