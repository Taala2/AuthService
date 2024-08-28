package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
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

	accessToken, err := GenerateAccessToken(userID, ip)
	if err != nil {
		http.Error(w, "Error generating access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		http.Error(w, "Error generating refresh token", http.StatusInternalServerError)
		return
	}

	// Save refreshToken in DB with hashed version

	tokens := TokenPair{
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

	// Validate access token
	_, err := ValidateAccessToken(req.AccessToken)
	if err != nil {
		http.Error(w, "Invalid access token", http.StatusUnauthorized)
		return
	}

	// Validate refresh token
	// Retrieve stored hash from DB and compare with provided refresh token

	// Generate new tokens
	ip := r.RemoteAddr
	newAccessToken, err := GenerateAccessToken("userID", ip)
	if err != nil {
		http.Error(w, "Error generating new access token", http.StatusInternalServerError)
		return
	}

	newRefreshToken, err := GenerateRefreshToken()
	if err != nil {
		http.Error(w, "Error generating new refresh token", http.StatusInternalServerError)
		return
	}

	// Update stored refresh token hash in DB

	tokens := TokenPair{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}
