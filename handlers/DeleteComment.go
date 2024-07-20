package handlers

import (
	"blog/database"
	"net/http"
	"strconv"
	"github.com/go-chi/chi"
)

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {

	CommentIDstring := chi.URLParam(r, "ID")
	CommentID,err := strconv.ParseUint(CommentIDstring,10,64)
	CommentID1 := uint(CommentID)
	if err != nil {
		return
	}

	UserIDclaim, ok := r.Context().Value(database.ContextUserID).(uint64)
	if !ok {
		http.Error(w, "Not able to claim UserID", http.StatusUnauthorized)
	}
	if database.CheckCommentbyID(CommentID1,UserIDclaim){
		err1 := database.DeleteComment(CommentID1)
		if err1 != nil {
			http.Error(w, "Failed to delete Comment", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
