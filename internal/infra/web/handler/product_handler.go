package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/application/usecase"
	"github.com/lidiagaldino/desafio-backend/internal/infra/web/utils"
)

type ProductHandler struct {
	pu *usecase.ProductUsecases
}

func NewProductHandler(pu *usecase.ProductUsecases) *ProductHandler {
	return &ProductHandler{
		pu: pu,
	}
}

func (h *ProductHandler) FindProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := h.pu.FindProductByIDUsecase.Execute(id)
	if err!= nil {
		utils.SendError(ctx, 404, "Product not found")
	}
	utils.SendSuccess(ctx, "find-product-by-id", product, 200)
}

func (h *ProductHandler) FindAllProducts(ctx *gin.Context) {
	products, err := h.pu.FindAllProductsUsecase.Execute()
	if err!= nil {
    utils.SendError(ctx, 404, "Products not found")
  }
	utils.SendSuccess(ctx, "find-all-productts", products, 200)
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	product := dto.ProductInputDTO{}
	err := ctx.BindJSON(&product)
	if err != nil {
		utils.SendError(ctx, 400, err.Error())
	}
	err = product.Validate()
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newProduct, err := h.pu.CreateProductUsecase.Execute(&product)
	if err != nil {
		utils.SendError(ctx, 500, "Error creating product: "+err.Error())
	}
  utils.SendSuccess(ctx, "create-product", newProduct, 201)
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	product := dto.ProductInputDTO{}
	id := ctx.Param("id")
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		utils.SendError(ctx, 400, "Invalid request body")
	}
	err = product.Validate()
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	newProduct, err := h.pu.UpdateProductUsecase.Execute(&product, id)
	log.Println(newProduct)
	if err != nil {
		utils.SendError(ctx, 500, "Error updating product"+err.Error())
	}
  utils.SendSuccess(ctx, "update-product", newProduct, 200)
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
  err := h.pu.DeleteProductUsecase.Execute(id)
  if err!= nil {
    utils.SendError(ctx, 500, "Error deleting product" + err.Error())
  }
  utils.SendSuccess(ctx, "delete-product", nil, 200)
}

