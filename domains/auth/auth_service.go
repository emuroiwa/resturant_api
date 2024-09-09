package auth

import (
	"restaurant-api/domains/auth/dto"
	"restaurant-api/models"
)

type AuthService interface {
	Register(req dto.RegisterRequest) (*models.User, error)
	Login(req dto.LoginRequest) (string, error)
}
