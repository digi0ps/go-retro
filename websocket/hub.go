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
	user    *client
}

// Hub acts as a interface for boards. Clients can push to mentioned channels to perform actions.
type Hub struct {
	boards     map[string]*board
	register   chan boardArg
	unregister chan boardArg
	broadcast  chan broadcastArg
}

// NewHub creates a new hub with correct defaults
func NewHub() *Hub {
	return &Hub{
		boards:     make(map[string]*board),
		register:   make(chan boardArg),
		unregister: make(chan boardArg),
		broadcast:  make(chan broadcastArg),
	}
}

// Run Starts listening to channels
func (h *Hub) Run() {
	logger.Info("Starting hub...")

	for {
		select {
		case arg := <-h.register:
			h.registerUser(arg.boardID, arg.user)
		case arg := <-h.unregister:
			h.unregisterUser(arg.boardID, arg.user)
		case arg := <-h.broadcast:
			h.broadcastMsg(arg.boardID, arg.message, arg.user)
		}
	}
}

// registerUser adds an user to a board (creating board if not already present)
func (h *Hub) registerUser(boardID string, user *client) {
	logger.Info("Registering user for board")
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
	logger.Info("Removing user from board")
	if target, ok := h.boards[boardID]; ok {
		delete(target.users, user)
		close(user.send)
	} else {
		logger.Error(fmt.Errorf("Board Unregister: %s board not found", boardID))
	}
}

// broadcastMsg broadcasts a message to a board
func (h *Hub) broadcastMsg(boardID string, msg []byte, broadcastUser *client) {
	if target, ok := h.boards[boardID]; ok {
		for user := range target.users {
			if broadcastUser != user {
				select {
				case user.send <- msg:
				default:
					h.unregisterUser(boardID, user)
				}
			}
		}
	} else {
		logger.Error(fmt.Errorf("Board Unregister: %s board not found", boardID))
	}
}
