package handlers

import (
	"blog/database"
	"blog/entities"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func GetBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	postTitle := chi.URLParam(r, "Title")
	post, err := database.GetBlogPostByTitle(postTitle)

	///////////
	blogID := post.ID
	var comments []entities.Comment
	db := database.GetDB()
	result := db.Where("blog_id = ?", blogID).Find(&comments)
	if result.Error != nil {
		return
	}
	post.Comments = comments
	/////////
	if err != nil {
		if err == entities.ErrRecordNotFound {
			http.Error(w, "Blog post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch blog post", http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(post)
}
