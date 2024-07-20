package routers

import (
	"blog/database"
	"blog/handlers"

	"github.com/go-chi/chi"
)

func GetRoutes(router chi.Router) {
	router.Use(database.JWTAuthMiddleware)
	router.Post("/blog", handlers.CreateBlogPostHandler)
	router.Get("/userProfile", handlers.UserProfileHandler)
	router.Get("/blog/{Title}", handlers.GetBlogPostHandler)
	router.Get("/blog/{Title}/comments", handlers.GetCommentsHandler)
	router.Put("/blog/{Title}", handlers.EditBlogPostHandler)
	router.Put("/EditComment/{ID}", handlers.EditCommentHandler)
	router.Delete("/blog/{Title}", handlers.DeleteBlogPostHandler)
	router.Get("/user/allblogs", handlers.ListAllBlogPostsByUserHandler)
	router.Get("/user/blogs", handlers.ListBlogPostsByUserHandler)
	router.Put("/blog/{Title}/tags", handlers.UpdateBlogPostTagsHandler)
	router.Post("/blog/{Title}/image", handlers.UploadImageInBlogPostHandler)
	router.Post("/blog/{Title}/comment", handlers.WriteCommentOnBlogPostHandler)
	router.Delete("/deleteComment/{ID}", handlers.DeleteCommentHandler)
}
