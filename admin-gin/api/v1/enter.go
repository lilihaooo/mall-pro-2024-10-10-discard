package v1

import (
	"admin-gin/api/v1/example"
	"admin-gin/api/v1/product"
	"admin-gin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	ProductApiGroup product.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
