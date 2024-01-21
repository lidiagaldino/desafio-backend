package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/desafio-backend/internal/application/usecase"
	"github.com/lidiagaldino/desafio-backend/internal/infra/web/handler"
)
func initializeRoutes(router *gin.Engine, pu *usecase.ProductUsecases) {
	productHandler := handler.NewProductHandler(pu)
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
	}
}
