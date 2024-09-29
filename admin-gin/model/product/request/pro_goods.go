package request

import "time"

type ProGoodsSearch struct {
	ID         uint `form:"g_id" json:"g_id"`
	CategoryID uint `form:"category_id" json:"category_id"`
	// 筛选
	PostCouponPriceFrom float64 `form:"qhj_f" json:"qhj_f"`
	PostCouponPriceTo   float64 `form:"qhj_t" json:"qhj_t"`
	CommissionRateFrom  int     `form:"cr_f" json:"cr_f"`
	CommissionRateTo    int     `form:"cr_t" json:"cr_t"`
	CommissionValue     float64 `form:"cv" json:"cv"`
	SaleNun             int     `form:"sale_num" json:"sale_num"`
	CouponValueForm     float64 `form:"couv_f" json:"couv_f"`
	CouponValueTo       float64 `form:"couv_t" json:"couv_t"`
	CouponExpirationDay int     `form:"qgt" json:"qgt"` // 券到期天数
	ExperienceScore     float64 `form:"tyf" json:"tyf"` // 体验分
	// 排序
	OrderBy   string `form:"order_by" json:"order_by"`
	OrderType string `form:"order_type" json:"order_type"`
	// 多选框
	IsChoice         int    `form:"is_choice" json:"is_choice"`                   //精选推荐
	TodayUp          int    `form:"jrsx" json:"jrsx"`                             //精选推荐
	DataFrom         int    `form:"data_from" json:"data_from"`                   // 数据来源
	AbsoluteLowPrice int    `form:"absolute_low_price" json:"absolute_low_price"` // 低价好卖
	BuyerReviews     int    `form:"buyer_reviews" json:"buyer_reviews"`           // 买家好评
	IsBrand          int    `form:"is_brand" json:"is_brand"`                     //品牌库
	Platform         uint   `form:"platform" json:"platform"`
	YunFeiXian       int    `form:"yfx" json:"yfx"`
	HuaBeiMX         int    `form:"hbmx" json:"hbmx"`
	PoSunBZ          int    `form:"psbz" json:"psbz"`
	Areas            string `form:"area" json:"area"` // ---------数组 以 逗号 分割
	ShopType         uint   `form:"shop" json:"shop"`
	ActiveGroup      string `form:"av_group" json:"av_group"`     // ----------数组 以 逗号 分割
	BiaoXianBD       int    `form:"bx_bdan" json:"bx_bdan"`       // 商品表现 榜单
	BiaoXianType     int    `form:"bx_type" json:"bx_type"`       // 商品表现 类型
	PushSuCai        string `form:"push_sucai" json:"push_sucai"` // -----------数组 以 逗号 分割

	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

type ProGoodsSearchV2 struct {
	ID         uint   `form:"g_id" json:"g_id"`
	Keyword    string `form:"keyword" json:"keyword"`
	CategoryID uint   `form:"category_id" json:"category_id"`
	// 筛选
	PostCouponPriceFrom float64 `form:"qhj_f" json:"qhj_f"`
	PostCouponPriceTo   float64 `form:"qhj_t" json:"qhj_t"`
	CommissionRateFrom  int     `form:"cr_f" json:"cr_f"`
	CommissionRateTo    int     `form:"cr_t" json:"cr_t"`
	CommissionValue     float64 `form:"cv" json:"cv"`
	SaleNun             int     `form:"sale_num" json:"sale_num"`
	CouponValueForm     float64 `form:"couv_f" json:"couv_f"`
	CouponValueTo       float64 `form:"couv_t" json:"couv_t"`
	CouponExpirationDay int     `form:"qgt" json:"qgt"` // 券到期天数
	ExperienceScore     float64 `form:"tyf" json:"tyf"` // 体验分
	// 排序
	OrderBy   string `form:"order_by" json:"order_by"`
	OrderType string `form:"order_type" json:"order_type"`
	// 多选框
	IsChoice         int    `form:"is_choice" json:"is_choice"`                   //精选推荐
	TodayUp          int    `form:"jrsx" json:"jrsx"`                             //精选推荐
	DataFrom         int    `form:"data_from" json:"data_from"`                   // 数据来源
	AbsoluteLowPrice int    `form:"absolute_low_price" json:"absolute_low_price"` // 低价好卖
	BuyerReviews     int    `form:"buyer_reviews" json:"buyer_reviews"`           // 买家好评
	IsBrand          int    `form:"is_brand" json:"is_brand"`                     //品牌库
	Platform         uint   `form:"platform" json:"platform"`
	YunFeiXian       int    `form:"yfx" json:"yfx"`
	HuaBeiMX         int    `form:"hbmx" json:"hbmx"`
	PoSunBZ          int    `form:"psbz" json:"psbz"`
	Areas            string `form:"area" json:"area"` // ---------数组 以 逗号 分割
	ShopType         uint   `form:"shop" json:"shop"`
	ActiveGroup      string `form:"av_group" json:"av_group"`     // ----------数组 以 逗号 分割
	BiaoXianBD       int    `form:"bx_bdan" json:"bx_bdan"`       // 商品表现 榜单
	BiaoXianType     int    `form:"bx_type" json:"bx_type"`       // 商品表现 类型
	PushSuCai        string `form:"push_sucai" json:"push_sucai"` // -----------数组 以 逗号 分割

	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

type GoodsInfo struct {
	ID uint `json:"id" form:"id"`
}

type GoodsImageUpload struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

type GoodsImageDelete struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

type SetGoodsImageCover struct {
	GoodsID uint `json:"goods_id,string" form:"goods_id" validate:"required"`
	ImageID uint `json:"image_id" form:"image_id" validate:"required"`
}

type GrabRecord struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

type Suggestion struct {
	Keyword string `json:"keyword" form:"keyword" validate:"required"`
}

type UpdateGoodsBaseInfo struct {
	Description    string  `json:"description"`
	Price          float64 `json:"price"  validate:"required"`
	CommissionRate int32   `json:"commission_rate"  validate:"required"`
	Tags           []uint  `json:"tags"`
	GoodsID        uint    `json:"goods_id"  validate:"required"`
}
type UpdateGoodsCouponInfo struct {
	GoodsID     uint      `json:"goods_id"  validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Amount      float64   `json:"amount" validate:"required"`
	MinAmount   float64   `json:"min_amount" validate:"required"`
	BeginTime   time.Time `json:"begin_time" validate:"required"`
	EndTime     time.Time `json:"end_time" validate:"required"`
	CouponTotal int       `json:"coupon_total" validate:"required"`
	CouponCover int       `json:"coupon_cover" validate:"required"`
}

// 用户收藏请求
type GoodsCollect struct {
	UserID  uint `json:"user_id" validate:"required"`
	GoodsID uint `json:"goods_id"  validate:"required"`
}

// 批量取消收藏
type BatchCollect struct {
	IDs []uint `json:"ids" validate:"required"`
}

// 用户推广请求
type GoodsPromotion struct {
	GoodsID uint `json:"goods_id"  validate:"required"`
}

// 我的收藏
type MyCollect struct {
	Page int `form:"page" json:"page"`
}

// 我的推广
type MyPromotion struct {
	Page int `form:"page" json:"page"`
}
