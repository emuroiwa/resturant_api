package routes

import (
	"restaurant-api/controllers"
	"restaurant-api/middlewares"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, dishController *controllers.DishController, authController *controllers.AuthController, ratingController *controllers.RatingController) {
	// Auth Routes
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)

	// Protected Routes
	dishGroup := e.Group("/dishes", middlewares.AuthMiddleware())
	{
		dishGroup.POST("", dishController.Create)
		dishGroup.GET("/:id", dishController.Show)
		dishGroup.GET("", dishController.Index)
		dishGroup.PUT("/:id", dishController.Update)
		dishGroup.DELETE("/:id", dishController.Delete)
		dishGroup.GET("/search", dishController.Search)

		dishGroup.POST("/:dish_id/ratings", ratingController.Create)
		dishGroup.GET("/:dish_id/ratings", ratingController.Show)
	}
	e.GET("/metrics", middlewares.PrometheusHandler())
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
