package goods

import (
	"admin-gin/kafka"
)

type Message struct {
	ID uint `json:"id,omitempty"`
	// 封面
	CoverImage string `json:"cover_image,omitempty"`
	// 基本信息
	Description     string  `json:"description,omitempty"`
	Price           float64 `json:"price,omitempty"`
	CommissionRate  int32   `json:"commission_rate,omitempty"`
	CommissionValue float64 `json:"commission_value,omitempty"`
	Tags            []uint  `json:"tags,omitempty"`
	BrandID         uint    `json:"brand_id,omitempty"`
	BrandName       string  `json:"brand_name,omitempty"`
	// 优惠券信息
	CouponID        uint    `json:"coupon_id,omitempty"`
	CouponAmount    float64 `json:"coupon_amount,omitempty"`
	CouponBeginTime string  `json:"coupon_begin_time,omitempty"`
	CouponEndTime   string  `json:"coupon_end_time,omitempty"`
	CouponTotal     int     `json:"coupon_total,omitempty"`
	CouponCover     int     `json:"coupon_cover,omitempty"`
	PostCouponPrice float64 `json:"post_coupon_price,omitempty"`
	// 添加时间
	CreatedAt string `json:"created_at,omitempty"`
}

func NewUpdateEsGoodsInfoSender(value Message) *kafka.Sender {
	return &kafka.Sender{
		Topic: "goods-info",
		Key:   "update_es_goods_info",
		Value: value,
	}
}
