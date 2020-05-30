package database

import (
	"context"
	"go-retro/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateBoard creates a board using the name given
func CreateBoard(db *mongo.Database, name string) (string, error) {
	ctx := context.TODO()

	newBoard := Board{
		ID:        primitive.NewObjectID(),
		Title:     name,
		Columns:   []Column{},
		CreatedAt: time.Now().UnixNano(),
	}

	insertResult, err := db.Collection("boards").InsertOne(ctx, newBoard)

	if err != nil {
		logger.Error(err)
	}

	objID := insertResult.InsertedID.(primitive.ObjectID)

	return objID.Hex(), err
}

// FindBoard finds a board by the ID taken in the argument
func FindBoard(db *mongo.Database, id string) (board Board, err error) {
	ctx := context.TODO()
	idHex, _ := primitive.ObjectIDFromHex(id)

	cursor := db.Collection("boards").FindOne(ctx, bson.M{"_id": idHex})
	err = cursor.Decode(&board)

	if err != nil {
		// Report: Board not found in database
		err = ErrorNotFound
		return
	}

	return board, nil
}

// DeleteBoard deletes a board by the string
func DeleteBoard(db *mongo.Database, id string) error {
	ctx := context.TODO()

	idHex, _ := primitive.ObjectIDFromHex(id)

	_, err := db.Collection("boards").DeleteOne(ctx, bson.M{"_id": idHex})

	return err
}
