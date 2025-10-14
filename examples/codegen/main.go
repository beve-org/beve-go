package main

//go:generate go run ../../cmd/bevegen/main.go -type=User,Product

import (
	"fmt"

	"github.com/beve-org/beve-go"
)

// User represents a user in the system
type User struct {
	ID       int64  `beve:"id"`
	Username string `beve:"username"`
	Email    string `beve:"email,omitempty"`
	Age      int    `beve:"age"`
	IsActive bool   `beve:"active"`
	// CreatedAt time.Time `beve:"created_at"` // TODO: Add extension type support
}

// Product represents a product
type Product struct {
	ID          int64   `beve:"id"`
	Name        string  `beve:"name"`
	Description string  `beve:"description,omitempty"`
	Price       float64 `beve:"price"`
	InStock     bool    `beve:"in_stock"`
	Quantity    int32   `beve:"quantity"`
}

func main() {
	// Create test user
	user := User{
		ID:       12345,
		Username: "john_doe",
		Email:    "john@example.com",
		Age:      30,
		IsActive: true,
	}

	// Create test product
	product := Product{
		ID:          67890,
		Name:        "Laptop",
		Description: "High-performance laptop",
		Price:       1299.99,
		InStock:     true,
		Quantity:    50,
	}

	// Test standard marshaling
	fmt.Println("=== Standard BEVE Marshaling ===")

	userData, err := beve.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshaling user: %v\n", err)
	} else {
		fmt.Printf("User marshaled: %d bytes\n", len(userData))
	}

	productData, err := beve.Marshal(product)
	if err != nil {
		fmt.Printf("Error marshaling product: %v\n", err)
	} else {
		fmt.Printf("Product marshaled: %d bytes\n", len(productData))
	}

	// After running `go generate`, this will test the generated methods:
	// userData, err := user.MarshalBEVE()
	// productData, err := product.MarshalBEVE()

	fmt.Println("\n✓ Code generation test struct created")
	fmt.Println("Run: go generate ./examples/codegen")
	fmt.Println("Then: go run ./examples/codegen")
}
