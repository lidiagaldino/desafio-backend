package usecase

import (
	"encoding/json"

	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/domain/entity"
	"github.com/lidiagaldino/desafio-backend/internal/domain/event"
	"github.com/lidiagaldino/desafio-backend/internal/domain/repository"
)

type CreateCategoryUsecase struct {
	categoryRepository repository.CategoryRepository
	sendMessage event.SendMessage
	arn string
}

func NewCreateCategoryUsecase(categoryRepository repository.CategoryRepository, sendMessage event.SendMessage, arn string) *CreateCategoryUsecase {
	return &CreateCategoryUsecase{
		categoryRepository: categoryRepository,
		sendMessage: sendMessage,
    arn: arn,
	}
}

func (uc *CreateCategoryUsecase) Execute(input *dto.CategoryInputDTO) (*dto.CategoryOutputDTO, error) {
	category := entity.Category{
		Title: input.Title,
		OwnerID: input.OwnerID,
		Description: input.Description,
	}
	createdCategory, err := uc.categoryRepository.Save(&category)
	if err != nil {
		return nil, err
	}

	dto := dto.CategoryOutputDTO{
		ID:          createdCategory.ID,
		Title:       createdCategory.Title,
		OwnerID:     createdCategory.OwnerID,
		Description: createdCategory.Description,
	}
	
	dtoBytes, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	err = uc.sendMessage.Publish(uc.arn, string(dtoBytes))
	if err != nil {
		return nil, err
	}
	return &dto, nil
}
