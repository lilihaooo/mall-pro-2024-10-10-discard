package request

import "time"

// 抓取商品请求
type GrabRequest struct {
	GoNum       int    `json:"go_num"`
	PageNum     int    `json:"page_num"`
	Option      int    `json:"option"`
	CategoryIDs []uint `json:"category_ids"`
}

// 模拟优惠券请求
type SimulateCouponRequest struct {
	Count        int       `json:"count" validate:"required,gte=10" label:"生成数量"`
	BeginDate    time.Time `json:"begin_date" validate:"required" label:"开始日期"` // UTC 时间
	EffectiveDay int       `json:"effective_day" validate:"required,gte=1" label:"有效期天数"`
	AmountFrom   float64   `json:"amount_from" validate:"required" label:"券额"`
	AmountTo     float64   `json:"amount_to" validate:"required" label:"券额"`
}
