package usecase

import (
	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/domain/repository"
)

type FindProductByIDUsecase struct {
	productRepository repository.ProductRepository
}

func NewFindProductByIDUsecase(productRepository repository.ProductRepository) *FindProductByIDUsecase {
	return &FindProductByIDUsecase{
		productRepository: productRepository,
	}
}

func (uc *FindProductByIDUsecase) Execute(id string) (*dto.ProductOutputDTO, error) {
	product, err := uc.productRepository.FindByID(id)
  if err!= nil {
    return nil, err
  }

  dto := dto.ProductOutputDTO{
    ID:          product.ID,
    Title:       product.Title,
    Description: product.Description,
    Price:       product.Price,
    CategoryID:  product.CategoryID,
    OwnerID:     product.OwnerID,
  }
  return &dto, nil
}
