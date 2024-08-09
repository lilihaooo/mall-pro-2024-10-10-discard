package product

type Category struct {
	ID       uint       `json:"id"`
	Level    string     `json:"level"`
	ParentID uint       `json:"parent_id"`
	Name     string     `json:"name"`
	Status   string     `json:"status"`
	Image    string     `json:"image"`
	Num      int        `json:"num"`
	Sort     int        `json:"sort"`
	Children []Category `json:"children" gorm:"foreignkey:ParentID"`
}

func (Category) TableName() string {
	return "category"
}
