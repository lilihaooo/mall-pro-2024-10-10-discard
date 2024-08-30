package product

type GoodsImage struct {
	GoodsID uint `json:"goods_id"`
	ImageID uint `json:"image_id"`
}

func (*GoodsImage) TableName() string {
	return "goods_image"
}
