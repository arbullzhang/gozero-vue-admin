package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"os"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/gorm"
	"gozero-vue-admin/service/user/api/internal/config"
	"gozero-vue-admin/service/user/api/internal/handler"
	"gozero-vue-admin/service/user/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	// 读取配置文件
	flag.Parse()
	var c config.Config
	// 不是从命令行启动的
	if len(os.Args) <= 1 {
		if path, err := os.Getwd(); err == nil {
			configFilePath := path + "/service/user/api/etc/user-api.yaml"
			conf.MustLoad(configFilePath, &c)
		}
	} else {
		conf.MustLoad(*configFile, &c)
	}

	// 初始化服务
	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 初始化日志
	global.ZapLog = c.Zap.Initialize()

	// 初始化数据库
	global.GormDB = gorm.Initialize("mysql", c.Mysql.DataSource, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.LogMode, c.Mysql.LogZap)
	if global.GormDB != nil {
		db, _ := global.GormDB.DB()
		c.ClearDBTimer.StartTimer()  // 定时清除数据库表中的日志信息

		defer db.Close()
	}

	// 初始化Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: 	  c.CacheRedis[0].Host,
		Password: "",
		DB:       0,
	})
	ping, err := redisClient.Ping().Result()
	if err != nil {
		global.ZapLog.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.ZapLog.Info("redis connect ping response:", zap.String("ping", ping))
		global.Redis = redisClient
	}

	//注册路由
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
