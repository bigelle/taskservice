package server

import "net/http"

// Server is an interface used for drop-in replacement from one framework to another
type Server interface {
	http.Handler
	GET(path string, handler http.HandlerFunc)
	POST(path string, handler http.HandlerFunc)
}

func MiddlewareCheckMethod(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("method not allowed"))
		}
		next.ServeHTTP(w,r)
	})
}
