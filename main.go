package main

import (
	"net/http"
	"sol-commerce/components"
	templrender "sol-commerce/x/templRender"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()

	ginEngine.SetTrustedProxies(nil)

	htmlRenderer := ginEngine.HTMLRender
	ginEngine.HTMLRender = &templrender.HTMLRenderer{FallbackRenderer: htmlRenderer}
	ginEngine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, " ", components.Base())
	})
	ginEngine.Static("/static", "./static")
	ginEngine.Static("/assets", "./assets")
	ginEngine.Run(":8080")
}
