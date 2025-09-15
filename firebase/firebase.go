package firebase

import (
	"context"
	"log"
	"os"
	"sms-verifier/config"

	"cloud.google.com/go/firestore"

	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitFirebase() {
	ctx := context.Background()

	// Read Firebase credentials from environment variable
	jsonCred := os.Getenv("FIREBASE_CRED_JSON")
	if jsonCred == "" {
		log.Fatal("FIREBASE_CRED_JSON not set")
	}

	// Initialize Firestore with credentials from JSON string
	opt := option.WithCredentialsJSON([]byte(jsonCred))
	client, err := firestore.NewClient(ctx, config.FirebaseProjectID, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}

	Client = client
	log.Println("Firestore initialized successfully")
}
