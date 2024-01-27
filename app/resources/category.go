package resources

import (
	"net/http"
	"server/app/repositories"
	"server/app/resources/adapters"

	"github.com/gin-gonic/gin"
)

func CreateCategoryHandler(c *gin.Context) {
	var data adapters.CreateCategoryRequestAdapter

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := repositories.Create(data)

	c.JSON(http.StatusOK, result)

}
