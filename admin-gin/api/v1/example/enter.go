package example

import "admin-gin/service"

type ApiGroup struct {
	CustomerApi
	FileUploadAndDownloadApi
}

var (
	customerService              = service.GroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.GroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
