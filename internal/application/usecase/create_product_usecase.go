package usecase

import (
	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/domain/entity"
	"github.com/lidiagaldino/desafio-backend/internal/domain/event"
	"github.com/lidiagaldino/desafio-backend/internal/domain/repository"
)

type CreateProductUsecase struct {
	productRepository repository.ProductRepository
	categoryRepository repository.CategoryRepository
	sendMessage event.SendMessage
	arn string
}

func NewCreateProductUsecase(productRepository repository.ProductRepository, sendMessage event.SendMessage, arn string, categoryRepository repository.CategoryRepository) *CreateProductUsecase {
	return &CreateProductUsecase{
		productRepository: productRepository,
		sendMessage: sendMessage,
		categoryRepository: categoryRepository,
		arn: arn,
	}
}

func (uc *CreateProductUsecase) Execute(input *dto.ProductInputDTO) (*dto.ProductOutputDTO, error) {
	_, err := uc.categoryRepository.FindByID(input.CategoryID)
	if err!= nil {
    return nil, err
  }

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
		err = uc.sendMessage.Publish(uc.arn, dto.OwnerID)
		if err!= nil {
	    return nil, err
	  }
  return &dto, nil
}
