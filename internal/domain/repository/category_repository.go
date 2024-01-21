package repository

import "github.com/lidiagaldino/desafio-backend/internal/domain/entity"

type CategoryRepository interface {
	Save(category *entity.Category) (*entity.Category, error)
	Update(category *entity.Category) (*entity.Category, error)
	FindByID(id string) (*entity.Category, error)
	FindAll() ([]entity.Category, error)
	Delete(id string) error
}
