package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	api "github.com/redcowe/personal_site_backend/api"
)

func main() {
	r := gin.Default()

	url := ""
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}

	setOrigin(mode, &url)

	r.Use(api.Cors(url))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/books", func(c *gin.Context) {

		c.JSON(http.StatusOK, api.GetCurrent().Resources)
	})

	r.GET("/read", func(c *gin.Context) {
		c.JSON(http.StatusOK, api.GetRead())
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func setOrigin(mode string, url *string) {
	if mode == "release" {
		*url = "https://api.jcowe.jp"
		log.Print("Using " + *url + " as origin")
		return
	}
	*url = "http://localhost:3000"
	log.Print("Using " + *url + " as origin")
}
