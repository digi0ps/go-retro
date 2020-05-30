package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt int64              `bson:"created_at", json:"created_at"`
}

type Column struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Cards     []Card             `bson:"cards" json:"cards"`
	CreatedAt int64              `bson:"created_at", json:"created_at"`
}

type Board struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Columns   []Column           `bson:"columns" json:"columns"`
	CreatedAt int64              `bson:"created_at", json:"created_at"`
}

// DatabaseService interface contains all the database methods used by this project
type DatabaseService interface {
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
