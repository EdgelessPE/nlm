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
			Msg:  "Invalid request: " + err.Error(),
			Data: nil,
		})
		return
	}

	res := pipeline.RunBotPipeline(req.Tasks, req.Force)
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code: 0,
		Msg:  "Bot pipeline run successfully",
		Data: res,
	})
}

func RunEptPipeline(c *gin.Context) {
	res := pipeline.RunEptPipeline()
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code: 0,
		Msg:  "Ept pipeline run successfully",
		Data: res,
	})
}

func CancelBotPipeline(c *gin.Context) {
	id := c.Param("id")
	err := pipeline.CancelBotPipeline(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Failed to cancel bot pipeline: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code: 0,
		Msg:  "Bot pipeline canceled successfully",
		Data: id,
	})
}

func CancelEptPipeline(c *gin.Context) {
	id := c.Param("id")
	err := pipeline.CancelEptPipeline(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Failed to cancel ept pipeline: " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code: 0,
		Msg:  "Ept pipeline canceled successfully",
		Data: id,
	})
}
