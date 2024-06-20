package product

import "github.com/gin-gonic/gin"

func ProductRouter(r *gin.Engine) {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	r.GET("/products", handler.GetAll)
	r.GET("/products/:id", handler.GetByID)
}
