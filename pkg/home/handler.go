package home

import (
	"net/http"

	"sol-commerce/components"

	"github.com/gin-gonic/gin"
)

func GetHomePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", components.HomePage())
	}
}
