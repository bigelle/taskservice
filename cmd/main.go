package main

import (
	"log/slog"
	"net/http"
	
	"github.com/bigelle/taskservice/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/create", handlers.HandleCreate)

	slog.Error(http.ListenAndServe(":8080", mux).Error())
}
