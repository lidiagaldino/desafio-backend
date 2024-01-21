package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/desafio-backend/internal/application/dto"
	"github.com/lidiagaldino/desafio-backend/internal/application/usecase"
	"github.com/lidiagaldino/desafio-backend/internal/infra/web/utils"
)

type CategoryHandler struct {
	cu *usecase.CategoryUsecases
}

func NewCategoryHandler(cu *usecase.CategoryUsecases) *CategoryHandler {
	return &CategoryHandler{
		cu: cu,
	}
}

func (h *CategoryHandler) FindCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
  category, err := h.cu.FindCategoryByIDUsecase.Execute(id)
  if err!= nil {
    utils.SendError(ctx, 404, "Category not found")
  }
  utils.SendSuccess(ctx, "find-category-by-id", category, 200)
}

func (h *CategoryHandler) FindAllCategories(ctx *gin.Context) {
	categories, err := h.cu.FindAllCategoriesUsecase.Execute()
  if err!= nil {
    utils.SendError(ctx, 404, "Categories not found")
  }
  utils.SendSuccess(ctx, "find-all-categories", categories, 200)
}


func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var input dto.CategoryInputDTO
if err := ctx.ShouldBindJSON(&input); err != nil {
	utils.SendError(ctx, 400, err.Error())
	return
}
err := input.Validate()
if err != nil {
	utils.SendError(ctx, http.StatusBadRequest, err.Error())
	return
}
category, err := h.cu.CreateCategoryUsecase.Execute(&input)
if err != nil {
	utils.SendError(ctx, 400, err.Error())
	return
}
  utils.SendSuccess(ctx, "create-category", category, 201)
}

func (h *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	var input dto.CategoryInputDTO
if err := ctx.ShouldBindJSON(&input); err != nil {
	utils.SendError(ctx, 400, err.Error())
	return
}
err := input.Validate()
if err != nil {
	utils.SendError(ctx, http.StatusBadRequest, err.Error()) // Fix: Add a comma after the closing parenthesis
	return
}
id := ctx.Param("id")
category, err := h.cu.UpdateCategoryUsecase.Execute(&input, id)
if err != nil {
	utils.SendError(ctx, 400, err.Error())
	return
}
  utils.SendSuccess(ctx, "update-category", category, 200)
}

func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
  err := h.cu.DeleteCategoryUsecase.Execute(id)
  if err!= nil {
    utils.SendError(ctx, 400, err.Error())
    return
  }
  utils.SendSuccess(ctx, "delete-category", nil, 200)
}
