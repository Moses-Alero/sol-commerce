package auth

import (
	"github.com/gin-gonic/gin"
)

type Auth interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type auth struct{}

func NewAuth() Auth {
	return &auth{}
}

func (a *auth) Register(ctx *gin.Context) {}

func (a *auth) Login(ctx *gin.Context) {}

func (a *auth) Logout(ctx *gin.Context) {}
