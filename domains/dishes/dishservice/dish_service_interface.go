package dishservice

import (
	"context"
	"restaurant-api/models"

	"github.com/google/uuid"
)

type DishService interface {
	CreateDish(dish *models.Dish) error
	GetDish(ctx context.Context, id uuid.UUID) (*models.Dish, error)
	ListDishes() ([]models.Dish, error)
	UpdateDish(dish *models.Dish) error
	SearchDishes(query string) ([]models.Dish, error)
	DeleteDishes(id uuid.UUID) error
}
