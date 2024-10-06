package gormmodel

type Tag struct {
	ID    uint    `gorm:"id" json:"id"`
	Title string  `gorm:"title" json:"title"`
	Goods []Goods `gorm:"many2many:goods_tag" json:"goods"`
}

func (Tag) TableName() string {
	return "tags"
}
