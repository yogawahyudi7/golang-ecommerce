package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"golang-ecommerce/dto"
	"golang-ecommerce/models"
	"golang-ecommerce/repositories"
	"golang-ecommerce/utils"
)

type AuthService interface {
	Signup(input dto.SignupRequest) (*dto.AuthResponse, error)
	Login(input dto.LoginRequest) (*dto.AuthResponse, error)
	Logout(tokenStr string) error
}

type authService struct {
	userRepo repositories.UserRepository
	jwtKey   []byte
	ttl      time.Duration
}

func NewAuthService(repo repositories.UserRepository, jwtKey []byte, ttl time.Duration) AuthService {
	return &authService{userRepo: repo, jwtKey: jwtKey, ttl: ttl}
}

func (s *authService) Signup(input dto.SignupRequest) (*dto.AuthResponse, error) {
	user := &models.User{ID: uuid.New(), Name: input.Name, Email: input.Email, Password: HashPassword(input.Password), Role: input.Role}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	token, jti, err := s.createToken(user)
	if err != nil {
		return nil, err
	}
	return &dto.AuthResponse{Token: token, Jti: jti}, nil
}

func (s *authService) Login(input dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if !CheckPasswordHash(input.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	token, jti, err := s.createToken(user)
	if err != nil {
		return nil, err
	}
	return &dto.AuthResponse{Token: token, Jti: jti}, nil
}

func (s *authService) Logout(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) { return s.jwtKey, nil })
	if err != nil || !token.Valid {
		return errors.New("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	jti, ok := claims["jti"].(string)
	expUnix := int64(claims["exp"].(float64))
	if !ok {
		return errors.New("missing jti")
	}
	ttl := time.Until(time.Unix(expUnix, 0))
	return utils.BlacklistToken(jti, ttl)
}

func (s *authService) createToken(user *models.User) (string, string, error) {
	jti := uuid.NewString()
	claims := jwt.RegisteredClaims{ID: jti, Subject: user.ID.String(), ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.ttl))}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.jwtKey)
	return signed, jti, err
}
