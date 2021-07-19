package request

import (
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/model"
)

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	utils.PageInfo
}
