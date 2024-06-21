package product

import (
	"fmt"
	"net/http"
	"sol-commerce/components"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetByCategory(ctx *gin.Context)
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
		ctx.HTML(http.StatusInternalServerError, " ", components.ProductsPage(products.Products)) //todo: fix this to use proper error message
		fmt.Println(err)
		return
	}
	ctx.HTML(http.StatusOK, "", components.ProductsPage(products.Products))
}

func (handler *productHandler) GetByID(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id"))
	product, err := handler.service.GetByID(ctx, ID)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "", components.ProductPage(product))
		return
	}
	ctx.HTML(http.StatusOK, "", components.ProductPage(product))
}

func (handler *productHandler) GetByCategory(ctx *gin.Context) {
	Category := ctx.Param("category")
	products, err := handler.service.GetByCategory(ctx, Category)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, " ", components.ProductsPage(products.Products))
		fmt.Println(err)
		return
	}
	ctx.HTML(http.StatusOK, "", components.Products(products.Products))
}
