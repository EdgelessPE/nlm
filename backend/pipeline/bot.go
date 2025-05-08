package pipeline

import (
	"encoding/json"
	"log"
	"nlm/context"
	"nlm/db"
	"nlm/model"
	"nlm/service"
	"time"

	"github.com/google/uuid"
)

var pipelineCtxBot *context.PipelineContext

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

		// 更新数据库的流水线状态
		db.DB.Model(&model.Pipeline{}).Where("id = ?", ctx.Id).Update("stage", "qa")

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

func RunBotPipeline(tasks []string, force bool) PipelineCreateResult {
	if pipelineCtxBot != nil {
		log.Printf("Pipeline %s already running", pipelineCtxBot.Id)
		return PipelineCreateResult{
			PipelineContext: *pipelineCtxBot,
			IsNewPipeline:   false,
		}
	}

	ctx := context.NewPipelineContext()
	pipelineCtxBot = &ctx
	pipeline := model.Pipeline{
		Base:      model.Base{ID: uuid.MustParse(ctx.Id)},
		ModelName: "bot",
		Status:    "running",
		Stage:     "bot",
	}
	db.DB.Create(&pipeline)
	go func() {
		log.Println("Running bot pipeline...")
		err := runner(&ctx, tasks, force)
		if err != nil {
			log.Println("Failed to run bot pipeline: ", err.Error())
			pipeline.Status = "failed"
			pipeline.ErrMsg = err.Error()
		} else {
			log.Println("Bot pipeline run successfully")
			pipeline.Status = "success"
		}
		pipeline.FinishedAt = time.Now()
		db.DB.Save(&pipeline)
		pipelineCtxBot = nil
	}()
	return PipelineCreateResult{
		PipelineContext: ctx,
		IsNewPipeline:   true,
	}
}
