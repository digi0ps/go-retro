package config

import (
	"github.com/spf13/viper"
)

// Config structure to store configuration
type Config struct {
	server   *ServerConfig
	postgres *PostgresConfig
}

var appConfig *Config

// Server config
func Server() *ServerConfig {
	return appConfig.server
}

// Postgres config
func Postgres() *PostgresConfig {
	return appConfig.postgres
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
		server:   loadServerConfig(),
		postgres: loadPostgresConfig(),
	}
}
