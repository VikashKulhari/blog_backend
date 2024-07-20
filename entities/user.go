package entities

import "gorm.io/gorm"

// User represents a user entity in the application.
// It includes fields for the user's username, email, password, associated blog posts, and comments.
type User struct {
	gorm.Model
	Username  string     `gorm:"not null" schema:"users"`            // Username of the user, must not be null
	Email     string     `gorm:"type:varchar(100);unique_index"`     // Email of the user, must be unique and have a maximum length of 100 characters
	Password  string     `gorm:"type:varchar(100)"`                  // Password of the user, stored as a hashed string with a maximum length of 100 characters
	BlogPosts []BlogPost `gorm:"foreignKey:UserID;onDelete:CASCADE"` // Associated blog posts, deleted with the user
	Comments  []Comment  `gorm:"foreignKey:UserID;onDelete:CASCADE"` // Associated comments, deleted with the user
}
