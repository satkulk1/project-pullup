package main

import (
	"encoding/json" // NEW
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

// Job describes the shape we return as JSON
type Job struct {
	ID       string `json:"id"`
	Company  string `json:"company"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Location string `json:"location,omitempty"`
}

func main() {
	r := chi.NewRouter()

	// CORS for your React dev server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	r.Use(c.Handler)

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, PullUp!")
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	// NEW: JSON API route
	r.Get("/api/jobs", jobsHandler)

	// PORT from env with default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("▶ listening on http://localhost%s ...", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

// jobsHandler writes a JSON array of jobs to the response
func jobsHandler(w http.ResponseWriter, r *http.Request) {
	jobs := []Job{
		{
			ID: "stripe-se2-remote", Company: "Stripe",
			Title: "Software Engineer II", URL: "https://careers.example.com/stripe/se2",
			Location: "Remote (US)",
		},
		{
			ID: "spotify-be-sf", Company: "Spotify",
			Title: "Backend Engineer", URL: "https://careers.example.com/spotify/be",
			Location: "San Francisco, CA",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(jobs) // ignoring error for now (we'll handle later)
}
