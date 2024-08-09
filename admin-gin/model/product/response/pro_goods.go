package response

type GoodsSearch struct {
	ID              uint    `json:"id"`
	Description     string  `json:"description"`
	ShopID          uint    `json:"shop_id"`
	ShopName        string  `json:"shop_name"`
	CoverImage      string  `json:"cover_image"`
	Status          int     `json:"status"`
	CategoryName    string  `json:"category_name"`
	SalesAll        int32   `json:"sales_all"`
	SalesMonth      int32   `json:"sales_month"`
	SalesDay        int32   `json:"sales_day"`
	Sales2Hour      int32   `json:"sales_2_hour"`
	PlatformName    string  `json:"platform_name"`
	MediaUID        *string `json:"media_uid"`
	MediaName       string  `json:"media_name"`
	Title           string  `json:"title"`
	CommissionRate  int32   `json:"commission_rate"`
	CommissionValue float64 `json:"commission_value"`
	PostCouponPrice float64 `json:"post_coupon_price"`
	CouponID        uint    `json:"coupon_id"`
	CouponAmount    float64 `json:"coupon_amount"`
	CouponEndTime   string  `json:"coupon_end_time"`
	CouponCover     int     `json:"coupon_cover"`
	CouponTotal     int     `json:"coupon_total"`
	Price           float64 `json:"price"`
	TrueSales       int32   `json:"true_sales"`
	Inventory       int32   `json:"inventory"`
	ExperienceScore float64 `json:"experience_score"`
	URL             string  `json:"url"`
	PushHot         int     `json:"push_hot"`
}

type GoodsSearchV2 struct {
	ID              uint    `json:"id"`
	Description     string  `json:"description"`
	ShopID          uint    `json:"shop_id"`
	ShopName        string  `json:"shop_name"`
	CoverImage      string  `json:"cover_image"`
	CategoryName    string  `json:"category_name"`
	SalesAll        int32   `json:"sales_all"`
	SalesDay        int32   `json:"sales_day"`
	Sales2Hour      int32   `json:"sales_2_hour"`
	PlatformName    string  `json:"platform_name"`
	MediaUID        *string `json:"media_uid"`
	MediaName       string  `json:"media_name"`
	Title           string  `json:"title"`
	CommissionRate  int32   `json:"commission_rate"`
	CommissionValue float64 `json:"commission_value"`
	PostCouponPrice float64 `json:"post_coupon_price"`
	CouponID        uint    `json:"coupon_id"`
	CouponAmount    float64 `json:"coupon_amount"`
	CouponEndTime   string  `json:"coupon_end_time"`
	CouponCover     int     `json:"coupon_cover"`
	CouponTotal     int     `json:"coupon_total"`
	Price           float64 `json:"price"`
	TrueSales       int32   `json:"true_sales"`
	Inventory       int32   `json:"inventory"`
	ExperienceScore float64 `json:"experience_score"`
	PushHot         int     `json:"push_hot"`
	Tags            []uint  `json:"tags"`
}
