package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sms-verifier/models"
	"sms-verifier/utils"
	"time"
)

type IncomingSMS struct {
	Body      string `json:"body"`
	Sender    string `json:"sender"`
	Timestamp int64  `json:"timestamp"`
}

func ReceiveSMS(w http.ResponseWriter, r *http.Request) {
	var sms IncomingSMS
	if err := json.NewDecoder(r.Body).Decode(&sms); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Use multi-bank parser
	parsed, err := utils.ParseBankMulti(sms.Body)
	if err != nil || parsed == nil {
		http.Error(w, "Failed to parse SMS: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Save transaction
	tx := &models.Transaction{
		TransactionID: parsed.TransactionID,
		Account:       parsed.Account,
		SenderName:    parsed.SenderName,
		Amount:        parsed.Amount,
		Date:          parsed.Date.Format(time.RFC3339),
		Balance:       parsed.Balance,
		RawSMS:        sms.Body,
		Status:        "pending",
	}

	if err := tx.Save(); err != nil {
		log.Println("Failed to save transaction:", err)
		http.Error(w, "Failed to save transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":         "ok",
		"transaction_id": tx.TransactionID,
	})
}
