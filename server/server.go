package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(portNum string, handler http.Handler) error {
	httpServer := &http.Server{
		Addr:              portNum,
		Handler:           handler,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       1 * time.Minute,
		MaxHeaderBytes:    1 << 20,
	}

	return httpServer.ListenAndServe()
}

func (s *Server) ShutDownServer() error {
	return s.httpServer.Close()
}
