package handler

import (
	"net/http"
	"nlm/service"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func RegisterNepRoutes(r *gin.RouterGroup) {
	nep := r.Group("/nep")
	nep.GET("/neps", GetNeps)
	nep.GET("/:scope/:name/releases", GetReleases)
}

func GetNeps(c *gin.Context) {
	offset, limit, err := GetOffsetAndLimit(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Failed to get offset and limit : " + err.Error(),
			Data: nil,
		})
		return
	}
	neps, err := service.GetNepsWithPagination(offset, limit, c.Query("q"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to get neps : " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code: 0,
		Msg:  "Success",
		Data: neps,
	})
}

func GetReleases(c *gin.Context) {
	scope := c.Param("scope")
	name := c.Param("name")
	releases, err := service.GetReleases(scope, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to get releases : " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code: 0,
		Msg:  "Success",
		Data: releases,
	})
}
