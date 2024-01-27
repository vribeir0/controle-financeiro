package repositories

import (
	"server/app/config"
	"server/app/models"
	"server/app/resources/adapters"
)

// TODO: Implementar struct de repositorio
func Create(data adapters.CreateCategoryRequestAdapter) models.Category {
	category := models.Category{Name: data.Name, Description: data.Description}
	result := config.DB.Create(&category)

	if result.Error != nil {
		panic(result.Error)
	}

	return category
}
