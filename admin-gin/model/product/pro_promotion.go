package product

import "time"

type Promotion struct {
	ID              uint       `json:"id"`
	UserID          uint       `json:"user_id"`
	GoodsID         uint       `json:"goods_id"`
	CreatedAt       time.Time  `json:"created_at"`
	Title           string     `json:"title"`
	ImageID         uint       `json:"image_id"`
	Image           Image      `json:"image"`
	PostCouponPrice float64    `json:"post_coupon_price"`
	CommissionValue float64    `json:"commission_value"`
	CommissionRate  int32      `json:"commission_rate"`
	CouponValue     float64    `json:"coupon_value"`
	CouponEndTime   *time.Time `json:"coupon_end_time"`
}

func (*Promotion) TableName() string {
	return "promotion"
}
