package controllers

import (
	"context"
	"net/http"
	"restaurant-api/domains/ratings/ratingservice"
	"restaurant-api/models"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RatingController struct {
	RatingService *ratingservice.RatingService
}

func NewRatingController(ratingService *ratingservice.RatingService) *RatingController {
	return &RatingController{
		RatingService: ratingService,
	}
}

// Create
// @Summary Create a new rating for a dish
// @Description Submit a rating for a specific dish
// @Tags Ratings
// @Accept json
// @Produce json
// @Param dish_id path string true "Dish ID"
// @Param rating body models.Rating true "Rating data"
// @Security ApiKeyAuth
// @Success 201 {string} string "Rating submitted successfully, it will be processed shortly."
// @Failure 400 {string} string "Invalid input or dish ID"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /dishes/{dish_id}/ratings [post]
func (rc *RatingController) Create(c echo.Context) error {
	var rating models.Rating
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, ok := claims["id"].(string)

	if !ok {
		return c.JSON(http.StatusUnauthorized, "User ID not found in token")
	}
	if err := c.Bind(&rating); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	dishID, err := uuid.Parse(c.Param("dish_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid dish ID")
	}

	err = rc.RatingService.Create(&rating, userID, dishID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "Rating submitted successfully, it will be processed shortly.")
}

// Show
// @Summary Get ratings for a dish
// @Description Retrieve all ratings for a specific dish by its UUID
// @Tags Ratings
// @Produce json
// @Param dish_id path string true "Dish ID"
// @Success 200 {array} models.Rating
// @Failure 400 {string} string "Invalid dish ID format"
// @Failure 500 {string} string "Internal server error"
// @Router /dishes/{dish_id}/ratings [get]
func (rc *RatingController) Show(c echo.Context) error {
	ctx := context.Background()
	dishID, err := uuid.Parse(c.Param("dish_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid dish ID format")
	}

	response, err := rc.RatingService.GetRatingsByDishID(ctx, dishID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
