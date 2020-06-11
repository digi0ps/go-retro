package websocket

import (
	"fmt"
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
}

// InitHandler is responsible for initialising websocket handler
func (s *SocketHandler) InitHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err)
		// Move error to middleware
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server error happened")
		return
	}

	boardID := mux.Vars(r)["board"]
	user := newClient(s.BoardHub, conn, boardID)

	args := boardArg{
		boardID: boardID,
		user:    user,
	}

	s.BoardHub.register <- args

	go user.readWorker()
	go user.writeWorker()
}
