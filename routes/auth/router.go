package auth

import (
	"awcoding.com/back/domain/auth"
	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewRoutesFactory(group *gin.RouterGroup) func(s auth.AuthService) {
	return func(s auth.AuthService) {
		group.POST("/sign-in", func(ctx *gin.Context) {
			signIn(ctx, s)
		})
	}
}
