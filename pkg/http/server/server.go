package server

import (
	"context"
	"gomoku/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	ctx    context.Context
	server *http.Server
}

func NewHttpServer(ctx context.Context, addr string, handler http.Handler) *Server {
	return &Server{
		ctx: ctx,
		server: &http.Server{
			Handler:           handler,
			Addr:              addr,
			ReadTimeout:       5 * time.Second,
			WriteTimeout:      10 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	logger.Infof("http server run on [%s]", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(signals ...os.Signal) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	<-c

	if err := s.server.Shutdown(s.ctx); err != nil {
		logger.Error(err)
	}
	return nil
}
