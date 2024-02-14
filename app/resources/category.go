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

func GetCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	result := repositories.Get(id)

	c.JSON(http.StatusOK, result)
}

func GetAllCategoryHandler(c *gin.Context) {
	result := repositories.GetAll()

	c.JSON(http.StatusOK, result)
}

func UpdateCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	var data adapters.CreateCategoryRequestAdapter

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := repositories.Update(id, data)

	c.JSON(http.StatusOK, result)
}

func DeleteCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	repositories.Delete(id)

	c.Status(http.StatusOK)
}
