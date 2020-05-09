package server

import (
	"fmt"
	"go-retro/config"
	"go-retro/logger"
	"net/http"
)

// InitServer sets up routes and starts listening on configured port
func InitServer() {
	conf := config.GetConfig()
	portStr := fmt.Sprintf(":%d", conf.GetPort())

	router := setupRoutes()

	server := &http.Server{
		Addr:    portStr,
		Handler: router,
	}

	logger.Error(server.ListenAndServe())
}