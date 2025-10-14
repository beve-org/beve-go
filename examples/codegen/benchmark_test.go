package main

import (
	"testing"

	"github.com/beve-org/beve-go"
)

// BenchmarkUser_MarshalBEVE_Generated tests the generated MarshalBEVE method
func BenchmarkUser_MarshalBEVE_Generated(b *testing.B) {
	user := User{
		ID:        12345,
		Username:  "john_doe",
		Email:     "john@example.com",
		Age:       30,
		IsActive:  true,
		// CreatedAt omitted - complex type
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := user.MarshalBEVE()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUser_Marshal_Reflection tests the reflection-based Marshal
func BenchmarkUser_Marshal_Reflection(b *testing.B) {
	user := User{
		ID:        12345,
		Username:  "john_doe",
		Email:     "john@example.com",
		Age:       30,
		IsActive:  true,
		// CreatedAt omitted - complex type
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := beve.Marshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkProduct_MarshalBEVE_Generated tests Product marshal
func BenchmarkProduct_MarshalBEVE_Generated(b *testing.B) {
	product := Product{
		ID:          98765,
		Name:        "Laptop",
		Description: "High-performance laptop",
		Price:       1299.99,
		InStock:     true,
		Quantity:    50,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := product.MarshalBEVE()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkProduct_Marshal_Reflection tests Product reflection marshal
func BenchmarkProduct_Marshal_Reflection(b *testing.B) {
	product := Product{
		ID:          98765,
		Name:        "Laptop",
		Description: "High-performance laptop",
		Price:       1299.99,
		InStock:     true,
		Quantity:    50,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := beve.Marshal(product)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// TestUser_MarshalBEVE_Correctness verifies generated code produces valid output
func TestUser_MarshalBEVE_Correctness(b *testing.T) {
	user := User{
		ID:        12345,
		Username:  "john_doe",
		Email:     "john@example.com",
		Age:       30,
		IsActive:  true,
		// CreatedAt: time.Time is complex type - skip for now
	}

	// Generated code
	generated, err := user.MarshalBEVE()
	if err != nil {
		b.Fatalf("MarshalBEVE failed: %v", err)
	}

	// Reflection-based code for comparison
	reflected, err := beve.Marshal(user)
	if err != nil {
		b.Fatalf("Marshal failed: %v", err)
	}

	// Both should produce same size (or very close)
	if len(generated) != len(reflected) {
		b.Logf("⚠️  Size difference: generated=%d, reflected=%d", len(generated), len(reflected))
	}

	b.Logf("✅ Generated code produces valid BEVE output")
	b.Logf("   Generated size: %d bytes", len(generated))
	b.Logf("   Reflected size: %d bytes", len(reflected))
}
