package model

import (
	"errors"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
)

type SysUser struct {
	global.BaseModel
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`                                                     // 用户UUID
	Username    string       `json:"userName" gorm:"comment:用户登录名"`                                                // 用户登录名
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`                                                    // 用户登录密码
	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                  // 用户昵称
	HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"` 								   // 用户角色ID

	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
}

//@function: Register
//@description: 用户注册
//@param: u SysUser
//@return: err error, userInter SysUser
func Register(u SysUser) (err error, userInter SysUser) {
	var user SysUser
	if !errors.Is(global.GormDB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GormDB.Create(&u).Error
	return err, u
}

//@function: Login
//@description: 用户登录
//@param: u *SysUser
//@return: err error, userInter *SysUser
func Login(u *SysUser) (err error, userInter *SysUser) {
	var user SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GormDB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

//@function: ChangePassword
//@description: 修改用户密码
//@param: u *SysUser, newPassword string
//@return: err error, userInter *SysUser
func ChangePassword(u *SysUser, newPassword string) (err error, userInter *SysUser) {
	var user SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GormDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64
func GetUserInfoList(info utils.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GormDB.Model(&SysUser{})
	var userList []SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error
func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.GormDB.Where("uuid = ?", uuid).First(&SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error
func DeleteUser(id float64) (err error) {
	var user SysUser
	err = global.GormDB.Where("id = ?", id).Delete(&user).Error
	return err
}

//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser SysUser
//@return: err error, user SysUser
func SetUserInfo(reqUser SysUser) (err error, user SysUser) {
	err = global.GormDB.Updates(&reqUser).Error
	return err, reqUser
}

//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *SysUser
func FindUserById(id int) (err error, user *SysUser) {
	var u SysUser
	err = global.GormDB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *SysUser
func FindUserByUuid(uuid string) (err error, user *SysUser) {
	var u SysUser
	if err = global.GormDB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
