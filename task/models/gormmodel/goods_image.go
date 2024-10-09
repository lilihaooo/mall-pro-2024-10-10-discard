package gormmodel

type GoodsImage struct {
	GoodsID uint `gorm:"column:goods_id;type:bigint;not null;comment:'商品ID'" json:"goods_id"`
	ImageID uint `gorm:"column:image_id;type:bigint;not null;comment:'图片ID'" json:"image_id"`
}

func (*GoodsImage) TableName() string {
	return "goods_image"
}
