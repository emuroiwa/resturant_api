package migrations

import (
	"restaurant-api/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Restaurant{}, &models.User{}, &models.Dish{}, &models.Rating{})
}
