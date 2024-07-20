package handlers

import (
	"blog/database"
	"blog/entities"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func UploadImageInBlogPostHandler(w http.ResponseWriter, r *http.Request) {

	postTitle := chi.URLParam(r, "Title")
	existingPost, err := database.GetBlogPostByTitle(postTitle)
	if err != nil {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	var req entities.Images
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	existingPost.Images = append(existingPost.Tags, req.Images...)
	fmt.Println("Image is uploaded...")
	w.WriteHeader(http.StatusCreated)
}
