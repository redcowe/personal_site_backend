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

	origin := ""
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}

	setOrigin(mode, &origin)

	r.Use(api.Cors(origin))

	r.GET("/books", func(c *gin.Context) {

		c.JSON(http.StatusOK, api.GetCurrent().Resources)
	})

	r.GET("/read", func(c *gin.Context) {
		c.JSON(http.StatusOK, api.GetRead())
	})
	r.Run()

}

func setOrigin(mode string, url *string) {
	if mode == "release" {
		*url = "https://jcowe.jp"
		log.Print("Using " + *url + " as origin")
		return
	}
	*url = "http://localhost:3000"
	log.Print("Using " + *url + " as origin")
}
