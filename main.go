package main

import (
	"log"
	"net/http"
	"sms-verifier/config"
	"sms-verifier/firebase"
	"sms-verifier/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	firebase.InitFirebase()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("Server running on port " + config.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.ServerPort, r))
}
