package handlers

import (
	"encoding/json"
	"net/http"
	"sms-verifier/models"
)

type VerifyRequest struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}

func VerifyDeposit(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	ok, err := models.VerifyTransaction(req.TransactionID, req.Amount)
	if err != nil {
		http.Error(w, "Verification failed", http.StatusInternalServerError)
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
