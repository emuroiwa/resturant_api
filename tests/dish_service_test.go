package tests

import (
	"context"
	"reflect"
	"restaurant-api/domains/dishes/dishservice"
	"restaurant-api/models"
	"restaurant-api/repositories"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func TestDishServiceImpl_CreateDish(t *testing.T) {
	type fields struct {
		DishRepo    repositories.DishRepository
		RedisClient *redis.Client
	}
	type args struct {
		dish *models.Dish
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &dishservice.DishServiceImpl{
				DishRepo:    tt.fields.DishRepo,
				RedisClient: tt.fields.RedisClient,
			}
			if err := ds.CreateDish(tt.args.dish); (err != nil) != tt.wantErr {
				t.Errorf("DishServiceImpl.CreateDish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDishServiceImpl_GetDish(t *testing.T) {
	type fields struct {
		DishRepo    repositories.DishRepository
		RedisClient *redis.Client
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Dish
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &dishservice.DishServiceImpl{
				DishRepo:    tt.fields.DishRepo,
				RedisClient: tt.fields.RedisClient,
			}
			got, err := ds.GetDish(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DishServiceImpl.GetDish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DishServiceImpl.GetDish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDishServiceImpl_ListDishes(t *testing.T) {
	type fields struct {
		DishRepo    repositories.DishRepository
		RedisClient *redis.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Dish
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &dishservice.DishServiceImpl{
				DishRepo:    tt.fields.DishRepo,
				RedisClient: tt.fields.RedisClient,
			}
			got, err := ds.ListDishes()
			if (err != nil) != tt.wantErr {
				t.Errorf("DishServiceImpl.ListDishes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DishServiceImpl.ListDishes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDishServiceImpl_UpdateDish(t *testing.T) {
	type fields struct {
		DishRepo    repositories.DishRepository
		RedisClient *redis.Client
	}
	type args struct {
		dish *models.Dish
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &dishservice.DishServiceImpl{
				DishRepo:    tt.fields.DishRepo,
				RedisClient: tt.fields.RedisClient,
			}
			if err := ds.UpdateDish(tt.args.dish); (err != nil) != tt.wantErr {
				t.Errorf("DishServiceImpl.UpdateDish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDishServiceImpl_SearchDishes(t *testing.T) {
	type fields struct {
		DishRepo    repositories.DishRepository
		RedisClient *redis.Client
	}
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Dish
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &dishservice.DishServiceImpl{
				DishRepo:    tt.fields.DishRepo,
				RedisClient: tt.fields.RedisClient,
			}
			got, err := ds.SearchDishes(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("DishServiceImpl.SearchDishes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DishServiceImpl.SearchDishes() = %v, want %v", got, tt.want)
			}
		})
	}
}
