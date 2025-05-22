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
	storage.GET("/list", GetStorages)
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

func GetStorages(c *gin.Context) {
	var params vo.GetStoragesParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Failed to bind query : " + err.Error(),
			Data: nil,
		})
		return
	}
	storages, total, err := service.GetStorages(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to get storages : " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code:  0,
		Msg:   "success",
		Data:  storages,
		Total: total,
	})
}
