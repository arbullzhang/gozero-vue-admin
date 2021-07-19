package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
	"gozero-vue-admin/common/cleardb"
	"gozero-vue-admin/common/log"
)

type Config struct {
	rest.RestConf

	System struct {
		DBType string
		UseMultipoint bool
	}

	Captcha struct {
		KeyLong   int
		ImgWidth  int
		ImgHeight int
	}

	Mysql struct {
		DataSource   string
		MaxIdleConns int
		MaxOpenConns int
		LogMode      bool
		LogZap 	     string
	}

	CacheRedis cache.CacheConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
		BufferTime   int64
	}
	Casbin struct
	 {
		ModelPath string
	}

	Zap log.ZapLog

	ClearDBTimer cleardb.Timer
}
