// domain/user/repository.go
package user

import "gorm.io/gorm"

// Repository defines CRUD operations on User
type Repository interface {
	Create(u *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository returns a new User repository backed by GORM
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(u *User) error {
	return r.db.Create(u).Error
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var u User
	err := r.db.Where("email = ?", email).First(&u).Error
	return &u, err
}

func (r *repository) FindByID(id string) (*User, error) {
	var u User
	err := r.db.First(&u, "id = ?", id).Error
	return &u, err
}
