package handlers

import (
	"blog/database"
	"net/http"

	"github.com/go-chi/chi"
)

func DeleteBlogPostHandler(w http.ResponseWriter, r *http.Request) {

	postTitle := chi.URLParam(r, "Title")

	err := database.DeleteBlogPost(postTitle)
	if err != nil {
		http.Error(w, "Failed to delete blog post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
