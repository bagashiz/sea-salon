package config

import (
	"github.com/joho/godotenv"
)

type (
	// Config holds the application configuration.
	Config struct {
		App *App
		DB  *DB
	}

	// App holds the application configuration.
	App struct {
		Name            string
		Env             string
		Host            string
		Port            string
		SessionLifetime string
	}

	// DB holds the database configuration.
	DB struct {
		Type     string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
)

// file is the name of the configuration file.
var file = ".env"

// New returns a new Config instance.
func New(getEnv func(string) string, path string) (*Config, error) {
	if getEnv("APP_ENV") != "production" {
		if err := godotenv.Load(path + file); err != nil {
			return nil, err
		}
	}

	return &Config{
		App: &App{
			Name:            getEnv("APP_NAME"),
			Env:             getEnv("APP_ENV"),
			Host:            getEnv("APP_HOST"),
			Port:            getEnv("APP_PORT"),
			SessionLifetime: getEnv("APP_SESSION_LIFETIME"),
		},
		DB: &DB{
			Type:     getEnv("DB_TYPE"),
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			Name:     getEnv("DB_NAME"),
		},
	}, nil
}
