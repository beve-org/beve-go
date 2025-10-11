package beve

import (
	"math"
	"reflect"
	"testing"
	"time"
)

// TestMarshalFloat64 tests float64 marshaling (currently 0% coverage)
func TestMarshalFloat64(t *testing.T) {
	tests := []struct {
		name  string
		input float64
	}{
		{"zero", 0.0},
		{"positive", 3.14159},
		{"negative", -2.71828},
		{"max", math.MaxFloat64},
		{"min", -math.MaxFloat64},
		{"smallest positive", math.SmallestNonzeroFloat64},
		{"inf", math.Inf(1)},
		{"negative inf", math.Inf(-1)},
		{"NaN", math.NaN()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Marshal(tt.input)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			var result float64
			if err := Unmarshal(data, &result); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			// Special handling for NaN
			if math.IsNaN(tt.input) {
				if !math.IsNaN(result) {
					t.Errorf("Expected NaN, got %v", result)
				}
			} else if result != tt.input {
				t.Errorf("Roundtrip failed: got %v, want %v", result, tt.input)
			}
		})
	}
}

// TestMarshalBytes tests []byte marshaling (currently 0% coverage)
func TestMarshalBytes(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
	}{
		{"empty", []byte{}},
		{"small", []byte("hello")},
		{"with nulls", []byte{0x00, 0x01, 0x02, 0xFF}},
		{"large", make([]byte, 1024)},
		{"binary data", []byte{0xDE, 0xAD, 0xBE, 0xEF}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Marshal(tt.input)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			var result []byte
			if err := Unmarshal(data, &result); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("Roundtrip failed: got %v, want %v", result, tt.input)
			}
		})
	}

	// Test nil separately (special case)
	t.Run("nil", func(t *testing.T) {
		var input []byte
		data, err := Marshal(input)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var result []byte
		if err := Unmarshal(data, &result); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}

		// nil slice becomes empty slice after roundtrip
		if result == nil {
			result = []byte{}
		}
		if input == nil {
			input = []byte{}
		}
		if !reflect.DeepEqual(result, input) {
			t.Errorf("Roundtrip failed: got %v, want %v", result, input)
		}
	})
}

