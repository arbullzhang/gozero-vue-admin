package request

import (
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/model"
)

// api分页条件查询及排序结构体
type SearchApiParams struct {
	model.SysApi    `optional`
	utils.PageInfo
	OrderKey string `json:"orderKey,optional"` // 排序
	Desc     bool   `json:"desc,optional"`     // 排序方式:升序false(默认)|降序true
}
