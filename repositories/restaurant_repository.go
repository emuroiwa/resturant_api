package repositories

import (
	"restaurant-api/models"

	"gorm.io/gorm"
)

type RestaurantRepository interface {
	Create(restaurant *models.Restaurant) error
	FindByID(id uint) (*models.Restaurant, error)
	FindAll() ([]models.Restaurant, error)
}

type RestaurantRepositoryImpl struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return &RestaurantRepositoryImpl{db: db}
}

func (r *RestaurantRepositoryImpl) Create(restaurant *models.Restaurant) error {
	return r.db.Create(restaurant).Error
}

func (r *RestaurantRepositoryImpl) FindByID(id uint) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	err := r.db.First(&restaurant, id).Error
	if err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (r *RestaurantRepositoryImpl) FindAll() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	err := r.db.Find(&restaurants).Error
	return restaurants, err
}
