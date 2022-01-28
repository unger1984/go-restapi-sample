package auth

import (
	"awcoding.com/back/domain/auth"
	"awcoding.com/back/domain/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewRoutesFactory(group *gin.RouterGroup) func(service auth.AuthService) {
	return func(service auth.AuthService) {
		group.POST("/sign-in", func(ctx *gin.Context) {
			var input signInInput

			if err := ctx.BindJSON(&input); err != nil {
				core.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
				return
			}

			res, err := service.SignIn(input.Login, input.Password)
			if err != nil {
				core.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
				return
			}

			ctx.JSON(http.StatusOK, res)
		})
	}
}
