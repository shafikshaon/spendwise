package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Database struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	SSLMode  string `mapstructure:"DB_SSLMODE"`
}

func LoadDatabase(v *viper.Viper) Database {
	host := v.GetString("DB_HOST")
	if host == "" {
		panic("DB_HOST is required but not set")
	}

	port := v.GetInt("DB_PORT")
	if port == 0 {
		panic("DB_PORT is required but not set")
	}

	user := v.GetString("DB_USER")
	if user == "" {
		panic("DB_USER is required but not set")
	}

	password := v.GetString("DB_PASSWORD")
	if password == "" {
		panic("DB_PASSWORD is required but not set")
	}

	name := v.GetString("DB_NAME")
	if name == "" {
		panic("DB_NAME is required but not set")
	}

	sslMode := v.GetString("DB_SSLMODE")
	if sslMode == "" {
		panic("DB_SSLMODE is required but not set")
	}

	// Validate SSL mode
	validSSLModes := []string{"disable", "require", "verify-ca", "verify-full"}
	valid := false
	for _, mode := range validSSLModes {
		if sslMode == mode {
			valid = true
			break
		}
	}
	if !valid {
		panic(fmt.Sprintf("DB_SSLMODE must be one of %v, got: %s", validSSLModes, sslMode))
	}

	return Database{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
		SSLMode:  sslMode,
	}
}
