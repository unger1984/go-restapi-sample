package routes

import (
	_ "awcoding.com/back/docs"
	core2 "awcoding.com/back/src/domain/core"
	"awcoding.com/back/src/domain/users"
	"awcoding.com/back/src/infrastructure/config"
	auth2 "awcoding.com/back/src/routes/auth"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

func NewHandler(s *core2.AppServices, cfg *config.Config) http.Handler {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.ForceConsoleColor()
	}

	router := gin.New()

	if cfg.Env != "production" && cfg.HttpServerConfig.Static != "" {
		router.Static("/uploads", cfg.HttpServerConfig.Static)
	}

	// For 404
	router.NoRoute(func(ctx *gin.Context) {
		core2.NewErrorResponse(ctx, http.StatusNotFound, errors.New("not found"))
	})

	// For logging
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			var statusColor, resetColor string
			if param.IsOutputColor() {
				statusColor = param.StatusCodeColor()
				resetColor = param.ResetColor()
			}
			if param.Latency > time.Minute {
				// Truncate in a golang < 1.8 safe way
				param.Latency = param.Latency - param.Latency%time.Second
			}
			// your custom format
			return fmt.Sprintf("%s %s%3d%s %s \"%s\" %13v %15s %s",
				param.TimeStamp.Format("02.01.2006 15:04:05"),
				statusColor, param.StatusCode, resetColor,
				param.Method,
				param.Path,
				param.Latency,
				param.ClientIP,
				param.ErrorMessage,
			)
		},
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
	auth2.NewRoutesFactory(authGroup)(s.AuthService)
	apiGroup := router.Group("/api")
	{
		apiGroup.Use(auth2.NewJWTMiddlewareFactory(s.AuthService))

		apiGroup.GET("test", func(ctx *gin.Context) {
			user := ctx.MustGet("user").(*users.User)
			ctx.JSON(http.StatusOK, user)
		})
	}

	return router
}
