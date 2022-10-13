package server

import (
	"net/http"
	"time"
)

type Server struct {
}

func NewServer() *Server {
	return new(Server)
}

func (s *Server) Run(addr string, handler http.Handler) error {
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	return srv.ListenAndServe()
}
