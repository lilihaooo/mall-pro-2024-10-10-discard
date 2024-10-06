package system

type SysSetting struct {
	ID                       uint   `gorm:"column:id;primaryKey;autoIncrement;type:bigint;comment:'系统设置ID'" json:"id"`
	UserID                   uint   `gorm:"column:user_id;type:bigint;not null;comment:'用户ID'" json:"user_id"`
	Weakness                 bool   `gorm:"column:weakness;type:boolean;default:false;comment:'是否启用色弱模式'" json:"weakness"`
	Grey                     bool   `gorm:"column:grey;type:boolean;default:false;comment:'是否启用灰色模式'" json:"grey"`
	PrimaryColor             string `gorm:"column:primary_color;type:varchar(100);not null;comment:'主色调'" validate:"required" json:"primaryColor"`
	ShowTabs                 bool   `gorm:"column:show_tabs;type:boolean;default:true;comment:'是否显示标签页'" json:"showTabs"`
	DarkMode                 string `gorm:"column:dark_mode;type:varchar(100);not null;comment:'暗黑模式'" validate:"required" json:"darkMode"`
	LayoutSideWidth          int    `gorm:"column:layout_side_width;type:int;not null;comment:'侧边栏宽度'" validate:"required" json:"layout_side_width"`
	LayoutSideCollapsedWidth int    `gorm:"column:layout_side_collapsed_width;type:int;not null;comment:'侧边栏折叠宽度'" validate:"required" json:"layout_side_collapsed_width"`
	LayoutSideItemHeight     int    `gorm:"column:layout_side_item_height;type:int;not null;comment:'侧边栏项高度'" validate:"required" json:"layout_side_item_height"`
	ShowWatermark            bool   `gorm:"column:show_watermark;type:boolean;default:false;comment:'是否显示水印'" json:"show_watermark"`
	SideMode                 string `gorm:"column:side_mode;type:varchar(100);not null;comment:'侧边栏模式'" validate:"required" json:"side_mode"`
}

func (SysSetting) TableName() string {
	return "sys_user_setting"
}
