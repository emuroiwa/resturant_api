package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rating struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Score  int       `gorm:"not null" validate:"required,min=1,max=5"`
	Review string    `gorm:"type:text" validate:"omitempty,max=500"`
	DishID uuid.UUID `gorm:"type:uuid;not null" validate:"required"`
	UserID uuid.UUID `gorm:"type:uuid;not null" validate:"required"`
}

// BeforeCreate hook to set ID before saving to the database
func (rating *Rating) BeforeCreate(tx *gorm.DB) (err error) {
	rating.ID = uuid.New()
	return
}
