package product

import (
	"net/http"
	"sol-commerce/db"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(ctx *gin.Context) ([]db.Product, error)
	GetByID(id int) (*db.Product, error)
}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}

func (p *repo) GetAll(ctx *gin.Context) ([]db.Product, error) {
	data, err := db.Request[[]db.Product](http.MethodGet, "products", nil)
	if err != nil {
		return []db.Product{}, nil
	}
	return *data, nil
}

func (p *repo) GetByID(id int) (*db.Product, error) {
	return nil, nil
}
