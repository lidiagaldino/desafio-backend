package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/desafio-backend/internal/application/usecase"
	"github.com/lidiagaldino/desafio-backend/internal/infra/web/handler"
)
func initializeRoutes(router *gin.Engine, pu *usecase.ProductUsecases, cu *usecase.CategoryUsecases) {
	productHandler := handler.NewProductHandler(pu)
	categoryHandler := handler.NewCategoryHandler(cu)
	basePath := "api/v1/"
	v1 := router.Group(basePath)
	{
		v1.GET("/", handler.HelloWorld)

		//Products routes
		v1.GET("/products", productHandler.FindAllProducts)
		v1.GET("/products/:id", productHandler.FindProductByID)
		v1.POST("/products", productHandler.CreateProduct)
		v1.PUT("/products/:id", productHandler.UpdateProduct)
		v1.DELETE("/products/:id", productHandler.DeleteProduct)

		//Categories routes
		v1.GET("/categories", categoryHandler.FindAllCategories)
		v1.GET("/categories/:id", categoryHandler.FindCategoryByID)
		v1.POST("/categories", categoryHandler.CreateCategory)
		v1.PUT("/categories/:id", categoryHandler.UpdateCategory)
		v1.DELETE("/categories/:id", categoryHandler.DeleteCategory)
	}
}
