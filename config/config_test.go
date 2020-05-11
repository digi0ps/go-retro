package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testConfigVars = map[string]string{
	"SERVER_PORT": "3000",

	"POSTGRES_HOST":     "psql",
	"POSTGRES_PORT":     "5432",
	"POSTGRES_USERNAME": "admin",
	"POSTGRES_PASSWORD": "password",
	"POSTGRES_DB":       "test",
	"POSTGRES_SSLMODE":  "enable",
}

func TestLoadConfig(t *testing.T) {
	for k, v := range testConfigVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	LoadConfig()

	assert := assert.New(t)

	t.Run("Test server config", func(t *testing.T) {
		assert.IsType(&ServerConfig{}, Server())
		assert.Equal(Server().GetPort(), 3000)
	})

	t.Run("Test database config", func(t *testing.T) {
		assert.IsType(&PostgresConfig{}, Postgres())

		expectedURI := "host=psql port=5432 user=admin dbname=test password=password sslmode=enable"
		assert.Equal(expectedURI, Postgres().GetAsString())
	})
}
