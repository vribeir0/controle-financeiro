package routes

import (
	"server/app/resources"

	"github.com/gin-gonic/gin"
)

func setupCategoryRoutes(rg *gin.RouterGroup) {
	category := rg.Group("/category")
	{
		category.POST("/", resources.CreateCategoryHandler)
		category.GET("/", resources.GetAllCategoryHandler)
		category.GET("/:id", resources.GetCategoryHandler)
		category.PUT("/:id", resources.UpdateCategoryHandler)
		category.DELETE("/:id", resources.DeleteCategoryHandler)
	}
}
