package repositories

import (
	"restaurant-api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RatingRepository struct {
	db *gorm.DB
}

// NewRatingRepository creates a new instance of the RatingRepository
func NewRatingRepository(db *gorm.DB) *RatingRepository {
	return &RatingRepository{db: db}
}

// GetRatingsByDishID retrieves all ratings for a given dish
func (repo *RatingRepository) GetRatingsByDishID(dishID uuid.UUID) ([]*models.Rating, error) {
	var ratings []*models.Rating
	err := repo.db.Where("dish_id = ?", dishID).Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

// Create adds a new rating for a dish
func (repo *RatingRepository) Create(rating *models.Rating) error {
	return repo.db.Create(&rating).Error
}

// UpdateRating updates an existing rating
func (repo *RatingRepository) UpdateRating(rating *models.Rating) error {
	return repo.db.Save(&rating).Error
}

// DeleteRating deletes a rating by its ID
func (repo *RatingRepository) DeleteRating(ratingID uuid.UUID) error {
	return repo.db.Delete(&models.Rating{}, "id = ?", ratingID).Error
}
