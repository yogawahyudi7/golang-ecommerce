package services

import (
	"errors"
	"golang-ecommerce/dto"
	"golang-ecommerce/repositories"
)

type UserService interface {
	GetProfile(userID string) (*dto.UserResponse, error)
	ListUsers() ([]dto.UserResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) GetProfile(userID string) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email}, nil
}

func (s *userService) ListUsers() ([]dto.UserResponse, error) {
	// example implementation
	return nil, errors.New("not implemented")
}
