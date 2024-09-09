package main

import (
	"restaurant-api/config"
	"restaurant-api/controllers"
	"restaurant-api/domains/auth"
	"restaurant-api/domains/dishes"
	"restaurant-api/domains/dishes/dishservice"
	"restaurant-api/domains/ratings"
	"restaurant-api/domains/ratings/ratingservice"
	"restaurant-api/middlewares"
	"restaurant-api/utils"

	_ "restaurant-api/docs"
	"restaurant-api/migrations"
	"restaurant-api/repositories"
	"restaurant-api/routes"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type CustomValidator struct {
	validator *validator.Validate
}

// Validate function for CustomValidator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// @title Restaurant API
// @version 1.0
// @description This is a sample server for a restaurant management system.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file") // Stops the program if the .env file cannot be loaded
	}

	// Connect to DB
	db := config.SetupDB()

	// Run Migrations
	migrations.Migrate(db)

	// Initialize Redis
	redisClient := utils.InitRedis()

	// Set up Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	middlewares.InitPrometheusMetrics()
	e.Use(middlewares.MetricsMiddleware)

	// Initialize RabbitMQ
	utils.SetupRabbitMQ()
	defer utils.Channel.Close()

	// Start RabbitMQ Consumer
	go ratings.StartConsumer(db)
	go dishes.StartConsumer(db)

	// Initialize Repositories
	userRepo := repositories.NewUserRepository(db)
	dishRepo := repositories.NewDishRepository(db)
	ratingRepo := repositories.NewRatingRepository(db)

	// Initialize Services
	dishService := dishservice.NewDishService(dishRepo, redisClient)
	ratingService := ratingservice.NewRatingService(ratingRepo, redisClient)
	authService := auth.NewAuthService(userRepo)

	// Initialize Controllers
	authController := controllers.NewAuthController(authService)
	dishController := controllers.NewDishController(dishService)
	ratingController := controllers.NewRatingController(ratingService)

	// Register Routes
	routes.RegisterRoutes(e, db, dishController, authController, ratingController)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
