package example

import (
	v1 "admin-gin/api/v1"
	"admin-gin/middleware"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	var apiGroup = "customer"
	customerRouter := Router.Group(apiGroup).Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group(apiGroup)
	exaCustomerApi := v1.ApiGroupApp.ExampleApiGroup.CustomerApi
	{
		common.RegisterRouteWithComment(customerRouter, "POST", "customer", "创建客户", apiGroup, exaCustomerApi.CreateExaCustomer)
		common.RegisterRouteWithComment(customerRouter, "PUT", "customer", "更新客户", apiGroup, exaCustomerApi.UpdateExaCustomer)
		common.RegisterRouteWithComment(customerRouter, "DELETE", "customer", "删除客户", apiGroup, exaCustomerApi.DeleteExaCustomer)

	}
	{

		common.RegisterRouteWithComment(customerRouterWithoutRecord, "GET", "customer", "获取单一客户信息", apiGroup, exaCustomerApi.GetExaCustomer)
		common.RegisterRouteWithComment(customerRouterWithoutRecord, "GET", "customerList", "获取客户列表", apiGroup, exaCustomerApi.GetExaCustomerList)

	}
}
