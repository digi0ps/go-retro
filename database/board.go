package database

import (
	"context"
	"go-retro/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Board struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	CreatedAt int64              `bson:"created_at"`
}

func AddBoard(db *mongo.Database, name string) error {
	ctx := context.TODO()

	newBoard := Board{
		ID:        primitive.NewObjectID(),
		Title:     name,
		CreatedAt: time.Now().UnixNano(),
	}

	insertResult, err := db.Collection("boards").InsertOne(ctx, newBoard)

	if err != nil {
		logger.Error(err)
	}

	logger.Debug(insertResult.InsertedID)
	return err
}

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

func DeleteBoard(db *mongo.Database, id string) error {
	ctx := context.TODO()

	idHex, _ := primitive.ObjectIDFromHex(id)

	_, err := db.Collection("boards").DeleteOne(ctx, bson.M{"_id": idHex})

	return err
}
