// 自动生成模板SysOperationRecord
package model

import (
	"time"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
)

// 如果含有time.Time 请自行import time包
type SysOperationRecord struct {
	global.BaseModel
	Ip           string        `json:"ip" gorm:"column:ip;comment:请求ip"`                             // 请求ip
	Method       string        `json:"method" gorm:"column:method;comment:请求方法"`                    // 请求方法
	Path         string        `json:"path" gorm:"column:path;comment:请求路径"`                        // 请求路径
	Status       int           `json:"status" gorm:"column:status;comment:请求状态"`                    // 请求状态
	Latency      time.Duration `json:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"` // 延迟
	Agent        string        `json:"agent" gorm:"column:agent;comment:代理"`                          // 代理
	ErrorMessage string        `json:"error_message" gorm:"column:error_message;comment:错误信息"`  	   // 错误信息
	Body         string        `json:"body" gorm:"type:longtext;column:body;comment:请求Body"`          // 请求Body
	Resp         string        `json:"resp" gorm:"type:longtext;column:resp;comment:响应Body"`          // 响应Body
	UserID       int           `json:"user_id" gorm:"column:user_id;comment:用户id"`                    // 用户id
	User         SysUser       `json:"user"`
}

type OperationRecordSearch struct {
	SysOperationRecord `optional`
	utils.PageInfo
}

//@function: CreateSysOperationRecord
//@description: 创建记录
//@param: sysOperationRecord SysOperationRecord
//@return: err error
func CreateSysOperationRecord(sysOperationRecord SysOperationRecord) (err error) {
	err = global.GormDB.Create(&sysOperationRecord).Error
	return err
}

//@function: DeleteSysOperationRecordByIds
//@description: 批量删除记录
//@param: ids utils.IdsReq
//@return: err error
func DeleteSysOperationRecordByIds(ids utils.IdsReq) (err error) {
	err = global.GormDB.Delete(&[]SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

//@function: DeleteSysOperationRecord
//@description: 删除操作记录
//@param: sysOperationRecord SysOperationRecord
//@return: err error
func DeleteSysOperationRecord(sysOperationRecord SysOperationRecord) (err error) {
	err = global.GormDB.Delete(&sysOperationRecord).Error
	return err
}

//@function: DeleteSysOperationRecord
//@description: 根据id获取单条操作记录
//@param: id uint
//@return: err error, sysOperationRecord SysOperationRecord
func GetSysOperationRecord(id uint) (err error, sysOperationRecord SysOperationRecord) {
	err = global.GormDB.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

//@function: GetSysOperationRecordInfoList
//@description: 分页获取操作记录列表
//@param: info OperationRecordSearch
//@return: err error, list interface{}, total int64
func GetSysOperationRecordInfoList(info OperationRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GormDB.Model(&SysOperationRecord{})
	var sysOperationRecords []SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return err, sysOperationRecords, total
}
