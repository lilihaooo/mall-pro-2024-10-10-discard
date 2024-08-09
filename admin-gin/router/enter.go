package router

import (
	"admin-gin/router/example"
	"admin-gin/router/product"
	"admin-gin/router/system"
)

type Group struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Product product.RouterGroup
}

var GroupApp = new(Group)
