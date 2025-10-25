package translatornative

import (
	"testing"
)

// TestFromJSONToJSON_Simple tests JSON→BEVE→JSON round-trip.
func TestFromJSONToJSON_Simple(t *testing.T) {
	tests := []struct {
		name string
		json string
	}{
		{
			name: "null",
			json: `null`,
		},
		{
			name: "boolean true",
			json: `true`,
		},
		{
			name: "boolean false",
			json: `false`,
		},
		{
			name: "number integer",
			json: `42`,
		},
		{
			name: "number float",
			json: `3.14`,
		},
		{
			name: "string",
			json: `"hello"`,
		},
		{
			name: "empty object",
			json: `{}`,
		},
		{
			name: "empty array",
			json: `[]`,
		},
		{
			name: "simple object",
			json: `{"name":"Alice","age":30,"active":true}`,
		},
		{
			name: "simple array",
			json: `[1,2,3,4,5]`,
		},
		{
			name: "nested object",
			json: `{"user":{"name":"Bob","email":"bob@example.com"},"active":true}`,
		},
		{
			name: "array of objects",
			json: `[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// JSON → BEVE
			beveData, err := FromJSON([]byte(tt.json))
			if err != nil {
				t.Fatalf("FromJSON failed: %v", err)
			}

			if len(beveData) == 0 {
				t.Fatal("BEVE data is empty")
			}

			t.Logf("JSON: %s", tt.json)
			t.Logf("BEVE: %d bytes: %x", len(beveData), beveData)

			// BEVE → JSON
			jsonData, err := ToJSON(beveData)
			if err != nil {
				t.Fatalf("ToJSON failed: %v", err)
			}

			if len(jsonData) == 0 {
				t.Fatal("JSON data is empty after round-trip")
			}

			t.Logf("JSON (round-trip): %s", string(jsonData))

			// Validate it's valid JSON
			if !ValidateJSON(jsonData) {
				t.Fatal("Round-trip JSON is invalid")
			}
		})
	}
}

// TestValidateJSON tests JSON validation.
func TestValidateJSON(t *testing.T) {
	tests := []struct {
		name  string
		json  string
		valid bool
	}{
		{"valid object", `{"key":"value"}`, true},
		{"valid array", `[1,2,3]`, true},
		{"valid string", `"test"`, true},
		{"invalid json", `{key:value}`, false},
		{"unterminated string", `"test`, false},
		{"empty", ``, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateJSON([]byte(tt.json))
			if result != tt.valid {
				t.Errorf("ValidateJSON(%q) = %v, want %v", tt.json, result, tt.valid)
			}
		})
	}
}

// TestValidateBEVE tests BEVE validation.
func TestValidateBEVE(t *testing.T) {
	tests := []struct {
		name  string
		beve  []byte
		valid bool
	}{
		{"valid BEVE", []byte{0x02, 0x04, 't', 'e', 's', 't'}, true}, // string "test"
		{"empty", []byte{}, false},
		{"invalid header", []byte{0xFF}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateBEVE(tt.beve)
			if result != tt.valid {
				t.Errorf("ValidateBEVE() = %v, want %v", result, tt.valid)
			}
		})
	}
}
