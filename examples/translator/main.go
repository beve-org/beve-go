package main

import (
	"fmt"
	"log"

	"github.com/beve-org/beve-go/translator"
)

func main() {
	// Example 1: Simple JSON to BEVE conversion
	fmt.Println("=== Example 1: JSON → BEVE ===")
	jsonStr := `{
		"id": 123,
		"name": "Alice",
		"email": "alice@example.com",
		"active": true,
		"score": 95.5
	}`

	beveData, err := translator.FromJSONString(jsonStr)
	if err != nil {
		log.Fatalf("Conversion failed: %v", err)
	}

	fmt.Printf("JSON size: %d bytes\n", len(jsonStr))
	fmt.Printf("BEVE size: %d bytes\n", len(beveData))
	fmt.Printf("Savings: %.1f%%\n\n", (1.0-float64(len(beveData))/float64(len(jsonStr)))*100)

	// Example 2: BEVE back to JSON
	fmt.Println("=== Example 2: BEVE → JSON ===")
	jsonOutput, err := translator.ToJSONIndent(beveData, "", "  ")
	if err != nil {
		log.Fatalf("Conversion failed: %v", err)
	}
	fmt.Println("Reconstructed JSON:")
	fmt.Println(jsonOutput)
	fmt.Println()

	// Example 3: Conversion with statistics
	fmt.Println("=== Example 3: Conversion Statistics ===")
	jsonData := []byte(`{
		"users": [
			{"id": 1, "name": "Alice"},
			{"id": 2, "name": "Bob"},
			{"id": 3, "name": "Charlie"}
		],
		"count": 3,
		"timestamp": 1699564800
	}`)

	beveResult, stats, err := translator.FromJSONWithStats(jsonData)
	if err != nil {
		log.Fatalf("Conversion failed: %v", err)
	}

	fmt.Printf("Original:  %d bytes (JSON)\n", stats.OriginalSize)
	fmt.Printf("Converted: %d bytes (BEVE)\n", stats.ConvertedSize)
	fmt.Printf("Ratio:     %.3f\n", stats.Ratio)
	fmt.Printf("Savings:   %.1f%%\n\n", stats.Savings*100)

	// Example 4: Validation
	fmt.Println("=== Example 4: Validation ===")
	validJSON := []byte(`{"key": "value"}`)
	invalidJSON := []byte(`{key: value}`)

	fmt.Printf("Valid JSON:   %v\n", translator.ValidateJSON(validJSON))
	fmt.Printf("Invalid JSON: %v\n", translator.ValidateJSON(invalidJSON))
	fmt.Printf("Valid BEVE:   %v\n\n", translator.ValidateBEVE(beveResult))

	// Example 5: Array conversion
	fmt.Println("=== Example 5: Array Conversion ===")
	arrayJSON := []byte(`[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]`)
	arrayBEVE, err := translator.FromJSON(arrayJSON)
	if err != nil {
		log.Fatalf("Array conversion failed: %v", err)
	}

	fmt.Printf("JSON array:  %d bytes\n", len(arrayJSON))
	fmt.Printf("BEVE array:  %d bytes\n", len(arrayBEVE))
	fmt.Printf("Savings:     %.1f%%\n\n", (1.0-float64(len(arrayBEVE))/float64(len(arrayJSON)))*100)

	// Example 6: Round-trip conversion
	fmt.Println("=== Example 6: Round-Trip Test ===")
	original := `{"message":"Hello, BEVE!"}`

	// JSON → BEVE
	beve1, _ := translator.FromJSONString(original)
	fmt.Printf("Step 1: JSON → BEVE (%d bytes)\n", len(beve1))

	// BEVE → JSON
	json1, _ := translator.ToJSONString(beve1)
	fmt.Printf("Step 2: BEVE → JSON (%d bytes)\n", len(json1))

	// JSON → BEVE again
	beve2, _ := translator.FromJSONString(json1)
	fmt.Printf("Step 3: JSON → BEVE (%d bytes)\n", len(beve2))

	// Verify consistency
	json2, _ := translator.ToJSONString(beve2)
	if json1 == json2 {
		fmt.Println("✅ Round-trip successful: data preserved!")
	} else {
		fmt.Println("❌ Round-trip failed: data corrupted")
	}
	fmt.Println()

	// Example 7: Complex nested structure
	fmt.Println("=== Example 7: Complex Nested Structure ===")
	complexJSON := `{
		"user": {
			"id": 12345,
			"profile": {
				"name": "John Doe",
				"age": 30,
				"contacts": {
					"email": "john@example.com",
					"phone": "+1234567890"
				}
			},
			"preferences": {
				"theme": "dark",
				"notifications": true,
				"language": "en"
			}
		},
		"metadata": {
			"created": "2025-10-15T10:30:00Z",
			"updated": "2025-10-15T11:45:00Z",
			"version": 3
		}
	}`

	complexBEVE, stats2, _ := translator.FromJSONWithStats([]byte(complexJSON))
	fmt.Printf("Complex structure conversion:\n")
	fmt.Printf("  JSON:  %d bytes\n", stats2.OriginalSize)
	fmt.Printf("  BEVE:  %d bytes\n", stats2.ConvertedSize)
	fmt.Printf("  Saved: %.1f%%\n\n", stats2.Savings*100)

	// Verify by converting back
	reconstructed, _ := translator.ToJSONIndent(complexBEVE, "", "  ")
	fmt.Println("Reconstructed structure:")
	fmt.Println(reconstructed)
}
