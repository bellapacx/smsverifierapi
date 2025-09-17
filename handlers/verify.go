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

type VerifyResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func VerifyDeposit(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	var resp VerifyResponse

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp.Status = "failed"
		resp.Message = "Invalid request: " + err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Use multi-bank parser
	parsed, err := utils.ParseBankMulti(req.Body)
	if err != nil || parsed == nil {
		resp.Status = "failed"
		resp.Message = "Failed to parse SMS: " + err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Verify the transaction in Firestore
	ok, err := models.VerifyTransaction(parsed.TransactionID, parsed.Amount)
	if err != nil {
		resp.Status = "failed"
		resp.Message = "Verification failed: " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	if ok {
		resp.Status = "success"
		resp.Message = "Transaction verified successfully"
	} else {
		resp.Status = "failed"
		resp.Message = "Transaction verification failed"
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
