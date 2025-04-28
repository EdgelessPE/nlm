package main

import (
	"log"
	"net/http"
	"nlm/constant"
	"nlm/db"
	"nlm/domain"
	"nlm/handler"
	"nlm/model"
	"nlm/service"
	"nlm/vo"

	"github.com/aurowora/compress"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	err := db.DB.AutoMigrate(&model.Nep{}, &model.Release{}, &model.Storage{})
	if err != nil {
		log.Fatalf("Failed to migrate nep table: %v", err)
	}

	// 初始化 nep
	domain.InitNepsWithBotTask()

	// 更新包信息
	service.RefreshMirrorPkgSoftware(false)

	// 启动定时任务
	service.InitCron()

	// err = pipeline.RunBotPipeline([]string{"scoop/curl"}, true)
	// if err != nil {
	// 	log.Fatalf("Failed to run bot pipeline: %v", err)
	// }

	// 新建服务器实例
	server := gin.Default()

	// 注册中间件
	server.Use(compress.Compress())
	// server.Use(middleware.MiddleWarePrintEncoding)

	// 注册路由
	r := server.Group(constant.API_PREFIX)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, vo.BaseResponse[string]{
			Code: 0,
			Msg:  "success",
			Data: "pong",
		})
	})
	handler.RegisterRoutes(r)

	// 启动服务器
	server.Run("0.0.0.0:3001")
}
