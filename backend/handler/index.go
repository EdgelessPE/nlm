package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	RegisterWebhookRoutes(r)
	RegisterNepRoutes(r)
	RegisterLogRoutes(r)
}
