package service

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}
func (s *Server) Run(port string, ip string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ip + ":" + port,
		Handler:      handler,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
