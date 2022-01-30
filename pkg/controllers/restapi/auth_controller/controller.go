package auth_controller

import (
	"awcoding.com/back/pkg/domain/usecases"
	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthController struct {
	AuthCases usecases.AuthCases
}

func NewAuthController(authCases usecases.AuthCases) *AuthController {
	return &AuthController{AuthCases: authCases}
}

func (a *AuthController) NewRoutes(g *gin.RouterGroup) {
	g.POST("/sign-in", a.signIn)
}
