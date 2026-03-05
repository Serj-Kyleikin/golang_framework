package http

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	addr string
	http *http.Server
}

func ConstructServer(addr string, handler http.Handler) *Server {
	return &Server{
		addr: addr,
		http: &http.Server{
			Addr:              addr,
			Handler:           handler,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       60 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
