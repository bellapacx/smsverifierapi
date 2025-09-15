package routes

import (
	"sms-verifier/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/sms", handlers.ReceiveSMS).Methods("POST")
	r.HandleFunc("/api/verify-deposit", handlers.VerifyDeposit).Methods("POST")
}
