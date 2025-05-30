package pipeline

import (
	"errors"
	"log"
	"nlm/context"
	"nlm/db"
	"nlm/model"
	"nlm/service"
	"nlm/utils"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

var pipelineCtxEpt *context.PipelineContext

func runEpt(ctx *context.PipelineContext) error {
	// 等待 10 秒
	time.Sleep(10 * time.Second)

	// 获取 GitHub Release
	release, err := utils.GetGitHubLatestRelease("EdgelessPE", "ept")
	if err != nil {
		return err
	}

	// 判断版本号是否存在
	latestVersion := strings.TrimPrefix(release.TagName, "v")
	log.Println("Latest version: ", latestVersion)
	var ept model.Ept
	db.DB.Model(&model.Ept{}).Where("version = ?", latestVersion).First(&ept)
	if ept.ID != uuid.Nil {
		return nil
	}

	// 准备 temp 目录
	tmpDir := "temp"
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return err
	}

	// 下载 Release
	fileName := release.Assets[0].Name
	url := release.Assets[0].BrowserDownloadURL
	filePath := filepath.Join(tmpDir, fileName)
	if err := utils.DownloadFile(url, filePath); err != nil {
		return err
	}

	// 计算文件的 blake3 值
	integrity, err := utils.GetBlake3HashFromFile(filePath)
	if err != nil {
		return err
	}

	// 获取文件大小
	fileStat, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	// 添加存储
	storageKey, err := service.AddStorage(filePath, false)
	if err != nil {
		return err
	}

	// 保存 ept 信息
	ept = model.Ept{
		Name:       fileName,
		Version:    latestVersion,
		StorageKey: storageKey,
		FileSize:   fileStat.Size(),
		Integrity:  integrity,
	}
	db.DB.Create(&ept)

	// 刷新镜像
	service.RefreshMirrorEptToolchain(false)

	return nil
}

func RunEptPipeline() PipelineCreateResult {
	if pipelineCtxEpt != nil {
		log.Printf("Pipeline %s already running", pipelineCtxEpt.Id)
		return PipelineCreateResult{
			PipelineContext: *pipelineCtxEpt,
			IsNewPipeline:   false,
		}
	}

	ctx := context.NewPipelineContext()
	pipelineCtxEpt = &ctx
	pipeline := model.Pipeline{
		Base:      model.Base{ID: uuid.MustParse(ctx.Id)},
		ModelName: "ept",
		Status:    "running",
	}
	db.DB.Create(&pipeline)
	go func() {
		log.Println("Running ept pipeline...")
		err := runEpt(&ctx)
		if err != nil {
			log.Println("Failed to run ept pipeline: ", err.Error())
			pipeline.Status = "failed"
			pipeline.ErrMsg = err.Error()
		} else {
			log.Println("Ept pipeline run successfully")
			pipeline.Status = "success"
		}
		pipeline.FinishedAt = time.Now()
		db.DB.Save(&pipeline)
		pipelineCtxEpt = nil
	}()
	return PipelineCreateResult{
		PipelineContext: ctx,
		IsNewPipeline:   true,
	}
}

func CancelEptPipeline(id string) error {
	if pipelineCtxEpt == nil {
		return errors.New("no ept pipeline running")
	}
	if pipelineCtxEpt.Id != id {
		return errors.New("pipeline id not match")
	}
	pipelineCtxEpt.Cancel()
	pipelineCtxEpt = nil
	go func() {
		time.Sleep(3 * time.Second)
		db.DB.Model(&model.Pipeline{}).Where("id = ?", id).Update("status", "canceled")
	}()
	return nil
}
