package user

import "github.com/google/uuid"

// SignupRequest defines payload for user signup
type SignupRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=admin pedagang pembeli"`
}

// LoginRequest defines payload for user login
type LoginRequest struct {
	Email    string `json:"email"   validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse returns JWT token and basic user info
type AuthResponse struct {
	Token string    `json:"token"`
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Role  string    `json:"role"`
}

// UserResponse returns user profile data
type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
}
