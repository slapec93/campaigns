package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// CampaignsList ...
func CampaignsList(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to API")
	json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
}
