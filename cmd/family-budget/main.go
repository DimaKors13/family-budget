package main

import (
	"family-budget/internal/config"
	"family-budget/internal/lib/logger"
	"family-budget/internal/storage/postgreSQL"
	"log/slog"
	"os"
)

const (
	envLocal string = "local"
	envDev   string = "dev"
	envProd  string = "prod"
)

func main() {

	//TODO: init CONFIG: cleanenv
	config := config.MustLoad()

	// TODO: init logger: slog
	log := setupLogger(config.Env)

	log.Info("Starting Family Budget App.", slog.String("env", config.Env))
	log.Debug("Debug messages in logger are enabled.")

	// TODO: init storage: postgresql
	storage, err := postgreSQL.New(&config.DBAccessInfo)
	if err != nil {
		log.Error("Failed to initialize storage with migration", logger.Err(err))
		os.Exit(1)
	}

	_ = storage

	// TODO run router: chi, chi.render

	// TODO: run server

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
