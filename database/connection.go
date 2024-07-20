package database

import (
	"blog/entities"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB


// init initializes the database connection and performs automatic migration for the BlogPost, User, and Comment entities.
// It sets up the connection string for a PostgreSQL database running on localhost and opens the connection using GORM.
// If there is an error during the connection, it panics with the error message. Once connected, it performs automatic
// migration for the specified entities to ensure the database schema is up-to-date.
func init() {
	var err error
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable search_path =blog"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to database....")
	db.AutoMigrate(&entities.BlogPost{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Comment{})
	fmt.Println("connected to database..")

}

// CreateBlogPost creates a new blog post record in the database.
// It takes a pointer to a BlogPost entity as an argument and returns an error.
// If the operation is successful, it returns nil; otherwise, it returns the encountered error.
func CreateBlogPost(blogPost *entities.BlogPost) error {
	return db.Create(blogPost).Error
}

// CreateComment creates a new comment record in the database.
// It takes a pointer to a Comment entity as an argument and returns an error.
// If the operation is successful, it returns nil; otherwise, it returns the encountered error.
func CreateComment(comment *entities.Comment) error {
	return db.Create(comment).Error
}


// GetBlogPostByTitle retrieves a blog post by its title if it is published.
// It takes a string argument representing the title of the blog post and returns
// a pointer to a BlogPost entity and an error.
// If the blog post is found and published, it returns a pointer to the BlogPost entity and nil error.
// If there is an error during the database query, it returns nil and the encountered error.
func GetBlogPostByTitle(postTitle string) (*entities.BlogPost, error) {
	var post entities.BlogPost
	err := db.Where("Title = ? AND Published=?", postTitle, true).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// DeleteBlogPost deletes a blog post from the database by its title.
// It takes a string argument representing the title of the blog post and returns an error.
// If the operation is successful, it returns nil; otherwise, it returns the encountered error.
func DeleteBlogPost(postTitle string) error {
	return db.Where("Title = ?", postTitle).Delete(&entities.BlogPost{}).Error
}

// DeleteComment deletes a comment from the database by its ID.
// It takes an unsigned integer argument representing the ID of the comment and returns an error.
// If the operation is successful, it returns nil; otherwise, it returns the encountered error.
func DeleteComment(CommentID uint) error {
	return db.Where("ID = ?", CommentID).Delete(&entities.Comment{}).Error
}


// GetCommentByID retrieves a comment by its ID.
// It takes an unsigned integer argument representing the ID of the comment and returns a pointer
// to a Comment entity and an error. If the comment is found, it returns a pointer to the Comment entity
// and nil error. If there is an error during the database query, it returns nil and the encountered error.
func GetCommentByID(commentID uint) (*entities.Comment, error) {
	var comment entities.Comment
	err := db.Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetCommentByIDforEditing retrieves a comment by its ID for editing purposes.
// It takes an unsigned integer argument representing the ID of the comment and returns a pointer
// to a Comment entity and an error. If the comment is found, it returns a pointer to the Comment entity
// and nil error. If there is an error during the database query, it returns nil and the encountered error.
func GetCommentByIDforEditing(commentID uint) (*entities.Comment, error) {
	var comment entities.Comment
	err := db.Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}


// UpdateComment updates an existing comment in the database.
// It takes a pointer to a Comment entity as an argument and returns an error.
// If the operation is successful, it returns nil; otherwise, it returns the encountered error.
func UpdateComment(comment *entities.Comment) error {
	return db.Save(comment).Error
}


// UpdateBlogPost updates an existing blog post in the database.
// It takes a pointer to a BlogPost entity as an argument and returns an error.
// If the operation is successful, it returns nil; otherwise, it returns the encountered error.
func UpdateBlogPost(post *entities.BlogPost) error {
	return db.Save(post).Error
}

// CheckBlogPostsByTitle checks if a blog post with a given title exists in the database.
// It takes a string argument representing the title of the blog post and returns a boolean.
// If the blog post is found, it returns true; otherwise, it returns false.
func CheckBlogPostsByTitle(title string) bool {
	var blogPost entities.BlogPost
	if err := db.Where("Title = ? ", title).First(&blogPost).Error; err != nil {
		return false
	}

	return true
}

// GetBlogPostsByUser retrieves all published blog posts by a user based on their username.
// It takes a string argument representing the username and returns a slice of BlogPost entities.
// If the user or the blog posts are not found, it prints an error message and returns an empty slice.
func GetBlogPostsByUser(username string) []entities.BlogPost {
	var user entities.User
	if err := db.Where("Username = ? ", username).First(&user).Error; err != nil {
		fmt.Println("User not found..")
	}
	var posts []entities.BlogPost
	if err := db.Find(&posts, "user_id=? AND Published=?", user.ID, true).Error; err != nil {
		fmt.Println("User found and unable to fetch posts")
	}

	return posts
}

// CheckCommentbyID checks if a comment with a given ID exists and belongs to a specific user.
// It takes two arguments: an unsigned integer representing the comment ID and an unsigned integer
// representing the user ID. It returns a boolean indicating whether the comment exists and belongs
// to the specified user. If the comment is found and belongs to the user, it returns true; otherwise, false.
func CheckCommentbyID(CID uint, UID uint64) bool {
	var comment entities.Comment
	if err := db.Where("ID = ? ", CID).First(&comment).Error; err != nil {
		fmt.Println("Comment not found..")
	}
	if comment.UserID == UID {
		return true
	}
	return false
}

// GetAllBlogPostsByUser retrieves all blog posts by a user based on their username.
// It takes a string argument representing the username and returns a slice of BlogPost entities.
// If the user or the blog posts are not found, it prints an error message and returns an empty slice.
func GetAllBlogPostsByUser(username string) []entities.BlogPost {
	var user entities.User
	if err := db.Where("Username = ? ", username).First(&user).Error; err != nil {
		fmt.Println("User not found..")
	}
	var posts []entities.BlogPost
	if err := db.Find(&posts, "user_id=?", user.ID).Error; err != nil {
		fmt.Println("User found and unable to fetch posts")
	}

	return posts
}

// SearchUserByUsername searches for a user by their username.
// It takes a string argument representing the username and returns a pointer to a User entity and an error.
// If the user is found, it returns a pointer to the User entity and nil error.
// If there is an error during the database query, it returns nil and the encountered error.
func SearchUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}


// AddBlogPostToUser adds a blog post to a user based on their username.
// It takes a string argument representing the username and a pointer to a BlogPost entity as arguments.
// It returns an error. If the operation is successful, it returns nil; otherwise, it returns the encountered error.
func AddBlogPostToUser(username string, blogPost *entities.BlogPost) error {
	user, err := SearchUserByUsername(username)
	if err != nil {
		return err
	}
	user.BlogPosts = append(user.BlogPosts, *blogPost)

	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetDB returns the database connection.
// It returns a pointer to a gorm.DB instance.
func GetDB() *gorm.DB {
	return db
}
