package server

import (
	"github.com/glebpepega/chanreader/internal/server/config"
	"github.com/glebpepega/chanreader/internal/server/logger"
	"log/slog"
	"net/http"
)

type Server struct {
	Addr   string
	Router *http.ServeMux
	Logger *slog.Logger
}

func New(cfg *config.Config) *Server {
	s := new(Server)

	s.configure(cfg)

	return s
}

func (s *Server) configure(cfg *config.Config) {
	s.Addr = cfg.BindAddr
	s.Router = http.NewServeMux()
	s.Logger = logger.New()
}

func (s *Server) Start() (err error) {
	return http.ListenAndServe(s.Addr, s.Router)
}
