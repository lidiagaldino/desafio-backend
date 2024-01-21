package usecase

type CategoryUsecases struct {
	FindCategoryByIDUsecase  *FindCategoryByIDUsecase
	FindAllCategoriesUsecase *FindAllCategoriesUsecase
	CreateCategoryUsecase    *CreateCategoryUsecase
	UpdateCategoryUsecase    *UpdateCategoryUsecase
	DeleteCategoryUsecase    *DeleteCategoryUsecase
}

func NewCategoryUsecases(
	findCategoryByIDUsecase FindCategoryByIDUsecase,
	findAllCategoriesUsecase FindAllCategoriesUsecase,
	createCategoryUsecase CreateCategoryUsecase,
	updateCategoryUsecase UpdateCategoryUsecase,
	deleteCategoryUsecase DeleteCategoryUsecase,
) *CategoryUsecases {
	return &CategoryUsecases{
		FindCategoryByIDUsecase:  &findCategoryByIDUsecase,
		FindAllCategoriesUsecase: &findAllCategoriesUsecase,
		CreateCategoryUsecase:    &createCategoryUsecase,
		UpdateCategoryUsecase:    &updateCategoryUsecase,
		DeleteCategoryUsecase:    &deleteCategoryUsecase,
	}
}
