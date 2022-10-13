package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Respond(w http.ResponseWriter, data any, code int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("respond error: %s", err)
	}
}

func InternalServerError(w http.ResponseWriter, msg string) {
	data := map[string]string{
		"error": msg,
	}
	Respond(w, data, http.StatusInternalServerError)
}

func BadRequest(w http.ResponseWriter, msg string) {
	data := map[string]string{
		"error":   "bad request",
		"message": msg,
	}
	Respond(w, data, http.StatusBadRequest)
}
