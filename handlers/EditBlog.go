package handlers

import (
	"blog/database"
	"blog/entities"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func EditBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	UserIDclaim, ok := r.Context().Value(database.ContextUserID).(uint)
	if !ok {
		http.Error(w, "Not able to claim UserID", http.StatusUnauthorized)
	}

	postTitle := chi.URLParam(r, "Title")

	existingPost, err := database.GetBlogPostByTitle(postTitle)
	if err != nil {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	var req entities.BlogPost
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	var title = req.Title
	if database.CheckBlogPostsByTitle(title){
		http.Error(w,"Title Already Exists", http.StatusBadRequest)
	}
	existingPost.Title = req.Title
	existingPost.Content = req.Content
	existingPost.Published = req.Published
	existingPost.UserID = uint64(UserIDclaim)

	existingPost.Tags = nil
	existingPost.Images = nil
	existingPost.Tags = append(existingPost.Tags, req.Tags...)
	existingPost.Images = append(existingPost.Images, req.Images...)
	err = database.UpdateBlogPost(existingPost)
	if err != nil {
		http.Error(w, "Failed to update blog post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
