package database

import (
	"context"
	"go-retro/config"
	"go-retro/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

type Board struct {
	Name      string
	CreatedAt int64
}

func AddBoard(db *mongo.Database, name string) error {
	ctx := context.TODO()

	newBoard := Board{
		Name:      name,
		CreatedAt: time.Now().UnixNano(),
	}

	insertResult, err := db.Collection("boards").InsertOne(ctx, newBoard)

	if err != nil {
		logger.Error(err)
	}

	logger.Debug(insertResult.InsertedID)
	return err
}
