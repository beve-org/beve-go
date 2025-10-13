package main

import (
	"encoding/json"
	"log"
	"time"

	beve "github.com/beve-org/beve-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// BeveMiddleware for Fiber framework
func BeveMiddleware(debug bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Handle BEVE request body
		contentType := string(c.Request().Header.ContentType())
		if contentType == "application/beve" {
			start := time.Now()

			// Decode BEVE body
			body := c.Body()
			var data interface{}
			if err := beve.Unmarshal(body, &data); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"error": "Invalid BEVE request",
				})
			}

			// Store decoded data in locals
			c.Locals("beveData", data)
			c.Locals("isBeveRequest", true)

			if debug {
				log.Printf("[BEVE] Request decoded in %v - %d bytes", time.Since(start), len(body))
			}
		}

		// Process request
		if err := c.Next(); err != nil {
			return err
		}

		// 2. Handle BEVE response encoding
		accepts := c.Accepts("application/beve")
		if accepts != "application/beve" {
			// Client doesn't want BEVE, return as-is
			return nil
		}

		// Get response body
		responseBody := c.Response().Body()
		if len(responseBody) == 0 {
			return nil
		}

		start := time.Now()

		// Parse JSON response
		var data interface{}
		if err := json.Unmarshal(responseBody, &data); err != nil {
			if debug {
				log.Printf("[BEVE] Failed to parse response as JSON: %v", err)
			}
			return nil // Return original response
		}

		// Encode as BEVE
		encoded, err := beve.Marshal(data)
		if err != nil {
			if debug {
				log.Printf("[BEVE] Failed to encode response: %v", err)
			}
			return nil // Return original response
		}

		// Replace response
		c.Response().Header.SetContentType("application/beve")
		c.Response().SetBody(encoded)

		if debug {
			savedBytes := len(responseBody) - len(encoded)
			savedPercent := float64(savedBytes) / float64(len(responseBody)) * 100
			log.Printf("[BEVE] Response encoded in %v - Saved %d bytes (%.1f%%)",
				time.Since(start), savedBytes, savedPercent)
		}

		return nil
	}
}

// Helper: Get decoded BEVE data from request
func GetBeveData(c *fiber.Ctx) (interface{}, bool) {
	if c.Locals("isBeveRequest") == true {
		if data := c.Locals("beveData"); data != nil {
			return data, true
		}
	}
	return nil, false
}

func main() {
	app := fiber.New(fiber.Config{
		AppName: "BEVE API Server",
	})

	// Logger middleware
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
	}))

	// BEVE middleware
	app.Use(BeveMiddleware(true))

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":    "BEVE API",
			"version": "1.0.0",
			"formats": []string{"json", "beve"},
		})
	})

	// Users API
	app.Get("/api/users", handleGetUsers)
	app.Post("/api/users", handleCreateUser)
	app.Get("/api/users/:id", handleGetUser)

	// Posts API
	app.Get("/api/posts", handleGetPosts)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"time":   time.Now(),
		})
	})

	log.Println("🚀 Server starting on :3000")
	log.Println("📦 BEVE encoding: enabled")
	log.Println("📊 Debug mode: enabled")
	log.Fatal(app.Listen(":3000"))
}

// Handlers
func handleGetUsers(c *fiber.Ctx) error {
	users := []fiber.Map{
		{
			"id":    1,
			"name":  "Alice Johnson",
			"email": "alice@example.com",
			"role":  "admin",
		},
		{
			"id":    2,
			"name":  "Bob Smith",
			"email": "bob@example.com",
			"role":  "user",
		},
		{
			"id":    3,
			"name":  "Charlie Brown",
			"email": "charlie@example.com",
			"role":  "user",
		},
	}

	return c.JSON(fiber.Map{
		"data":  users,
		"count": len(users),
	})
}

func handleCreateUser(c *fiber.Ctx) error {
	// Try to get BEVE data first
	if data, ok := GetBeveData(c); ok {
		log.Println("✅ Received BEVE request:", data)

		// Process BEVE data
		user := data.(map[string]interface{})
		user["id"] = 123 // Assign ID
		user["createdAt"] = time.Now()

		return c.Status(201).JSON(user)
	}

	// Fallback to JSON parsing
	var user map[string]interface{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user["id"] = 123
	user["createdAt"] = time.Now()

	return c.Status(201).JSON(user)
}

func handleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// Mock user data
	user := fiber.Map{
		"id":    id,
		"name":  "Alice Johnson",
		"email": "alice@example.com",
		"role":  "admin",
		"profile": fiber.Map{
			"avatar": "https://example.com/avatar.jpg",
			"bio":    "Software Engineer",
		},
		"preferences": fiber.Map{
			"theme":    "dark",
			"language": "en",
		},
	}

	return c.JSON(user)
}

func handleGetPosts(c *fiber.Ctx) error {
	posts := []fiber.Map{
		{
			"id":        1,
			"title":     "Introduction to BEVE",
			"author":    "Alice",
			"content":   "BEVE is a fast binary encoding format...",
			"tags":      []string{"beve", "encoding", "performance"},
			"createdAt": time.Now().Add(-24 * time.Hour),
		},
		{
			"id":        2,
			"title":     "BEVE vs JSON Performance",
			"author":    "Bob",
			"content":   "Comparing BEVE and JSON encoding speeds...",
			"tags":      []string{"beve", "json", "benchmark"},
			"createdAt": time.Now().Add(-12 * time.Hour),
		},
		{
			"id":        3,
			"title":     "Building APIs with BEVE",
			"author":    "Charlie",
			"content":   "Learn how to integrate BEVE into your API...",
			"tags":      []string{"beve", "api", "tutorial"},
			"createdAt": time.Now().Add(-6 * time.Hour),
		},
	}

	return c.JSON(fiber.Map{
		"data":  posts,
		"count": len(posts),
	})
}
