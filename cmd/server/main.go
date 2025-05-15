package main

import (
	"log"
	"net/http"
	// "friends-forever/internal/db"
	"friends-forever/web/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Initialize the database
	// if err := db.InitDB("game.db"); err != nil {
	// 	log.Fatalf("Could not initialize database: %v", err)
	// }

	// Create a new router
	r := chi.NewRouter()

	handlers.RegisterRoutes(r)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	// // Define routes
	// r.Get("/hello", handlers.HelloHandler)
	// r.Get("/events", handlers.EventsHandler)

	// Start the server
	log.Println("Server running at http://localhost:3001")
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}