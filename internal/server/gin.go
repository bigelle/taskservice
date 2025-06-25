package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServerGin implements Server
type ServerGin struct {
	e *gin.Engine
}

func NewServerGin() *ServerGin {
	return &ServerGin{e: gin.Default()}
}

func (s *ServerGin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.e.ServeHTTP(w, r)
}

func (s *ServerGin) GET(path string, handler http.HandlerFunc) {
	s.e.GET(path, gin.WrapF(handler))
}

func (s *ServerGin) POST(path string, handler http.HandlerFunc) {
	s.e.POST(path, gin.WrapF(handler))
}

