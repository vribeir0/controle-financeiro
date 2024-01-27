package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Run() {
	getRoutes()

	router.Run(":8080")
}

func getRoutes() {
	api := router.Group("/api")
	setupCategoryRoutes(api)

}
