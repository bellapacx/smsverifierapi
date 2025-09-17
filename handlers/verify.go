package handlers

import (
	"encoding/json"
	"net/http"
	"sms-verifier/models"
	"sms-verifier/utils"
)

type VerifyRequest struct {
	Body string `json:"body"` // full SMS text
}

type VerifyResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func VerifyDeposit(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	var resp VerifyResponse

	w.Header().Set("Content-Type", "application/json")

	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp.Status = "failed"
		resp.Message = "Invalid request: " + err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Parse SMS
	parsed, err := utils.ParseBankMulti(req.Body)
	if err != nil || parsed == nil {
		resp.Status = "failed"
		resp.Message = "Failed to parse SMS: " + err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Check if transaction exists in Firestore
	exists, err := models.TransactionExists(parsed.TransactionID)
	if err != nil {
		resp.Status = "failed"
		resp.Message = "Error checking transaction: " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	if exists {
		resp.Status = "success"
		resp.Message = "Transaction exists"
	} else {
		resp.Status = "failed"
		resp.Message = "Transaction not found"
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
