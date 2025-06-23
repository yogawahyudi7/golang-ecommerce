package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BuyerID    uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID  uuid.UUID `gorm:"type:uuid;not null;index"`
	Quantity   int       `gorm:"not null"`
	TotalPrice float64   `gorm:"not null"`
	Status     string    `gorm:"size:20;not null"` // pending, paid, shipped, delivered
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
