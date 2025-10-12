package beve

import (
	"testing"
)

func TestWASMScenario(t *testing.T) {
	// Simulate what WASM does
	data := map[string]interface{}{
		"id":     int64(123),
		"name":   "Alice",
		"active": true,
		"metadata": map[string]interface{}{
			"created": "2025-10-12",
			"tags":    []interface{}{"admin", "user"},
		},
	}

	// Marshal
	encoded, err := Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Encoded %d bytes", len(encoded))

	// Try unmarshal as map[string]RawMessage
	var rawMap map[string]RawMessage
	if err := Unmarshal(encoded, &rawMap); err != nil {
		t.Fatalf("Unmarshal to map[string]RawMessage failed: %v", err)
	}
	t.Logf("Decoded map with %d keys", len(rawMap))

	// Try to decode each value
	for k, v := range rawMap {
		t.Logf("Key: %s, RawMessage: %d bytes", k, len(v))

		// Try to decode the RawMessage
		var decoded interface{}
		if err := Unmarshal(v, &decoded); err != nil {
			t.Errorf("Failed to unmarshal RawMessage for key %s: %v", k, err)
		} else {
			t.Logf("  Decoded: %v", decoded)
		}
	}
}

func TestWASMArrayScenario(t *testing.T) {
	// Test array unmarshaling to interface{}
	data := []interface{}{"admin", "user", int64(123)}

	// Marshal array
	encoded, err := Marshal(data)
	if err != nil {
		t.Fatalf("Marshal array failed: %v", err)
	}
	t.Logf("Encoded array: %d bytes", len(encoded))

	// Unmarshal to interface{}
	var decoded interface{}
	if err := Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal array to interface{} failed: %v", err)
	}
	t.Logf("Decoded array: %v", decoded)

	// Verify it's a slice
	slice, ok := decoded.([]interface{})
	if !ok {
		t.Fatalf("Expected []interface{}, got %T", decoded)
	}
	if len(slice) != 3 {
		t.Fatalf("Expected 3 elements, got %d", len(slice))
	}
	t.Logf("Array elements: %v, %v, %v", slice[0], slice[1], slice[2])
}

func TestWASMNestedArrayScenario(t *testing.T) {
	// Test nested structure with arrays
	data := map[string]interface{}{
		"name":  "Alice",
		"email": []interface{}{"alice@example.com", "alice@work.com"},
		"tags":  []interface{}{"admin", "user"},
	}

	// Marshal
	encoded, err := Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Encoded %d bytes", len(encoded))

	// Unmarshal to interface{}
	var decoded interface{}
	if err := Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal to interface{} failed: %v", err)
	}
	t.Logf("Decoded: %v", decoded)

	// Verify structure
	result, ok := decoded.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected map[string]interface{}, got %T", decoded)
	}

	// Check email array
	email, ok := result["email"].([]interface{})
	if !ok {
		t.Fatalf("Expected email to be []interface{}, got %T", result["email"])
	}
	if len(email) != 2 {
		t.Fatalf("Expected 2 email addresses, got %d", len(email))
	}
	t.Logf("Email addresses: %v", email)
}
