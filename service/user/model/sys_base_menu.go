package model

import (
	"errors"
	"gorm.io/gorm"
	"gozero-vue-admin/common/global"
)

type SysBaseMenu struct {
	global.BaseModel
	MenuLevel     uint                   `json:"-"`
	ParentId      string                 `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Path          string                 `json:"path" gorm:"comment:路由path"`        // 路由path
	Name          string                 `json:"name" gorm:"comment:路由name"`        // 路由name
	Hidden        bool                   `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component     string                 `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort          int                    `json:"sort" gorm:"comment:排序标记"`          // 排序标记
	Meta          `json:"meta" gorm:"comment:附加属性"`                                 // 附加属性
	SysAuthoritys []SysAuthority         `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu          `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter `json:"parameters"`
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm:"comment:菜单名"`                // 菜单名
	Icon        string `json:"icon" gorm:"comment:菜单图标"`                // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`         // 自动关闭tab
}

type SysBaseMenuParameter struct {
	global.BaseModel
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"` // 地址栏携带参数为params还是query
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`            // 地址栏携带参数的key
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`            // 地址栏携带参数的值
}

//@function: DeleteBaseMenu
//@description: 删除基础路由
//@param: id float64
//@return: err error
func DeleteBaseMenu(id float64) (err error) {
	err = global.GormDB.Preload("Parameters").Where("parent_id = ?", id).First(&SysBaseMenu{}).Error
	if err != nil {
		var menu SysBaseMenu
		db := global.GormDB.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		err = global.GormDB.Delete(&SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		if len(menu.SysAuthoritys) > 0 {
			err = global.GormDB.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys)
		} else {
			err = db.Error
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

//@function: UpdateBaseMenu
//@description: 更新路由
//@param: menu SysBaseMenu
//@return: err error
func UpdateBaseMenu(menu SysBaseMenu) (err error) {
	var oldMenu SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err = global.GormDB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				global.ZapLog.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.ZapLog.Debug(txErr.Error())
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k, _ := range menu.Parameters {
				menu.Parameters[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				global.ZapLog.Debug(txErr.Error())
				return txErr
			}
		}

		txErr = db.Updates(upDateMap).Error
		if txErr != nil {
			global.ZapLog.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

//@function: GetBaseMenuById
//@description: 返回当前选中menu
//@param: id float64
//@return: err error, menu SysBaseMenu
func GetBaseMenuById(id float64) (err error, menu SysBaseMenu) {
	err = global.GormDB.Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}
