package handlers

import (
	"blog/database"
	"encoding/json"
	"net/http"
)

func ListBlogPostsByUserHandler(w http.ResponseWriter, r *http.Request) {
	UsernameClaim, ok := r.Context().Value(database.ContextUsername).(string)
	if !ok {
		http.Error(w, "Not able to claim UserID", http.StatusUnauthorized)
	}
	posts := database.GetBlogPostsByUser(UsernameClaim)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
