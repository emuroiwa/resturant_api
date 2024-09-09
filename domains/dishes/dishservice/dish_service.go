package dishservice

import (
	"context"
	"encoding/json"
	"errors"
	"restaurant-api/domains/rabbitmq"
	"restaurant-api/models"
	"restaurant-api/repositories"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type DishServiceImpl struct {
	DishRepo    repositories.DishRepository
	RedisClient *redis.Client
}

func NewDishService(dishRepo repositories.DishRepository, redisClient *redis.Client) DishService {
	return &DishServiceImpl{
		DishRepo:    dishRepo,
		RedisClient: redisClient,
	}
}

func (ds *DishServiceImpl) CreateDish(dish *models.Dish) error {
	ratingJSON, err := json.Marshal(dish)
	if err != nil {
		return errors.New("failed to serialize dish")
	}

	err = rabbitmq.PublishMessage(string(ratingJSON), "dish_queue")
	if err != nil {
		return errors.New("failed to publish dish to queue")
	}

	return nil
}

func (ds *DishServiceImpl) GetDish(ctx context.Context, id uuid.UUID) (*models.Dish, error) {
	cacheKey := "dish_" + id.String()

	cachedResponse, err := ds.RedisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		dish, err := ds.DishRepo.FindByID(id)
		if err != nil {
			return nil, errors.New("dish not found")
		}

		responseBody, err := json.Marshal(dish)
		if err != nil {
			return nil, errors.New("failed to serialize response")
		}

		err = ds.RedisClient.Set(ctx, cacheKey, responseBody, time.Minute*10).Err()
		if err != nil {
			return nil, errors.New("failed to cache response")
		}

		return dish, nil
	} else if err != nil {
		return nil, errors.New("failed to get cached response")
	}

	var dish models.Dish
	err = json.Unmarshal([]byte(cachedResponse), &dish)
	if err != nil {
		return nil, errors.New("failed to deserialize cached response")
	}

	return &dish, nil
}

func (ds *DishServiceImpl) ListDishes() ([]models.Dish, error) {
	return ds.DishRepo.FindAll()
}

func (ds *DishServiceImpl) UpdateDish(dish *models.Dish) error {
	return ds.DishRepo.Update(dish)
}

func (ds *DishServiceImpl) SearchDishes(query string) ([]models.Dish, error) {
	return ds.DishRepo.SearchDishes(query)
}

func (ds *DishServiceImpl) DeleteDishes(id uuid.UUID) error {
	return ds.DishRepo.Delete(id)
}
