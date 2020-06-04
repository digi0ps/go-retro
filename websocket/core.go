package websocket

import (
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

type hub struct {
	board string
	users []*client
}

type client struct {
	board *hub
	id    string
	conn  *websocket.Conn
}

func (c *client) close() {
	index := -1

	for i, user := range c.board.users {
		if user.id == c.id {
			index = i
		}
	}

	c.board.users = append(c.board.users[:index], c.board.users[index+1:]...)

	c.conn.Close()
}

var mainHub = &hub{
	board: "main",
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

	current := &client{
		board: mainHub,
		id:    fmt.Sprintf("%d", time.Now().Unix()),
		conn:  conn,
	}

	fmt.Println(current)

	mainHub.users = append(mainHub.users, current)

	defer current.close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			break
		}

		type wsmsg struct {
			Name string `json:"name"`
		}

		// var jsonMsg wsmsg

		// if err := json.Unmarshal(message, &jsonMsg); err != nil {
		// 	logger.Error(err)
		// 	return
		// }

		fmt.Printf("Message: %s, Type: %v", message, mt)

		for _, userConn := range current.board.users {
			if userConn.id == current.id {
				continue
			}

			msg := fmt.Sprintf("user#%v: %s", current.id, message)

			go func(conn *websocket.Conn, msg string) {
				b := []byte(msg)
				if err := conn.WriteMessage(mt, b); err != nil {
					logger.Error(err)
				}
			}(userConn.conn, msg)
		}
	}

}
