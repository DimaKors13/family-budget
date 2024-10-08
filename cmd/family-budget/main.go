// Пакет main реализует работу http-сервера приложения по учету доходов и расходов и его API.
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

// Варианты запуска приложения
const (
	envLocal string = "local"
	envDev   string = "dev"
	envProd  string = "prod"
)

// main главный метод приложения, осуществляющий его запуск.
//
// Запуск http-сервера выполняется в несколько этапов:
//  1. Инициализация config файла
//  2. Инициация логгера
//  3. Установка соединения с базой данных, запуск миграций при установке соединения
//  4. Инициализация роутера
//     4.1 Объявление middleware хендлеров
//     4.2 Объявление основных хендлеров запросов клиента
//  5. Инициализация и запуск сервера
func main() {

	config := config.MustLoad()

	log := setupLogger(config.Env)

	log.Info("Starting Family Budget App.", slog.String("env", config.Env))
	log.Debug("Debug messages in logger are enabled.")

	storage, err := postgreSQL.New(&config.DBAccessInfo)
	if err != nil {
		log.Error("Failed to initialize storage with migration", logger.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(httpLogger.New(log))
	router.Use(middleware.Recoverer)

	router.Post("/account", handlers.PostAccount(log, storage))

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

// setupLogger выполняет инициализацию логгера slog и его настройку по данным config файла.
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
