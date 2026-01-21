package main

import (
	"fmt"
	"net/http"

	"validator/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	port := "3333"

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		MaxAge:         300,
	}).Handler)

	r.Get("/", routes.HomeRoute)

	r.Get("/validate", routes.ValidateRoute)

	fmt.Println("Server running on port " + port)

	http.ListenAndServe(":"+port, r)
}
