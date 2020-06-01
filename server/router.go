package server

import (
	"go-retro/database"
	"go-retro/handler"
	"go-retro/websocket"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(retroDatabase database.Service) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler.Ping).Methods(http.MethodGet)

	// Retro API handlers
	handlers := handler.RetroHandler{
		Database: retroDatabase,
	}

	r.HandleFunc("/api/board", handlers.GetBoard).Methods(http.MethodGet)
	r.HandleFunc("/api/board", handlers.PutBoard).Methods(http.MethodPut)

	// Websocket API Handlers
	r.HandleFunc("/ws", websocket.InitHandler)

	return r
}
