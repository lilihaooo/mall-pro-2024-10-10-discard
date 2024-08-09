package service

import (
	"admin-gin/service/example"
	"admin-gin/service/product"
	"admin-gin/service/system"
)

type Group struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	ProductServiceGroup product.ServiceGroup
}

var GroupApp = new(Group)
