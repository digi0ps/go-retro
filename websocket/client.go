package websocket

import (
	"errors"
	"fmt"
	"go-retro/logger"

	"github.com/gorilla/websocket"
)

type client struct {
	hub     *Hub
	conn    *websocket.Conn
	boardID string
	send    chan []byte
}

func newClient(hub *Hub, conn *websocket.Conn, boardID string) *client {
	return &client{
		hub:     hub,
		conn:    conn,
		boardID: boardID,
		send:    make(chan []byte, 500),
	}
}

func (c *client) readWorker() {
	defer func() {
		fmt.Println("Unregistering client from readWorker.")
		c.hub.unregister <- boardArg{boardID: c.boardID, user: c}
		c.conn.Close()
	}()

	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			break
		}

		fmt.Printf("Message: %s, Type: %v", message, mt)
		arg := broadcastArg{
			boardID: c.boardID,
			message: message,
		}
		c.hub.broadcast <- arg
	}
}

func (c *client) writeWorker() {
	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				fmt.Println("Send channel closed")
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			fmt.Println("Broadcasting message to ", c)
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Println("Error ", err)
				return
			}

			if len(c.send) > 0 {
				logger.Error(errors.New("Unsent messages in channel"))
			}
		}
	}
}
