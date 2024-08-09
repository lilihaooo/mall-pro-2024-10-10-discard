package product

import "time"

type GrabLog struct {
	ID             uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	GrabRecordUuid string    `json:"grab_record_uuid" gorm:"grab_record_uuid"`
	Level          int       `json:"level" gorm:"type:varchar(255)"`
	Msg            string    `json:"msg" gorm:"type:text"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}
