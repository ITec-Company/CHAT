package server

import (
	"context"
	"net/http"
	"time"

	"itec.chat/internal/handlers"
	"itec.chat/pkg/logging"
)

type Server struct {
	Logger     logging.Logger
	httpServer *http.Server
}

func NewServer(logger logging.Logger, router handlers.Router, host, port string) *Server {
	srv := &http.Server{
		Addr:           host + ":" + port,
		Handler:        router.Router,
		MaxHeaderBytes: 1 << 20, //1 Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return &Server{
		Logger:     logger,
		httpServer: srv,
	}
}

func (s *Server) Start() error {

	s.Logger.Infof("Server starts at %s", s.httpServer.Addr)
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
