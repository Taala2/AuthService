package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Taala2/auth-service/utils"
	"github.com/Taala2/auth-service/models"

)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/token", GetTokenHandler).Methods("GET")
	r.HandleFunc("/refresh", RefreshTokenHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	ip := r.RemoteAddr

	accessToken, err := utils.GenerateAccessToken(userID, ip)
	if err != nil {
		http.Error(w, "Error generating access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		http.Error(w, "Error generating refresh token", http.StatusInternalServerError)
		return
	}

	tokens := models.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}


	_, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		http.Error(w, "Invalid access token", http.StatusUnauthorized)
		return
	}


	ip := r.RemoteAddr
	newAccessToken, err := utils.GenerateAccessToken("userID", ip)
	if err != nil {
		http.Error(w, "Error generating new access token", http.StatusInternalServerError)
		return
	}

	newRefreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		http.Error(w, "Error generating new refresh token", http.StatusInternalServerError)
		return
	}

	// Update stored refresh token hash in DB

	tokens := models.TokenPair{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}
