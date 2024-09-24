package product

import (
	"admin-gin/service/product/goods_grab"
	"admin-gin/service/product/simulate"
)

type ServiceGroup struct {
	CategoryService
	GoodsService
	goods_grab.GoodsGrabService
	simulate.SimService
	ShopService
	SearchService
}
