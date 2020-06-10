package websocket

import (
	"fmt"
	"go-retro/logger"
)

type board struct {
	boardID string
	users   map[*client]bool
}

type boardArg struct {
	boardID string
	user    *client
}

type broadcastArg struct {
	boardID string
	message []byte
}

type Hub struct {
	boards     map[string]*board
	register   chan boardArg
	unregister chan boardArg
	broadcast  chan broadcastArg
}

func newHub() *Hub {
	return &Hub{
		boards:     make(map[string]*board),
		register:   make(chan boardArg),
		unregister: make(chan boardArg),
		broadcast:  make(chan broadcastArg),
	}
}

func (h *Hub) run() {
	for {
		select {
		case arg := <-h.register:
			h.registerUser(arg.boardID, arg.user)
		case arg := <-h.unregister:
			h.unregisterUser(arg.boardID, arg.user)
		case arg := <-h.broadcast:
			h.broadcastMsg(arg.boardID, arg.message)
		}
	}
}

// registerUser adds an user to a board (creating board if not already present)
func (h *Hub) registerUser(boardID string, user *client) {
	if target, ok := h.boards[boardID]; ok {
		target.users[user] = true
	} else {
		newUsers := map[*client]bool{
			user: true,
		}

		newBoard := &board{
			boardID: boardID,
			users:   newUsers,
		}

		h.boards[boardID] = newBoard
	}
}

// unregisterUser removes an user from a board (deleting board if no user present)
func (h *Hub) unregisterUser(boardID string, user *client) {
	if target, ok := h.boards[boardID]; ok {
		delete(target.users, user)
	} else {
		logger.Error(fmt.Errorf("Board Unregister: %s board not found", boardID))
	}
}

// broadcastMsg broadcasts a message to a board
func (h *Hub) broadcastMsg(boardID string, msg []byte) {
	if target, ok := h.boards[boardID]; ok {
		for user := range target.users {
			fmt.Println("Sending msg to: ", user)
			user.send <- msg
		}
	} else {
		logger.Error(fmt.Errorf("Board Unregister: %s board not found", boardID))
	}
}
