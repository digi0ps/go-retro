package websocket

import (
	"errors"
	"fmt"
	"go-retro/logger"
	"time"

	"github.com/gorilla/websocket"
)

var (
	writeWait = 5 * time.Second
	readWait  = 60 * time.Second
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
		logger.Info("Unregsitering client from readWorker")
		c.hub.unregister <- boardArg{boardID: c.boardID, user: c}
		c.conn.Close()
	}()

	for {
		c.conn.SetReadDeadline(time.Now().Add(readWait))
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			break
		}

		arg := broadcastArg{
			boardID: c.boardID,
			message: message,
			user:    c,
		}
		c.hub.broadcast <- arg
	}
}

func (c *client) writeWorker() {
	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				logger.Info("Closing send channel")
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

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
