package product

type Platform struct {
	ID   uint   `gorm:"column:id;primaryKey;autoIncrement;type:bigint;comment:'平台ID'" json:"id"`
	Name string `gorm:"column:name;type:varchar(10);not null;comment:'平台名称'" json:"name"`
}

func (*Platform) TableName() string {
	return "platform"
}
