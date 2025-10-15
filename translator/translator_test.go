package translator

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/beve-org/beve-go"
)

// TestFromJSON_BasicTypes tests conversion of primitive JSON types
func TestFromJSON_BasicTypes(t *testing.T) {
	tests := []struct {
		name     string
		jsonStr  string
		wantErr  bool
		validate func(t *testing.T, beveData []byte)
	}{
		{
			name:    "null value",
			jsonStr: "null",
			validate: func(t *testing.T, beveData []byte) {
				if len(beveData) != 1 || beveData[0] != 0x00 {
					t.Errorf("Expected null (0x00), got %x", beveData)
				}
			},
		},
		{
			name:    "boolean true",
			jsonStr: "true",
			validate: func(t *testing.T, beveData []byte) {
				if len(beveData) != 1 || beveData[0] != 0x18 {
					t.Errorf("Expected true (0x18), got %x", beveData)
				}
			},
		},
		{
			name:    "boolean false",
			jsonStr: "false",
			validate: func(t *testing.T, beveData []byte) {
				if len(beveData) != 1 || beveData[0] != 0x08 {
					t.Errorf("Expected false (0x08), got %x", beveData)
				}
			},
		},
		{
			name:    "positive integer",
			jsonStr: "42",
			validate: func(t *testing.T, beveData []byte) {
				// Should encode as small uint
				var result interface{}
				if err := beve.Unmarshal(beveData, &result); err != nil {
					t.Fatalf("Unmarshal failed: %v", err)
				}
				num, ok := result.(float64) // JSON numbers decode as float64
				if !ok || num != 42 {
					t.Errorf("Expected 42, got %v (%T)", result, result)
				}
			},
		},
		{
			name:    "negative integer",
			jsonStr: "-123",
			validate: func(t *testing.T, beveData []byte) {
				var result interface{}
				if err := beve.Unmarshal(beveData, &result); err != nil {
					t.Fatalf("Unmarshal failed: %v", err)
				}
				num, ok := result.(float64)
				if !ok || num != -123 {
					t.Errorf("Expected -123, got %v", result)
				}
			},
		},
		{
			name:    "float",
			jsonStr: "3.14159",
			validate: func(t *testing.T, beveData []byte) {
				var result interface{}
				if err := beve.Unmarshal(beveData, &result); err != nil {
					t.Fatalf("Unmarshal failed: %v", err)
				}
				num, ok := result.(float64)
				if !ok || num < 3.14 || num > 3.15 {
					t.Errorf("Expected ~3.14, got %v", result)
				}
			},
		},
		{
			name:    "simple string",
			jsonStr: `"hello"`,
			validate: func(t *testing.T, beveData []byte) {
				var result string
				if err := beve.Unmarshal(beveData, &result); err != nil {
					t.Fatalf("Unmarshal failed: %v", err)
				}
				if result != "hello" {
					t.Errorf("Expected 'hello', got '%s'", result)
				}
			},
		},
		{
			name:    "empty string",
			jsonStr: `""`,
			validate: func(t *testing.T, beveData []byte) {
				var result string
				if err := beve.Unmarshal(beveData, &result); err != nil {
					t.Fatalf("Unmarshal failed: %v", err)
				}
				if result != "" {
					t.Errorf("Expected empty string, got '%s'", result)
				}
			},
		},
		{
			name:    "string with unicode",
			jsonStr: `"Hello 世界 🌍"`,
			validate: func(t *testing.T, beveData []byte) {
				var result string
				if err := beve.Unmarshal(beveData, &result); err != nil {
					t.Fatalf("Unmarshal failed: %v", err)
				}
				if result != "Hello 世界 🌍" {
					t.Errorf("Unicode mismatch: got '%s'", result)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beveData, err := FromJSON([]byte(tt.jsonStr))

			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("FromJSON failed: %v", err)
			}

			if tt.validate != nil {
				tt.validate(t, beveData)
			}
		})
	}
}

