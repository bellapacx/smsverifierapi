package main

import (
	"log"
	"net/http"
	"sms-verifier/config"
	"sms-verifier/firebase"
	"sms-verifier/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Load config & init Firebase
	config.LoadConfig()
	firebase.InitFirebase()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	// Apply CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedOrigins([]string{"*"}), // ⚠️ for dev, replace "*" with your frontend domain in prod
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
	)

	log.Println("Server running on port " + config.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.ServerPort, corsHandler(r)))
}
