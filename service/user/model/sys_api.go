package model

import (
	"errors"
	"gorm.io/gorm"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
)

type SysApi struct {
	global.BaseModel
	Path        string `json:"path" gorm:"comment:api路径"`                    // api路径
	Description string `json:"description" gorm:"comment:api中文描述"`           // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`                 // api组
	Method      string `json:"method" gorm:"default:POST" gorm:"comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

//@function: CreateApi
//@description: 新增基础api
//@param: api SysApi
//@return: err error
func CreateApi(api SysApi) (err error) {
	if !errors.Is(global.GormDB.Where("path = ? AND method = ?", api.Path, api.Method).First(&SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GormDB.Create(&api).Error
}

//@function: DeleteApi
//@description: 删除基础api
//@param: api SysApi
//@return: err error
func DeleteApi(api SysApi) (err error) {
	err = global.GormDB.Delete(&api).Error
	ClearCasbin(1, api.Path, api.Method)
	return err
}

//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api SysApi, info request.PageInfo, order string, desc bool
//@return: err error
func GetAPIInfoList(api SysApi, info utils.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GormDB.Model(&SysApi{})
	var apiList []SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []SysApi
func GetAllApis() (err error, apis []SysApi) {
	err = global.GormDB.Find(&apis).Error
	return
}

//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api SysApi
func GetApiById(id float64) (err error, api SysApi) {
	err = global.GormDB.Where("id = ?", id).First(&api).Error
	return
}


//@function: UpdateApi
//@description: 根据id更新api
//@param: api SysApi
//@return: err error
func UpdateApi(api SysApi) (err error) {
	var oldA SysApi
	err = global.GormDB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GormDB.Where("path = ? AND method = ?", api.Path, api.Method).First(&SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.GormDB.Save(&api).Error
		}
	}
	return err
}

//@function: DeleteApis
//@description: 删除选中API
//@param: apis []SysApi
//@return: err error
func DeleteApisByIds(ids utils.IdsReq) (err error) {
	err = global.GormDB.Delete(&[]SysApi{}, "id in ?", ids.Ids).Error
	return err
}
