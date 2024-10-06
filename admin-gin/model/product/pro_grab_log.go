package product

import "time"

type GrabLog struct {
	ID             uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint;comment:'唯一ID'" json:"id"`
	GrabRecordUuid string    `gorm:"column:grab_record_uuid;type:varchar(64);not null;comment:'抓取记录UUID'" json:"grab_record_uuid"`
	Level          int       `gorm:"column:level;type:int;comment:'基本'" json:"level"`
	Msg            string    `gorm:"column:msg;type:varchar(255);comment:'抓取返回信息'" json:"msg"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime;type:datetime;comment:'记录创建时间'" json:"created_at"`
}
