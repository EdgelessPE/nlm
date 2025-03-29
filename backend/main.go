package main

import (
	"net/http"
	"nlm/db"
	"nlm/domain"
	"nlm/handler"
	"nlm/model"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	db.DB.AutoMigrate(&model.Nep{}, &model.Release{})

	// 初始化 nep
	domain.InitNepsWithBotTask()

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

	handler.RegisterRoutes(r)
	server.Run("0.0.0.0:3001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
