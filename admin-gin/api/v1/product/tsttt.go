package product

import (
	"gorm.io/gorm"
	"time"
)

type Goods struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ShopID   uint    `json:"shop_id"`
	Images   []Image `json:"images"`
	CouponID uint    `json:"coupon_id"`
	Tags     []Tag   `gorm:"many2many:goods_tag" json:"tags"`
}

type Tag struct {
	ID    uint    `gorm:"id" json:"id"`
	Title string  `gorm:"title" json:"title"`
	Goods []Goods `gorm:"many2many:goods_tag" json:"goods"`
}

type Coupon struct {
	gorm.Model
	Title       string    `json:"title"`        // 优惠券标题
	Amount      float64   `json:"amount"`       // 券金额
	MinAmount   float64   `json:"min_amount"`   // 的最小消费金额
	StartTime   time.Time `json:"start_time"`   // 开始时间
	EndTime     time.Time `json:"end_time"`     // 结束时间
	ShopID      uint      `json:"shop_id"`      // 店铺ID
	CouponTotal int64     `json:"coupon_total"` // 优惠券总数
	CouponCover int64     `json:"coupon_cover"` // 领券量
	Status      uint      `json:"status"`       // 状态
}

type Image struct {
	ID      uint   `json:"id"`
	GoodsID int    `json:"goods_id"`
	Url     string `json:"url"`
}
