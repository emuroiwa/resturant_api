package tests

import (
	"context"
	"reflect"
	"restaurant-api/domains/ratings/ratingservice"
	"restaurant-api/models"
	"restaurant-api/repositories"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func TestRatingService_GetRatingsByDishID(t *testing.T) {
	type fields struct {
		RatingRepo  *repositories.RatingRepository
		RedisClient *redis.Client
	}
	type args struct {
		ctx    context.Context
		dishID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ratingservice.RatingService{
				RatingRepo:  tt.fields.RatingRepo,
				RedisClient: tt.fields.RedisClient,
			}
			got, err := rs.GetRatingsByDishID(tt.args.ctx, tt.args.dishID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RatingService.GetRatingsByDishID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RatingService.GetRatingsByDishID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRatingService_Create(t *testing.T) {
	type fields struct {
		RatingRepo  *repositories.RatingRepository
		RedisClient *redis.Client
	}
	type args struct {
		rating *models.Rating
		userID string
		dishID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ratingservice.RatingService{
				RatingRepo:  tt.fields.RatingRepo,
				RedisClient: tt.fields.RedisClient,
			}
			if err := rs.Create(tt.args.rating, tt.args.userID, tt.args.dishID); (err != nil) != tt.wantErr {
				t.Errorf("RatingService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
