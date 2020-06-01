package websocket

import (
	"encoding/json"
	"fmt"
	"go-retro/logger"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout: 3 * time.Second,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
}

// InitHandler is responsible for initialising websocket handler
func InitHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err)
		// Move error to middleware
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server error happened")
		return
	}

	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			return
		}

		type wsmsg struct {
			Name string `json:"name"`
		}

		var jsonMsg wsmsg

		if err := json.Unmarshal(message, &jsonMsg); err != nil {
			logger.Error(err)
			return
		}

		fmt.Printf("Message: %s, Type: %v", jsonMsg, mt)

		if err := conn.WriteMessage(mt, message); err != nil {
			logger.Error(err)
			return
		}
	}

}
