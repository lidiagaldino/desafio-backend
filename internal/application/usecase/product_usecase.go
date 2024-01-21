package usecase

type ProductUsecases struct {
	FindProductByIDUsecase *FindProductByIDUsecase
	FindAllProductsUsecase *FindAllProductsUsecase
	CreateProductUsecase   *CreateProductUsecase
	UpdateProductUsecase   *UpdateProductUsecase
	DeleteProductUsecase   *DeleteProductUsecase
}

func NewProductUsecases(
	findProductByIDUsecase FindProductByIDUsecase,
	findAllProductsUsecase FindAllProductsUsecase,
	createProductUsecase CreateProductUsecase,
	updateProductUsecase UpdateProductUsecase,
	deleteProductUsecase DeleteProductUsecase,
) *ProductUsecases {
	return &ProductUsecases{
		FindProductByIDUsecase: &findProductByIDUsecase,
		FindAllProductsUsecase: &findAllProductsUsecase,
		CreateProductUsecase:   &createProductUsecase,
		UpdateProductUsecase:   &updateProductUsecase,
		DeleteProductUsecase:   &deleteProductUsecase,
	}
}
