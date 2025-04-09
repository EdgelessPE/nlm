package service

import (
	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New()

	// 每天 0 点清理
	c.AddFunc("0 0 * * *", func() {
		// 清理过期日志
		CleanLogs()
		// 清理过期临时存储
		CleanTempStorage()
	})

	c.Start()
}
