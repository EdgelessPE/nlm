package handler

import (
	"net/http"
	"nlm/service"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func RegisterStorageRoutes(r *gin.RouterGroup) {
	storage := r.Group("/storage")
	storage.GET("/url/:uuid", GetStorageUrl)
}

func GetStorageUrl(c *gin.Context) {
	uuid := c.Param("uuid")
	url, err := service.GetStorageUrl(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[string]{
		Code: http.StatusOK,
		Msg:  "success",
		Data: url,
	})
}
