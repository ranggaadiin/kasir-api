package repositories

import (
	"fmt"
	"kasir-api/models"
)

type CategoryRepository struct {
	data []models.Category
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		data: []models.Category{
			{ID: 1, Nama: "Makanan", Description: "Kategori untuk makanan"},
			{ID: 2, Nama: "Minuman", Description: "Kategori untuk minuman"},
		},
	}
}

func (r *CategoryRepository) GetAll() []models.Category {
	return r.data
}

func (r *CategoryRepository) GetByID(id int) (*models.Category, error) {
	for _, c := range r.data {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("category not found")
}

func (r *CategoryRepository) Create(category models.Category) models.Category {
	category.ID = len(r.data) + 1
	r.data = append(r.data, category)
	return category
}
