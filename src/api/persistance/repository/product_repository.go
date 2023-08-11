package repository

import (
	"github.com/rodrigoherera/know-vegan-service/src/api/domain"
	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(product *domain.Product, photo *domain.Photo) error
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) CreateProduct(product *domain.Product, photo *domain.Photo) error {
	return pr.db.Create(&product).Error
}
