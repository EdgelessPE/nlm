package pipeline

import (
	"encoding/json"
	"log"
	"nlm/context"
	"nlm/service"
)

func runner(ctx *context.PipelineContext, tasks []string, force bool) error {
	// 生成 bot 数据库
	log.Println("Generating bot database...")
	neps, err := service.BotGenerateDatabase()
	if err != nil {
		return err
	}
	log.Println("Generated bot database with", len(neps), "records")

	// 运行 bot
	log.Println("Running bot...")
	botBuilds, err := service.BotRun(ctx, tasks, force)
	if err != nil {
		log.Println("Bot run failed: ", err.Error())
		return err
	}
	log.Println("Bot run successfully")
	botBuildsJson, err := json.Marshal(botBuilds)
	if err != nil {
		return err
	}
	log.Println("Bot builds: ", string(botBuildsJson))

	if len(botBuilds) > 0 {
		// 准备 qa
		log.Println("Preparing qa...")
		err = service.QaPreparePackages(botBuilds)
		if err != nil {
			return err
		}
		log.Println("Qa prepared successfully with", len(botBuilds), "packages")

		// 运行 qa
		log.Println("Running qa...")
		_, err = service.QaRun(ctx, botBuilds)
		if err != nil {
			return err
		}
		log.Println("Qa run successfully")
	} else {
		log.Println("No bot builds found")
	}

	// 刷新软件包索引
	service.RefreshMirrorPkgSoftware(true)

	return nil
}

func RunBotPipeline(tasks []string, force bool) context.PipelineContext {
	ctx := context.NewPipelineContext()
	go func() {
		log.Println("Running bot pipeline...")
		err := runner(&ctx, tasks, force)
		if err != nil {
			log.Fatalf("Failed to run bot pipeline: %v", err)
		}
		log.Println("Bot pipeline run successfully")
	}()
	return ctx
}
