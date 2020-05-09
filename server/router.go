package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", Ping).Methods(http.MethodGet)

	return r
}
