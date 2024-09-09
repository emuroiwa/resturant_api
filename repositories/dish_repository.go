package repositories

import (
	"restaurant-api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DishRepository interface {
	Create(dish *models.Dish) error
	FindByID(id uuid.UUID) (*models.Dish, error)
	FindAll() ([]models.Dish, error)
	Update(dish *models.Dish) error
	Delete(id uuid.UUID) error
	SearchDishes(query string) ([]models.Dish, error)
}

type DishRepositoryImpl struct {
	db *gorm.DB
}

func NewDishRepository(db *gorm.DB) DishRepository {
	return &DishRepositoryImpl{db: db}
}

func (r *DishRepositoryImpl) Create(dish *models.Dish) error {
	return r.db.Create(dish).Error
}

func (r *DishRepositoryImpl) FindByID(id uuid.UUID) (*models.Dish, error) {
	var dish models.Dish
	err := r.db.First(&dish, id).Error
	if err != nil {
		return nil, err
	}
	return &dish, nil
}

func (r *DishRepositoryImpl) FindAll() ([]models.Dish, error) {
	var dishes []models.Dish
	err := r.db.Find(&dishes).Error
	return dishes, err
}

func (r *DishRepositoryImpl) Update(dish *models.Dish) error {
	return r.db.Save(dish).Error
}

func (r *DishRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Dish{}, id).Error
}

func (r *DishRepositoryImpl) SearchDishes(query string) ([]models.Dish, error) {
	var dishes []models.Dish
	if err := r.db.Where("name ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&dishes).Error; err != nil {
		return nil, err
	}
	return dishes, nil
}
