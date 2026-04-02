package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB DbConfig
}

type DbConfig struct {
	DSN string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("no config file found")
	}
	return &Config{
		DB: DbConfig{
			DSN: os.Getenv("DB_DSN"),
		},
	}
}
