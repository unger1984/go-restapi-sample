package routes

import (
	_ "awcoding.com/back/docs"
	"awcoding.com/back/domain/core"
	"awcoding.com/back/domain/users"
	"awcoding.com/back/infrastructure/config"
	"awcoding.com/back/routes/auth"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

func NewHandler(s *core.AppServices, cfg *config.Config) http.Handler {
	if cfg.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.ForceConsoleColor()
	}

	router := gin.New()

	if cfg.ENV != "production" && cfg.HttpServerConfig.Static != "" {
		router.Static("/upload", cfg.HttpServerConfig.Static)
	}

	// For 404
	router.NoRoute(func(ctx *gin.Context) {
		core.NewErrorResponse(ctx, http.StatusNotFound, errors.New("not found"))
	})
	// For logging
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("[%s] - \"%s %s\" %d %s %s %s",
			param.TimeStamp.Format(time.Kitchen),
			param.Method,
			param.Path,
			param.StatusCode,
			param.ClientIP,
			param.Latency,
			param.ErrorMessage,
		)
	}))
	// For recovery
	router.Use(gin.Recovery())

	// For CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	// For Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
