package usecase

import (
	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/domain/event"
	"github.com/lidiagaldino/desafio-backend/internal/domain/repository"
)

type UpdateProductUsecase struct {
	productRepository repository.ProductRepository
	categoryRepository repository.CategoryRepository
	sendMessage event.SendMessage
	arn string
}

func NewUpdateProductUsecase(productRepository repository.ProductRepository, sendMessage event.SendMessage, arn string, categoryRepository repository.CategoryRepository) *UpdateProductUsecase {
	return &UpdateProductUsecase{
    productRepository: productRepository,
		sendMessage: sendMessage,
    categoryRepository: categoryRepository,
    arn: arn,
  }
}

func (uc *UpdateProductUsecase) Execute(input *dto.ProductInputDTO, id string) (*dto.ProductOutputDTO, error) {
	_,err := uc.categoryRepository.FindByID(input.CategoryID)
	if err!= nil {
    return nil, err
  }
	
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
	err = uc.sendMessage.Publish(uc.arn, dto.OwnerID)
	if err!= nil {
    return nil, err
  }
	return &dto, nil
}
