package entities

import (
	"errors"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found, new user created")

// BlogPost represents a blog post entity in the application.
// It includes fields for the post's title, content, publication status, user ID, tags, images, and associated comments.
type BlogPost struct {
	gorm.Model
	Title     string         `gorm:"not null"`                  // Title of the blog post, must not be null
	Content   string         `gorm:"not null" json:"Content"`   // Content of the blog post, must not be null
	Published bool           `gorm:"not null" json:"Published"` // Publication status of the blog post, must not be null
	UserID    uint64         // ID of the user who created the blog post
	Tags      pq.StringArray `gorm:"type:text[]"`                         // Array of tags associated with the blog post
	Images    pq.StringArray `gorm:"type:text[]"`                         // Array of image URLs associated with the blog post
	Comments  []Comment      `gorm:"foreignKey:BlogID; onDelete:CASCADE"` // Associated comments, deleted with the blog post
}
