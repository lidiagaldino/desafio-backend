package usecase

import (
	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/domain/repository"
)

type FindCategoryByIDUsecase struct {
	categoryRepository repository.CategoryRepository
}

func NewFindCategoryByIDUsecase(categoryRepository repository.CategoryRepository) *FindCategoryByIDUsecase {
	return &FindCategoryByIDUsecase{
		categoryRepository: categoryRepository,
	}
}

func (uc *FindCategoryByIDUsecase) Execute(id string) (*dto.CategoryOutputDTO, error) {
	category, err := uc.categoryRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	dto := dto.CategoryOutputDTO{
		ID:          category.ID,
		Title:       category.Title,
		OwnerID:     category.OwnerID,
		Description: category.Description,
	}
	return &dto, nil
}
