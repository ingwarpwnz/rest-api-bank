package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ingwarpwnz/rest-api-bank/internal/api/http/response"
	"github.com/ingwarpwnz/rest-api-bank/internal/service"
)

type AccountHandler struct {
	service *service.AccountService
}

func NewAccountHandler(s *service.AccountService) *AccountHandler {
	return &AccountHandler{
		service: s,
	}
}

func (h *AccountHandler) Create() http.HandlerFunc {
	type createRequest struct {
		Balance float64 `json:"balance"`
	}
	req := new(createRequest)

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.InternalServerError(w, "internal server error")
			log.Printf("account handler when create(), readAll: %s", err)
			return
		}

		err = json.Unmarshal(body, &req)
		if err != nil {
			response.BadRequest(w, "the 'balance' must be an float")
			log.Printf("account handler when create(), unmarshal: %s", err)
			return
		}

		data, err := h.service.Create(req.Balance)
		response.Respond(w, data, http.StatusOK)
	}
}

func (h *AccountHandler) Detail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		acc := h.service.FindById(id)
		response.Respond(w, map[string]any{
			"data": acc,
		}, http.StatusOK)
	}
}

func (h *AccountHandler) Transfer() http.HandlerFunc {
	type transferRequest struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}
	req := new(transferRequest)

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.InternalServerError(w, "internal server error")
			log.Printf("account handler when transfer(), readAll: %s", err)
			return
		}

		err = json.Unmarshal(body, &req)
		if err != nil {
			response.BadRequest(w, "the 'from', 'to' must be an uuid, and 'amount' must be an float")
			log.Printf("account handler when transfer(), unmarshal: %s", err)
			return
		}

		from, to, err := h.service.Transfer(req.From, req.To, req.Amount)
		if err != nil {
			response.BadRequest(w, err.Error())
			return
		}

		response.Respond(w, map[string]any{
			"data": map[string]any{
				"sender":    from,
				"recipient": to,
			},
		}, http.StatusOK)
	}
}
