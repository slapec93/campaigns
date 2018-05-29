package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slapec93/campaigns/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.CampaignsList).Methods("GET")
	log.Fatal(http.ListenAndServe(":3001", router))
}
