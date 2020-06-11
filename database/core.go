package database

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// ErrorNotFound is thrown when an item is not found
	ErrorNotFound = errors.New("Item not found")
)

// Card represents the card in the board
type Card struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt int64              `bson:"created_at" json:"created_at"`
}

// Column represents a group of cards displayed vertically in UI
type Column struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Cards     []Card             `bson:"cards" json:"cards"`
	CreatedAt int64              `bson:"created_at" json:"created_at"`
}

// Board represents an entire group of cards where people can collaborate
type Board struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Columns   []Column           `bson:"columns" json:"columns"`
	CreatedAt int64              `bson:"created_at" json:"created_at"`
}

// Service interface contains all the database methods used by this project
type Service interface {
	// Connection
	OpenConnection() error
	CloseConnection() error
	// Board
	CreateBoard(boardName string) (string, error)
	FindBoard(boardID string) (board Board, err error)
	DeleteBoard(boardID string) error
	// Column
	CreateColumn(boardID, columnName string) (string, error)
	UpdateColumn(boardID, columnID, newName string) error
	DeleteColumn(boardID, columnID string) error
	// Cards
	CreateCard(boardID, columnID, content string) (string, error)
	UpdateCard(boardID, columnID, cardID, newContent string) error
	DeleteCard(boardID, columnID, cardID string) error
}
