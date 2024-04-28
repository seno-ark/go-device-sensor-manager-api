package main

import (
	"context"
	"fmt"
	apiv1 "go-api/internal/api/v1"
	"go-api/internal/repositories/postgres"
	"go-api/pkg/config"
	"go-api/pkg/database"
	"go-api/pkg/util"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"

	_ "go-api/cmd/api/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Device-Sensor API
// @version 1.0
// @description This is a simple api server to manage devices and sensors

// @host localhost:9000
// @BasePath
func main() {
	conf := config.GetConfig()

	db, err := database.Postgres(conf)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	validate := validator.New()
	util.RegisterCustomValidator(validate)

	repository := postgres.NewRepository(db)
	handlerV1 := apiv1.NewHandler(validate, repository)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))

	r.Mount("/v1", handlerV1.Routes())

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm fine, thanks")
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.Port),
		Handler: r,
	}

	go func() {
		slog.Info("Starting HTTP server...", "port", conf.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("ListenAndServe: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("Shutting down HTTP server...")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed: %v", err)
	}

	slog.Info("HTTP server gracefully stopped.")
}
