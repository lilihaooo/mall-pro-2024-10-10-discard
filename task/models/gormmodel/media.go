package gormmodel

type Media struct {
	UID           string `gorm:"uid" json:"uid"`                         // Uid
	Certification string `gorm:"certification" json:"certification"`     // 认证类型
	GoodAtCid     string `gorm:"good_at_cid" json:"good_at_cid"`         // 擅长
	IsAlbumTalent string `gorm:"is_album_talent" json:"is_album_talent"` // is 专辑达人
	IsGroupLeader string `gorm:"is_group_leader" json:"is_group_leader"` // 是否是团长xx
	IsSelection   string `gorm:"is_selection" json:"is_selection"`       // ...
	MediaHead     string `gorm:"media_head" json:"media_head"`           // 头像
	MediaName     string `gorm:"media_name" json:"media_name"`           // 名称
	Qq            string `gorm:"qq" json:"qq"`                           // QQ
	UserLevel     string `gorm:"user_level" json:"user_level"`           // 级别
}
