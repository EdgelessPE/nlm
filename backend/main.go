package main

import (
	"net/http"
	"nlm/db"
	"nlm/handler"
	"nlm/model"
	"nlm/service"
	"nlm/vo"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	db.DB.AutoMigrate(&model.Nep{}, &model.Release{})

	service.AddRelease("Google", "Chrome", "139.0.0.0", "I", time.Now(), "114514")

	// 启动服务器
	server := gin.Default()
	r := server.Group("/api")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, vo.BaseResponse[string]{
			Code: 0,
			Msg:  "success",
			Data: "pong",
		})
	})

	handler.RegisterWebhookRoutes(r)
	server.Run("0.0.0.0:3001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
