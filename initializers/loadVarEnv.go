package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadVarEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
