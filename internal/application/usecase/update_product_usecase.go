package usecase

import (
	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/domain/repository"
)

type UpdateProductUsecase struct {
	productRepository repository.ProductRepository
}

func NewUpdateProductUsecase(productRepository repository.ProductRepository) *UpdateProductUsecase {
	return &UpdateProductUsecase{
    productRepository: productRepository,
  }
}

func (uc *UpdateProductUsecase) Execute(input *dto.ProductInputDTO, id string) (*dto.ProductOutputDTO, error) {
	product, err := uc.productRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	product.Title = input.Title
	product.Description = input.Description
	product.Price = input.Price
	product.CategoryID = input.CategoryID
	product.OwnerID = input.OwnerID

	updatedProduct, err := uc.productRepository.Update(product)
	if err != nil {
		return nil, err
	}

	dto := dto.ProductOutputDTO{
		ID:          updatedProduct.ID,
		Title:       updatedProduct.Title,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		CategoryID:  updatedProduct.CategoryID,
		OwnerID:     updatedProduct.OwnerID,
	}
	return &dto, nil
}
