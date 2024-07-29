package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DBUser        string
	DBPassword    string
	DBAddress     string
	DBName        string
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		ServerAddress: fmt.Sprintf("%s:%s", getEnv("HOST", "localhost"), getEnv("PORT", "8080")),
		DBUser:        getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", "password"),
		DBAddress:     fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:        getEnv("DB_NAME", "ecomm"),
	}
}

func getEnv(k, fallback string) string {
	v := os.Getenv(k)
	if v == "" {
		return fallback
	}
	return v
}
