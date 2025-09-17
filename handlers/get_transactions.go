// handlers/get_transactions.go
package handlers

import (
	"encoding/json"
	"net/http"
	"sms-verifier/models"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := models.GetAllTransactions()
	if err != nil {
		http.Error(w, "Failed to fetch transactions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
