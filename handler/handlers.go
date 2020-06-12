package handler

import (
	"fmt"
	"go-retro/database"
	"go-retro/logger"
	"net/http"
)

// RetroHandler groups all the rest handlers as method so they can share the dependencies
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
	boardID := r.URL.Query().Get("id")

	board, err := retro.Database.FindBoard(boardID)
	if err == database.ErrorNotFound {
		writeErrorResponse(w, "Board entity not found", http.StatusNotFound)
		return
	}

	writeSuccessResponse(w, board)
	return
}

// PutBoard handler
func (retro *RetroHandler) PutBoard(w http.ResponseWriter, r *http.Request) {
	logger.Info("Entering Put Board handler")
	boardTitle := r.FormValue("title")

	boardID, err := retro.Database.CreateBoard(boardTitle)
	if err != nil {
		writeErrorResponse(w, "Board could not be created", http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"board": boardID,
	}

	writeSuccessResponse(w, res)
}
