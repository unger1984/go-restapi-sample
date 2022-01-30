package auth_controller

import (
	"awcoding.com/back/pkg/controllers/restapi/response"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (a *AuthController) VerifyJWT(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		response.NewErrorResponse(ctx, http.StatusUnauthorized, errors.New("empty auth_controller header"))
		return
	}
	token := strings.TrimPrefix(header, "Bearer ")
	user, err := a.AuthCases.GetByToken(token)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	ctx.Set("user", user)
}
