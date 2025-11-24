package config

import (
	"github.com/spf13/viper"
)

type Redis struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     int    `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

func LoadRedis(v *viper.Viper) Redis {
	host := v.GetString("REDIS_HOST")
	if host == "" {
		panic("REDIS_HOST is required but not set")
	}

	port := v.GetInt("REDIS_PORT")
	if port == 0 {
		panic("REDIS_PORT is required but not set")
	}

	password := v.GetString("REDIS_PASSWORD")
	// Password can be empty for development, so no validation

	db := v.GetInt("REDIS_DB")
	// DB defaults to 0 if not set, which is valid

	return Redis{
		Host:     host,
		Port:     port,
		Password: password,
		DB:       db,
	}
}
