package product

import (
	"admin-gin/global"
)

type GrabRecord struct {
	global.GVA_MODEL
	Uuid        string `json:"uuid"`
	UserID      uint   `json:"user_id" gorm:"user_id"`
	CategoryNum int    `json:"category_num" gorm:"category_num"`
	GoNum       int    `json:"go_num" gorm:"go_num"`
	PageNum     int    `json:"page_num" gorm:"page_num"`
	RunTime     string `json:"run_time" gorm:"run_time"`
	GrabNum     int64  `json:"grab_num" gorm:"grab_num"`
}
