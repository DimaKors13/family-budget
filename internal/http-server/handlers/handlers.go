package handlers

import (
	"family-budget/internal/http-server/api"
	"family-budget/internal/http-server/data"
	"family-budget/internal/lib/logger"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type dataStorage interface {
	AddAccount(name string) (int, error)
}

func PostAccount(log *slog.Logger, storage dataStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log := log.With(
			slog.String("handler", "postAccount"),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		var account data.Account
		err := render.DecodeJSON(r.Body, &account)
		if err != nil {
			log.Error("failed to decode request body: %w", logger.Err(err))
			render.JSON(w, r, api.Error("failed to decode request body"))
			return
		}

		log.Info("Request body decoded")

		if err := validator.New().Struct(account); err != nil {
			log.Error("invalid request: %w", logger.Err(err))
			validateErr := err.(validator.ValidationErrors)
			render.JSON(w, r, api.ValidationError(validateErr))
			return
		}

		accountId, err := storage.AddAccount(account.Name)
		if err != nil {
			log.Error("failed to create account in storage: %w", logger.Err(err))
			render.JSON(w, r, api.Error("failed to create account in storage"))
			return
		}

		log.Info("New account created. accountId = " + strconv.Itoa(accountId))
		render.JSON(w, r, api.OK())
	}
}
