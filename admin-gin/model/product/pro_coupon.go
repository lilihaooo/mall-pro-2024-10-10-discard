package product

import (
	"admin-gin/model/common"
	"time"
)

type Coupon struct {
	common.GVA_MODEL
	Title       string    `json:"title" gorm:"type:varchar(255);not null;comment:'优惠券标题'"`
	Amount      float64   `json:"amount" gorm:"type:decimal(10,2);not null;comment:'券金额'"`
	MinAmount   float64   `json:"min_amount" gorm:"type:decimal(10,2);not null;comment:'最小消费金额'"`
	BeginTime   time.Time `json:"begin_time" gorm:"type:datetime;not null;comment:'开始时间'"`
	EndTime     time.Time `json:"end_time" gorm:"type:datetime;not null;comment:'结束时间'"`
	CouponTotal int       `json:"coupon_total" gorm:"type:int;not null;comment:'优惠券总数'"`
	CouponCover int       `json:"coupon_cover" gorm:"type:int;not null;comment:'领券数量'"`
	Status      uint      `json:"status" gorm:"type:tinyint(1);not null;comment:'状态（0: 停用, 1: 启用）'"`
}
