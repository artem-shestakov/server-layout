package server

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Handler    http.Handler
	HttpServer http.Server
	Logger     *logrus.Logger
}

// New server create Server object
func NewServer(address string, handler http.Handler, logger *logrus.Logger) *Server {
	srv := http.Server{
		Addr:    address,
		Handler: handler,
	}
	logger.SetFormatter(&logrus.JSONFormatter{})
	return &Server{
		Handler:    handler,
		HttpServer: srv,
		Logger:     logger,
	}
}

// Run http server
func (s *Server) Run() {
	s.Logger.Infof("Server starting at %s", s.HttpServer.Addr)
	s.HttpServer.ListenAndServe()
}

// Stop http server
func (s *Server) Stop(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	if err := s.HttpServer.Shutdown(ctx); err != nil {
		s.Logger.Fatalf("Error with sop serve %s", err.Error())
	}
	cancel()
	s.Logger.Infoln("Server stoped")
}
