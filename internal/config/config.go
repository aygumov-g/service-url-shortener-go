package config

import (
	"os"
)

type Config struct {
	App AppConfig
	SCC ShortCodeConfig
	DB  DBConfig
}

type AppConfig struct {
	Port   string
	Domain string
}

type ShortCodeConfig struct {
	Alphabet string
	Secret   string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() *Config {
	return &Config{
		App: AppConfig{
			Port:   os.Getenv("APP_PORT"),
			Domain: os.Getenv("APP_DOMAIN"),
		},
		SCC: ShortCodeConfig{
			Alphabet: os.Getenv("SHORT_CODE_ALPHABET"),
			Secret:   os.Getenv("SHORT_CODE_SECRET"),
		},
		DB: DBConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Name:     os.Getenv("POSTGRES_DB"),
			SSLMode:  "disable",
		},
	}
}
