package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DataBaseConnection string
}

func GetEnv() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		DataBaseConnection: os.Getenv("DB_URI"),
	}

	return config

}
