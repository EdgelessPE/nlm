package handler

import (
	"io"
	"log"
	"net/http"
	"strings"

	"nlm/config"
	"nlm/trigger"
	"nlm/utils"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func RegisterWebhookRoutes(r *gin.RouterGroup) {
	r.POST("/webhook", TriggerWebhook)
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
	log.Println("Triggered by GitHub Webhook, delivery id:", c.Request.Header.Get("X-GitHub-Delivery"))

	ok, err := utils.VerifySignature(config.ENV.WEBHOOK_TOKEN, string(body), signature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to verify signature",
			Data: nil,
		})
		return
	}
	if !ok {
		c.JSON(http.StatusUnauthorized, vo.BaseResponse[any]{
			Code: 401,
			Msg:  "Invalid signature",
			Data: nil,
		})
		return
	}

	key, err := trigger.TriggerWebhook(c.Request.Header.Get("X-GitHub-Event"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to trigger webhook: " + err.Error(),
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
