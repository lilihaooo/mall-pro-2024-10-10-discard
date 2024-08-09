package simulate

import (
	"admin-gin/global"
	"admin-gin/model/product"
	"admin-gin/utils"
	"sync"
)

var (
	coupons []product.Coupon
	tags    []product.Tag
	media   []product.Media

	mu          sync.Mutex
	initialized bool
)

func initSimulateBaseData() {
	mu.Lock()
	defer mu.Unlock()
	if initialized {
		return
	}

	err := global.XTK_DB.Find(&coupons).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	err = global.XTK_DB.Find(&tags).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	err = global.XTK_DB.Find(&media).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	initialized = true
}

func CleanSimulateBaseData() {
	mu.Lock()
	defer mu.Unlock()
	coupons = nil
	tags = nil
	media = nil
	initialized = false
	global.GVA_LOG.Info("清理模拟包基础数据")
}

func GoodsBindCoupon(g *product.Goods) *product.Goods {
	if !initialized {
		initSimulateBaseData()
	}
	l := len(coupons)
	if l == 0 {
		return g
	}
	index := utils.RandomInt(0, l)
	if coupons[index].Amount < g.Price {
		g.CouponID = coupons[index].ID
		postCouponPrice := utils.RoundToOneDecimal(g.Price - coupons[index].Amount)
		g.PostCouponPrice = postCouponPrice
		g.CommissionValue = utils.RoundToOneDecimal(float64(g.CommissionRate) * postCouponPrice / 100)
	}
	return g
}

func MakeRandomTags() (ta []product.Tag) {
	if !initialized {
		initSimulateBaseData()
	}
	l := len(tags)
	if l == 0 {
		return
	}
	tagNum := utils.RandomInt(0, 5)
	tagMap := make(map[int]product.Tag)
	for i := 0; i < tagNum; i++ {
		index := utils.RandomInt(0, l)
		tagMap[index] = tags[index]
	}
	for _, v := range tagMap {
		ta = append(ta, v)
	}
	return ta
}

func MakeRandomMediaUID() *string {
	if !initialized {
		initSimulateBaseData()
	}
	l := len(media)
	if l == 0 {
		return nil
	}
	lmIndex := utils.RandomInt(0, l)
	return &media[lmIndex].UID
}
