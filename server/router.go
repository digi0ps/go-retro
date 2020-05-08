package server

import (
	"fmt"
	"go-retro/config"
	"go-retro/logger"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/ping", Ping)
}

// InitServer sets up routes and starts listening on configured port
func InitServer() {
	conf := config.GetConfig()
	portStr := fmt.Sprintf(":%d", conf.GetPort())

	setupRoutes()

	err := http.ListenAndServe(portStr, nil)

	if err != nil {
		logger.Error(err)
	}
}
