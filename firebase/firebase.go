package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	ctx := context.Background()

	credJSON := os.Getenv("FIREBASE_CRED_JSON")
	if credJSON == "" {
		log.Fatal("FIREBASE_CRED_JSON not set")
	}

	// Create config with project ID
	config := &firebase.Config{
		ProjectID: "supermarket-pos-c8339",
	}

	opt := option.WithCredentialsJSON([]byte(credJSON))
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	App = app
	log.Println("Firebase initialized successfully")
}
