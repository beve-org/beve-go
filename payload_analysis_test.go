package beve_test

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"

	"github.com/beve-org/beve-go"
	"github.com/bytedance/sonic"
	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"
)

// SmallStruct for payload analysis
type SmallStruct struct {
	ID     int64   `json:"id" beve:"id" msgpack:"id" cbor:"id"`
	Name   string  `json:"name" beve:"name" msgpack:"name" cbor:"name"`
	Active bool    `json:"active" beve:"active" msgpack:"active" cbor:"active"`
	Score  float64 `json:"score" beve:"score" msgpack:"score" cbor:"score"`
}

// MediumStruct with more fields
type MediumStruct struct {
	ID       int64             `json:"id" beve:"id" msgpack:"id" cbor:"id"`
	Name     string            `json:"name" beve:"name" msgpack:"name" cbor:"name"`
	Email    string            `json:"email" beve:"email" msgpack:"email" cbor:"email"`
	Age      int               `json:"age" beve:"age" msgpack:"age" cbor:"age"`
	Balance  float64           `json:"balance" beve:"balance" msgpack:"balance" cbor:"balance"`
	Active   bool              `json:"active" beve:"active" msgpack:"active" cbor:"active"`
	Tags     []string          `json:"tags" beve:"tags" msgpack:"tags" cbor:"tags"`
	Metadata map[string]string `json:"metadata" beve:"metadata" msgpack:"metadata" cbor:"metadata"`
}

// LargeStruct with many fields
type LargeStruct struct {
	ID        int64             `json:"id" beve:"id" msgpack:"id" cbor:"id"`
	FirstName string            `json:"first_name" beve:"first_name" msgpack:"first_name" cbor:"first_name"`
	LastName  string            `json:"last_name" beve:"last_name" msgpack:"last_name" cbor:"last_name"`
	Email     string            `json:"email" beve:"email" msgpack:"email" cbor:"email"`
	Phone     string            `json:"phone" beve:"phone" msgpack:"phone" cbor:"phone"`
	Address   string            `json:"address" beve:"address" msgpack:"address" cbor:"address"`
	City      string            `json:"city" beve:"city" msgpack:"city" cbor:"city"`
	Country   string            `json:"country" beve:"country" msgpack:"country" cbor:"country"`
	Age       int               `json:"age" beve:"age" msgpack:"age" cbor:"age"`
	Balance   float64           `json:"balance" beve:"balance" msgpack:"balance" cbor:"balance"`
	Score     float64           `json:"score" beve:"score" msgpack:"score" cbor:"score"`
	Active    bool              `json:"active" beve:"active" msgpack:"active" cbor:"active"`
	Premium   bool              `json:"premium" beve:"premium" msgpack:"premium" cbor:"premium"`
	Tags      []string          `json:"tags" beve:"tags" msgpack:"tags" cbor:"tags"`
	Metadata  map[string]string `json:"metadata" beve:"metadata" msgpack:"metadata" cbor:"metadata"`
}

