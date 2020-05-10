package config

import (
	"github.com/spf13/viper"
)

// ServerConfig contains config related to server
type ServerConfig struct {
	port int
}

// GetPort gets port
func (s *ServerConfig) GetPort() int {
	return s.port
}

func loadServerConfig() *ServerConfig {
	return &ServerConfig{
		port: viper.GetInt("SERVER_PORT"),
	}
}
