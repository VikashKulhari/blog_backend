package handlers

import (
	"blog/database"
	"blog/entities"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type UpdateTagsRequest struct {
	Tags []string `json:"Tags"`
}

func UpdateBlogPostTagsHandler(w http.ResponseWriter, r *http.Request) {
	postTitle := chi.URLParam(r, "Title")

	existingPost, err := database.GetBlogPostByTitle(postTitle)
	if err != nil {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	var req entities.Tags
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	existingPost.Tags = nil
	existingPost.Tags = append(existingPost.Tags, req.Tags...)

	err = database.UpdateBlogPost(existingPost)
	if err != nil {
		http.Error(w, "Failed to update blog post tags", http.StatusInternalServerError)
		return
	}
	fmt.Println("Tags for the post with title " + existingPost.Title + " has been updated.")
	w.WriteHeader(http.StatusOK)
}
