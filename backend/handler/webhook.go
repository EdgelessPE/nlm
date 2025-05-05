package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"nlm/config"
	"nlm/pipeline"
	"nlm/utils"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func RegisterWebhookRoutes(r *gin.RouterGroup) {
	webhook := r.Group("/webhook")
	webhook.POST("/trigger", TriggerWebhook)
}

func TriggerWebhook(c *gin.Context) {
	// 读出 body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to read request body",
			Data: nil,
		})
		return
	}

	// 校验签名
	signature := strings.TrimPrefix(c.Request.Header.Get("X-Hub-Signature-256"), "sha256=")
	if signature == "" {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Failed to read signature from request header",
			Data: nil,
		})
		return
	}
	log.Println("signature", signature)

	if !utils.VerifySignature(config.ENV.WEBHOOK_TOKEN, string(body), signature) {
		c.JSON(http.StatusUnauthorized, vo.BaseResponse[any]{
			Code: 401,
			Msg:  "Invalid token",
			Data: nil,
		})
		return
	}

	// 触发 webhook
	var req vo.GitHubWebhookRequest
	if err := json.Unmarshal(body, &req); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Invalid request body",
			Data: nil,
		})
		return
	}

	key, err := triggerWebhook(req.Event)
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

func triggerWebhook(event string) (string, error) {
	log.Println("Triggering webhook with event:", event)
	switch event {
	case "release":
		ctx := pipeline.RunEptPipeline()
		return ctx.Id, nil
	default:
		return "", fmt.Errorf("invalid webhook event: %s", event)
	}
}
