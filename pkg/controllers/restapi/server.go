package restapi

import (
	_ "awcoding.com/back/docs"
	"awcoding.com/back/pkg/controllers/restapi/auth_controller"
	"awcoding.com/back/pkg/controllers/restapi/response"
	"awcoding.com/back/pkg/domain/entities"
	"awcoding.com/back/pkg/domain/usecases"
	"awcoding.com/back/pkg/infrastructure/config"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

type Server struct {
	UserCases usecases.UserCases
	AuthCases usecases.AuthCases
	Config    config.Config
}

func NewServer(userCases usecases.UserCases, authCases usecases.AuthCases, cfg config.Config) *Server {
	return &Server{
		AuthCases: authCases,
		UserCases: userCases,
		Config:    cfg,
	}
}

func (s *Server) NewdHandler() http.Handler {
	if s.Config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.ForceConsoleColor()
	}

	router := gin.New()

	if s.Config.Env != "production" && s.Config.HttpServerConfig.Static != "" {
		router.Static("/uploads", s.Config.HttpServerConfig.Static)
	}

	// For 404
	router.NoRoute(func(ctx *gin.Context) {
		response.NewErrorResponse(ctx, http.StatusNotFound, errors.New("not found"))
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
			return fmt.Sprintf("%s %s%3d%s %6s \"%s\" %13v %15s %s",
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

	authController := auth_controller.NewAuthController(s.AuthCases)
	authGroup := router.Group("/auth")
	authController.NewRoutes(authGroup)

	apiGroup := router.Group("/api")
	{
		apiGroup.Use(authController.VerifyJWT)

		apiGroup.GET("test", func(ctx *gin.Context) {
			user := ctx.MustGet("user").(*entities.User)
			ctx.JSON(http.StatusOK, user)
		})
	}

	return router
}
