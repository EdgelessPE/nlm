package handler

import (
	"net/http"
	"nlm/service"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func RegisterEptRoutes(r *gin.RouterGroup) {
	ept := r.Group("/ept")
	ept.GET("/list", GetEpts)
}

func GetEpts(c *gin.Context) {
	var params vo.GetEptsParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Failed to bind query : " + err.Error(),
			Data: nil,
		})
		return
	}
	epts, total, err := service.GetEpts(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to get epts : " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code:  0,
		Msg:   "Success",
		Data:  epts,
		Total: total,
	})
}