// TestEncodeErrors tests error paths in Encode function
func TestEncodeErrors(t *testing.T) {
	tests := []struct {
		name      string
		input     interface{}
		wantError bool
	}{
		{"channel", make(chan int), true},
		{"function", func() {}, true},
		{"complex64", complex(1, 2), true},
		{"complex128", complex128(1 + 2i), true},
		{"nil interface", nil, false}, // nil is valid
		{"pointer to channel", new(chan int), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Marshal(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("Marshal error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

// TestEncodeNilPointers tests nil pointer handling
func TestEncodeNilPointers(t *testing.T) {
	type TestStruct struct {
		Name  string `beve:"name"`
		Value *int   `beve:"value,omitempty"`
		Count int    `beve:"count"`
	}

	tests := []struct {
		name  string
		input interface{}
	}{
		{"nil pointer", (*int)(nil)},
		{"struct with nil fields", TestStruct{Name: "test", Value: nil, Count: 42}},
		{"nil slice pointer", (*[]string)(nil)},
		{"nil map pointer", (*map[string]int)(nil)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Marshal(tt.input)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			// Just verify marshaling succeeds
			if len(data) == 0 {
				t.Error("Expected non-empty encoding")
			}
		})
	}
}

// TestMarshalEmptyCollections tests empty slices, maps, arrays
func TestMarshalEmptyCollections(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
	}{
		{"empty slice", []int{}},
		{"empty map", map[string]int{}},
		{"empty array", [0]int{}},
		{"empty struct", struct{}{}},
		{"nil slice", []int(nil)},
		{"nil map", map[string]int(nil)},
		{"slice of empty structs", []struct{}{{}, {}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Marshal(tt.input)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			if len(data) == 0 {
				t.Errorf("Expected non-empty encoding, got empty")
			}
		})
	}
}

// TestMarshalMaxValues tests boundary values
func TestMarshalMaxValues(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
	}{
		{"max int8", int8(127)},
		{"min int8", int8(-128)},
		{"max int16", int16(32767)},
		{"min int16", int16(-32768)},
		{"max int32", int32(2147483647)},
		{"min int32", int32(-2147483648)},
		{"max int64", int64(9223372036854775807)},
		{"min int64", int64(-9223372036854775808)},
		{"max uint8", uint8(255)},
		{"max uint16", uint16(65535)},
		{"max uint32", uint32(4294967295)},
		{"max uint64", uint64(18446744073709551615)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Marshal(tt.input)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			// Unmarshal and verify
			result := reflect.New(reflect.TypeOf(tt.input))
			if err := Unmarshal(data, result.Interface()); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			if !reflect.DeepEqual(result.Elem().Interface(), tt.input) {
				t.Errorf("Roundtrip failed: got %v, want %v", result.Elem().Interface(), tt.input)
			}
		})
	}
}

// TestMarshalLargeNestedStruct tests deeply nested structures
func TestMarshalLargeNestedStruct(t *testing.T) {
	type Level3 struct {
		Value int
		Data  []string
	}
	type Level2 struct {
		Name  string
		Items []Level3
	}
	type Level1 struct {
		ID       int
		Children []Level2
		Metadata map[string]interface{}
	}

	input := Level1{
		ID: 1,
		Children: []Level2{
			{
				Name: "child1",
				Items: []Level3{
					{Value: 10, Data: []string{"a", "b"}},
					{Value: 20, Data: []string{"c", "d"}},
				},
			},
			{
				Name: "child2",
				Items: []Level3{
					{Value: 30, Data: []string{"e", "f"}},
				},
			},
		},
		Metadata: map[string]interface{}{
			"version": "1.0",
			"count":   42,
		},
	}

	data, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result Level1
	if err := Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Basic validation
	if result.ID != input.ID {
		t.Errorf("ID mismatch: got %v, want %v", result.ID, input.ID)
	}
	if len(result.Children) != len(input.Children) {
		t.Errorf("Children count mismatch: got %v, want %v", len(result.Children), len(input.Children))
	}
}

// TestMarshalTimeTime tests time.Time encoding (will use reflection for now)
func TestMarshalTimeTime(t *testing.T) {
	t.Skip("time.Time native support not yet implemented - covered by COMPETITIVE_ANALYSIS.md")

	now := time.Now()
	tests := []struct {
		name  string
		input time.Time
	}{
		{"now", now},
		{"zero", time.Time{}},
		{"unix epoch", time.Unix(0, 0)},
		{"far future", time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Marshal(tt.input)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			var result time.Time
			if err := Unmarshal(data, &result); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			// Compare Unix timestamps (nanosecond precision may vary)
			if !result.Equal(tt.input) {
				t.Errorf("Roundtrip failed: got %v, want %v", result, tt.input)
			}
		})
	}
}

// TestUnsupportedError tests UnsupportedError type
func TestUnsupportedError(t *testing.T) {
	err := &UnsupportedError{Msg: "test error"}
	expected := "beve: test error"
	if err.Error() != expected {
		t.Errorf("Error() = %v, want %v", err.Error(), expected)
	}
}

// TestIndent tests Indent function (no-op for binary format)
func TestIndent(t *testing.T) {
	enc := NewEncoder(nil)
	enc.Indent("", "  ") // Should not panic (it's a no-op for binary)
}

// TestSetEscapeHTML tests SetEscapeHTML (no-op for binary format)
func TestSetEscapeHTML(t *testing.T) {
	enc := NewEncoder(nil)
	enc.SetEscapeHTML(true)
	enc.SetEscapeHTML(false)
	// No panic means success (it's a no-op)
}

// TestNewEncoder tests NewEncoder creation
func TestNewEncoder(t *testing.T) {
	enc := NewEncoder(nil)
	if enc == nil {
		t.Fatal("NewEncoder returned nil")
	}

	// Test encoding with new encoder
	data := struct{ Name string }{Name: "test"}
	result, err := enc.Encode(data)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Expected non-empty encoding")
	}
}
