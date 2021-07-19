package model

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
)

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string                 `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}

//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: err error, treeMap map[string][]SysMenu
func getMenuTreeMap(authorityId string) (err error, treeMap map[string][]SysMenu) {
	var allMenus []SysMenu
	treeMap = make(map[string][]SysMenu)
	err = global.GormDB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: authorityId string
//@return: err error, menus []SysMenu
func GetMenuTree(authorityId string) (err error, menus []SysMenu) {
	err, menuTree := getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *SysMenu, treeMap map[string][]SysMenu
//@return: err error
func getChildrenList(menu *SysMenu, treeMap map[string][]SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@function: GetInfoList
//@description: 获取路由分页
//@return: err error, list interface{}, total int64
func GetInfoList() (err error, list interface{}, total int64) {
	var menuList []SysBaseMenu
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *SysBaseMenu, treeMap map[string][]SysBaseMenu
//@return: err error
func getBaseChildrenList(menu *SysBaseMenu, treeMap map[string][]SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@function: AddBaseMenu
//@description: 添加基础路由
//@param: menu SysBaseMenu
//@return: error
func AddBaseMenu(menu SysBaseMenu) error {
	if !errors.Is(global.GormDB.Where("name = ?", menu.Name).First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.GormDB.Create(&menu).Error
}

//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: err error, treeMap map[string][]SysBaseMenu
func getBaseMenuTreeMap() (err error, treeMap map[string][]SysBaseMenu) {
	var allMenus []SysBaseMenu
	treeMap = make(map[string][]SysBaseMenu)
	err = global.GormDB.Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: err error, menus []SysBaseMenu
func GetBaseMenuTree() (err error, menus []SysBaseMenu) {
	err, treeMap := getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []SysBaseMenu, authorityId string
//@return: err error
func AddMenuAuthority(menus []SysBaseMenu, authorityId string) (err error) {
	var auth SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = SetMenuAuthority(&auth)
	return err
}

//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: err error, menus []SysMenu
func GetMenuAuthority(info *utils.GetAuthorityId) (err error, menus []SysMenu) {
	err = global.GormDB.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	//err = global.GormDB.Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}