package usecase

import (
	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/domain/repository"
)

type FindAllProductsUsecase struct {
	productRepository repository.ProductRepository
}

func NewFindAllProductsUsecase(productRepository repository.ProductRepository) *FindAllProductsUsecase {
	return &FindAllProductsUsecase{
		productRepository: productRepository,
	}
}

func (uc *FindAllProductsUsecase) Execute() ([]*dto.ProductOutputDTO, error) {
	products, err := uc.productRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var dtos []*dto.ProductOutputDTO
	for _, product := range products {
		dtos = append(dtos, &dto.ProductOutputDTO{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
			CategoryID:  product.CategoryID,
			OwnerID:     product.OwnerID,
		})
	}
	return dtos, nil
}
