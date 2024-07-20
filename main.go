package main

import (
	"blog/handlers"
	"blog/routers"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func main() {
	router := chi.NewRouter()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"}, // Allow requests from Client Side 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for CORS preflight requests
	})
	router.Use(corsHandler.Handler)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	
	//routes with AUTH
	router.Group(routers.GetRoutes)

	//do not need auth on below requests
	router.Post("/signup", handlers.SignUp)
	router.Post("/login", handlers.Login)
	port := ":8080"
	log.Println("Server started on 8080")
	log.Fatal(http.ListenAndServe(port, router))

}
