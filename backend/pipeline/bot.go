package pipeline

import (
	"encoding/json"
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
	botBuilds, err := service.BotRun(ctx, tasks, force)
	if err != nil {
		println("Bot run failed: ", err.Error())
		return err
	}
	println("Bot run successfully")
	botBuildsJson, err := json.Marshal(botBuilds)
	if err != nil {
		return err
	}
	println("Bot builds: ", string(botBuildsJson))

	return nil
}
