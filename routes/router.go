package routes

import (
	"awcoding.com/back/domain/core"
	"awcoding.com/back/domain/users"
	"awcoding.com/back/routes/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewHandler(s *core.AppServices) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	authGroup := router.Group("/auth")
	auth.NewRoutesFactory(authGroup)(s.AuthService)
	apiGroup := router.Group("/api")
	{
		apiGroup.Use(auth.NewJWTMiddlewareFactory(s.AuthService))

		apiGroup.GET("test", func(ctx *gin.Context) {
			user := ctx.MustGet("user").(*users.User)
			ctx.JSON(http.StatusOK, user)
		})
	}

	return router
}
