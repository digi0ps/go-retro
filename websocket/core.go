package websocket

import (
	"fmt"
	"go-retro/database"
	"go-retro/logger"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout: 3 * time.Second,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
}

// SocketHandler implements handler for websocket
type SocketHandler struct {
	BoardHub *Hub
	Database database.Service
}

// InitHandler is responsible for initialising websocket handler
func (s *SocketHandler) InitHandler(w http.ResponseWriter, r *http.Request) {
	boardID := mux.Vars(r)["board"]

	_, err := s.Database.FindBoard(boardID)
	if err == database.ErrorNotFound {
		http.NotFound(w, r)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server error happened")
		return
	}

	user := newClient(s.BoardHub, conn, boardID)

	args := boardArg{
		boardID: boardID,
		user:    user,
	}

	s.BoardHub.register <- args

	go user.readWorker()
	go user.writeWorker()
}
