package server

import (
	"fmt"
	"go-retro/config"
	"go-retro/logger"
	"net/http"
)

// InitServer sets up routes and starts listening on configured port
func InitServer() {
	portStr := fmt.Sprintf(":%d", config.Server().GetPort())

	router := setupRoutes()

	server := &http.Server{
		Addr:    portStr,
		Handler: router,
	}

	logger.Error(server.ListenAndServe())
}
