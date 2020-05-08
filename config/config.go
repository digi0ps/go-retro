package config

// Config structure to store configuration
type Config struct {
	port int
	env  string
}

// GetPort port
func (c *Config) GetPort() int {
	return c.port
}

// GetEnv env
func (c *Config) GetEnv() string {
	return c.env
}

// GetConfig returns a config
func GetConfig() *Config {
	return &Config{
		port: 8080,
		env:  "dev",
	}
}
