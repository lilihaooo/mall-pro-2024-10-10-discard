package product

import (
	"time"
)

type Goods struct {
	ID              uint      `gorm:"column:id;primaryKey;autoIncrement:true;type:bigint;comment:'ID'" json:"id"`
	Description     string    `gorm:"column:description;type:varchar(255);comment:'商品描述'" json:"description"`
	ShopID          uint      `gorm:"column:shop_id;type:bigint;comment:'店铺ID'" json:"shop_id"`
	Shop            Shop      `gorm:"foreignKey:ShopID;references:ID;" json:"shop"`
	CouponID        uint      `gorm:"column:coupon_id;type:bigint;comment:'优惠券ID'" json:"coupon_id"`
	Coupon          Coupon    `gorm:"foreignKey:CouponID;references:ID;" json:"coupon"`
	Images          []Image   `gorm:"many2many:goods_image;" json:"images"`
	CoverImageID    *uint     `gorm:"column:cover_image_id;not null;type:bigint;comment:'封面图片ID'" json:"cover_image_id"`
	Image           Image     `gorm:"foreignKey:CoverImageID;references:ID;comment:'封面图片'" json:"image"`
	Status          int       `gorm:"column:status;type:int;comment:'商品状态'" json:"status"`
	CategoryID      uint      `gorm:"column:category_id;type:bigint;comment:'分类ID'" json:"category_id"`
	Category        Category  `gorm:"foreignKey:CategoryID;references:ID;" json:"category"`
	SalesAll        int32     `gorm:"column:sales_all;type:int;comment:'全部销量'" json:"sales_all"`
	SalesMonth      int32     `gorm:"column:sales_month;type:int;comment:'月销量'" json:"sales_month"`
	SalesDay        int32     `gorm:"column:sales_day;type:int;comment:'24小时销量'" json:"sales_day"`
	Sales2Hour      int32     `gorm:"column:sales_2_hour;type:int;comment:'2小时销量'" json:"sales_2_hour"`
	Title           string    `gorm:"column:title;type:varchar(255);not null;comment:'商品标题'" json:"title"`
	CommissionRate  int32     `gorm:"column:commission_rate;type:int;comment:'佣金比例'" json:"commission_rate"`
	CommissionValue float64   `gorm:"column:commission_value;type:decimal(10,2);comment:'佣金金额'" json:"commission_value"`
	Price           float64   `gorm:"column:price;type:decimal(10,2);not null;comment:'商品价格'" json:"price"`
	PostCouponPrice float64   `gorm:"column:post_coupon_price;type:decimal(10,2);comment:'优惠券后的最终价格'" json:"post_coupon_price"`
	Inventory       int32     `gorm:"column:inventory;type:int;comment:'商品库存'" json:"inventory"`
	ExperienceScore float64   `gorm:"column:experience_score;type:decimal(3,1);comment:'体验分'" json:"experience_score"`
	Tags            []Tag     `gorm:"many2many:goods_tag;" json:"tags"`
	BrandID         uint64    `gorm:"column:brand_id;type:bigint;comment:'品牌ID'" json:"brand_id"`
	Brand           Brand     `gorm:"foreignKey:BrandID;references:ID;comment:'品牌信息'" json:"brand"`
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime;comment:'创建时间'" json:"created_at"`
	PlatformID      uint      `gorm:"column:platform_id;type:bigint;comment:'平台ID'" json:"platform_id"`
	Platform        Platform  `gorm:"foreignKey:PlatformID;references:ID;not null;comment:'平台信息'" json:"platform"`
	DataFrom        int       `gorm:"column:data_from;type:int;not null;comment:'商品数据来源'" json:"data_from"`
	PushHot         int       `gorm:"column:push_hot;type:int;comment:'热销商品标识'" json:"push_hot"`
	MediaUID        *string   `gorm:"column:media_uid;type:varchar(64);comment:'媒体ID'" json:"media_uid"`
	Media           Media     `gorm:"foreignKey:MediaUID;references:UID;comment:'媒体信息'" json:"media"`
	IsToEs          int       `gorm:"column:is_to_es;type:tinyint(1);default:0;comment:'是否同步到ES (0: 否, 1: 是)'" json:"is_to_es"`
}

func (*Goods) TableName() string {
	return "goods"
}

type EsGoods struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	CategoryID      uint    `json:"category_id"`
	CategoryName    string  `json:"category_name"`
	CoverImage      string  `json:"cover_image"`
	ShopName        string  `json:"shop_name"`
	Price           float64 `json:"price"`
	Status          int     `json:"status"`
	CouponID        uint    `json:"coupon_id"`
	CouponBeginTime string  `json:"coupon_begin_time"`
	CouponEndTime   string  `json:"coupon_end_time"`
	CouponAmount    float64 `json:"coupon_amount"`
	CouponTotal     int64   `json:"coupon_total"`
	CouponCover     int64   `json:"coupon_cover"`
	SalesAll        int32   `json:"sales_all"`
	SalesMonth      int32   `json:"sales_month"`
	SalesDay        int32   `json:"sales_day"`
	Sales2Hour      int32   `json:"sales_2_hour"`
	PlatformName    string  `json:"platform_name"`
	PlatformID      uint    `json:"platform_id"`
	MediaUID        string  `json:"media_uid"`
	MediaName       string  `json:"media_name"`
	CommissionRate  int32   `json:"commission_rate"`
	CommissionValue float64 `json:"commission_value"`
	PostCouponPrice float64 `json:"post_coupon_price"`
	Inventory       int32   `json:"inventory"`
	ExperienceScore float64 `json:"experience_score"`
	CreatedAt       string  `json:"created_at"`
	PushHot         int     `json:"push_hot"`
	GoodsDataFrom   int     `json:"goods_data_from"`
	BrandName       string  `json:"brand_name"`
	Tags            []uint  `json:"tags"`
	IsToEs          int     `gorm:"column:is_to_es" json:"is_to_es"`
	DataFrom        int     `json:"data_from"`
}
