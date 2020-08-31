package server

import (
	"go-retro/database"
	"go-retro/handler"
	"go-retro/middlewares"
	"go-retro/websocket"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(retroDatabase database.Service, boardHub *websocket.Hub) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler.Ping).Methods(http.MethodGet)

	// Retro API handlers
	handlers := handler.RetroHandler{
		Database: retroDatabase,
	}

	r.HandleFunc("/api/board", handlers.GetBoard).Methods(http.MethodGet)
	r.HandleFunc("/api/board", handlers.PutBoard).Methods(http.MethodPut, http.MethodOptions)
	r.Use(middlewares.CorsMiddleware)

	// Websocket API Handlers
	socks := websocket.SocketHandler{
		BoardHub: boardHub,
		Database: retroDatabase,
	}

	r.HandleFunc("/ws/{board}", socks.InitHandler)

	return r
}
