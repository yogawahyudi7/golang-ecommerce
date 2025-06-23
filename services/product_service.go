package services

import (
	"golang-ecommerce/dto"
	"golang-ecommerce/models"
	"golang-ecommerce/repositories"

	"github.com/google/uuid"
)

type ProductService interface {
	Create(input dto.CreateProductRequest) (*dto.ProductResponse, error)
	ListBySeller(sellerID string) ([]dto.ProductResponse, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(input dto.CreateProductRequest) (*dto.ProductResponse, error) {
	p := &models.Product{
		ID:          uuid.New(),
		SellerID:    input.SellerID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}
	if err := s.repo.Create(p); err != nil {
		return nil, err
	}
	return &dto.ProductResponse{ID: p.ID, Name: p.Name, Price: p.Price, Stock: p.Stock}, nil
}

func (s *productService) ListBySeller(sellerID string) ([]dto.ProductResponse, error) {
	items, err := s.repo.ListBySeller(sellerID)
	if err != nil {
		return nil, err
	}
	var res []dto.ProductResponse
	for _, p := range items {
		res = append(res, dto.ProductResponse{ID: p.ID, Name: p.Name, Price: p.Price, Stock: p.Stock})
	}
	return res, nil
}
