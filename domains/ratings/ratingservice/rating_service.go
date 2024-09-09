package ratingservice

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

type RatingService struct {
	RatingRepo  *repositories.RatingRepository
	RedisClient *redis.Client
}

const (
	MinRatingScore = 1
	MaxRatingScore = 5
)

func NewRatingService(ratingRepo *repositories.RatingRepository, redisClient *redis.Client) *RatingService {
	return &RatingService{
		RatingRepo:  ratingRepo,
		RedisClient: redisClient,
	}
}

func (rs *RatingService) Create(rating *models.Rating, userID string, dishID uuid.UUID) error {
	if rating.Score < MinRatingScore || rating.Score > MaxRatingScore {
		return errors.New("rating score must be between 1 and 5")
	}

	rating.ID = uuid.New()
	rating.DishID = dishID
	rating.UserID = uuid.MustParse(userID)

	ratingJSON, err := json.Marshal(rating)
	if err != nil {
		return errors.New("failed to serialize rating")
	}

	err = rabbitmq.PublishMessage(string(ratingJSON), "rating_queue")
	if err != nil {
		return errors.New("failed to publish rating to queue")
	}

	return nil
}

func (rs *RatingService) GetRatingsByDishID(ctx context.Context, dishID uuid.UUID) (map[string]interface{}, error) {
	cacheKey := "rating_" + dishID.String()
	cachedResponse, err := rs.RedisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// Cache miss, fetch from DB
		ratings, err := rs.RatingRepo.GetRatingsByDishID(dishID)
		if err != nil {
			return nil, errors.New("ratings not found")
		}

		var totalRating int
		var ratingCount int
		for _, rating := range ratings {
			totalRating += rating.Score
			ratingCount++
		}

		avgRating := 0.0
		if ratingCount > 0 {
			avgRating = float64(totalRating) / float64(ratingCount)
		}

		response := map[string]interface{}{
			"average_rating": avgRating,
			"total_ratings":  ratingCount,
			"ratings":        ratings,
		}

		responseBody, err := json.Marshal(response)
		if err != nil {
			return nil, errors.New("failed to serialize response")
		}

		// Cache the response
		err = rs.RedisClient.Set(ctx, cacheKey, responseBody, time.Minute*10).Err()
		if err != nil {
			return nil, errors.New("failed to cache response")
		}

		return response, nil
	} else if err != nil {
		return nil, errors.New("failed to get cached response")
	}

	// Cache hit
	var response map[string]interface{}
	err = json.Unmarshal([]byte(cachedResponse), &response)
	if err != nil {
		return nil, errors.New("failed to unmarshal cached response")
	}

	return response, nil
}
