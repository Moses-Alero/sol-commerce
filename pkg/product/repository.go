package product

import (
	"fmt"
	"net/http"
	"sol-commerce/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(ctx *gin.Context) ([]db.Product, error)
	GetByID(ctx *gin.Context, ID int) (*db.Product, error)
}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}

func (p *repo) GetAll(ctx *gin.Context) ([]db.Product, error) {
	products, err := db.Request[[]db.Product](http.MethodGet, "products", nil)
	if err != nil {
		return []db.Product{}, nil
	}
	return *products, nil
}

func (p *repo) GetByID(ctx *gin.Context, ID int) (*db.Product, error) {
	path := "products/" + strconv.Itoa(ID)
	product, err := db.Request[db.Product](http.MethodGet, path, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return product, nil
}
