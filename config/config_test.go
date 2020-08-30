package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testConfigVars = map[string]string{
	"SERVER_PORT": "3000",

	"MONGO_HOST":     "local",
	"MONGO_PORT":     "2222",
	"MONGO_DB":       "test",
	"MONGO_USER":     "admin",
	"MONGO_PASSWORD": "password",
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
		assert.IsType(&MongoConfig{}, Mongo())

		expectedURI := "mongodb://admin:password@local:2222"
		assert.Equal(expectedURI, Mongo().GetURI())

		assert.Equal("test", Mongo().GetDatabase())
	})
}
