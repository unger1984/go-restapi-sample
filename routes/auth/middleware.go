package auth

import (
	"awcoding.com/back/domain/auth"
	"awcoding.com/back/domain/core"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func NewJWTMiddlewareFactory(s auth.AuthService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			core.NewErrorResponse(ctx, http.StatusUnauthorized, errors.New("empty auth header"))
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		user, err := s.GetByToken(token)
		if err != nil {
			core.NewErrorResponse(ctx, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		ctx.Set("user", user)
	}
}
