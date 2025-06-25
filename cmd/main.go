package main

import (
	"log/slog"
	"net/http"

	"github.com/bigelle/taskservice/internal/handlers"
	"github.com/bigelle/taskservice/internal/server"
)

func main() {
	var serv server.Server
	
	serv = server.NewServerMux()
	// works with pretty much every net/http compatible framework:
	// serv = server.NewServerGin()
	// serv = server.NewServerGorilla()

	serv.POST("/create", handlers.HandleCreate)
	serv.GET("/view", handlers.HandleView)
	serv.POST("/update", handlers.HandleUpdate)
	serv.POST("/delete", handlers.HandleDelete)

	slog.Info("Listening and serving on :8080")
	slog.Error(http.ListenAndServe(":8080", serv).Error())
}
