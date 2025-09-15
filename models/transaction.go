package models

import (
	"context"
	"log"
	"sms-verifier/firebase"
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

// Save transaction to Firestore
func (t *Transaction) Save() error {
	ctx := context.Background()
	_, err := firebase.Client.Collection("transactions").Doc(t.TransactionID).Set(ctx, t)
	if err != nil {
		log.Println("Failed to save transaction:", err)
		return err
	}
	return nil
}

// Verify transaction
func VerifyTransaction(transactionID string, amount float64) (bool, error) {
	ctx := context.Background()
	doc, err := firebase.Client.Collection("transactions").Doc(transactionID).Get(ctx)
	if err != nil {
		return false, err
	}

	var t Transaction
	if err := doc.DataTo(&t); err != nil {
		return false, err
	}

	if t.Status == "pending" && t.Amount == amount {
		t.Status = "verified"
		_, err := firebase.Client.Collection("transactions").Doc(transactionID).Set(ctx, t)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
