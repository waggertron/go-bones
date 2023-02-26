package transport

import (
	"context"
	"net/http"
)

type HttpServerCfg struct {
	Name string
	Port string
}

type HttpServer struct {
	name   string
	port   string
	mux    *http.ServeMux
	ctx    context.Context
	cancel context.CancelFunc
	server http.Server
}

func NewHttpServer(cfg HttpServerCfg) *HttpServer {
	ctx, cancel := context.WithCancel(context.Background())
	server := http.Server{
		Addr:    cfg.Port,
		Handler: http.NewServeMux(),
	}
	s := HttpServer{
		name:   cfg.Name,
		ctx:    ctx,
		cancel: cancel,
		server: server,
	}

	return &s
}

func (s *HttpServer) Start() (context.CancelFunc, error) {
	return s.cancel, s.server.ListenAndServe()
}

func (s *HttpServer) Stop() error {
	return s.server.Shutdown(s.ctx)
}
