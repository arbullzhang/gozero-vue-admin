package cleardb

import (
	"fmt"
	"gozero-vue-admin/common/global"
)

type Detail struct {
	TableName    string `mapstructure:"tableName" json:"TableName" yaml:"tableName"`          // 需要清理的表名
	CompareField string `mapstructure:"compareField" json:"CompareField" yaml:"compareField"` // 需要比较时间的字段
	Interval     string `mapstructure:"interval" json:"Interval" yaml:"interval"`             // 时间间隔
}

type Timer struct {
	Start  bool     `mapstructure:"start" json:"Start" yaml:"start"` // 是否启用
	Spec   string   `mapstructure:"spec" json:"Spec" yaml:"spec"`    // CRON表达式
	Detail []Detail `mapstructure:"detail" json:"Detail" yaml:"detail"`
}

func (t *Timer) StartTimer() {
	if t.Start {
		for _, detail := range t.Detail {
			fmt.Println(detail)
			go func(detail Detail) {
				global.Timer.AddTaskByFunc("ClearDB", t.Spec, func() {
					err := ClearTable(global.GormDB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(detail)
		}
	}
}
