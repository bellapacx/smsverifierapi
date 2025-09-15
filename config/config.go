package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ServerPort        string
	FirebaseProjectID string
	FirebaseCred      string
	APIKey            string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	ServerPort = os.Getenv("SERVER_PORT")
	FirebaseProjectID = os.Getenv("FIREBASE_PROJECT_ID")
	FirebaseCred = os.Getenv("FIREBASE_CRED")
	APIKey = os.Getenv("API_KEY")

	if ServerPort == "" || FirebaseProjectID == "" || FirebaseCred == "" {
		log.Fatal("Missing required environment variables")
	}
}
