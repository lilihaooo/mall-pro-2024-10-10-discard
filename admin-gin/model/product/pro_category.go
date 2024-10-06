package product

type Category struct {
	ID       uint       `json:"id" gorm:"primaryKey;autoIncrement;comment:'ID'"`
	Level    string     `json:"level" gorm:"type:varchar(1);not null;comment:'类别的层级'"`
	ParentID uint       `json:"parent_id" gorm:"type:bigint;not null;comment:'父类别的ID'"`
	Name     string     `json:"name" gorm:"type:varchar(100);not null;comment:'类别名称'"`
	Status   string     `json:"status" gorm:"type:varchar(1);not null;comment:'类别的状态'"`
	Image    string     `json:"image" gorm:"type:varchar(255);comment:'类别图片URL'"`
	Num      int        `json:"num" gorm:"type:int;comment:'类别下的商品数量'"`
	Sort     int        `json:"sort" gorm:"type:int;comment:'类别的排序优先级'"`
	Children []Category `json:"children" gorm:"foreignkey:ParentID;"`
}

func (Category) TableName() string {
	return "category"
}
