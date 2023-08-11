package repository

import (
	"github.com/rodrigoherera/know-vegan-service/src/api/domain"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	CreateCategory(category *domain.Category)
	GetAllCategories(offset, limit int64) (*[]domain.Category, int64)
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (cr *CategoryRepository) CreateCategory(category *domain.Category) {
	cr.db.Create(&category)
}

func (cr *CategoryRepository) GetAllCategories(offset, limit int64) (*[]domain.Category, int64) {
	var (
		category *[]domain.Category
		total    int64
	)

	if err := cr.db.Model(&domain.Category{}).
		Count(&total).Error; err != nil {
		if err != nil {
			println(err.Error())
			return nil, 0
		}
	}

	err := cr.db.Preload("Products.Tags").
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&category).Error
	if err != nil {
		println(err.Error())
		return nil, 0
	}
	return category, total
}
