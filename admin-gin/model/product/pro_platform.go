package product

type Platform struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `gorm:"name" json:"name"`
}

func (*Platform) TableName() string {
	return "platform"
}
