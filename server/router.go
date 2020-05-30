package server

import (
	"go-retro/database"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(retroDatabase database.Service) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", Ping).Methods(http.MethodGet)

	handlers := RetroHandler{
		db: retroDatabase,
	}
	r.HandleFunc("/api/board", handlers.GetBoard).Methods(http.MethodGet)

	return r
}
