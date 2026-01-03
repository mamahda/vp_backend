package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
		return
	}

	log.Println(".env file loaded successfully")
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetJWT() string {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return "yoursecretkey"
	}
	return key;
}
