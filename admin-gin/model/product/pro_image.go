package product

type Image struct {
	ID  uint   `gorm:"column:id;primaryKey;autoIncrement;type:bigint;comment:'图片ID'" json:"id"`
	Url string `gorm:"column:url;type:varchar(255);not null;comment:'图片地址'" json:"url"`
}
