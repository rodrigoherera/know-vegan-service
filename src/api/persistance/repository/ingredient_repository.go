package repository

import (
	"github.com/rodrigoherera/know-vegan-service/src/api/domain"
	"gorm.io/gorm"
)

type IIngredientRepository interface {
	CreateIngredient(ingredient *domain.Ingredient) error
	GetAll(offset, limit int64) (*[]domain.Ingredient, int64)
}

type IngredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) *IngredientRepository {
	return &IngredientRepository{
		db: db,
	}
}

func (ir *IngredientRepository) CreateIngredient(ingredient *domain.Ingredient) error {
	return ir.db.Create(&ingredient).Error
}

func (ir *IngredientRepository) GetAll(offset, limit int64) (*[]domain.Ingredient, int64) {
	var (
		ingredient []domain.Ingredient
		total      int64
	)

	if err := ir.db.Model(&domain.Ingredient{}).
		Count(&total).Error; err != nil {
		if err != nil {
			print(err.Error())
			return nil, 0
		}
	}

	err := ir.db.Preload("Tags").
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&ingredient).Error

	if err != nil {
		println(err.Error())
		return nil, 0
	}

	return &ingredient, total
}