// TestFromJSON_Arrays tests JSON array conversion
func TestFromJSON_Arrays(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		wantErr bool
	}{
		{
			name:    "empty array",
			jsonStr: "[]",
		},
		{
			name:    "number array",
			jsonStr: "[1, 2, 3, 4, 5]",
		},
		{
			name:    "string array",
			jsonStr: `["apple", "banana", "cherry"]`,
		},
		{
			name:    "mixed array",
			jsonStr: `[1, "two", true, null, 3.14]`,
		},
		{
			name:    "nested array",
			jsonStr: `[[1, 2], [3, 4], [5, 6]]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beveData, err := FromJSON([]byte(tt.jsonStr))
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("FromJSON failed: %v", err)
			}

			// Verify round-trip
			var result interface{}
			if err := beve.Unmarshal(beveData, &result); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			// Re-encode to JSON and compare
			jsonData, err := json.Marshal(result)
			if err != nil {
				t.Fatalf("JSON marshal failed: %v", err)
			}

			// Normalize both JSON strings for comparison
			var expected, actual interface{}
			json.Unmarshal([]byte(tt.jsonStr), &expected)
			json.Unmarshal(jsonData, &actual)

			expectedJSON, _ := json.Marshal(expected)
			actualJSON, _ := json.Marshal(actual)

			if string(expectedJSON) != string(actualJSON) {
				t.Errorf("Round-trip mismatch:\nExpected: %s\nGot:      %s", expectedJSON, actualJSON)
			}
		})
	}
}

// TestFromJSON_Objects tests JSON object conversion
func TestFromJSON_Objects(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		wantErr bool
	}{
		{
			name:    "empty object",
			jsonStr: "{}",
		},
		{
			name:    "simple object",
			jsonStr: `{"name": "John", "age": 30}`,
		},
		{
			name:    "nested object",
			jsonStr: `{"user": {"name": "Alice", "email": "alice@example.com"}}`,
		},
		{
			name:    "object with array",
			jsonStr: `{"items": [1, 2, 3], "count": 3}`,
		},
		{
			name:    "complex nested structure",
			jsonStr: `{"id": 1, "data": {"values": [10, 20, 30], "metadata": {"created": true}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beveData, err := FromJSON([]byte(tt.jsonStr))
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("FromJSON failed: %v", err)
			}

			// Verify round-trip
			jsonData, err := ToJSON(beveData)
			if err != nil {
				t.Fatalf("ToJSON failed: %v", err)
			}

			// Compare normalized JSON
			var expected, actual interface{}
			json.Unmarshal([]byte(tt.jsonStr), &expected)
			json.Unmarshal(jsonData, &actual)

			expectedJSON, _ := json.Marshal(expected)
			actualJSON, _ := json.Marshal(actual)

			if string(expectedJSON) != string(actualJSON) {
				t.Errorf("Round-trip mismatch:\nExpected: %s\nGot:      %s", expectedJSON, actualJSON)
			}
		})
	}
}

// TestToJSON_RoundTrip tests BEVE → JSON → BEVE round-trip
func TestToJSON_RoundTrip(t *testing.T) {
	testData := []interface{}{
		nil,
		true,
		false,
		int64(42),
		float64(3.14),
		"hello world",
		[]interface{}{1, 2, 3},
		map[string]interface{}{"key": "value"},
	}

	for i, original := range testData {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			// Encode to BEVE
			beveData, err := beve.Marshal(original)
			if err != nil {
				t.Fatalf("BEVE marshal failed: %v", err)
			}

			// Convert to JSON
			jsonData, err := ToJSON(beveData)
			if err != nil {
				t.Fatalf("ToJSON failed: %v", err)
			}

			// Convert back to BEVE
			beveData2, err := FromJSON(jsonData)
			if err != nil {
				t.Fatalf("FromJSON failed: %v", err)
			}

			// Decode and compare
			var result interface{}
			if err := beve.Unmarshal(beveData2, &result); err != nil {
				t.Fatalf("BEVE unmarshal failed: %v", err)
			}

			// Compare values
			originalJSON, _ := json.Marshal(original)
			resultJSON, _ := json.Marshal(result)

			if string(originalJSON) != string(resultJSON) {
				t.Errorf("Round-trip mismatch:\nOriginal: %s\nResult:   %s", originalJSON, resultJSON)
			}
		})
	}
}

// TestFromJSONString tests the string convenience wrapper
func TestFromJSONString(t *testing.T) {
	jsonStr := `{"message": "hello"}`

	beveData, err := FromJSONString(jsonStr)
	if err != nil {
		t.Fatalf("FromJSONString failed: %v", err)
	}

	var result map[string]interface{}
	if err := beve.Unmarshal(beveData, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result["message"] != "hello" {
		t.Errorf("Expected 'hello', got '%v'", result["message"])
	}
}

// TestToJSONString tests string output
func TestToJSONString(t *testing.T) {
	data := map[string]interface{}{
		"id":   123,
		"name": "Test",
	}

	beveData, err := beve.Marshal(data)
	if err != nil {
		t.Fatalf("BEVE marshal failed: %v", err)
	}

	jsonStr, err := ToJSONString(beveData)
	if err != nil {
		t.Fatalf("ToJSONString failed: %v", err)
	}

	if !strings.Contains(jsonStr, "123") || !strings.Contains(jsonStr, "Test") {
		t.Errorf("JSON string missing expected content: %s", jsonStr)
	}
}

// TestToJSONIndent tests pretty-printed JSON
func TestToJSONIndent(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "Alice",
			"age":  30,
		},
	}

	beveData, err := beve.Marshal(data)
	if err != nil {
		t.Fatalf("BEVE marshal failed: %v", err)
	}

	jsonStr, err := ToJSONIndent(beveData, "", "  ")
	if err != nil {
		t.Fatalf("ToJSONIndent failed: %v", err)
	}

	// Should contain newlines and indentation
	if !strings.Contains(jsonStr, "\n") || !strings.Contains(jsonStr, "  ") {
		t.Error("JSON output is not indented")
	}

	t.Logf("Indented JSON:\n%s", jsonStr)
}

