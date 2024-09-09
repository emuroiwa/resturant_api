package controllers

import (
	"net/http"
	"restaurant-api/domains/auth"
	"restaurant-api/domains/auth/dto"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService auth.AuthService
}

var validate = validator.New()

func NewAuthController(authService auth.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register handles user registration
func (a *AuthController) Register(c echo.Context) error {
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := a.authService.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

// Login handles user login
func (a *AuthController) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed"})
	}

	token, err := a.authService.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
