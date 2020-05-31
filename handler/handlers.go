package handler

import (
	"fmt"
	"go-retro/database"
	"go-retro/logger"
	"net/http"
)

type RetroHandler struct {
	Database database.Service
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

	board, err := retro.Database.FindBoard(boardID)
	if err != nil {
		writeErrorResponse(w, "Board entity not found", http.StatusNotFound)
		return
	}

	writeSuccessResponse(w, board)
	return
}
