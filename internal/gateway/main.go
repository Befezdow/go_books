package gateway

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Gateway struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func NewGateway(config *Config) *Gateway {
	return &Gateway{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Gateway) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting gateway")

	s.configureRouter()

	s.logger.Info(fmt.Sprintf("Running gateway on %s port", s.config.BindAddr))

	if err := http.ListenAndServe(s.config.BindAddr, s.router); err != nil {
		return err
	}

	return nil
}

func (s *Gateway) configureLogger() error {
	var level, err = logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Gateway) configureRouter() {
	s.router.HandleFunc("/books", s.handleGetBooks()).Methods("GET")
	s.router.HandleFunc("/books", s.handleCreateBook()).Methods("POST")
}
