package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func MiddleWarePrintEncoding(c *gin.Context) {
	c.Next()
	log.Printf("Content-Encoding: %s", c.Writer.Header().Get("Content-Encoding"))
}
