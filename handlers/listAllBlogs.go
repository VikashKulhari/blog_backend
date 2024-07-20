package handlers

import (
	"blog/database"
	"encoding/json"
	"fmt"
	"net/http"
)

func ListAllBlogPostsByUserHandler(w http.ResponseWriter, r *http.Request) {
	Usernameclaim, ok := r.Context().Value(database.ContextUsername).(string)
	if !ok {
		http.Error(w, "Not able to claim UserID", http.StatusUnauthorized)
	}
	posts := database.GetAllBlogPostsByUser(Usernameclaim)
	fmt.Println(posts)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
