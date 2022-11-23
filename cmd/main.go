package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/hugobally/mimiko_api/internal/authentication/jwt"
	"github.com/hugobally/mimiko_api/internal/authentication/login"
	"github.com/hugobally/mimiko_api/internal/authentication/spotify"
	"github.com/hugobally/mimiko_api/internal/config"
	"github.com/hugobally/mimiko_api/internal/db"
	"github.com/hugobally/mimiko_api/internal/graph"
	"github.com/hugobally/mimiko_api/internal/server"
	"github.com/hugobally/mimiko_api/internal/shared"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func initSharedServices() *shared.Services {
	configInstance := config.New()
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	services := shared.NewServices()

	services.SetConfig(configInstance)
	services.SetHttpClient(httpClient)
	services.SetJwtUtil(jwt.NewJwtHandler(configInstance.Auth.JwtKey))
	services.SetSpotify(spotify.New(configInstance, httpClient))
	services.SetDatabase(db.NewClient(configInstance.Env == "DEV"))

	return services
}

func main() {
	services := initSharedServices()

	lf, err := os.OpenFile("server.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer lf.Close()

	logger := log.New(io.MultiWriter(lf, os.Stdout), "backend ", log.LstdFlags|log.Lshortfile)
	services.SetLogger(logger)

	cfg := services.Config

	mux := http.NewServeMux()

	login.NewHandler(services).SetupRoutes(mux)
	graph.NewHandler(services).SetupRoutes(mux)

	addr := fmt.Sprintf("%v:%d", cfg.Server.Host, cfg.Server.Port)
	srv := server.New(handlers.LoggingHandler(os.Stdout, mux), addr)

	services.Logger.Printf("server starting at %v", addr)
	err = srv.ListenAndServeTLS(cfg.Tls.Cert, cfg.Tls.Key)

	if err != nil {
		services.Logger.Fatalf("server failed to start: %v", err)
	}
}
