package handlers

import (
	"blog/database"
	"blog/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func EditCommentHandler(w http.ResponseWriter, r *http.Request) {
	UserIDclaim, ok := r.Context().Value(database.ContextUserID).(uint64)
	if !ok {
		http.Error(w, "Not able to claim UserID", http.StatusUnauthorized)
	}

	CommentIDstring := chi.URLParam(r, "ID")
	CommentID, err := strconv.ParseUint(CommentIDstring, 10, 64)
	if err != nil {
		http.Error(w, "Parsing of Comment ID failed", http.StatusExpectationFailed)
		return
	}
	CommentID1 := uint(CommentID)
	existingComment, err1 := database.GetCommentByIDforEditing(CommentID1)
	if err1 != nil {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	var req entities.Comment
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	fmt.Println(req.Content, "working")
	if existingComment.UserID == uint64(UserIDclaim) {
		existingComment.Content = req.Content
	}

	err = database.UpdateComment(existingComment)
	if err != nil {
		http.Error(w, "Failed to update blog post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
