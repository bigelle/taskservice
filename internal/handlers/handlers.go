package handlers

import "net/http"

type HandlerSet map[string]http.Handler

func New() HandlerSet {
	return HandlerSet{
		"/create": CreateHandler{},
		//"/view":   nil,
		//"/update": nil,
		//"/delete": nil,
	}
}
