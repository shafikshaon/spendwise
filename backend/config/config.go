package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Database Database
}

func Load() (*Config, error) {
	v := viper.New()

	// Support both .env file and environment variables
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	// Read .env if it exists (ignore error for K8s deployment)
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	cfg := &Config{
		Server:   LoadServer(v),
		Database: LoadDatabase(v),
	}

	return cfg, nil
}
