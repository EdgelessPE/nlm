package main

import (
	"net/http"
	"nlm/handler"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func main() {
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
