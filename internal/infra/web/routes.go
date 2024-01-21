package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/desafio-backend/internal/infra/web/handler"
)
func initializeRoutes(router *gin.Engine) {
	basePath := "api/v1/"
	v1 := router.Group(basePath)
	{
		v1.GET("/", handler.HelloWorld)
	}
}
