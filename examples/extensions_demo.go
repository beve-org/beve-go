package main

import (
	"fmt"
	"time"

	beve "github.com/beve-org/beve-go"
)

// User represents a sample user struct
type User struct {
	Name  string `beve:"name"`
	Email string `beve:"email"`
	Age   int    `beve:"age"`
}

// Post represents a blog post with nested comments
type Post struct {
	Title    string    `beve:"title"`
	Author   string    `beve:"author"`
	Created  time.Time `beve:"created"`
	Comments []Comment `beve:"comments"`
}

// Comment represents a user comment
type Comment struct {
	Author string `beve:"author"`
	Text   string `beve:"text"`
	Likes  int    `beve:"likes"`
}

func main() {
	fmt.Println("🚀 BEVE Extensions Demo")
	fmt.Println("=" + string(make([]byte, 50)) + "=")

	// Example 1: Typed Object Arrays (Extension 1)
	fmt.Println("\n📦 Example 1: Typed Object Arrays (48% smaller)")
	demoTypedArrays()

	// Example 2: Timestamps (Extension 4)
	fmt.Println("\n⏰ Example 2: Nanosecond Timestamps")
	demoTimestamps()

	// Example 3: UUIDs (Extension 8)
	fmt.Println("\n🆔 Example 3: Binary UUIDs (50% smaller)")
	demoUUIDs()

	// Example 4: Field Index (Extension 0)
	fmt.Println("\n🔍 Example 4: O(1) Field Access")
	demoFieldIndex()

	// Example 5: Auto-Detection
	fmt.Println("\n🎯 Example 5: Auto-Detection")
	demoAutoDetection()
}

func demoTypedArrays() {
	users := []User{
		{"Alice", "alice@example.com", 30},
		{"Bob", "bob@example.com", 25},
		{"Charlie", "charlie@example.com", 35},
	}

	// Standard BEVE encoding
	standardData, _ := beve.Marshal(users)
	fmt.Printf("  Standard BEVE: %d bytes\n", len(standardData))

	// Typed array encoding (Extension 1)
	typedData, _ := beve.MarshalTyped(users)
	fmt.Printf("  Typed Array:   %d bytes\n", len(typedData))

	savings := float64(len(standardData)-len(typedData)) / float64(len(standardData)) * 100
	fmt.Printf("  💰 Savings:     %.1f%%\n", savings)

	// Decode
	var decoded []map[string]interface{}
	beve.UnmarshalAuto(typedData, &decoded)
	fmt.Printf("  ✅ Decoded:     %d users\n", len(decoded))
}

func demoTimestamps() {
	now := time.Now()

	// Encode timestamp with nanosecond precision
	data, _ := beve.MarshalTimestamp(now)
	fmt.Printf("  Encoded size:  %d bytes (UTC)\n", len(data))

	// Decode
	decoded, _ := beve.UnmarshalTimestamp(data)
	fmt.Printf("  Original:      %v\n", now.Format(time.RFC3339Nano))
	fmt.Printf("  Decoded:       %v\n", decoded.Format(time.RFC3339Nano))
	fmt.Printf("  ✅ Exact match: %v\n", now.Equal(decoded))
}

func demoUUIDs() {
	// Sample UUID v4
	uuid := [16]byte{
		0x55, 0x0e, 0x84, 0x00, 0xe2, 0x9b, 0x41, 0xd4,
		0xa7, 0x16, 0x44, 0x66, 0x55, 0x44, 0x00, 0x00,
	}

	// String representation
	uuidString := "550e8400-e29b-41d4-a716-446655440000"
	fmt.Printf("  String:        %s (%d bytes)\n", uuidString, len(uuidString))

	// Binary encoding
	data, _ := beve.MarshalUUID(uuid)
	fmt.Printf("  Binary BEVE:   %d bytes\n", len(data))

	savings := float64(len(uuidString)-len(data)) / float64(len(uuidString)) * 100
	fmt.Printf("  💰 Savings:     %.1f%%\n", savings)

	// Decode
	decoded, _ := beve.UnmarshalUUIDString(data)
	fmt.Printf("  ✅ Decoded:     %s\n", decoded)
}

func demoFieldIndex() {
	// Large object
	obj := map[string]interface{}{
		"id":       12345,
		"name":     "Alice Smith",
		"email":    "alice@example.com",
		"phone":    "+1-555-0123",
		"address":  "123 Main St, City, State",
		"profile":  map[string]interface{}{"bio": "Software engineer", "skills": []string{"Go", "Rust", "C++"}},
		"settings": map[string]interface{}{"theme": "dark", "notifications": true},
		"metadata": map[string]interface{}{"created": time.Now(), "updated": time.Now()},
	}

	// Encode with field index
	data, _ := beve.EncodeIndexedObject(obj)
	fmt.Printf("  Indexed object: %d bytes\n", len(data))

	// Read single field (O(1) lookup, no full decode)
	email, _ := beve.ReadFieldByName(data, "email")
	fmt.Printf("  ✅ Read 'email': %v (without decoding other fields)\n", email)

	// Full decode for comparison
	decoded, _ := beve.DecodeIndexedObject(data)
	fmt.Printf("  Full decode:    %d fields\n", len(decoded))
}

func demoAutoDetection() {
	// Different formats
	users := []User{
		{"Alice", "alice@example.com", 30},
		{"Bob", "bob@example.com", 25},
	}

	// Encode with auto-detection
	data, _ := beve.MarshalAuto(users)
	fmt.Printf("  Encoded:       %d bytes\n", len(data))

	// Detect format
	encoding := beve.DetectEncoding(data)
	fmt.Printf("  Detected:      %s\n", encoding)

	isExtension := beve.IsExtension(data)
	fmt.Printf("  Is extension:  %v\n", isExtension)

	if isExtension {
		extID, _ := beve.GetExtensionID(data)
		fmt.Printf("  Extension ID:  %d\n", extID)
	}

	// Decode (auto-detects format)
	var decoded []map[string]interface{}
	beve.Unmarshal(data, &decoded)
	fmt.Printf("  ✅ Decoded:     %d users\n", len(decoded))
}
