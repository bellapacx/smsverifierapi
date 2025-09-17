package models

import (
	"context"
	"log"
	"sms-verifier/firebase"

	"google.golang.org/api/iterator"
)

type Transaction struct {
	TransactionID string  `json:"transaction_id"`
	Account       string  `json:"account"`
	SenderName    string  `json:"sender_name"`
	Amount        float64 `json:"amount"`
	Date          string  `json:"date"`
	Balance       float64 `json:"balance"`
	RawSMS        string  `json:"raw_sms"`
	Status        string  `json:"status"` // pending / verified
}

// Save saves the transaction to Firestore
func (t *Transaction) Save() error {
	ctx := context.Background()
	client, err := firebase.App.Firestore(ctx)
	if err != nil {
		log.Println("Failed to create Firestore client:", err)
		return err
	}
	defer client.Close()

	_, err = client.Collection("transactions").Doc(t.TransactionID).Set(ctx, t)
	if err != nil {
		log.Println("Failed to save transaction:", err)
		return err
	}
	return nil
}

// VerifyTransaction verifies a transaction by ID and amount
func VerifyTransaction(transactionID string, amount float64) (bool, error) {
	ctx := context.Background()
	client, err := firebase.App.Firestore(ctx)
	if err != nil {
		return false, err
	}
	defer client.Close()

	doc, err := client.Collection("transactions").Doc(transactionID).Get(ctx)
	if err != nil {
		return false, err
	}

	var t Transaction
	if err := doc.DataTo(&t); err != nil {
		return false, err
	}

	if t.Status == "pending" && t.Amount == amount {
		t.Status = "verified"
		_, err := client.Collection("transactions").Doc(transactionID).Set(ctx, t)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

// models/transaction.go
func GetAllTransactions() ([]Transaction, error) {
	var transactions []Transaction
	ctx := context.Background()

	client, err := firebase.App.Firestore(ctx)
	if err != nil {
		log.Println("Failed to create Firestore client:", err)
		return nil, err
	}
	defer client.Close()

	iter := client.Collection("transactions").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("Failed to iterate transactions:", err)
			return nil, err
		}

		var tx Transaction
		if err := doc.DataTo(&tx); err != nil {
			log.Println("Failed to parse transaction:", err)
			continue
		}
		transactions = append(transactions, tx)
	}

	return transactions, nil
}
func TransactionExists(transactionID string) (bool, error) {
	ctx := context.Background()
	client, err := firebase.App.Firestore(ctx)
	if err != nil {
		return false, err
	}
	defer client.Close()

	doc, err := client.Collection("transactions").Doc(transactionID).Get(ctx)
	if err != nil {
		return false, nil
	}

	return doc.Exists(), nil
}
