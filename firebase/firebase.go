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

	opt := option.WithCredentialsJSON([]byte(credJSON))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}
	App = app
	log.Println("Firebase initialized successfully")
}
