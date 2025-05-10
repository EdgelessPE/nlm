package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func PrintEncoding(c *gin.Context) {
	c.Next()
	log.Printf("Content-Encoding: %s", c.Writer.Header().Get("Content-Encoding"))
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	}
}
