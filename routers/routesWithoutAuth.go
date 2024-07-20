package routers

import (
	"blog/handlers"

	"github.com/go-chi/chi"
)

func GetRoutes1(router *chi.Mux) {
	router.Post("/signup", handlers.SignUp)
	router.Post("/login", handlers.Login)
	
}