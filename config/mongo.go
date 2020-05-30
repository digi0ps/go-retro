package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// MongoConfig contains config related to Mongo
type MongoConfig struct {
	host     string
	port     int
	user     string
	password string
	db       string
	sslmode  string
}

// GetURI gets the config parsed in Go database format
func (config *MongoConfig) GetURI() (uri string) {
	uri = fmt.Sprintf("mongodb://%s:%d", config.host, config.port)

	return
}

// GetDatabase gets database name
func (config *MongoConfig) GetDatabase() string {
	return config.db
}

func loadMongoConfig() *MongoConfig {
	return &MongoConfig{
		host: viper.GetString("MONGO_HOST"),
		port: viper.GetInt("MONGO_PORT"),
		db:   viper.GetString("MONGO_DB"),
	}
}
