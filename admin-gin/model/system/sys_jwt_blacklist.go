package system

import (
	"admin-gin/model/common"
)

type JwtBlacklist struct {
	common.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
