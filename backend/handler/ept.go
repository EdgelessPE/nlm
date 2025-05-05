package handler

import (
	"net/http"
	"nlm/service"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func RegisterEptRoutes(r *gin.RouterGroup) {
	ep := r.Group("/ept")
	ep.GET("/toolchain", GetEptToolchain)
}

func GetEptToolchain(c *gin.Context) {
	ept, err := service.GetEptToolchain()
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
	}
	c.JSON(http.StatusOK, ept)
}
