package main

import (
	"git_copilot_test/internal/config"
	"git_copilot_test/internal/handler"
	"git_copilot_test/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.New()
	services := service.New()
	handlers := handler.New(cfg, services)

	server := &http.Server{
		Addr:         cfg.Server.Host + cfg.Server.Port,
		Handler:      handlers.InitRoutes(),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	log.Println("Starting server on", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
