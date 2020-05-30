package server

import (
	"encoding/json"
	"fmt"
	"go-retro/database"
	"go-retro/logger"
	"net/http"
)

type RetroHandler struct {
	db database.Service
}

// Ping handler
func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		logger.Info(r)
		fmt.Fprintf(w, "Pong\n")
	} else {
		http.NotFound(w, r)
	}
}

// GetBoard handler
func (retro *RetroHandler) GetBoard(w http.ResponseWriter, r *http.Request) {
	logger.Info("Entering Get Board handler")
	boardID := r.URL.Query().Get("id")

	board, err := retro.db.FindBoard(boardID)
	if err != nil {
		// TODO: Better error
		http.NotFound(w, r)
		return
	}

	// TODO: Move to separate response
	body, err := json.Marshal(board)
	if err != nil {
		logger.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}
