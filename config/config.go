package config

import "github.com/spf13/viper"

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
	viper.SetConfigFile("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	appConfig = &Config{
		server: loadServerConfig(),
	}
}
