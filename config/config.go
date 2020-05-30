package config

import (
	"github.com/spf13/viper"
)

// Config structure to store configuration
type Config struct {
	server *ServerConfig
	mongo  *MongoConfig
}

var appConfig *Config

// Server config
func Server() *ServerConfig {
	return appConfig.server
}

// Mongo config
func Mongo() *MongoConfig {
	return appConfig.mongo
}

// LoadConfig returns a config
func LoadConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	appConfig = &Config{
		server: loadServerConfig(),
		mongo:  loadMongoConfig(),
	}
}
