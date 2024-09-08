package main

import (
	"family-budget/internal/config"
	"family-budget/internal/http-server/handlers"
	httpLogger "family-budget/internal/http-server/middleware/logger"
	"family-budget/internal/lib/logger"
	"family-budget/internal/storage/postgreSQL"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal string = "local"
	envDev   string = "dev"
	envProd  string = "prod"
)

func main() {

	//init CONFIG: cleanenv
	config := config.MustLoad()

	//init logger: slog
	log := setupLogger(config.Env)

	log.Info("Starting Family Budget App.", slog.String("env", config.Env))
	log.Debug("Debug messages in logger are enabled.")

	//init storage: postgresql
	storage, err := postgreSQL.New(&config.DBAccessInfo)
	if err != nil {
		log.Error("Failed to initialize storage with migration", logger.Err(err))
		os.Exit(1)
	}

	//run router
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(httpLogger.New(log))
	router.Use(middleware.Recoverer)

	router.Post("/account", handlers.PostAccount(log, storage))
	//TODO router.Get("/account", )
	//TODO router.Post("/category", )
	//TODO router.Get("/category", )
	//TODO router.Post("/flow-record", )
	//TODO router.Get("/flow-record", )

	//run server
	srv := http.Server{
		Addr:         "localhost:8081",
		Handler:      router,
		ReadTimeout:  config.HTTPServer.TimeOut,
		WriteTimeout: config.HTTPServer.TimeOut,
		IdleTimeout:  config.HTTPServer.IdleTimeout,
	}

	log.Info("Starting server")

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
