package request

import (
	"admin-gin/model/common/request"
	"admin-gin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
