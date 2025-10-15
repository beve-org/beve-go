// Package codegen provides integration tests for bevegen code generation
package main

import (
	"testing"

	"github.com/beve-org/beve-go"
)

// TestGeneratedCodeRoundTrip tests that generated MarshalBEVE works with Unmarshal
func TestGeneratedCodeRoundTrip(t *testing.T) {
	user := User{
		ID:       12345,
		Username: "john_doe",
		Email:    "john@example.com",
		Age:      30,
		IsActive: true,
	}

	// Marshal with generated code
	data, err := user.MarshalBEVE()
	if err != nil {
		t.Fatalf("MarshalBEVE failed: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("Marshaled data is empty")
	}

	// Unmarshal with standard beve
	var decoded User
	if err := beve.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Verify fields
	if decoded.ID != user.ID {
		t.Errorf("ID mismatch: got %d, want %d", decoded.ID, user.ID)
	}
	if decoded.Username != user.Username {
		t.Errorf("Username mismatch: got %s, want %s", decoded.Username, user.Username)
	}
	if decoded.Email != user.Email {
		t.Errorf("Email mismatch: got %s, want %s", decoded.Email, user.Email)
	}
	if decoded.Age != user.Age {
		t.Errorf("Age mismatch: got %d, want %d", decoded.Age, user.Age)
	}
	if decoded.IsActive != user.IsActive {
		t.Errorf("IsActive mismatch: got %v, want %v", decoded.IsActive, user.IsActive)
	}
}

// TestGeneratedCodeOmitEmpty tests omitempty tag handling
func TestGeneratedCodeOmitEmpty(t *testing.T) {
	// User without email (should be omitted)
	userWithoutEmail := User{
		ID:       1,
		Username: "minimal",
		Email:    "", // Empty, should be skipped
		Age:      25,
		IsActive: true,
	}

	dataWithout, err := userWithoutEmail.MarshalBEVE()
	if err != nil {
		t.Fatalf("MarshalBEVE without email failed: %v", err)
	}
	t.Logf("userWithoutEmail.Email = %q (len=%d)", userWithoutEmail.Email, len(userWithoutEmail.Email))
	t.Logf("Data without email (len=%d): %x", len(dataWithout), dataWithout)

	// User with email
	userWithEmail := User{
		ID:       1,
		Username: "minimal",
		Email:    "user@example.com",
		Age:      25,
		IsActive: true,
	}

	dataWith, err := userWithEmail.MarshalBEVE()
	if err != nil {
		t.Fatalf("MarshalBEVE with email failed: %v", err)
	}

	// Data with email should be larger
	if len(dataWith) <= len(dataWithout) {
		t.Errorf("Expected larger data with email: got %d bytes, without email %d bytes",
			len(dataWith), len(dataWithout))
	}

	// Verify unmarshal
	var decoded User
	if err := beve.Unmarshal(dataWithout, &decoded); err != nil {
		t.Logf("Data without email (len=%d): %x", len(dataWithout), dataWithout)
		t.Fatalf("Unmarshal without email failed: %v", err)
	}

	if decoded.Email != "" {
		t.Errorf("Expected empty email, got: %s", decoded.Email)
	}

	// Also verify unmarshal with email works
	var decodedWith User
	if err := beve.Unmarshal(dataWith, &decodedWith); err != nil {
		t.Logf("Data with email (len=%d): %x", len(dataWith), dataWith)
		t.Fatalf("Unmarshal with email failed: %v", err)
	}
	if decodedWith.Email != userWithEmail.Email {
		t.Errorf("Email mismatch: got %s, want %s", decodedWith.Email, userWithEmail.Email)
	}
}

// TestProductGeneration tests Product struct code generation
func TestProductGeneration(t *testing.T) {
	product := Product{
		ID:          67890,
		Name:        "Laptop",
		Description: "High-performance laptop",
		Price:       1299.99,
		InStock:     true,
		Quantity:    50,
	}

	// Marshal with generated code
	data, err := product.MarshalBEVE()
	if err != nil {
		t.Fatalf("MarshalBEVE failed: %v", err)
	}

	// Unmarshal
	var decoded Product
	if err := beve.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Verify fields
	if decoded.ID != product.ID {
		t.Errorf("ID mismatch: got %d, want %d", decoded.ID, product.ID)
	}
	if decoded.Name != product.Name {
		t.Errorf("Name mismatch: got %s, want %s", decoded.Name, product.Name)
	}
	if decoded.Description != product.Description {
		t.Errorf("Description mismatch: got %s, want %s", decoded.Description, product.Description)
	}
	if decoded.Price != product.Price {
		t.Errorf("Price mismatch: got %f, want %f", decoded.Price, product.Price)
	}
	if decoded.InStock != product.InStock {
		t.Errorf("InStock mismatch: got %v, want %v", decoded.InStock, product.InStock)
	}
	if decoded.Quantity != product.Quantity {
		t.Errorf("Quantity mismatch: got %d, want %d", decoded.Quantity, product.Quantity)
	}
}

// TestProductOmitEmpty tests omitempty with Product
func TestProductOmitEmpty(t *testing.T) {
	// Product without description
	productWithout := Product{
		ID:          100,
		Name:        "Widget",
		Description: "", // Should be omitted
		Price:       29.99,
		InStock:     true,
		Quantity:    500,
	}

	dataWithout, err := productWithout.MarshalBEVE()
	if err != nil {
		t.Fatalf("MarshalBEVE without description failed: %v", err)
	}

	// Product with description
	productWith := Product{
		ID:          100,
		Name:        "Widget",
		Description: "A useful widget",
		Price:       29.99,
		InStock:     true,
		Quantity:    500,
	}

	dataWith, err := productWith.MarshalBEVE()
	if err != nil {
		t.Fatalf("MarshalBEVE with description failed: %v", err)
	}

	// Data with description should be larger
	if len(dataWith) <= len(dataWithout) {
		t.Errorf("Expected larger data with description: got %d bytes, without %d bytes",
			len(dataWith), len(dataWithout))
	}
}

// TestZeroValues tests handling of zero values
func TestZeroValues(t *testing.T) {
	user := User{
		ID:       0,
		Username: "",
		Email:    "",
		Age:      0,
		IsActive: false,
	}

	data, err := user.MarshalBEVE()
	if err != nil {
		t.Fatalf("MarshalBEVE with zero values failed: %v", err)
	}

	var decoded User
	if err := beve.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal with zero values failed: %v", err)
	}

	// All fields should be zero
	if decoded.ID != 0 {
		t.Errorf("Expected ID=0, got %d", decoded.ID)
	}
	if decoded.Username != "" {
		t.Errorf("Expected empty Username, got %s", decoded.Username)
	}
	if decoded.Age != 0 {
		t.Errorf("Expected Age=0, got %d", decoded.Age)
	}
	if decoded.IsActive {
		t.Error("Expected IsActive=false, got true")
	}
}
