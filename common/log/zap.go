package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
	"gozero-vue-admin/common/directory"
	"gozero-vue-admin/common/rotatelogs"
)

type ZapLog struct {
	Level         string `mapstructure:"level" json:"Level" yaml:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"Format" yaml:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"Prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"Director"  yaml:"director"`                 // 日志文件夹
	LinkName      string `mapstructure:"link-name" json:"LinkName" yaml:"link-name"`                // 软链接名称
	ShowLine      bool   `mapstructure:"show-line" json:"ShowLine" yaml:"showLine"`                 // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"EncodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"StacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"LogInConsole" yaml:"log-in-console"`  // 输出控制台
}

var level zapcore.Level

func (log *ZapLog) Initialize() (logger *zap.Logger) {
	if ok, _ := directory.PathExists(log.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", log.Director)
		_ = os.Mkdir(log.Director, os.ModePerm)
	}

	switch log.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(log.getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(log.getEncoderCore())
	}
	if log.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	logger.WithOptions(zap.Hooks(ZapLogHandler)) //增加日志钩子处理函数

	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func (log *ZapLog) getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  log.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     log.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case log.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case log.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case log.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case log.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func (log *ZapLog) getEncoder() zapcore.Encoder {
	if log.Format == "json" {
		return zapcore.NewJSONEncoder(log.getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(log.getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func (log *ZapLog) getEncoderCore() (core zapcore.Core) {
	writer, err := rotatelogs.GetWriteSyncer(log.Director, log.LinkName, log.LogInConsole) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(log.getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func (log *ZapLog) CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(log.Prefix + "2021/07/15 - 12:00:00.000"))
}