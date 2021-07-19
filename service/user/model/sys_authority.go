package model

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"time"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
)

type SysAuthority struct {
	CreatedAt       time.Time      // 创建时间
	UpdatedAt       time.Time      // 更新时间
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	DefaultRouter   string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

type SysAuthorityCopyResponse struct {
	Authority      SysAuthority `json:"authority"`
	OldAuthorityId string       `json:"oldAuthorityId"` // 旧角色ID
}

//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth SysAuthority
//@return: err error, authority SysAuthority
func CreateAuthority(auth SysAuthority) (err error, authority SysAuthority) {
	var authorityBox SysAuthority
	if !errors.Is(global.GormDB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = global.GormDB.Create(&auth).Error
	return err, auth
}

//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: err error, authority SysAuthority
func CopyAuthority(copyInfo SysAuthorityCopyResponse) (err error, authority SysAuthority) {
	var authorityBox SysAuthority
	if !errors.Is(global.GormDB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []SysAuthority{}
	err, menus := GetMenuAuthority(&utils.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	var baseMenu []SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.GormDB.Create(&copyInfo.Authority).Error

	paths := GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}

//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth SysAuthority
//@return: err error, authority SysAuthority
func UpdateAuthority(auth SysAuthority) (err error, authority SysAuthority) {
	err = global.GormDB.Where("authority_id = ?", auth.AuthorityId).First(&SysAuthority{}).Updates(&auth).Error
	return err, auth
}

//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *SysAuthority
//@return: err error
func DeleteAuthority(auth *SysAuthority) (err error) {
	if !errors.Is(global.GormDB.Where("authority_id = ?", auth.AuthorityId).First(&SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GormDB.Where("parent_id = ?", auth.AuthorityId).First(&SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.GormDB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if len(auth.SysBaseMenus) > 0 {
		err = global.GormDB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		//err = db.Association("SysBaseMenus").Delete(&auth)
	} else {
		err = db.Error
	}
	ClearCasbin(0, auth.AuthorityId)
	return err
}

//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64
func GetAuthorityInfoList(info utils.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GormDB
	var authority []SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth SysAuthority
//@return: err error, sa SysAuthority
func GetAuthorityInfo(auth SysAuthority) (err error, sa SysAuthority) {
	err = global.GormDB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

//@function: SetDataAuthority
//@description: 设置角色资源权限
//@param: auth SysAuthority
//@return: error
func SetDataAuthority(auth SysAuthority) error {
	var s SysAuthority
	global.GormDB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GormDB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *SysAuthority
//@return: error
func SetMenuAuthority(auth *SysAuthority) error {
	var s SysAuthority
	global.GormDB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GormDB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *SysAuthority
//@return: err error
func findChildrenAuthority(authority *SysAuthority) (err error) {
	err = global.GormDB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}