// TestValidateJSON tests JSON validation
func TestValidateJSON(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
	}{
		{"valid object", `{"key": "value"}`, true},
		{"valid array", `[1, 2, 3]`, true},
		{"valid null", `null`, true},
		{"invalid missing quote", `{"key: "value"}`, false},
		{"invalid trailing comma", `{"key": "value",}`, false},
		{"empty input", ``, false},
		{"invalid json", `{key: value}`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateJSON([]byte(tt.input))
			if result != tt.valid {
				t.Errorf("ValidateJSON(%q) = %v, want %v", tt.input, result, tt.valid)
			}
		})
	}
}

// TestValidateBEVE tests BEVE validation
func TestValidateBEVE(t *testing.T) {
	// Valid BEVE data
	validData, _ := beve.Marshal(map[string]interface{}{"key": "value"})

	if !ValidateBEVE(validData) {
		t.Error("Valid BEVE data rejected")
	}

	// Invalid BEVE data
	invalidData := []byte{0xFF, 0xFF, 0xFF}
	if ValidateBEVE(invalidData) {
		t.Error("Invalid BEVE data accepted")
	}

	// Empty data
	if ValidateBEVE([]byte{}) {
		t.Error("Empty data accepted")
	}
}

// TestFromJSONWithStats tests conversion statistics
func TestFromJSONWithStats(t *testing.T) {
	jsonData := []byte(`{"id": 12345, "name": "Test User", "active": true}`)

	beveData, stats, err := FromJSONWithStats(jsonData)
	if err != nil {
		t.Fatalf("FromJSONWithStats failed: %v", err)
	}

	if stats.OriginalSize != len(jsonData) {
		t.Errorf("OriginalSize = %d, want %d", stats.OriginalSize, len(jsonData))
	}

	if stats.ConvertedSize != len(beveData) {
		t.Errorf("ConvertedSize = %d, want %d", stats.ConvertedSize, len(beveData))
	}

	if stats.Ratio <= 0 {
		t.Error("Ratio should be positive")
	}

	// BEVE should typically be smaller than JSON
	if stats.Savings <= 0 {
		t.Logf("Warning: BEVE not smaller than JSON (savings: %.2f%%)", stats.Savings*100)
	}

	t.Logf("Conversion stats: %d bytes → %d bytes (%.1f%% savings)",
		stats.OriginalSize, stats.ConvertedSize, stats.Savings*100)
}

// TestToJSONWithStats tests reverse conversion statistics
func TestToJSONWithStats(t *testing.T) {
	data := map[string]interface{}{
		"id":   12345,
		"name": "Test",
	}

	beveData, err := beve.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	jsonData, stats, err := ToJSONWithStats(beveData)
	if err != nil {
		t.Fatalf("ToJSONWithStats failed: %v", err)
	}

	if stats.OriginalSize != len(beveData) {
		t.Errorf("OriginalSize = %d, want %d", stats.OriginalSize, len(beveData))
	}

	if stats.ConvertedSize != len(jsonData) {
		t.Errorf("ConvertedSize = %d, want %d", stats.ConvertedSize, len(jsonData))
	}

	// JSON is typically larger, so savings will be negative
	t.Logf("JSON overhead: %d bytes → %d bytes (%.1f%% larger)",
		stats.OriginalSize, stats.ConvertedSize, -stats.Savings*100)
}

// TestErrorHandling tests error cases
func TestErrorHandling(t *testing.T) {
	t.Run("empty JSON", func(t *testing.T) {
		_, err := FromJSON([]byte{})
		if err == nil {
			t.Error("Expected error for empty JSON")
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		_, err := FromJSON([]byte(`{invalid}`))
		if err == nil {
			t.Error("Expected error for invalid JSON")
		}
	})

	t.Run("empty BEVE", func(t *testing.T) {
		_, err := ToJSON([]byte{})
		if err == nil {
			t.Error("Expected error for empty BEVE")
		}
	})

	t.Run("invalid BEVE", func(t *testing.T) {
		_, err := ToJSON([]byte{0xFF, 0xFF})
		if err == nil {
			t.Error("Expected error for invalid BEVE")
		}
	})
}

// TestLargeData tests conversion of larger structures
func TestLargeData(t *testing.T) {
	// Create a moderately large JSON structure
	data := make(map[string]interface{})
	for i := 0; i < 100; i++ {
		data[string(rune('a'+i%26))+string(rune('0'+i/26))] = map[string]interface{}{
			"id":    i,
			"value": float64(i) * 3.14,
			"text":  "Lorem ipsum dolor sit amet",
			"flag":  i%2 == 0,
		}
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("JSON marshal failed: %v", err)
	}

	t.Logf("Test data size: %d bytes JSON", len(jsonData))

	// Convert to BEVE
	beveData, stats, err := FromJSONWithStats(jsonData)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}

	t.Logf("BEVE size: %d bytes (%.1f%% savings)", len(beveData), stats.Savings*100)

	// Convert back to JSON
	jsonData2, err := ToJSON(beveData)
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	// Verify round-trip
	var original, result interface{}
	json.Unmarshal(jsonData, &original)
	json.Unmarshal(jsonData2, &result)

	originalJSON, _ := json.Marshal(original)
	resultJSON, _ := json.Marshal(result)

	if string(originalJSON) != string(resultJSON) {
		t.Error("Large data round-trip failed")
	}
}
