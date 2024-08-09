package product

import (
	"admin-gin/global"
	"time"
)

type Coupon struct {
	global.GVA_MODEL
	Title       string    `json:"title"`        // 优惠券标题
	Amount      float64   `json:"amount"`       // 券金额
	MinAmount   float64   `json:"min_amount"`   // 的最小消费金额
	StartTime   time.Time `json:"start_time"`   // 开始时间
	EndTime     time.Time `json:"end_time"`     // 结束时间
	CouponTotal int       `json:"coupon_total"` // 优惠券总数
	CouponCover int       `json:"coupon_cover"` // 领券量
	Status      uint      `json:"status"`       // 状态
}
