package config

import (
	"log"

	"github.com/joho/godotenv"
)


func Load() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file, %v", err)
	}
}
