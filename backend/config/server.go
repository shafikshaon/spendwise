package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Server struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ServerPort    int    `mapstructure:"SERVER_PORT"`
	ServerMode    string `mapstructure:"SERVER_MODE"`
}

func LoadServer(v *viper.Viper) Server {
	address := v.GetString("SERVER_ADDRESS")
	if address == "" {
		panic("SERVER_ADDRESS is required but not set")
	}

	port := v.GetInt("SERVER_PORT")
	if port == 0 {
		panic("SERVER_PORT is required but not set")
	}

	mode := v.GetString("SERVER_MODE")
	if mode == "" {
		panic("SERVER_MODE is required but not set")
	}

	// Validate mode
	if mode != "dev" && mode != "prod" && mode != "test" {
		panic(fmt.Sprintf("SERVER_MODE must be 'dev', 'prod', or 'test', got: %s", mode))
	}

	return Server{
		ServerAddress: address,
		ServerPort:    port,
		ServerMode:    mode,
	}
}
