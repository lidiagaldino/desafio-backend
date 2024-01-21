package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/desafio-backend/internal/application/usecase"
)

func Initialize(pu *usecase.ProductUsecases) {
	r := gin.Default()
	initializeRoutes(r, pu)
	r.Run()
}
