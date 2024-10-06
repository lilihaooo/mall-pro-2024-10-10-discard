package product

import (
	"admin-gin/model/common"
)

type GrabRecord struct {
	common.GVA_MODEL
	Uuid        string `gorm:"column:uuid;type:varchar(64);not null;comment:'UUID'" json:"uuid"`
	UserID      uint   `gorm:"column:user_id;type:bigint;not null;comment:'用户ID'" json:"user_id"`
	CategoryNum int    `gorm:"column:category_num;type:int;not null;comment:'分类数量'" json:"category_num"`
	GoNum       int    `gorm:"column:go_num;type:int;not null;comment:'并发数量'" json:"go_num"`
	PageNum     int    `gorm:"column:page_num;type:int;not null;comment:'分页数量'" json:"page_num"`
	RunTime     string `gorm:"column:run_time;type:varchar(100);comment:'运行时间'" json:"run_time"`
	GrabNum     int64  `gorm:"column:grab_num;type:bigint;comment:'抓取记录数量'" json:"grab_num"`
}
