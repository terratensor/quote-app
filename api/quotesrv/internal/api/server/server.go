package server

import (
	"context"
	"gitgub.com/terratensor/quote-app/api/quotesrv/internal/app/repos/quote"
	"net/http"
	"time"
)

type Server struct {
	srv http.Server
	qs  *quote.Quotes
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
}

func (s *Server) Start(qs *quote.Quotes) {
	s.qs = qs
	go s.srv.ListenAndServe()
}
