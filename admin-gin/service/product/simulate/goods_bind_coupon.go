package simulate

import (
	"admin-gin/global"
	"admin-gin/model/product"
	"strconv"
)

func (*SimService) GoodsBindCoupon() {
	// 获取没有优惠券的商品 100 个
	var goods []product.Goods
	err := global.XTK_DB.Where("coupon_id = 0").Limit(1000).Find(&goods).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	count := 0

	for _, g := range goods {
		g = *GoodsBindCoupon(&g)
		err = global.XTK_DB.Updates(&g).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
		}
		count++
	}

	// 清理模拟的基础数据
	CleanSimulateBaseData()
	global.GVA_LOG.Info("为" + strconv.Itoa(count) + "条商品成功添加了优惠券")
}
