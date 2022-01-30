package auth

import (
	domain "awcoding.com/back/src/domain/auth"
	"awcoding.com/back/src/domain/core"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      Вход
// @Description  Вход по логину и паролю
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body signInInput true  "SignIn information"
// @Success      200  {object}  domain.Auth
// @Failure      400  {object}  core.errorResponse
// @Failure      500  {object}  core.errorResponse
// @Router       /auth/sign-in [post]
func signIn(ctx *gin.Context, s domain.AuthService) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		core.NewErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	res, err := s.SignIn(input.Login, input.Password)
	if err != nil {
		if err.Error() == "login and password incorrect" {
			core.NewErrorResponse(ctx, http.StatusUnauthorized, err)
		} else {
			core.NewErrorResponse(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, res)
}
