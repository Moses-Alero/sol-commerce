package home

import "github.com/gin-gonic/gin"

func GetHome(r *gin.Engine) {
	r.GET("/", GetHomePage())
}
