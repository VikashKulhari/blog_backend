package handlers

import (
	"blog/database"
	"blog/entities"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var req entities.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("req=", req)
	var existingUser entities.User
	if err := db.Where("Username = ? OR Email = ?", req.Username, req.Email).First(&existingUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Verify the password
	if existingUser.Password != req.Password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	fmt.Println(existingUser)
	// If credentials are valid, generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID":   existingUser.ID,
		"Username": req.Username,
		"Email":    req.Email,
		"Password": req.Password,
	})

	fmt.Println(existingUser.ID)
	// Sign the token with a secret key to create a string representation of the token

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return the token to the client
	fmt.Println(req.Username + " with userID " + fmt.Sprint(existingUser.ID) +
		" logged in successfully In our System.")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	w.WriteHeader(http.StatusOK)

}
