package usecase

import "github.com/lidiagaldino/desafio-backend/internal/domain/repository"

type DeleteProductUsecase struct {
	productRepository repository.ProductRepository
}

func NewDeleteProductUsecase(productRepository repository.ProductRepository) *DeleteProductUsecase {
	return &DeleteProductUsecase{
		productRepository: productRepository,
	}
}

func (c *DeleteProductUsecase) Execute(id string) error {
	_, err := c.productRepository.FindByID(id)
	if err!= nil {
    return err
  }
	err = c.productRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
