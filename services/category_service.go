package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() []models.Category {
	return s.repo.GetAll()
}

func (s *CategoryService) GetCategoryByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) CreateCategory(category models.Category) models.Category {
	return s.repo.Create(category)
}
