package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/ingwarpwnz/rest-api-bank/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("main api: %s", err)
	}
}

func run() error {
	// create new handler instance
	router := mux.NewRouter()
	handler := server.NewHandler(router)
	handler.Init() // Initialize routers

	srv := server.NewServer()
	addr := ":8080"
	log.Printf("app listening on port %s", addr)
	return srv.Run(addr, handler)
}
