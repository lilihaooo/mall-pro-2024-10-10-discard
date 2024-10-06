package product

import "time"

type Collect struct {
	ID              uint       `json:"id" gorm:"primaryKey;autoIncrement;comment:'ID'"`
	UserID          uint       `json:"user_id" gorm:"type:bigint;not null;comment:'用户ID'"`
	GoodsID         uint       `json:"goods_id" gorm:"type:bigint;not null;comment:'商品ID'"`
	CreatedAt       time.Time  `json:"created_at" gorm:"type:datetime;comment:'创建时间'"`
	Title           string     `json:"title" gorm:"type:varchar(255);not null;comment:'商品标题'"`
	ImageID         uint       `json:"image_id" gorm:"type:bigint;comment:'图片ID'"`
	Image           Image      `json:"image" gorm:"foreignKey:ImageID;"`
	PostCouponPrice float64    `json:"post_coupon_price" gorm:"type:decimal(10,2);comment:'优惠券后的价格'"`
	CommissionValue float64    `json:"commission_value" gorm:"type:decimal(10,2);comment:'佣金金额'"`
	CommissionRate  int32      `json:"commission_rate" gorm:"type:int;comment:'佣金比率（整数）'"`
	CouponValue     float64    `json:"coupon_value" gorm:"type:decimal(10,2);comment:'优惠券金额'"`
	CouponEndTime   *time.Time `json:"coupon_end_time" gorm:"type:datetime;comment:'优惠券结束时间'"`
}

func (*Collect) TableName() string {
	return "collect"
}
