package product

type Media struct {
	UID           string `gorm:"column:uid;type:varchar(64);not null;comment:'UUID'" json:"uid"`
	Certification string `gorm:"column:certification;type:varchar(100);comment:'认证类型'" json:"certification"`
	GoodAtCid     string `gorm:"column:good_at_cid;type:varchar(100);comment:'擅长分类'" json:"good_at_cid"`
	IsAlbumTalent string `gorm:"column:is_album_talent;type:varchar(10);comment:'是否为专辑达人'" json:"is_album_talent"`
	IsGroupLeader string `gorm:"column:is_group_leader;type:varchar(1);comment:'是否为团长'" json:"is_group_leader"`
	IsSelection   string `gorm:"column:is_selection;type:varchar(1);comment:'是否为精选'" json:"is_selection"`
	MediaHead     string `gorm:"column:media_head;type:varchar(255);comment:'媒体头像URL'" json:"media_head"`
	MediaName     string `gorm:"column:media_name;type:varchar(100);comment:'媒体名称'" json:"media_name"`
	Qq            string `gorm:"column:qq;type:varchar(20);comment:'媒体QQ号'" json:"qq"`
	UserLevel     string `gorm:"column:user_level;type:varchar(10);comment:'媒体用户级别'" json:"user_level"`
}
