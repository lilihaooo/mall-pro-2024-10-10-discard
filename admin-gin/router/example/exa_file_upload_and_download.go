package example

import (
	v1 "admin-gin/api/v1"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {

	var apiGroup = "fileUploadAndDownload"

	fileUploadAndDownloadRouter := Router.Group(apiGroup)
	exaFileUploadAndDownloadApi := v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	{
		common.RegisterRouteWithComment(fileUploadAndDownloadRouter, "POST", "upload", "上传文件", apiGroup, exaFileUploadAndDownloadApi.UploadFile)
		common.RegisterRouteWithComment(fileUploadAndDownloadRouter, "POST", "getFileList", "获取上传文件列表", apiGroup, exaFileUploadAndDownloadApi.GetFileList)
		common.RegisterRouteWithComment(fileUploadAndDownloadRouter, "POST", "deleteFile", "删除指定文件", apiGroup, exaFileUploadAndDownloadApi.DeleteFile)
		common.RegisterRouteWithComment(fileUploadAndDownloadRouter, "POST", "editFileName", "编辑文件名或者备注", apiGroup, exaFileUploadAndDownloadApi.EditFileName)
		common.RegisterRouteWithComment(fileUploadAndDownloadRouter, "POST", "breakpointContinue", "断点续传", apiGroup, exaFileUploadAndDownloadApi.BreakpointContinue)
		common.RegisterRouteWithComment(fileUploadAndDownloadRouter, "GET", "findFile", "查询当前文件成功的切片", apiGroup, exaFileUploadAndDownloadApi.FindFile)
		common.RegisterRouteWithComment(fileUploadAndDownloadRouter, "POST", "breakpointContinueFinish", "切片传输完成", apiGroup, exaFileUploadAndDownloadApi.BreakpointContinueFinish)
		common.RegisterRouteWithComment(fileUploadAndDownloadRouter, "POST", "removeChunk", "移除切片", apiGroup, exaFileUploadAndDownloadApi.RemoveChunk)

	}
}
