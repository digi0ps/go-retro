package database

import (
	"context"
	"go-retro/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateColumn creates a column in a board
func CreateColumn(db *mongo.Database, boardID, columnName string) (string, error) {
	ctx := context.TODO()

	boardIDHex, _ := primitive.ObjectIDFromHex(boardID)

	newColumn := Column{
		ID:        primitive.NewObjectID(),
		Name:      columnName,
		Cards:     []Card{},
		CreatedAt: time.Now().UnixNano(),
	}

	filterBson := bson.M{"_id": boardIDHex}
	updateBson := bson.M{"$addToSet": bson.M{"columns": newColumn}}
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := db.Collection("boards").FindOneAndUpdate(ctx, filterBson, updateBson, &opts)
	if result.Err() != nil {
		logger.Error(result.Err())
	}

	id := newColumn.ID.Hex()

	return id, result.Err()
}

// UpdateColumn updates a column's name
func UpdateColumn(db *mongo.Database, boardID string, columnID string, newName string) error {
	ctx := context.TODO()

	boardIDHex, _ := primitive.ObjectIDFromHex(boardID)
	columnIDHex, _ := primitive.ObjectIDFromHex(columnID)

	filterBson := bson.M{"_id": boardIDHex, "columns._id": columnIDHex}
	updateBson := bson.M{"$set": bson.M{"columns.$.name": newName}}

	result := db.Collection("boards").FindOneAndUpdate(ctx, filterBson, updateBson)
	if result.Err() != nil {
		logger.Error(result.Err())
	}

	return result.Err()
}

// DeleteColumn deletes a column with the boardID and columnID
func DeleteColumn(db *mongo.Database, boardID string, columnID string) error {
	ctx := context.TODO()

	boardIDHex, _ := primitive.ObjectIDFromHex(boardID)
	columnIDHex, _ := primitive.ObjectIDFromHex(columnID)

	filterBson := bson.M{"_id": boardIDHex}
	updateBson := bson.M{"$pull": bson.M{"columns": bson.M{"_id": columnIDHex}}}

	result := db.Collection("boards").FindOneAndUpdate(ctx, filterBson, updateBson)
	if result.Err() != nil {
		logger.Error(result.Err())
	}

	return result.Err()
}
