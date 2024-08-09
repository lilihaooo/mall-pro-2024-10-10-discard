package product

import (
	"admin-gin/global"
	"admin-gin/model/product"
	"admin-gin/model/product/request"
)

type ShopService struct{}

func (*ShopService) ShopList(req request.ProShopList) (list []product.Shop, count int64, err error) {
	query := global.XTK_DB.Model(product.Shop{})
	err = query.Count(&count).Error
	if err != nil {
		return
	}
	offset := (req.Page - 1) * req.PageSize
	err = query.Order("id desc").Limit(req.PageSize).Offset(offset).Find(&list).Error
	return
}
