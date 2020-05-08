package server

import (
	"fmt"
	"go-retro/logger"
	"net/http"
)

// Ping handler
func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		logger.Info(r)
		fmt.Fprintf(w, "Pong\n")
	} else {
		http.NotFound(w, r)
	}
}
