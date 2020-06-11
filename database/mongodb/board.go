package mongodb

import (
	"context"
	"go-retro/database"
	"go-retro/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateBoard creates a board using the name given
func (mdb *MongoDatabase) CreateBoard(name string) (string, error) {
	ctx := context.TODO()

	newBoard := database.Board{
		ID:        primitive.NewObjectID(),
		Title:     name,
		Columns:   []database.Column{},
		CreatedAt: time.Now().UnixNano(),
	}

	insertResult, err := mdb.db.Collection("boards").InsertOne(ctx, newBoard)

	if err != nil {
		logger.Error(err)
	}

	objID := insertResult.InsertedID.(primitive.ObjectID)

	return objID.Hex(), err
}

// FindBoard finds a board by the ID taken in the argument
func (mdb *MongoDatabase) FindBoard(id string) (board database.Board, err error) {
	ctx := context.TODO()
	idHex, _ := primitive.ObjectIDFromHex(id)

	cursor := mdb.db.Collection("boards").FindOne(ctx, bson.M{"_id": idHex})

	err = cursor.Decode(&board)
	if err != nil {
		// Report: Board not found in database
		err = database.ErrorNotFound
	}

	return
}

// DeleteBoard deletes a board by the string
func (mdb *MongoDatabase) DeleteBoard(id string) error {
	ctx := context.TODO()

	idHex, _ := primitive.ObjectIDFromHex(id)

	_, err := mdb.db.Collection("boards").DeleteOne(ctx, bson.M{"_id": idHex})

	return err
}
