package entities

import "gorm.io/gorm"

// Comment represents a comment entity in the application.
// It includes fields for the comment's content, the ID of the associated blog post, and the ID of the user who made the comment.
type Comment struct {
	gorm.Model
	Content string `gorm:"not null"` // Content of the comment, must not be null
	BlogID  uint64 // ID of the blog post the comment is associated with
	UserID  uint64 // ID of the user who made the comment
}
