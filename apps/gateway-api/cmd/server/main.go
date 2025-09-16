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
	jobs := sampleJobs() // NEW: fetch from helper function

	company := r.URL.Query().Get("company")
	if company != "" {
		filtered := []Job{}
		for _, job := range jobs {
			if job.Company == company {
				filtered = append(filtered, job)
			}
		}
		jobs = filtered
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(jobs)
}
