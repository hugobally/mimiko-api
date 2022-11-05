package main

import (
	"fmt"
	"github.com/hugobally/mimiko/backend/static"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hugobally/mimiko/backend/auth/jwt"
	"github.com/hugobally/mimiko/backend/auth/login"
	"github.com/hugobally/mimiko/backend/client"
	"github.com/hugobally/mimiko/backend/config"
	"github.com/hugobally/mimiko/backend/prisma"
	"github.com/hugobally/mimiko/backend/shared"
	"github.com/hugobally/mimiko/backend/spotify"

	"github.com/gorilla/handlers"
	"github.com/hugobally/mimiko/backend/api"
	"github.com/hugobally/mimiko/backend/server"
)

func initSharedServices() *shared.Services {
	cfg := config.New()
	httpClient := client.New()
	prismaClient := prisma.New(nil)
	jwtUtil := jwt.NewJwtHandler(cfg.Auth.JwtKey)

	svcs := shared.NewServices()

	svcs.SetConfig(cfg)
	svcs.SetHttpClient(httpClient)
	svcs.SetPrisma(prismaClient)
	svcs.SetJwtUtil(jwtUtil)

	sp := spotify.New(cfg, httpClient)
	svcs.SetSpotify(sp)

	return svcs
}

func main() {
	svcs := initSharedServices()

	// TODO
	lf, err := os.OpenFile("server.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer lf.Close()
	logger := log.New(io.MultiWriter(lf, os.Stdout), "backend ", log.LstdFlags|log.Lshortfile)
	svcs.SetLogger(logger)
	//

	cfg := svcs.Config

	mux := http.NewServeMux()

	login.NewHandler(svcs).SetupRoutes(mux)
	api.NewHandler(svcs).SetupRoutes(mux)
	static.NewHandler(svcs).SetupRoutes(mux)

	addr := fmt.Sprintf("%v:%d", cfg.Server.Host, cfg.Server.Port)
	srv := server.New(handlers.LoggingHandler(os.Stdout, mux), addr)

	svcs.Logger.Printf("server starting at %v", addr)
	err = srv.ListenAndServeTLS(cfg.Tls.Cert, cfg.Tls.Key)
	//err = srv.ListenAndServe()

	if err != nil {
		svcs.Logger.Fatalf("server failed to start: %v", err)
	}
}
