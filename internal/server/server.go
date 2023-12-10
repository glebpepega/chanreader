package server

import (
	"bytes"
	"fmt"
	"github.com/glebpepega/chanreader/internal/server/config"
	"github.com/glebpepega/chanreader/internal/server/handler"
	"github.com/glebpepega/chanreader/internal/server/logger"
	"log/slog"
	"net/http"
)

type Server struct {
	Addr    string
	WebHook string
	ApiURL  string
	Router  *http.ServeMux
	Logger  *slog.Logger
}

func New(cfg *config.Config) *Server {
	s := new(Server)

	s.configure(cfg)

	return s
}

func (s *Server) configure(cfg *config.Config) {
	s.Addr = cfg.Addr
	s.WebHook = cfg.WebHook
	s.ApiURL = cfg.ApiUrl
	s.Router = http.NewServeMux()
	s.Logger = logger.New()
}

func (s *Server) setWebHook() (err error) {
	body := bytes.NewBuffer([]byte(fmt.Sprintf(
		`{"url":"%v"}`,
		s.WebHook,
	)))

	resp, err := http.Post(s.ApiURL+"/setWebhook", "application/json", body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return
}

func (s *Server) setupHandlers() {
	s.Router.HandleFunc("/", handler.New(s.Logger, s.ApiURL))
}

func (s *Server) Start() (err error) {
	err = s.setWebHook()
	s.setupHandlers()

	s.Logger.Info("server started running")

	err = http.ListenAndServe(s.Addr, s.Router)

	return
}
