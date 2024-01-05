package main

import (
	"server/config"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	config.Connect()
	return r
}

func main() {
	router := setupRouter()

	router.Run()

}
