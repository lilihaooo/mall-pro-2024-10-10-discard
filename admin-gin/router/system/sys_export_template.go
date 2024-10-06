package system

import (
	"admin-gin/api/v1"
	"admin-gin/middleware"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type SysExportTemplateRouter struct {
}

// InitSysExportTemplateRouter 初始化 导出模板 路由信息
func (s *SysExportTemplateRouter) InitSysExportTemplateRouter(Router *gin.RouterGroup) {
	apiGroup := "sysExportTemplate"
	sysExportTemplateRouter := Router.Group(apiGroup).Use(middleware.OperationRecord())
	sysExportTemplateRouterWithoutRecord := Router.Group(apiGroup)
	var sysExportTemplateApi = v1.ApiGroupApp.SystemApiGroup.SysExportTemplateApi
	{
		common.RegisterRouteWithComment(sysExportTemplateRouter, "POST", "createSysExportTemplate", "新建导出模板", apiGroup, sysExportTemplateApi.CreateSysExportTemplate)
		common.RegisterRouteWithComment(sysExportTemplateRouter, "DELETE", "deleteSysExportTemplate", "删除导出模板", apiGroup, sysExportTemplateApi.DeleteSysExportTemplate)
		common.RegisterRouteWithComment(sysExportTemplateRouter, "DELETE", "deleteSysExportTemplateByIds", "批量删除导出模板", apiGroup, sysExportTemplateApi.DeleteSysExportTemplateByIds)
		common.RegisterRouteWithComment(sysExportTemplateRouter, "PUT", "updateSysExportTemplate", "更新导出模板", apiGroup, sysExportTemplateApi.UpdateSysExportTemplate)
		common.RegisterRouteWithComment(sysExportTemplateRouter, "POST", "importExcel", "导入Excel", apiGroup, sysExportTemplateApi.ImportExcel)

	}
	{
		common.RegisterRouteWithComment(sysExportTemplateRouterWithoutRecord, "GET", "findSysExportTemplate", "根据ID获取导出模板", apiGroup, sysExportTemplateApi.FindSysExportTemplate)
		common.RegisterRouteWithComment(sysExportTemplateRouterWithoutRecord, "GET", "getSysExportTemplateList", "获取导出模板列表", apiGroup, sysExportTemplateApi.GetSysExportTemplateList)
		common.RegisterRouteWithComment(sysExportTemplateRouterWithoutRecord, "GET", "exportExcel", "导出表格", apiGroup, sysExportTemplateApi.ExportExcel)
		common.RegisterRouteWithComment(sysExportTemplateRouterWithoutRecord, "GET", "exportTemplate", "导出模板", apiGroup, sysExportTemplateApi.ExportTemplate)

	}
}
