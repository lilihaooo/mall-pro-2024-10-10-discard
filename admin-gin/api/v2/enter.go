package v2

import (
	"admin-gin/api/v2/product"
)

type ApiGroup struct {
	ProductApiGroup product.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
