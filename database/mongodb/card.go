package mongodb

import (
	"context"
	"go-retro/database"
	"go-retro/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateCard creates a new card in a column
func (mdb *MongoDatabase) CreateCard(boardID, columnID, content string) (string, error) {
	ctx := context.TODO()

	newCard := database.Card{
		ID:        primitive.NewObjectID(),
		Content:   content,
		CreatedAt: time.Now().Unix(),
	}

	boardIDHex, _ := primitive.ObjectIDFromHex(boardID)
	columnIDHex, _ := primitive.ObjectIDFromHex(columnID)

	filterBson := bson.M{"_id": boardIDHex, "columns._id": columnIDHex}
	updateBson := bson.M{"$addToSet": bson.M{"columns.$.cards": newCard}}

	result := mdb.db.Collection("boards").FindOneAndUpdate(ctx, filterBson, updateBson)
	if result.Err() != nil {
		logger.Error(result.Err())
	}

	return newCard.ID.Hex(), result.Err()
}

// UpdateCard updates a card's name
func (mdb *MongoDatabase) UpdateCard(boardID, columnID, cardID, newContent string) error {
	ctx := context.TODO()

	boardIDHex, _ := primitive.ObjectIDFromHex(boardID)
	columnIDHex, _ := primitive.ObjectIDFromHex(columnID)
	cardIDHex, _ := primitive.ObjectIDFromHex(cardID)

	filterBson := bson.M{"_id": boardIDHex, "columns._id": columnIDHex}
	updateBson := bson.M{"$set": bson.M{"columns.$.cards.$[c].content": newContent}}
	opts := options.FindOneAndUpdate().SetArrayFilters(
		options.ArrayFilters{
			Filters: []interface{}{bson.M{"c._id": cardIDHex}},
		},
	)

	result := mdb.db.Collection("boards").FindOneAndUpdate(ctx, filterBson, updateBson, opts)
	if result.Err() != nil {
		logger.Error(result.Err())
	}

	return result.Err()
}

// DeleteCard deletes a card
func (mdb *MongoDatabase) DeleteCard(boardID, columnID, cardID string) error {
	ctx := context.TODO()

	boardIDHex, _ := primitive.ObjectIDFromHex(boardID)
	columnIDHex, _ := primitive.ObjectIDFromHex(columnID)
	cardIDHex, _ := primitive.ObjectIDFromHex(cardID)

	filterBson := bson.M{"_id": boardIDHex, "columns._id": columnIDHex}
	updateBson := bson.M{"$pull": bson.M{"columns.$.cards": bson.M{"_id": cardIDHex}}}

	result := mdb.db.Collection("boards").FindOneAndUpdate(ctx, filterBson, updateBson)
	if result.Err() != nil {
		logger.Error(result.Err())
	}

	return result.Err()
}
