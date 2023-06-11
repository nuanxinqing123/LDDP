package cron

import (
	"github.com/robfig/cron/v3"
	"time"
)

var err error
var c *cron.Cron

// Task 定时任务
func Task() error {
	// 刷新并启用启动任务
	c = cron.New(cron.WithLocation(time.FixedZone("CST", 8*3600))) // 设置时区

	// 定时任务区域

	// 定时任务结束
	if err != nil {
		return err
	}
	c.Start()
	return nil
}
