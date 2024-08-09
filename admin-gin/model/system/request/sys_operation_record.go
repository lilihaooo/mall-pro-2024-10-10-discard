package request

import (
	"admin-gin/model/common/request"
	"admin-gin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
