package handlers

import (
	"blog/database"
	"blog/entities"
	"encoding/json"
	"fmt"
	"net/http"
)

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	UsernameClaim, ok := r.Context().Value(database.ContextUsername).(string)
	fmt.Println("username claimed",UsernameClaim)
	if !ok {
		http.Error(w, "Not able to claim UserID", http.StatusUnauthorized)
	}
	var existingUser entities.User
	if err := db.Where("Username = ?", UsernameClaim).First(&existingUser).Error; err == nil {
		fmt.Println("Userexists")
	}

	w.WriteHeader(http.StatusOK)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userData := entities.User{
		Username: UsernameClaim,
		Email:    existingUser.Email,
		Password: existingUser.Password,
	}
	userData.ID = existingUser.ID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}
