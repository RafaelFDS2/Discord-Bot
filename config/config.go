package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environments struct {
	ApiPort string `mapstructure:"API_PORT" envconfig:"API_PORT"`
	DSN     string `mapstructure:"DSN" envconfig:"DSN"`
}

// LoadEnvVars load the environment variables
func LoadEnvVars() (*Environments, error) {
	godotenv.Load()
	c := &Environments{}
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}
