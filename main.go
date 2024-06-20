package main

import (
	"sol-commerce/pkg/product"
	templrender "sol-commerce/x/templRender"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()

	ginEngine.SetTrustedProxies(nil)

	htmlRenderer := ginEngine.HTMLRender
	ginEngine.HTMLRender = &templrender.HTMLRenderer{FallbackRenderer: htmlRenderer}

	ginEngine.Static("/static", "./static")
	ginEngine.Static("/assets", "./assets")

	//routers
	product.ProductRouter(ginEngine)

	ginEngine.Run(":8080")
}
