package pipeline

import (
	"nlm/context"
	"nlm/service"
)

func RunBotPipeline(tasks []string, force bool) error {
	ctx := context.NewPipelineContext()

	// 生成 bot 数据库
	println("Generating bot database...")
	neps, err := service.BotGenerateDatabase()
	if err != nil {
		return err
	}
	println("Generated bot database with", len(neps), "records")

	// 运行 bot
	println("Running bot...")
	_, err = service.BotRun(ctx, tasks, force)
	if err != nil {
		return err
	}
	println("Bot run successfully")

	return nil
}
