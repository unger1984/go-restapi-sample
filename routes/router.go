package routes

import (
	"awcoding.com/back/routes/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewHandler() http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	apiGroup := router.Group("/api")
	{
		authGroup := apiGroup.Group("/auth")
		auth.NewRoutesFactory(authGroup)()
	}

	return router
}
