package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

var ctx = context.Background()

func CacheMiddleware(redisClient *redis.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cacheKey := c.Request().RequestURI

			cachedResponse, err := redisClient.Get(ctx, cacheKey).Result()
			if err == redis.Nil {
				return next(c)
			} else if err != nil {
				log.Printf("Failed to get cached response: %v", err)
				return next(c)
			}
			return c.JSONBlob(http.StatusOK, []byte(cachedResponse))
		}
	}
}

func CacheResponse(redisClient *redis.Client, cacheKey string, responseBody []byte) {
	err := redisClient.Set(ctx, cacheKey, responseBody, time.Minute*10).Err()
	if err != nil {
		log.Printf("Failed to cache response: %v", err)
	}
}
