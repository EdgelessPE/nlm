package handler

import (
	"net/http"
	"nlm/constant"
	"nlm/service"

	"github.com/gin-gonic/gin"
)

func RegisterMirrorRoutes(r *gin.RouterGroup) {
	r.GET(constant.ServicePathHello, func(c *gin.Context) {
		c.JSON(http.StatusOK, service.MirrorHello())
	})

	r.GET(constant.ServicePathPkgSoftware, func(c *gin.Context) {
		c.JSON(http.StatusOK, service.MirrorPkgSoftware())
	})
}
