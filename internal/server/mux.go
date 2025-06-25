package server

import (
	"net/http"

)

// A default net/http implementation for Server
type ServerMux struct {
	mux *http.ServeMux
}

func NewServerMux() *ServerMux {
	return &ServerMux{mux: http.NewServeMux()}
}

func (m *ServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

func (m *ServerMux) GET(path string, handler http.HandlerFunc) {
	m.mux.Handle(path, MiddlewareCheckMethod("GET", handler))
}

func (m *ServerMux) POST(path string, handler http.HandlerFunc) {
	m.mux.Handle(path, MiddlewareCheckMethod("POST", handler))
}

