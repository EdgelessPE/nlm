package handler

import (
	"fmt"
	"log"
	"net/http"

	"nlm/config"
	"nlm/pipeline"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func RegisterWebhookRoutes(r *gin.RouterGroup) {
	webhook := r.Group("/webhook")
	webhook.POST("/trigger", TriggerWebhook)
}

func TriggerWebhook(c *gin.Context) {
	var req vo.WebhookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Invalid request body",
			Data: nil,
		})
		return
	}

	if req.Token != config.ENV.WEBHOOK_TOKEN {
		c.JSON(http.StatusUnauthorized, vo.BaseResponse[any]{
			Code: 401,
			Msg:  "Invalid token",
			Data: nil,
		})
		return
	}

	key, err := triggerWebhook(req.Key, req.Params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to trigger webhook",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, vo.BaseResponse[string]{
		Code: 0,
		Msg:  "Webhook triggered successfully",
		Data: key,
	})
}

func triggerWebhook(key string, _ any) (string, error) {
	log.Println("Triggering webhook with key:", key)
	switch key {
	case "ept":
		ctx := pipeline.RunEptPipeline()
		return ctx.Id, nil
	default:
		return "", fmt.Errorf("invalid webhook key: %s", key)
	}
}
