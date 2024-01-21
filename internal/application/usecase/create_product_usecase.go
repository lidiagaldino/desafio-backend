package usecase

import (
	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/domain/entity"
	"github.com/lidiagaldino/desafio-backend/internal/domain/repository"
)

type CreateProductUsecase struct {
	productRepository repository.ProductRepository
}

func NewCreateProductUsecase(productRepository repository.ProductRepository) *CreateProductUsecase {
	return &CreateProductUsecase{
		productRepository: productRepository,
	}
}

func (uc *CreateProductUsecase) Execute(input *dto.ProductInputDTO) (*dto.ProductOutputDTO, error) {
	product := entity.Product{
		Title:       input.Title,
		Price:       input.Price,
		Description: input.Description,
		OwnerID:     input.OwnerID,
		CategoryID:  input.CategoryID,
	}
	createdProduct, err := uc.productRepository.Save(&product)
  if err!= nil {
    return nil, err
  }

	dto := dto.ProductOutputDTO{
		ID:          createdProduct.ID,
    Title:       createdProduct.Title,
    Description: createdProduct.Description,
    Price:       createdProduct.Price,
    CategoryID:  createdProduct.CategoryID,
		OwnerID:     createdProduct.OwnerID,
	}
  return &dto, nil
}