package trigger

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

var (
	cronScheduler *cron.Cron
	once          sync.Once
)

// InitCron 初始化 cron 调度器
func InitCron() {
	once.Do(func() {
		cronScheduler = cron.New(cron.WithSeconds())
		cronScheduler.Start()
	})
}

// AddCronJob 添加一个 cron 任务
func AddCronJob(spec string, key string) error {
	if cronScheduler == nil {
		InitCron()
	}

	_, err := cronScheduler.AddFunc(spec, func() {
		// 执行 webhook 触发
		_, err := TriggerWebhook(key)
		if err != nil {
			fmt.Printf("Failed to execute cron job for key %s: %v\n", key, err)
		}
	})

	if err != nil {
		return fmt.Errorf("invalid cron spec: %v", err)
	}

	return nil
}

// RemoveCronJob 移除一个 cron 任务
func RemoveCronJob(key string) {
	if cronScheduler == nil {
		return
	}

	// 获取所有任务
	entries := cronScheduler.Entries()
	for _, entry := range entries {
		// 通过 key 匹配任务并移除
		if entry.Job != nil {
			cronScheduler.Remove(entry.ID)
		}
	}
}

// StopCron 停止所有 cron 任务
func StopCron() {
	if cronScheduler != nil {
		cronScheduler.Stop()
	}
}
