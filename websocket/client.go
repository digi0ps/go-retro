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
	hub           *Hub
	conn          *websocket.Conn
	actionHandler func(e Event) error
	boardID       string
	send          chan []byte
}

func newClient(hub *Hub, conn *websocket.Conn, boardID string, ah func(e Event) error) *client {
	return &client{
		hub:           hub,
		conn:          conn,
		actionHandler: ah,
		boardID:       boardID,
		send:          make(chan []byte, 500),
	}
}

func (c *client) readWorker() {
	defer func() {
		logger.Info("Unregsitering client from readWorker")
		c.hub.unregister <- boardArg{boardID: c.boardID, user: c}
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			break
		}

		event, err := unmarshallMessage(message)
		if err != nil {
			logger.Error(fmt.Errorf("Message not right format"))
			continue
		}

		logger.Debug(event)

		arg := broadcastArg{
			boardID: c.boardID,
			message: message,
			user:    c,
		}
		c.hub.broadcast <- arg

		if err := c.actionHandler(event); err != nil {
			wrapErr := errors.New("[ACTION] " + err.Error())
			logger.Error(wrapErr)
		}
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
