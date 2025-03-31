package service

import (
	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New()

	// 每天 0 点清理日志
	c.AddFunc("0 0 * * *", func() {
		CleanLogs()
	})

	c.Start()
}
