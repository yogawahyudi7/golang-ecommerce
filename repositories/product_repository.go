package repositories

import (
	"golang-ecommerce/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindByID(id string) (*models.Product, error)
	ListBySeller(sellerID string) ([]models.Product, error)
	Update(product *models.Product) error
	Delete(id string) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindByID(id string) (*models.Product, error) {
	var p models.Product
	if err := r.db.First(&p, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepository) ListBySeller(sellerID string) ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Where("seller_id = ?", sellerID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id string) error {
	return r.db.Delete(&models.Product{}, "id = ?", id).Error
}
