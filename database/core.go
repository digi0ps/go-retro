package database

import (
	"context"
	"errors"
	"go-retro/config"
	"go-retro/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrorNotFound = errors.New("Item not found")
)

// OpenMongoConnection gives mongo connection
func OpenMongoConnection() (*mongo.Database, func()) {
	uri := config.Mongo().GetURI()
	db := config.Mongo().GetDatabase()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error(err)
	}

	logger.Info("Connected to MongoDB")

	closeConnection := func() {
		err := client.Disconnect(ctx)
		if err != nil {
			logger.Error(err)
		}
	}

	return client.Database(db), closeConnection
}
