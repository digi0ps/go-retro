package server

import (
	"go-retro/database"
	"go-retro/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(retroDatabase database.Service) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler.Ping).Methods(http.MethodGet)

	handlers := handler.RetroHandler{
		Database: retroDatabase,
	}

	r.HandleFunc("/api/board", handlers.GetBoard).Methods(http.MethodGet)
	r.HandleFunc("/api/board", handlers.PutBoard).Methods(http.MethodPut)

	return r
}
