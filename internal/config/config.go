package config

import (
	"os"
)

type Config struct {
	App App
	SCC ShortCode
	DB  Postgres
}

func Load() *Config {
	return &Config{
		App: App{
			Port:   os.Getenv("APP_PORT"),
			Domain: os.Getenv("APP_DOMAIN"),
		},
		SCC: ShortCode{
			Alphabet: os.Getenv("SHORT_CODE_ALPHABET"),
			Secret:   os.Getenv("SHORT_CODE_SECRET"),
		},
		DB: Postgres{
			dBHost:     os.Getenv("POSTGRES_HOST"),
			dBUser:     os.Getenv("POSTGRES_USER"),
			dBPassword: os.Getenv("POSTGRES_PASSWORD"),
			dBName:     os.Getenv("POSTGRES_DB"),
		},
	}
}
