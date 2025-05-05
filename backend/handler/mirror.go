package handler

import (
	"net/http"
	"nlm/constant"
	"nlm/service"
	"nlm/vo"

	"github.com/gin-gonic/gin"
)

func RegisterMirrorRoutes(r *gin.RouterGroup) {
	r.GET(constant.ServicePathHello, func(c *gin.Context) {
		c.JSON(http.StatusOK, service.MirrorHello())
	})

	r.GET(constant.ServicePathPkgSoftware, func(c *gin.Context) {
		c.JSON(http.StatusOK, service.MirrorPkgSoftware())
	})

	r.GET(constant.ServicePathEptToolchain, func(c *gin.Context) {
		c.JSON(http.StatusOK, service.MirrorEptToolchain())
	})

	r.GET(constant.ServicePathRedirect, func(c *gin.Context) {
		url, err := service.MirrorRedirect(c.Param("scope"), c.Param("software"), c.Param("file_name"))
		if err != nil {
			c.JSON(http.StatusNotFound, vo.BaseResponse[any]{
				Code: 404,
				Msg:  "Failed to redirect: " + err.Error(),
				Data: nil,
			})
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, url)
	})
}
