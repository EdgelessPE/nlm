package trigger

import (
	"nlm/pipeline"
	"nlm/service"

	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New()

	// 每天 0 点清理
	c.AddFunc("0 0 * * *", func() {
		// 清理过期日志
		service.CleanLogs()
		// 清理过期临时存储
		service.CleanTempStorage()
	})

	// 每天凌晨 4 点执行 Bot 工作流
	c.AddFunc("0 4 * * *", func() {
		pipeline.RunBotPipeline([]string{"scoop/curl"}, true)
	})

	// 每天下午 16 点执行 Ept 工作流
	c.AddFunc("0 16 * * *", func() {
		pipeline.RunEptPipeline()
	})

	c.Start()
}
