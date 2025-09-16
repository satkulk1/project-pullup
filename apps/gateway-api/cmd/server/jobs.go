package main

// Job describes the shape we return as JSON
type Job struct {
	ID       string `json:"id"`
	Company  string `json:"company"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Location string `json:"location,omitempty"`
}

// sampleJobs returns a hardcoded slice of Job for now
func sampleJobs() []Job {
	return []Job{
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
}
