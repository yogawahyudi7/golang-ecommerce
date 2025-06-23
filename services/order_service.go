package services

import (
	"golang-ecommerce/dto"
	"golang-ecommerce/messaging"
	"golang-ecommerce/models"
	"golang-ecommerce/repositories"

	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(input dto.CreateOrderRequest) (*dto.OrderResponse, error)
	ListByBuyer(buyerID string) ([]dto.OrderResponse, error)
}

type orderService struct {
	repo      repositories.OrderRepository
	publisher messaging.Publisher
}

func NewOrderService(repo repositories.OrderRepository, pub messaging.Publisher) OrderService {
	return &orderService{repo: repo, publisher: pub}
}

func (s *orderService) CreateOrder(input dto.CreateOrderRequest) (*dto.OrderResponse, error) {
	o := &models.Order{
		ID:         uuid.New(),
		BuyerID:    input.BuyerID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		TotalPrice: input.Price * float64(input.Quantity),
		Status:     "pending",
	}
	if err := s.repo.Create(o); err != nil {
		return nil, err
	}
	if err := s.publisher.Publish("order.created", o); err != nil {
		return nil, err
	}
	return &dto.OrderResponse{ID: o.ID, Status: o.Status, TotalPrice: o.TotalPrice}, nil
}

func (s *orderService) ListByBuyer(buyerID string) ([]dto.OrderResponse, error) {
	orders, err := s.repo.ListByBuyer(buyerID)
	if err != nil {
		return nil, err
	}
	res := make([]dto.OrderResponse, len(orders))
	for i, o := range orders {
		res[i] = dto.OrderResponse{ID: o.ID, Status: o.Status, TotalPrice: o.TotalPrice}
	}
	return res, nil
}
