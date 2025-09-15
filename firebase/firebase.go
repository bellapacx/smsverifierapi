package firebase

import (
	"context"
	"fmt"
	"log"
	"sms-verifier/config"

	"cloud.google.com/go/firestore"

	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitFirebase() {
	ctx := context.Background()
	sa := option.WithCredentialsFile(config.FirebaseCred)

	client, err := firestore.NewClient(ctx, config.FirebaseProjectID, sa)
	if err != nil {
		log.Fatalf("Error initializing Firestore: %v", err)
	}

	Client = client
	fmt.Println("Firestore initialized successfully")
}
