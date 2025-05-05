package pipeline

import (
	"log"
	"nlm/context"
	"nlm/db"
	"nlm/model"
	"nlm/service"
	"nlm/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func runEpt(ctx *context.PipelineContext) error {
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

	return nil
}

func RunEptPipeline() context.PipelineContext {
	ctx := context.NewPipelineContext()
	go func() {
		log.Println("Running ept pipeline...")
		err := runEpt(&ctx)
		if err != nil {
			log.Fatalf("Failed to run ept pipeline: %v", err)
		}
		log.Println("Ept pipeline run successfully")
	}()
	return ctx
}
