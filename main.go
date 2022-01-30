package main

import (
	database2 "awcoding.com/back/src/data/database"
	datUsers "awcoding.com/back/src/data/users"
	"awcoding.com/back/src/domain/auth"
	"awcoding.com/back/src/domain/core"
	"awcoding.com/back/src/domain/users"
	"awcoding.com/back/src/infrastructure/config"
	"awcoding.com/back/src/routes"
	"context"
	"errors"
	"flag"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	httpServer *http.Server
}

func (s *server) Run(config config.HttpServerConfig, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           config.Host + ":" + config.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServeTLS(config.CertPem, config.CertKey)
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// @title           Swagger Back API
// @version         1.0
// @description     This is a sample REST API server.
// @BasePath  /
// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	flgConfigPath := flag.String("config", "", "Path to config")
	var configPath string
	if len(*flgConfigPath) == 0 {
		configPath = "./config.yaml"
	} else {
		configPath = *flgConfigPath
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := database2.ConnectPostgresDB(cfg.DBConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := database2.ApplyMigrations(db.DB, cfg.DBConfig.MigrationsPath); err != nil {
		logrus.Fatal(err)
	}

	usersRepository := datUsers.NewRepository(db)
	userService := users.NewService(usersRepository)
	authService := auth.NewService(userService, cfg)
	appServices := core.NewAppService(userService, authService)

	handler := routes.NewHandler(appServices, cfg)
	srv := new(server)

	go func() {
		if err := srv.Run(cfg.HttpServerConfig, handler); err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Println("Graceful shutting down...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatal("Server forced to shutdown: %s", err)
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("error occured on db connection close: %s", err.Error())
	}
}
