package usecase

import (
	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/domain/repository"
)

type FindAllCategoriesUsecase struct {
	categoryRepository repository.CategoryRepository
}

func NewFindAllCategoriesUsecase(categoryRepository repository.CategoryRepository) *FindAllCategoriesUsecase {
	return &FindAllCategoriesUsecase{
		categoryRepository: categoryRepository,
	}
}

func (uc *FindAllCategoriesUsecase) Execute() ([]*dto.CategoryOutputDTO, error) {
	categories, err := uc.categoryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var dtos []*dto.CategoryOutputDTO
	for _, category := range categories {
		dto := dto.CategoryOutputDTO{
			ID:          category.ID,
			Title:       category.Title,
			OwnerID:     category.OwnerID,
			Description: category.Description,
		}
		dtos = append(dtos, &dto)
	}
	return dtos, nil
}
