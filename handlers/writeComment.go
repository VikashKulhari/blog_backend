package handlers

import (
	"blog/database"
	"blog/entities"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func WriteCommentOnBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	UserIDclaim, ok := r.Context().Value(database.ContextUserID).(uint64)
	if !ok {
		http.Error(w, "Not able to claim UserID", http.StatusUnauthorized)
	}
	postTitle := chi.URLParam(r, "Title")
	existingPost, err := database.GetBlogPostByTitle(postTitle)
	if err != nil {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}
	var req entities.Comment
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	comment := entities.Comment{
		Content: req.Content,
		BlogID:  uint64(existingPost.ID),
		UserID:  uint64(UserIDclaim),
	}
	err = database.CreateComment(&comment)
	if err != nil {
		http.Error(w, "Failed to add comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
