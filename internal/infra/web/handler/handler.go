package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/desafio-backend/internal/infra/web/utils"
)

func HelloWorld(ctx *gin.Context) {
	utils.SendSuccess(ctx, "Hello World", nil, 200)
}
