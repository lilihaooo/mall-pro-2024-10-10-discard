package product

import "time"

type Promotion struct {
	ID              uint       `gorm:"column:id;primaryKey;autoIncrement;type:bigint;comment:'ID'" json:"id"`
	UserID          uint       `gorm:"column:user_id;type:bigint;not null;comment:'用户ID'" json:"user_id"`
	GoodsID         uint       `gorm:"column:goods_id;type:bigint;not null;comment:'商品ID'" json:"goods_id"`
	CreatedAt       time.Time  `gorm:"column:created_at;type:datetime;autoCreateTime;comment:'创建时间'" json:"created_at"`
	Title           string     `gorm:"column:title;type:varchar(255);not null;comment:'促销活动标题'" json:"title"`
	ImageID         uint       `gorm:"column:image_id;type:bigint;comment:'图片ID'" json:"image_id"`
	Image           Image      `gorm:"foreignKey:ImageID;references:ID;" json:"image"`
	PostCouponPrice float64    `gorm:"column:post_coupon_price;type:decimal(10,2);not null;comment:'券后价格'" json:"post_coupon_price"`
	CommissionValue float64    `gorm:"column:commission_value;type:decimal(10,2);not null;comment:'佣金金额'" json:"commission_value"`
	CommissionRate  int32      `gorm:"column:commission_rate;type:int;not null;comment:'佣金比例'" json:"commission_rate"`
	CouponValue     float64    `gorm:"column:coupon_value;type:decimal(10,2);comment:'优惠券金额'" json:"coupon_value"`
	CouponEndTime   *time.Time `gorm:"column:coupon_end_time;type:datetime;comment:'优惠券结束时间'" json:"coupon_end_time"`
}

func (*Promotion) TableName() string {
	return "promotion"
}
