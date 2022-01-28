package main

import (
	"awcoding.com/back/infrastructure/config"
	"awcoding.com/back/routes"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	httpServer *http.Server
}

func (s *server) Run(config config.HttpServerConfig, handler http.Handler, isProduction bool) error {
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

func main() {
	cfg := config.GetInstance()

	if err := config.Load(cfg); err != nil {
		log.Fatal(err)
	}

	handler := routes.NewHandler()

	srv := new(server)

	go func() {
		if err := srv.Run(cfg.HttpServerConfig, handler, cfg.ENV == "production"); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Graceful shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

	//if err := db.Close(); err != nil {
	//	log.Fatalf("error occured on db connection close: %s", err.Error())
	//}
}
