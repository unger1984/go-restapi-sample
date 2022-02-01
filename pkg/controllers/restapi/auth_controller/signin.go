package auth_controller

import (
	"awcoding.com/back/pkg/controllers/restapi/response"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      Вход
// @Description  Вход по логину и паролю
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body SignInInput true  "SignIn information"
// @Success      200  {object}  entities.Auth
// @Failure      400  {object}  response.errorResponse
// @Failure      500  {object}  response.errorResponse
// @Router       /auth/sign-in [post]
func (a *AuthController) signIn(ctx *gin.Context) {
	var input SignInInput

	if err := ctx.BindJSON(&input); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	res, err := a.AuthCases.SignIn(input.Login, input.Password)
	if err != nil {
		if err.Error() == "login and password incorrect" {
			response.NewErrorResponse(ctx, http.StatusUnauthorized, err)
		} else {
			response.NewErrorResponse(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, res)
}
