package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gozero-vue-admin/common/timer"
)

var (
	GormDB  *gorm.DB
	Redis   *redis.Client
	ZapLog  *zap.Logger
	Timer   timer.Timer = timer.NewTimerTask()
)


