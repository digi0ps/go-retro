package mongodb

import (
	"context"
	"go-retro/config"
	"go-retro/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDatabase holds the methods for interacting with MongoDB
type MongoDatabase struct {
	db *mongo.Database
}

// OpenConnection gives mongo connection
func (mdb *MongoDatabase) OpenConnection() error {
	uri := config.Mongo().GetURI()
	db := config.Mongo().GetDatabase()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error(err)
		return err
	}

	mdb.db = client.Database(db)
	return nil
}

// CloseConnection closes the connection
func (mdb *MongoDatabase) CloseConnection() error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	err := mdb.db.Client().Disconnect(ctx)

	return err
}
