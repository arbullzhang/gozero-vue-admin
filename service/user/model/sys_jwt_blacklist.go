package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"gozero-vue-admin/common/global"
)

type JwtBlacklist struct {
	global.BaseModel
	Jwt string `gorm:"type:text;comment:jwt"`
}

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList JwtBlacklist
//@return: err error
func JsonInBlacklist(jwtList JwtBlacklist) (err error) {
	err = global.GormDB.Create(&jwtList).Error
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool
func IsBlacklist(jwt string) bool {
	err := global.GormDB.Where("jwt = ?", jwt).First(&JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string
func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.Redis.Get(userName).Result()
	return err, redisJWT
}

//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error
func SetRedisJWT(jwt string, userName string, expiresTime int64) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(expiresTime) * time.Second
	err = global.Redis.Set(userName, jwt, timer).Err()
	return err
}