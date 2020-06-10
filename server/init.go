package server

import (
	"fmt"
	"go-retro/config"
	"go-retro/database/mongodb"
	"go-retro/logger"
	"go-retro/websocket"
	"net/http"
)

// InitServer sets up routes and starts listening on configured port
func InitServer() {
	portStr := fmt.Sprintf(":%d", config.Server().GetPort())
	logger.Info(fmt.Sprint("Running server on localhost", portStr))

	goRetroDatabase := &mongodb.MongoDatabase{}
	goRetroDatabase.OpenConnection()
	defer goRetroDatabase.CloseConnection()

	boardHub := websocket.NewHub()
	go boardHub.Run()

	router := setupRoutes(goRetroDatabase, boardHub)

	server := &http.Server{
		Addr:    portStr,
		Handler: router,
	}

	logger.Error(server.ListenAndServe())
}
