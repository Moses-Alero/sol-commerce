package product

import (
	"fmt"
	"math/rand"
	"net/http"
	"sol-commerce/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(ctx *gin.Context) (db.Products, error)
	GetByID(ctx *gin.Context, ID int) (*db.Product, error)
	GetByCategory(ctx *gin.Context, Category string) (db.Products, error)
}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}

func (p *repo) GetAll(ctx *gin.Context) (db.Products, error) {
	skip := strconv.Itoa(rand.Intn(120))
	path := "products?skip=" + skip
	products, err := db.Request[db.Products](http.MethodGet, path, nil)
	if err != nil {
		return db.Products{}, nil
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

func (p *repo) GetByCategory(ctx *gin.Context, Category string) (db.Products, error) {
	skip := strconv.Itoa(rand.Intn(4))
	path := "products/category/" + Category + "?limit=3&skip=" + skip
	fmt.Println(path)
	products, err := db.Request[db.Products](http.MethodGet, path, nil)
	if err != nil {
		fmt.Println(err.Error())
		return db.Products{}, nil
	}
	return *products, nil
}
