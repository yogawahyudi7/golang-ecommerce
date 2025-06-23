package repositories

import (
	"golang-ecommerce/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	FindByID(id string) (*models.Order, error)
	ListByBuyer(buyerID string) ([]models.Order, error)
	Update(order *models.Order) error
	Delete(id string) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) FindByID(id string) (*models.Order, error) {
	var o models.Order
	if err := r.db.First(&o, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *orderRepository) ListByBuyer(buyerID string) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Where("buyer_id = ?", buyerID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) Delete(id string) error {
	return r.db.Delete(&models.Order{}, "id = ?", id).Error
}
