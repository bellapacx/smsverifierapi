package firebase

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	App    *firebase.App
	Client *firestore.Client
	once   sync.Once
)

// InitFirebase initializes Firebase App and Firestore client
func InitFirebase() {
	once.Do(func() {
		ctx := context.Background()

		credJSON := os.Getenv("FIREBASE_CRED_JSON")
		if credJSON == "" {
			log.Fatal("FIREBASE_CRED_JSON not set")
		}

		// Use the raw JSON as downloaded from Firebase Console
		opt := option.WithCredentialsJSON([]byte(credJSON))

		// Provide ProjectID in config
		config := &firebase.Config{
			ProjectID: "supermarket-pos-c8339",
		}

		app, err := firebase.NewApp(ctx, config, opt)
		if err != nil {
			log.Fatalf("Failed to initialize Firebase App: %v", err)
		}
		App = app

		client, err := App.Firestore(ctx)
		if err != nil {
			log.Fatalf("Failed to create Firestore client: %v", err)
		}
		Client = client

		log.Println("Firebase and Firestore initialized successfully")
	})
}

// CloseFirestore closes the Firestore client
func CloseFirestore() {
	if Client != nil {
		Client.Close()
	}
}
