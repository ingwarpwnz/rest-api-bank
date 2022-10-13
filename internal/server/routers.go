package server

import (
	"net/http"

	"github.com/ingwarpwnz/rest-api-bank/internal/api/http/controller"
	"github.com/ingwarpwnz/rest-api-bank/internal/repository/db"
	"github.com/ingwarpwnz/rest-api-bank/internal/repository/memory"
	"github.com/ingwarpwnz/rest-api-bank/internal/service"
)

func (h *Handler) Init() {
	storage := db.NewMemoryStorage()
	accountRepo := memory.NewAccountRepository(storage)
	accountService := service.NewAccountService(accountRepo)
	accountHandler := controller.NewAccountHandler(accountService)
	h.router.HandleFunc("/account", accountHandler.Create()).Methods(http.MethodPost)
	h.router.HandleFunc("/account/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", accountHandler.Detail()).Methods(http.MethodGet)
	h.router.HandleFunc("/account/transfer", accountHandler.Transfer()).Methods(http.MethodPost)
}
