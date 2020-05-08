package config

import "testing"

func TestGetConfig(t *testing.T) {
	config := GetConfig()

	t.Run("getPort should get port", func(t *testing.T) {
		if expected := config.port; config.GetPort() != expected {
			t.Error("Config GetPort Failed.")
		}
	})

	t.Run("getEnv should get env", func(t *testing.T) {
		if expected := config.env; config.GetEnv() != expected {
			t.Error("Config getEnv Failed.")
		}
	})
}
