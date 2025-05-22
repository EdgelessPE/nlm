package handler

import (
	"net/http"
	"nlm/service"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func RegisterNepRoutes(r *gin.RouterGroup) {
	nep := r.Group("/nep")
	nep.GET("/list", GetNeps)
	nep.GET("/releases", GetReleases)
	nep.GET("/scopes", GetScopes)
	nep.GET("/release_versions", GetReleaseVersions)
}

func GetNeps(c *gin.Context) {
	var params vo.NepParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Failed to bind query : " + err.Error(),
			Data: nil,
		})
		return
	}
	neps, total, err := service.GetNeps(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to get neps : " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code:  0,
		Msg:   "Success",
		Data:  neps,
		Total: total,
	})
}

func GetReleases(c *gin.Context) {
	var params vo.ReleaseParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{
			Code: 400,
			Msg:  "Failed to bind query : " + err.Error(),
			Data: nil,
		})
		return
	}
	releases, total, err := service.GetReleases(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to get releases : " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code:  0,
		Msg:   "Success",
		Data:  releases,
		Total: total,
	})
}

func GetScopes(c *gin.Context) {
	scopes, err := service.GetScopes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to get scopes : " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code: 0,
		Msg:  "Success",
		Data: scopes,
	})
}

func GetReleaseVersions(c *gin.Context) {
	nepId := c.Query("id")
	versions, err := service.GetReleaseVersions(nepId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{
			Code: 500,
			Msg:  "Failed to get release versions : " + err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.BaseResponse[any]{
		Code: 0,
		Msg:  "Success",
		Data: versions,
	})
}
