package controllers

import (
	"context"
	"net/http"
	dishservice "restaurant-api/domains/dishes/dishservice"
	"restaurant-api/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DishController struct {
	DishService dishservice.DishService
}

func NewDishController(dishService dishservice.DishService) *DishController {
	return &DishController{
		DishService: dishService,
	}
}

// Create
// @Summary Create a new dish
// @Description Submit a new dish that will be processed
// @Tags Dishes
// @Accept json
// @Produce json
// @Param dish body models.Dish true "Dish data"
// @Success 201 {string} string "Dish submitted successfully, it will be processed shortly."
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /dishes [post]
func (controller *DishController) Create(c echo.Context) error {
	dish := new(models.Dish)
	if err := c.Bind(dish); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := controller.DishService.CreateDish(dish); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Dish submitted successfully, it will be processed shortly.")
}

// Show
// @Summary Get a dish by ID
// @Description Retrieve the details of a specific dish by its UUID
// @Tags Dishes
// @Produce json
// @Param id path string true "Dish ID"
// @Success 200 {object} models.Dish
// @Failure 400 {string} string "Invalid UUID format"
// @Failure 404 {string} string "Dish not found"
// @Router /dishes/{id} [get]
func (controller *DishController) Show(c echo.Context) error {
	ctx := context.Background()

	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID format")
	}

	dish, err := controller.DishService.GetDish(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, dish)
}

// Index
// @Summary List all dishes
// @Description Retrieve a list of all available dishes
// @Tags Dishes
// @Produce json
// @Success 200 {array} models.Dish
// @Failure 500 {string} string "Internal server error"
// @Router /dishes [get]
func (controller *DishController) Index(c echo.Context) error {
	dishes, err := controller.DishService.ListDishes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dishes)
}

// Update
// @Summary Update a dish
// @Description Update the details of an existing dish
// @Tags Dishes
// @Accept json
// @Produce json
// @Param dish body models.Dish true "Dish data"
// @Success 204 {object} models.Dish
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /dishes [put]
func (controller *DishController) Update(c echo.Context) error {
	dish := new(models.Dish)
	if err := c.Bind(dish); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := controller.DishService.UpdateDish(dish); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, dish)
}

// Search
// @Summary Search for dishes
// @Description Search for dishes by name or other criteria
// @Tags Dishes
// @Produce json
// @Param query query string true "Search query"
// @Success 200 {array} models.Dish
// @Failure 400 {string} string "Query parameter is required"
// @Failure 500 {string} string "Internal server error"
// @Router /dishes/search [get]
func (controller *DishController) Search(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, "Query parameter is required")
	}

	dishes, err := controller.DishService.SearchDishes(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dishes)
}

// Delete
// @Summary Delete a dish by ID
// @Description Remove a specific dish by its UUID
// @Tags Dishes
// @Param id path string true "Dish ID"
// @Success 204 {string} string "No content"
// @Failure 400 {string} string "Invalid UUID format"
// @Failure 404 {string} string "Dish not found"
// @Router /dishes/{id} [delete]
func (controller *DishController) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID format")
	}

	res := controller.DishService.DeleteDishes(id)
	if res != nil {
		return c.JSON(http.StatusNotFound, res.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
