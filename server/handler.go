package server

import (
	"fmt"
	"go-retro/logger"
	"net/http"
)

// Ping handler
func Ping(w http.ResponseWriter, r *http.Request) {
	logger.Info(r)
	fmt.Fprintf(w, "Pong\n")
}
