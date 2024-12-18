package http

import (
	"context"
	"email-service/pkg/log"
	"errors"
	"net/http"
	"time"
)

const (
	serverReadTimeout     = 5 * time.Second
	serverWriteTimeout    = 55 * time.Second
	serverShutdownTimeout = 60 * time.Second
)

type server struct {
	srv *http.Server
	log log.MultiLogger
}

func NewServer(port string, handler http.Handler, log log.MultiLogger) *server {
	return &server{
		srv: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  serverReadTimeout,
			WriteTimeout: serverWriteTimeout,
		},
		log: log,
	}
}

func (s *server) Start() {
	go func() {
		s.log.Info().Printf("starting HTTP server in %s", s.srv.Addr)
		err := s.srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Fatal().Print("error on start server")
		}
	}()
}

func (s *server) Shutdown() {
	s.log.Info().Println("shutting down HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.log.Error().Printf("could not shutdown in %d:%v", serverShutdownTimeout, err)
	}
	s.log.Info().Println("server HTTP gracefully stopped")
}
