package product

import (
	"sol-commerce/db"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetAll(ctx *gin.Context) ([]db.Product, error)
	GetByID(ctx *gin.Context, ID int) (db.Product, error)
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
	products, err := p.r.GetAll(ctx)
	if err != nil {
		return []db.Product{}, err
	}
	return products, nil
}

func (p *productService) GetByID(ctx *gin.Context, ID int) (db.Product, error) {
	product, err := p.r.GetByID(ctx, ID)
	if err != nil {
		return db.Product{}, err
	}
	return *product, nil
}
