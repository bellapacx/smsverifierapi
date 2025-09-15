package handlers

import (
	"encoding/json"
	"net/http"
	"sms-verifier/models"
	"sms-verifier/utils"
)

type VerifyRequest struct {
	Body string `json:"body"` // the full SMS text
}

func VerifyDeposit(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Parse the SMS
	parsed, err := utils.ParseVerifySMS(req.Body)
	if err != nil {
		http.Error(w, "Failed to parse SMS: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verify the transaction in Firestore
	ok, err := models.VerifyTransaction(parsed.TransactionID, parsed.Amount)
	if err != nil {
		http.Error(w, "Verification failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	status := "failed"
	if ok {
		status = "success"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": status,
	})
}
