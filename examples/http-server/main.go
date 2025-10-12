package httpserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/beve-org/beve-go"
)

// User represents a user entity
type User struct {
	ID        int       `beve:"id" json:"id"`
	Name      string    `beve:"name" json:"name"`
	Email     string    `beve:"email" json:"email"`
	CreatedAt time.Time `beve:"created_at" json:"created_at"`
	Active    bool      `beve:"active" json:"active"`
}

// Response represents an API response
type Response struct {
	Success bool        `beve:"success" json:"success"`
	Data    interface{} `beve:"data,omitempty" json:"data,omitempty"`
	Error   string      `beve:"error,omitempty" json:"error,omitempty"`
}

// MIME type constant
const MIMETypeBEVE = "application/beve"

// respondBEVE sends a BEVE-encoded response
func respondBEVE(w http.ResponseWriter, statusCode int, data interface{}) {
	encoded, err := beve.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", MIMETypeBEVE)
	w.WriteHeader(statusCode)
	w.Write(encoded)
}

// getUserHandler returns a single user
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:        1,
		Name:      "Alice Johnson",
		Email:     "alice@example.com",
		CreatedAt: time.Now(),
		Active:    true,
	}

	response := Response{
		Success: true,
		Data:    user,
	}

	respondBEVE(w, http.StatusOK, response)
}

// getUsersHandler returns multiple users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Alice Johnson", Email: "alice@example.com", CreatedAt: time.Now(), Active: true},
		{ID: 2, Name: "Bob Smith", Email: "bob@example.com", CreatedAt: time.Now(), Active: true},
		{ID: 3, Name: "Carol White", Email: "carol@example.com", CreatedAt: time.Now(), Active: false},
	}

	response := Response{
		Success: true,
		Data:    users,
	}

	respondBEVE(w, http.StatusOK, response)
}

// createUserHandler creates a new user from BEVE-encoded request body
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Check Content-Type
	if r.Header.Get("Content-Type") != MIMETypeBEVE {
		respondBEVE(w, http.StatusBadRequest, Response{
			Success: false,
			Error:   "Content-Type must be application/beve",
		})
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondBEVE(w, http.StatusBadRequest, Response{
			Success: false,
			Error:   "Failed to read request body",
		})
		return
	}
	defer r.Body.Close()

	// Unmarshal BEVE data
	var user User
	if err := beve.Unmarshal(body, &user); err != nil {
		respondBEVE(w, http.StatusBadRequest, Response{
			Success: false,
			Error:   fmt.Sprintf("Invalid BEVE data: %v", err),
		})
		return
	}

	// Set server-side fields
	user.ID = 4 // In real app, would be from database
	user.CreatedAt = time.Now()

	// Return created user
	response := Response{
		Success: true,
		Data:    user,
	}

	respondBEVE(w, http.StatusCreated, response)
}

// healthHandler returns server health status
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: true,
		Data: map[string]interface{}{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
			"format":    "BEVE",
		},
	}

	respondBEVE(w, http.StatusOK, response)
}

// loggingMiddleware logs each request
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.Header.Get("Content-Type"))
		next(w, r)
		log.Printf("Completed in %v", time.Since(start))
	}
}

func main() {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/health", loggingMiddleware(healthHandler))
	mux.HandleFunc("/api/user", loggingMiddleware(getUserHandler))
	mux.HandleFunc("/api/users", loggingMiddleware(getUsersHandler))
	mux.HandleFunc("/api/users/create", loggingMiddleware(createUserHandler))

	// Start server
	addr := ":8080"
	log.Printf("🚀 BEVE HTTP Server starting on %s", addr)
	log.Printf("📡 MIME Type: %s", MIMETypeBEVE)
	log.Println("\nEndpoints:")
	log.Println("  GET  /health            - Health check")
	log.Println("  GET  /api/user          - Get single user")
	log.Println("  GET  /api/users         - Get all users")
	log.Println("  POST /api/users/create  - Create user (requires application/beve)")
	log.Println("\nExample usage:")
	log.Println("  curl -H 'Accept: application/beve' http://localhost:8080/api/user")

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
