package beve

import (
	"testing"
)

// Struct with more primitive fields (good for cache)
type CachedStruct struct {
	ID       int     `beve:"id"`
	Name     string  `beve:"name"`
	Age      int     `beve:"age"`
	Email    string  `beve:"email"`
	Score    float64 `beve:"score"`
	IsActive bool    `beve:"is_active"`
	Balance  float64 `beve:"balance"`
	Count    int     `beve:"count"`
}

func BenchmarkCachedStruct_BEVE_Marshal(b *testing.B) {
	s := CachedStruct{
		ID:       42,
		Name:     "Test User",
		Age:      25,
		Email:    "test@example.com",
		Score:    98.5,
		IsActive: true,
		Balance:  1234.56,
		Count:    100,
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data, _ := Marshal(s)
		benchBytesSink = data
	}
}

// Comparison: Same struct with slice (slower)
type UserWithSlice struct {
	ID       int      `beve:"id"`
	Name     string   `beve:"name"`
	Age      int      `beve:"age"`
	Email    string   `beve:"email"`
	Score    float64  `beve:"score"`
	IsActive bool     `beve:"is_active"`
	Balance  float64  `beve:"balance"`
	Count    int      `beve:"count"`
	Tags     []string `beve:"tags"`
}

func BenchmarkUserWithSlice_BEVE_Marshal(b *testing.B) {
	s := UserWithSlice{
		ID:       42,
		Name:     "Test User",
		Age:      25,
		Email:    "test@example.com",
		Score:    98.5,
		IsActive: true,
		Balance:  1234.56,
		Count:    100,
		Tags:     []string{"premium", "verified", "active"},
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data, _ := Marshal(s)
		benchBytesSink = data
	}
}
