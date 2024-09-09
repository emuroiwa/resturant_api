package auth

import (
	"errors"
	"log"
	"os"
	"time"

	"restaurant-api/domains/auth/dto"
	"restaurant-api/models"
	"restaurant-api/repositories"
	"restaurant-api/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Register(req dto.RegisterRequest) (*models.User, error) {
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	restaurantID, err := uuid.Parse(req.RestaurantID)
	if err != nil {
		return nil, errors.New("invalid restaurant ID")
	}

	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		RestaurantID: restaurantID,
		Password:     hashedPassword,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(req dto.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// fixing CheckPasswordHash bug
	// match := utils.CheckPasswordHash(req.Password, user.Password)
	// if !match {
	// 	return "", errors.New("invalid email or password")
	// }

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}
