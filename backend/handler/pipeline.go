package handler

import (
	"net/http"
	"nlm/pipeline"
	"nlm/service"
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

	// 取消 pipeline
	cancel := pipeline.Group("/cancel")
	cancel.DELETE("/bot/:id", CancelBotPipeline)
	cancel.DELETE("/ept/:id", CancelEptPipeline)

	// 获取 pipeline
	pipeline.GET("/list", GetPipelines)
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
	if res.IsNewPipeline {
		c.JSON(http.StatusOK, vo.BaseResponse[any]{
			Code: 0,
			Msg:  "Bot pipeline run successfully",
			Data: res.PipelineContext.Id,
		})
	} else {
		c.JSON(http.StatusOK, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Bot pipeline already running",
			Data: res.PipelineContext.Id,
		})
	}
}
func RunEptPipeline(c *gin.Context) {
	res := pipeline.RunEptPipeline()
	if res.IsNewPipeline {
		c.JSON(http.StatusOK, vo.BaseResponse[any]{
			Code: 0,
			Msg:  "Ept pipeline run successfully",
			Data: res.PipelineContext.Id,
		})
	} else {
		c.JSON(http.StatusOK, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Ept pipeline already running",
			Data: res.PipelineContext.Id,
		})
	}
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

func GetPipelines(c *gin.Context) {
	var params vo.PipelineParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Invalid request: " + err.Error(),
			Data: nil,
		})
		return
	}

	pipelines, total, err := service.GetPipelines(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code:  500,
			Msg:   "Failed to get pipelines: " + err.Error(),
			Data:  nil,
			Total: total,
		})
		return
	}

	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code:  0,
		Msg:   "Pipelines fetched successfully",
		Data:  pipelines,
		Total: total,
	})
}
