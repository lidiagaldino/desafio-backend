package usecase

import (
	"fmt"

	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/domain/event"
	"github.com/lidiagaldino/desafio-backend/internal/domain/repository"
)

type UpdateCategoryUsecase struct {
	categoryRepository repository.CategoryRepository
	sendMessage event.SendMessage
	arn string
}

func NewUpdateCategoryUsecase(categoryRepository repository.CategoryRepository, sendMessage event.SendMessage, arn string) *UpdateCategoryUsecase {
	return &UpdateCategoryUsecase{
		categoryRepository: categoryRepository,
		sendMessage: sendMessage,
    arn: arn,
	}
}

func (uc *UpdateCategoryUsecase) Execute(input *dto.CategoryInputDTO, id string) (*dto.CategoryOutputDTO, error) {
	category, err := uc.categoryRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	category.Title = input.Title
	category.OwnerID = input.OwnerID
	category.Description = input.Description

	updatedCategory, err := uc.categoryRepository.Update(category)
	if err != nil {
		return nil, err
	}

	dto := dto.CategoryOutputDTO{
		ID:          updatedCategory.ID,
		Title:       updatedCategory.Title,
		OwnerID:     updatedCategory.OwnerID,
		Description: updatedCategory.Description,
	}
	err = uc.sendMessage.Publish(uc.arn, fmt.Sprintf("%v", dto))
	if err!= nil {
    return nil, err
  }
	return &dto, nil
}
