package handlers

import (
	"blog/database"
	"blog/entities"
	"encoding/json"
	"net/http"
)

func CreateBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	UserIDclaim, ok := r.Context().Value(database.ContextUserID).(uint64)
	if !ok {
		http.Error(w, "Not able to claim UserID", http.StatusUnauthorized)
	}
	var req entities.BlogPost
	if r.Body == http.NoBody {
		http.Error(w, "Empty Body", http.StatusBadRequest)
		return
	}
	var title = req.Title
	if !database.CheckBlogPostsByTitle(title) {
		http.Error(w, "Title Already Exists", http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	req.UserID = UserIDclaim
	err := database.CreateBlogPost(&req)
	if err != nil {
		http.Error(w, "Failed to create blog post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
//firstCommit