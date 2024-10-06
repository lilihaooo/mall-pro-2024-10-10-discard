package gormmodel

import (
	"time"
)

type Goods struct {
	ID              uint      `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"` // id
	Description     string    `gorm:"column:description;comment:描述" json:"description"`             // 描述
	ShopID          uint      `gorm:"column:shop_id;comment:店铺id" json:"shop_id"`                   // 店铺id
	Shop            Shop      `json:"shop"`
	CouponID        uint      `gorm:"column:coupon_id" json:"coupon_id"`
	Coupon          Coupon    `json:"coupon"`
	Images          []Image   `gorm:"many2many:goods_image" json:"images" `
	CoverImageID    uint      `gorm:"column:cover_image_id" json:"cover_image_id"`
	Image           Image     `gorm:"foreignKey:CoverImageID" json:"image"`
	Status          int       `gorm:"column:status;comment:状态" json:"status"`             // 状态
	CategoryID      uint      `gorm:"column:category_id;comment:分类id" json:"category_id"` // 分类id
	Category        Category  `json:"category"`
	SalesAll        int32     `gorm:"column:sales_all;comment:全部销量" json:"sales_all"`                 // 全部销量
	SalesMonth      int32     `gorm:"column:sales_month" json:"sales_month"`                          // 月销量
	SalesDay        int32     `gorm:"column:sales_day" json:"sales_day"`                              // 24小时销量
	Sales2Hour      int32     `gorm:"column:sales_2_hour" json:"sales_2_hour"`                        // 2小时销量
	Title           string    `gorm:"column:title;comment:标题" json:"title"`                           // 标题
	CommissionRate  int32     `gorm:"column:commission_rate;comment:佣金比例" json:"commission_rate"`     // 佣金比例
	CommissionValue float64   `gorm:"column:commission_value;comment:佣金金额" json:"commission_value"`   // 佣金比例
	Price           float64   `gorm:"column:price;comment:价格" json:"price"`                           // 价格
	PostCouponPrice float64   `gorm:"column:post_coupon_price;comment:最终价格" json:"post_coupon_price"` // 券后价
	Inventory       int32     `gorm:"column:inventory;comment:库存" json:"inventory"`                   // 库存
	ExperienceScore float64   `gorm:"column:experience_score;comment:体验分" json:"experience_score"`    // 体验分
	Tags            []Tag     `gorm:"many2many:goods_tag" json:"tags"`
	BrandID         uint64    `gorm:"column:brand_id" json:"brand_id"`
	Brand           Brand     `json:"brand"`
	CreatedAt       time.Time `json:"created_at"`
	PlatformID      uint      `gorm:"column:platform_id" json:"platform_id"`
	Platform        Platform  `json:"platform"`
	DataFrom        int       `gorm:"column:data_from" json:"data_from"`
	PushHot         int       `gorm:"column:push_hot" json:"push_hot"`
	MediaUID        *string   `gorm:"column:media_uid;comment:媒体ID" json:"media_uid"` // 团长ID
	Media           Media     `gorm:"foreignKey:MediaUID;references:UID" json:"media"`
	IsToEs          int       `gorm:"column:is_to_es" json:"is_to_es"`
}

// TableName Product's table name
func (*Goods) TableName() string {
	return "goods"
}
