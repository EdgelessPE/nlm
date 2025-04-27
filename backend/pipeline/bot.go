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

	// 准备 qa
	println("Preparing qa...")
	err = service.QaPreparePackages(botBuilds)
	if err != nil {
		return err
	}
	println("Qa prepared successfully with", len(botBuilds), "packages")

	// 运行 qa
	println("Running qa...")
	_, err = service.QaRun(ctx, botBuilds)
	if err != nil {
		return err
	}
	println("Qa run successfully")

	// 刷新软件包索引
	service.RefreshMirrorPkgSoftware(true)

	return nil
}
