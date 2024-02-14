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

func Get(id string) models.Category {
	var category models.Category
	result := config.DB.First(&category, id)

	if result.Error != nil {
		panic(result.Error)
	}

	return category
}

func GetAll() []models.Category {
	var categories []models.Category
	result := config.DB.Find(&categories)

	if result.Error != nil {
		panic(result.Error)
	}

	return categories
}

func Update(id string, data adapters.CreateCategoryRequestAdapter) models.Category {
	var category models.Category
	result := config.DB.First(&category, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = config.DB.Model(&category).Updates(data)

	if result.Error != nil {
		panic(result.Error)
	}

	return category
}

func Delete(id string) {
	var category models.Category
	result := config.DB.First(&category, id)

	if result.Error != nil {
		panic(result.Error)
	}

	config.DB.Delete(&category)
}
