package product

type Shop struct {
	ID             uint    `gorm:"column:id;primaryKey" json:"id"`
	Turnover       int64   `gorm:"column:turnover" json:"turnover"` // 店铺销量
	Name           string  `gorm:"column:name" json:"name"`
	Logo           string  `gorm:"column:logo" json:"logo"`
	LogisticsScore float64 `gorm:"logistics_score" json:"logistics_score"` // 物流分
	ProductScore   float64 `gorm:"product_score" json:"product_score"`     // 商品分
	ServiceScore   float64 `gorm:"service_score" json:"service_score"`     // 服务分
}

func (*Shop) TableName() string {
	return "shop"
}
