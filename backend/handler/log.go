package handler

import (
	"net/http"
	"nlm/service"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func RegisterLogRoutes(r *gin.RouterGroup) {
	log := r.Group("/log")
	log.GET("/stream/:moduleName/:pipelineId", StreamLog)
}

func StreamLog(c *gin.Context) {
	pipelineId := c.Param("pipelineId")
	moduleName := c.Param("moduleName")

	if pipelineId == "" || moduleName == "" {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Missing pipelineId or moduleName",
			Data: nil,
		})
		return
	}

	service.StreamLog(c, pipelineId, moduleName)
}
