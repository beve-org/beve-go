package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	beve "github.com/beve-org/beve-go"
)

// BeveMiddleware handles automatic BEVE encoding/decoding
type BeveMiddleware struct {
	// Enable debug logging
	Debug bool
	// Fallback to JSON for non-BEVE clients
	FallbackJSON bool
}

// Handler wraps an http.Handler with BEVE support
func (m *BeveMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Decode BEVE request body if present
		if err := m.decodeRequest(r); err != nil {
			http.Error(w, "Invalid BEVE request: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Wrap response writer to encode response as BEVE
		bw := &beveResponseWriter{
			ResponseWriter: w,
			request:        r,
			middleware:     m,
		}

		next.ServeHTTP(bw, r)
	})
}

// decodeRequest decodes BEVE request body
func (m *BeveMiddleware) decodeRequest(r *http.Request) error {
	contentType := r.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/beve") {
		// Not a BEVE request
		return nil
	}

	start := time.Now()

	// Read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Body.Close()

	if m.Debug {
		log.Printf("[BEVE Request] %s %s - Body size: %d bytes", r.Method, r.URL.Path, len(body))
	}

	// Decode BEVE
	var data interface{}
	if err := beve.Unmarshal(body, &data); err != nil {
		return err
	}

	// Re-encode as JSON for standard handlers
	// (handlers can read the decoded data from body)
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Replace body with JSON version
	r.Body = io.NopCloser(bytes.NewReader(jsonBody))
	r.ContentLength = int64(len(jsonBody))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-Original-Format", "beve")

	if m.Debug {
		log.Printf("[BEVE Request] Decoded in %v - JSON size: %d bytes", time.Since(start), len(jsonBody))
	}

	return nil
}

// beveResponseWriter wraps http.ResponseWriter to encode responses as BEVE
type beveResponseWriter struct {
	http.ResponseWriter
	request    *http.Request
	middleware *BeveMiddleware
	buffer     bytes.Buffer
	statusCode int
	written    bool
}

func (w *beveResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	// Don't write yet, we need to intercept the body first
}

func (w *beveResponseWriter) Write(b []byte) (int, error) {
	if !w.written {
		w.written = true

		// Check if client accepts BEVE
		accept := w.request.Header.Get("Accept")
		wantsBEVE := strings.Contains(accept, "application/beve")

		if wantsBEVE {
			// Try to encode as BEVE
			if encoded, err := w.encodeBEVE(b); err == nil {
				if w.middleware.Debug {
					log.Printf("[BEVE Response] %s %s - Original: %d bytes, BEVE: %d bytes (%.1f%% smaller)",
						w.request.Method, w.request.URL.Path,
						len(b), len(encoded),
						float64(len(b)-len(encoded))/float64(len(b))*100,
					)
				}

				// Write BEVE response
				w.ResponseWriter.Header().Set("Content-Type", "application/beve")
				if w.statusCode != 0 {
					w.ResponseWriter.WriteHeader(w.statusCode)
				}
				return w.ResponseWriter.Write(encoded)
			} else if w.middleware.Debug {
				log.Printf("[BEVE Response] Encoding failed: %v, falling back to JSON", err)
			}
		}

		// Fallback to original format (JSON)
		if w.statusCode != 0 {
			w.ResponseWriter.WriteHeader(w.statusCode)
		}
	}

	return w.ResponseWriter.Write(b)
}

func (w *beveResponseWriter) encodeBEVE(data []byte) ([]byte, error) {
	// Parse JSON data
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}

	// Encode as BEVE
	return beve.Marshal(v)
}

// Example usage with standard net/http
func main() {
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/users", handleUsers)
	mux.HandleFunc("/api/posts", handlePosts)

	// Wrap with BEVE middleware
	beveMiddleware := &BeveMiddleware{
		Debug:        true,
		FallbackJSON: true,
	}

	handler := beveMiddleware.Handler(mux)

	log.Println("🚀 Server starting on :3000")
	log.Println("📦 BEVE encoding enabled")
	log.Fatal(http.ListenAndServe(":3000", handler))
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Return users list
		users := []map[string]interface{}{
			{"id": 1, "name": "Alice", "email": "alice@example.com"},
			{"id": 2, "name": "Bob", "email": "bob@example.com"},
		}
		respondJSON(w, users)

	case http.MethodPost:
		// Create user
		var user map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user["id"] = 123 // Assign ID
		respondJSON(w, user)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	posts := []map[string]interface{}{
		{"id": 1, "title": "Hello BEVE", "author": "Alice"},
		{"id": 2, "title": "BEVE vs JSON", "author": "Bob"},
	}
	respondJSON(w, posts)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
