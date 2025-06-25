package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ServerGorilla implements Server
type ServerGorilla struct {
	mux *mux.Router
}

func NewServerGorilla() *ServerGorilla {
	return &ServerGorilla{mux: mux.NewRouter()}
}

func (s *ServerGorilla) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *ServerGorilla) GET(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler).Methods("GET")
}

func (s *ServerGorilla) POST(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler).Methods("POST")
}
