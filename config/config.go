package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config structure to store configuration
type Config struct {
	server *ServerConfig
}

var appConfig *Config

// Server config
func Server() *ServerConfig {
	return appConfig.server
}

// LoadConfig returns a config
func LoadConfig() {
	viper.SetConfigFile("application.yml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Error reading config file"))
	}

	appConfig = &Config{
		server: loadServerConfig(),
	}
}
