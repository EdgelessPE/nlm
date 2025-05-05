package handler

import (
	"net/http"
	"nlm/pipeline"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

type RunBotPipelineRequest struct {
	Tasks []string `json:"tasks"`
	Force bool     `json:"force"`
}

func RegisterPipelineRoutes(r *gin.RouterGroup) {
	pipeline := r.Group("/pipeline")

	// 执行 pipeline
	run := pipeline.Group("/run")
	run.POST("/bot", RunBotPipeline)
	run.POST("/ept", RunEptPipeline)
}

func RunBotPipeline(c *gin.Context) {
	var req RunBotPipelineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Invalid request",
			Data: nil,
		})
		return
	}

	ctx := pipeline.RunBotPipeline(req.Tasks, req.Force)
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code: 0,
		Msg:  "Bot pipeline run successfully",
		Data: ctx.Id,
	})
}

func RunEptPipeline(c *gin.Context) {
	ctx := pipeline.RunEptPipeline()
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code: 0,
		Msg:  "Ept pipeline run successfully",
		Data: ctx.Id,
	})
}
