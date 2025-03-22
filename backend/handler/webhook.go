package handler

import (
	"net/http"

	"nlm/trigger"
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

	key, err := trigger.TriggerWebhook(req.Key)
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
