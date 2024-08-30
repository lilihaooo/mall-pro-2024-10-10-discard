package product

import "admin-gin/service"

type ApiGroup struct {
	CategoryApi
	GoodsApi
	GrobApi
	ShopApi
	TagApi
}

var (
	categoryService = service.GroupApp.ProductServiceGroup.CategoryService
	goodsService    = service.GroupApp.ProductServiceGroup.GoodsService
	shopService     = service.GroupApp.ProductServiceGroup.ShopService
)
