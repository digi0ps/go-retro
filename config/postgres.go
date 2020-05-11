package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// PostgresConfig contains config related to postgres
type PostgresConfig struct {
	host     string
	port     int
	user     string
	password string
	db       string
	sslmode  string
}

// GetAsString gets the config parsed in Go database format
func (config *PostgresConfig) GetAsString() (uri string) {
	uri = fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.host, config.port, config.user, config.db, config.password, config.sslmode,
	)

	return
}

func loadPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		host:     viper.GetString("POSTGRES_HOST"),
		port:     viper.GetInt("POSTGRES_PORT"),
		user:     viper.GetString("POSTGRES_USERNAME"),
		password: viper.GetString("POSTGRES_PASSWORD"),
		db:       viper.GetString("POSTGRES_DB"),
		sslmode:  viper.GetString("POSTGRES_SSLMODE"),
	}
}
