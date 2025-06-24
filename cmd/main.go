package main

import (
	"log/slog"
	"net/http"
	
	"github.com/bigelle/taskservice/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	
	handlers := handlers.New()
	for path, handler := range handlers {
		mux.Handle(path, handler)
	}

	slog.Error(http.ListenAndServe(":8080", mux).Error())
}
