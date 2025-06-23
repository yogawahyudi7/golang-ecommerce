package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is the GORM model for a user record
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100;unique;not null"`
	Password  string    `gorm:"size:255;not null"`
	Role      string    `gorm:"size:20;not null"` // admin, pedagang, pembeli
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
