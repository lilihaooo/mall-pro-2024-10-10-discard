package product

import "admin-gin/service"

type ApiGroup struct {
	GoodsApi
}

var (
	goodsService = service.GroupApp.ProductServiceGroup.GoodsService
)
