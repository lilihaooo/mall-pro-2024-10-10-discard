package gormmodel

import (
	"task/models/common"
	"time"
)

type Coupon struct {
	common.GVA_MODEL
	Title       string    `json:"title"`        // 优惠券标题
	Amount      float64   `json:"amount"`       // 券金额
	MinAmount   float64   `json:"min_amount"`   // 的最小消费金额
	BeginTime   time.Time `json:"begin_time"`   // 开始时间
	EndTime     time.Time `json:"end_time"`     // 结束时间
	ShopID      uint      `json:"shop_id"`      // 店铺ID
	CouponTotal int64     `json:"coupon_total"` // 优惠券总数
	CouponCover int64     `json:"coupon_cover"` // 领券量
	Status      uint      `json:"status"`       // 状态
}

func (Coupon) TableName() string {
	return "coupons"
}
