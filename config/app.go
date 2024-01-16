package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBConnection           string `mapstructure:"DBConnection"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func GetEnv() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[Error] loading .env file encounter problem")
	}

	config := Config{
		DBConnection:  os.Getenv("DBConnection"),
		ServerAddress: os.Getenv("ServerAddress"),
	}

	return config

}
