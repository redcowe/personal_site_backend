package api

import (
	"github.com/gin-gonic/gin"
)

func Cors(origin string) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")
		c.Next()
	}
}
