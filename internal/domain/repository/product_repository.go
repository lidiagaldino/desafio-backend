package repository

import "github.com/lidiagaldino/desafio-backend/internal/domain/entity" 

type ProductRepository interface {
	Save(product *entity.Product) (*entity.Product, error) 
	Update(product *entity.Product) (*entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
	Delete(id string) error
}
