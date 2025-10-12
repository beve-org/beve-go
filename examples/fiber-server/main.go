package fiberserver

import (
	"fmt"
	"log"
	"time"

	"github.com/beve-org/beve-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

// SendBEVE sends a BEVE-encoded response
func SendBEVE(c *fiber.Ctx, statusCode int, data interface{}) error {
	encoded, err := beve.Marshal(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to encode response",
		})
	}

	c.Set("Content-Type", MIMETypeBEVE)
	return c.Status(statusCode).Send(encoded)
}

// GetUser returns a single user
func GetUser(c *fiber.Ctx) error {
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

	return SendBEVE(c, fiber.StatusOK, response)
}

// GetUsers returns multiple users
func GetUsers(c *fiber.Ctx) error {
	users := []User{
		{ID: 1, Name: "Alice Johnson", Email: "alice@example.com", CreatedAt: time.Now(), Active: true},
		{ID: 2, Name: "Bob Smith", Email: "bob@example.com", CreatedAt: time.Now(), Active: true},
		{ID: 3, Name: "Carol White", Email: "carol@example.com", CreatedAt: time.Now(), Active: false},
	}

	response := Response{
		Success: true,
		Data:    users,
	}

	return SendBEVE(c, fiber.StatusOK, response)
}

// CreateUser creates a new user from BEVE-encoded request body
func CreateUser(c *fiber.Ctx) error {
	// Check Content-Type
	if c.Get("Content-Type") != MIMETypeBEVE {
		return SendBEVE(c, fiber.StatusBadRequest, Response{
			Success: false,
			Error:   "Content-Type must be application/beve",
		})
	}

	// Get request body
	body := c.Body()

	// Unmarshal BEVE data
	var user User
	if err := beve.Unmarshal(body, &user); err != nil {
		return SendBEVE(c, fiber.StatusBadRequest, Response{
			Success: false,
			Error:   fmt.Sprintf("Invalid BEVE data: %v", err),
		})
	}

	// Set server-side fields
	user.ID = 4 // In real app, would be from database
	user.CreatedAt = time.Now()

	// Return created user
	response := Response{
		Success: true,
		Data:    user,
	}

	return SendBEVE(c, fiber.StatusCreated, response)
}

// Health check endpoint
func Health(c *fiber.Ctx) error {
	response := Response{
		Success: true,
		Data: fiber.Map{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
			"format":    "BEVE",
			"framework": "Fiber v2",
		},
	}

	return SendBEVE(c, fiber.StatusOK, response)
}

// Stats endpoint showing request stats
func Stats(c *fiber.Ctx) error {
	response := Response{
		Success: true,
		Data: fiber.Map{
			"routes_count":  c.App().HandlersCount(),
			"fiber_version": fiber.Version,
			"server":        "Fiber v2 + BEVE",
		},
	}

	return SendBEVE(c, fiber.StatusOK, response)
}

// BEVEMiddleware adds BEVE support indicator to context
func BEVEMiddleware(c *fiber.Ctx) error {
	c.Locals("beve_enabled", true)
	return c.Next()
}

func main() {
	// Create Fiber app with custom config
	app := fiber.New(fiber.Config{
		AppName:               "BEVE Fiber Server",
		DisableStartupMessage: false,
		EnablePrintRoutes:     true,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(BEVEMiddleware)

	// Routes
	api := app.Group("/api")

	// Health and stats
	app.Get("/health", Health)
	app.Get("/stats", Stats)

	// User endpoints
	api.Get("/user", GetUser)
	api.Get("/users", GetUsers)
	api.Post("/users/create", CreateUser)

	// Start server
	log.Printf("🚀 BEVE Fiber Server")
	log.Printf("📡 MIME Type: %s", MIMETypeBEVE)
	log.Printf("⚡ Framework: Fiber v2")
	log.Println("\nEndpoints:")
	log.Println("  GET  /health            - Health check")
	log.Println("  GET  /stats             - Server stats")
	log.Println("  GET  /api/user          - Get single user")
	log.Println("  GET  /api/users         - Get all users")
	log.Println("  POST /api/users/create  - Create user (requires application/beve)")

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
