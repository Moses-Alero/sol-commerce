package product

import (
	"sol-commerce/db"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetAll(ctx *gin.Context) ([]db.Product, error)
	// GetByID(id int) (db.Product, error)
}

type productService struct {
	r Repository
}

func NewService(repo Repository) ProductService {
	return &productService{
		r: repo,
	}
}

func (p *productService) GetAll(ctx *gin.Context) ([]db.Product, error) {
	data, err := p.r.GetAll(ctx)
	if err != nil {
		return []db.Product{}, err
	}
	return data, nil
}
