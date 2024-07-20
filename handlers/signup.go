package handlers

import (
	"blog/database"
	"blog/entities"
	"encoding/json"
	"net/http"
)

// SignUpRequest represents the expected JSON payload for the SignUp handler.
type SignUpRequest struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}



//To improve security, it's crucial to hash passwords with a salt before storing them in the database. 
//you can modify the SignUp function to hash passwords using the bcrypt package

// SignUp handles user registration requests.
// It decodes the JSON request body into a SignUpRequest struct, checks if the username is already taken,
// and if not, creates a new user in the database. Appropriate HTTP status codes and messages are returned
// based on the outcome of these operations.
func SignUp(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var req SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if the username is already taken
	var existingUser entities.User
	if err := db.Where("Username = ?", req.Username).First(&existingUser).Error; err == nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Create a new user
	newUser := entities.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := db.Create(&newUser).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
