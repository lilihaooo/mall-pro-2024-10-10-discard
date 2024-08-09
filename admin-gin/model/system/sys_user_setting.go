package system

type SysSetting struct {
	ID                       uint   `json:"id"  gorm:"column:id"`
	UserID                   uint   `json:"user_id"  gorm:"column:user_id"`
	Weakness                 bool   `json:"weakness" gorm:"column:weakness"`
	Grey                     bool   `json:"grey" gorm:"column:grey"`
	PrimaryColor             string `validate:"required" json:"primaryColor" gorm:"column:primary_color;type:varchar(255)"`
	ShowTabs                 bool   `json:"showTabs" gorm:"column:show_tabs"`
	DarkMode                 string `validate:"required" json:"darkMode" gorm:"column:dark_mode"`
	LayoutSideWidth          int    `validate:"required" json:"layout_side_width" gorm:"column:layout_side_width"`
	LayoutSideCollapsedWidth int    `validate:"required" json:"layout_side_collapsed_width" gorm:"column:layout_side_collapsed_width"`
	LayoutSideItemHeight     int    `validate:"required" json:"layout_side_item_height" gorm:"column:layout_side_item_height"`
	ShowWatermark            bool   `json:"show_watermark" gorm:"column:show_watermark"`
	SideMode                 string `validate:"required" json:"side_mode" gorm:"column:side_mode"`
}

func (SysSetting) TableName() string {
	return "sys_user_setting"
}
