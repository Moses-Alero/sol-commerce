package product

import (
	"fmt"
	"net/http"
	"sol-commerce/components"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetAll(ctx *gin.Context)
}

type productHandler struct {
	service ProductService
}

func NewHandler(service ProductService) ProductHandler {
	return &productHandler{
		service: service,
	}
}

func (handler *productHandler) GetAll(ctx *gin.Context) {
	products, err := handler.service.GetAll(ctx)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, " ", components.Base(products))
		fmt.Println(err)
		return
	}
	ctx.HTML(http.StatusOK, "", components.Base(products))
}
