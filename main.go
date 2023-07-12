package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	books "github.com/redcowe/personal_site_backend/api"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})
	r.GET("/books", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		c.JSON(http.StatusOK, books.Get().Resources)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