func TestPayloadSizeAnalysis(t *testing.T) {
	sep := "================================================================================"
	line := "--------------------------------------------------------------------------------"

	fmt.Println("\n" + sep)
	fmt.Println("PAYLOAD SIZE ANALYSIS")
	fmt.Println(sep)

	// Small struct
	small := SmallStruct{
		ID:     123,
		Name:   "John Doe",
		Active: true,
		Score:  95.5,
	}

	fmt.Println("\n1. SMALL STRUCT (4 fields)")
	fmt.Println(line)
	analyzePayloadSize(t, "Small", small)

	// Medium struct
	medium := MediumStruct{
		ID:      12345,
		Name:    "John Doe",
		Email:   "john@example.com",
		Age:     30,
		Balance: 1000.50,
		Active:  true,
		Tags:    []string{"premium", "verified", "active"},
		Metadata: map[string]string{
			"role":       "admin",
			"department": "engineering",
			"level":      "senior",
		},
	}

	fmt.Println("\n2. MEDIUM STRUCT (8 fields)")
	fmt.Println(line)
	analyzePayloadSize(t, "Medium", medium)

	// Large struct
	large := LargeStruct{
		ID:        12345,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     "+1-555-0123",
		Address:   "123 Main St, Apt 4B",
		City:      "New York",
		Country:   "USA",
		Age:       30,
		Balance:   1000.50,
		Score:     95.5,
		Active:    true,
		Premium:   true,
		Tags:      []string{"premium", "verified", "active", "partner"},
		Metadata: map[string]string{
			"role":       "admin",
			"department": "engineering",
			"level":      "senior",
			"team":       "platform",
			"manager":    "Jane Smith",
		},
	}

	fmt.Println("\n3. LARGE STRUCT (15 fields)")
	fmt.Println(line)
	analyzePayloadSize(t, "Large", large)

	// Array of structs
	smallArray := make([]SmallStruct, 10)
	for i := 0; i < 10; i++ {
		smallArray[i] = SmallStruct{
			ID:     int64(100 + i),
			Name:   fmt.Sprintf("User%d", i),
			Active: i%2 == 0,
			Score:  float64(80 + i),
		}
	}

	fmt.Println("\n4. ARRAY OF 10 SMALL STRUCTS")
	fmt.Println(line)
	analyzePayloadSize(t, "Array", smallArray)

	fmt.Println("\n" + sep)
	fmt.Println("ANALYSIS COMPLETE")
	fmt.Println(sep + "\n")
}

func analyzePayloadSize(t *testing.T, name string, data interface{}) {
	// BEVE
	beveData, err := beve.Marshal(data)
	if err != nil {
		t.Fatalf("BEVE marshal failed: %v", err)
	}

	// JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("JSON marshal failed: %v", err)
	}

	// Sonic
	sonicData, err := sonic.Marshal(data)
	if err != nil {
		t.Fatalf("Sonic marshal failed: %v", err)
	}

	// MessagePack
	msgpackData, err := msgpack.Marshal(data)
	if err != nil {
		t.Fatalf("MessagePack marshal failed: %v", err)
	}

	// CBOR
	cborData, err := cbor.Marshal(data)
	if err != nil {
		t.Fatalf("CBOR marshal failed: %v", err)
	}

	// Calculate sizes
	beveSize := len(beveData)
	jsonSize := len(jsonData)
	sonicSize := len(sonicData)
	msgpackSize := len(msgpackData)
	cborSize := len(cborData)

	// Find baseline (smallest)
	baseline := msgpackSize
	if cborSize < baseline {
		baseline = cborSize
	}

	// Print results
	fmt.Printf("%-15s %6d bytes  (baseline: %.2fx)\n", "BEVE:", beveSize, float64(beveSize)/float64(baseline))
	fmt.Printf("%-15s %6d bytes  (baseline: %.2fx)\n", "JSON:", jsonSize, float64(jsonSize)/float64(baseline))
	fmt.Printf("%-15s %6d bytes  (baseline: %.2fx)\n", "Sonic:", sonicSize, float64(sonicSize)/float64(baseline))
	fmt.Printf("%-15s %6d bytes  (baseline: %.2fx)\n", "MessagePack:", msgpackSize, float64(msgpackSize)/float64(baseline))
	fmt.Printf("%-15s %6d bytes  (baseline: %.2fx)\n", "CBOR:", cborSize, float64(cborSize)/float64(baseline))

	fmt.Printf("\nBEVE overhead:\n")
	fmt.Printf("  vs MessagePack: %+d bytes (%.1f%% larger)\n", beveSize-msgpackSize, (float64(beveSize)/float64(msgpackSize)-1)*100)
	fmt.Printf("  vs CBOR:        %+d bytes (%.1f%% larger)\n", beveSize-cborSize, (float64(beveSize)/float64(cborSize)-1)*100)
	fmt.Printf("  vs JSON:        %+d bytes (%.1f%% %s)\n", beveSize-jsonSize, math.Abs(float64(beveSize)/float64(jsonSize)-1)*100,
		map[bool]string{true: "smaller", false: "larger"}[beveSize < jsonSize])
}
