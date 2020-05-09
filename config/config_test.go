package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testConfigVars = map[string]string{
	"SERVER_PORT": "3000",
}

func TestLoadConfig(t *testing.T) {
	for k, v := range testConfigVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	LoadConfig()

	t.Run("test server config", func(t *testing.T) {
		assert.Equal(t, Server().GetPort(), 3000)
	})
}
