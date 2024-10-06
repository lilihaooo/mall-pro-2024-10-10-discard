package system

import (
	v1 "admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	apiGroup := "sysOperationRecord"
	operationRecordRouter := Router.Group(apiGroup)
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		common.RegisterRouteWithComment(operationRecordRouter, "POST", "createSysOperationRecord", "新建SysOperationRecord", apiGroup, authorityMenuApi.CreateSysOperationRecord)
		common.RegisterRouteWithComment(operationRecordRouter, "DELETE", "deleteSysOperationRecord", "删除SysOperationRecord", apiGroup, authorityMenuApi.DeleteSysOperationRecord)
		common.RegisterRouteWithComment(operationRecordRouter, "DELETE", "deleteSysOperationRecordByIds", "批量删除SysOperationRecord", apiGroup, authorityMenuApi.DeleteSysOperationRecordByIds)
		common.RegisterRouteWithComment(operationRecordRouter, "GET", "findSysOperationRecord", "根据ID获取SysOperationRecord", apiGroup, authorityMenuApi.FindSysOperationRecord)
		common.RegisterRouteWithComment(operationRecordRouter, "GET", "getSysOperationRecordList", "获取SysOperationRecord列表", apiGroup, authorityMenuApi.GetSysOperationRecordList)
	}
}
