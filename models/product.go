package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SellerID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Name        string    `gorm:"size:150;not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null"`
	Stock       int       `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
