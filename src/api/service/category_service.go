package service

import (
	"github.com/rodrigoherera/know-vegan-service/src/api/domain"
	"github.com/rodrigoherera/know-vegan-service/src/api/persistance/repository"
)

type ICategoryService interface {
	CreateCategory(category *domain.Category)
	GetAllCategories(offset, limit int64) (*[]domain.Category, int64)
}

type CategoryService struct {
	categoryRepository repository.ICategoryRepository
}

func NewCategoryService(categoryRepository repository.ICategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (cs *CategoryService) CreateCategory(category *domain.Category) {
	cs.categoryRepository.CreateCategory(category)
}

func (cs *CategoryService) GetAllCategories(offset, limit int64) (*[]domain.Category, int64) {
	return cs.categoryRepository.GetAllCategories(offset, limit)
}
