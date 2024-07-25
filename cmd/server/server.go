package server

import (
	"log"
	"net/http"

	"github.com/cesc1802/onboarding-and-volunteer-service/cmd/config"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config) *Server {
	r := feature.NewRouter()

	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.ServerAddress,
			Handler: r,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Starting server on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
