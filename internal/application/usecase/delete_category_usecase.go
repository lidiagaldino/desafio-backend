package usecase

import "github.com/lidiagaldino/desafio-backend/internal/domain/repository"

type DeleteCategoryUsecase struct {
	categoryRepository repository.CategoryRepository
}

func NewDeleteCategoryUsecase(categoryRepository repository.CategoryRepository) *DeleteCategoryUsecase {
	return &DeleteCategoryUsecase{
		categoryRepository: categoryRepository,
	}
}

func (c *DeleteCategoryUsecase) Execute(id string) error {
	_, err := c.categoryRepository.FindByID(id)
	if err != nil {
		return err
	}
	err = c.categoryRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
