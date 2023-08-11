package service

import (
	"github.com/rodrigoherera/know-vegan-service/src/api/domain"
	"github.com/rodrigoherera/know-vegan-service/src/api/persistance/repository"
)

type IIngredientService interface {
	CreateIngredient(ingredient *domain.Ingredient) error
	GetAll(offset, limit int64) (*[]domain.Ingredient, int64)
}

type IngredientService struct {
	ingredientRepository repository.IIngredientRepository
}

func NewIngredientService(ingredientRepository repository.IIngredientRepository) *IngredientService {
	return &IngredientService{
		ingredientRepository: ingredientRepository,
	}
}

func (is *IngredientService) CreateIngredient(ingredient *domain.Ingredient) error {
	return is.ingredientRepository.CreateIngredient(ingredient)
}

func (is *IngredientService) GetAll(offset, limit int64) (*[]domain.Ingredient, int64) {
	return is.ingredientRepository.GetAll(offset, limit)
}
