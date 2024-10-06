package product

type Shop struct {
	ID             uint   `gorm:"column:id;primaryKey;autoIncrement;type:bigint;comment:'店铺ID'" json:"id"`
	Turnover       int64  `gorm:"column:turnover;type:bigint;not null;comment:'营业额'" json:"turnover"`
	Name           string `gorm:"column:name;type:varchar(100);not null;comment:'店铺名称'" json:"name"`
	Logo           string `gorm:"column:logo;type:varchar(255);comment:'店铺Logo'" json:"logo"`
	LogisticsScore int    `gorm:"column:logistics_score;type:int;not null;comment:'物流评分'" json:"logistics_score"`
	ProductScore   int    `gorm:"column:product_score;type:int;not null;comment:'商品评分'" json:"product_score"`
	ServiceScore   int    `gorm:"column:service_score;type:int;not null;comment:'服务评分'" json:"service_score"`
}

func (*Shop) TableName() string {
	return "shop"
}
