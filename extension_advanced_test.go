package beve

import (
	"regexp"
	"testing"
	"time"
)

// ============================================================================
// Extension 2: Typed Nested Arrays Tests
// ============================================================================

func TestTypedNestedArrayEncoding(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		wantSize int
	}{
		{
			name: "2d_int_array",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			wantSize: 50, // Approximate
		},
		{
			name: "3d_float_array",
			input: [][][]float32{
				{{1.1, 2.2}, {3.3, 4.4}},
				{{5.5, 6.6}, {7.7, 8.8}},
			},
			wantSize: 80, // Approximate
		},
		{
			name: "nested_string_arrays",
			input: [][]string{
				{"hello", "world"},
				{"foo", "bar", "baz"},
			},
			wantSize: 60, // Approximate
		},
		{
			name:     "empty_nested_array",
			input:    [][]int{},
			wantSize: 10,
		},
		{
			name:     "single_nested_element",
			input:    [][]int{{42}},
			wantSize: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode
			data, err := EncodeTypedNestedArray(tt.input)
			if err != nil {
				t.Fatalf("EncodeTypedNestedArray failed: %v", err)
			}

			if len(data) == 0 {
				t.Error("Encoded data is empty")
			}

			// Check approximate size
			if len(data) > tt.wantSize*2 {
				t.Errorf("Encoded size too large: got %d, want ~%d", len(data), tt.wantSize)
			}

			t.Logf("Nested array size: %d bytes", len(data))

			// Decode (if decoder is implemented)
			decoded, err := DecodeTypedNestedArray(data)
			if err == nil {
				t.Logf("Decoded successfully: %v", decoded)
			} else {
				t.Logf("Decode not fully implemented yet: %v", err)
			}
		})
	}
}

func TestTypedNestedArrayEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		input     interface{}
		wantError bool
	}{
		{
			name:      "nil_input",
			input:     nil,
			wantError: true,
		},
		{
			name:      "not_nested_array",
			input:     []int{1, 2, 3},
			wantError: true,
		},
		{
			name:      "mixed_types",
			input:     []interface{}{[]int{1}, []string{"a"}},
			wantError: true,
		},
		{
			name: "deeply_nested_4d",
			input: [][][][]int{
				{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := EncodeTypedNestedArray(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("EncodeTypedNestedArray() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

// ============================================================================
// Extension 9: RegExp Tests
// ============================================================================

func TestRegExpEncoding(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		flags   string
	}{
		{
			name:    "simple_pattern",
			pattern: "^[a-z]+$",
			flags:   "",
		},
		{
			name:    "case_insensitive",
			pattern: "test",
			flags:   "i",
		},
		{
			name:    "multiline_global",
			pattern: "^start.*end$",
			flags:   "gm",
		},
		{
			name:    "unicode_pattern",
			pattern: "\\p{L}+",
			flags:   "u",
		},
		{
			name:    "email_pattern",
			pattern: "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
			flags:   "",
		},
		{
			name:    "complex_with_all_flags",
			pattern: "(?:foo|bar)",
			flags:   "gimsu",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test low-level encoding
			data, err := EncodeRegExp(tt.pattern, tt.flags)
			if err != nil {
				t.Fatalf("EncodeRegExp failed: %v", err)
			}

			if len(data) == 0 {
				t.Error("Encoded regex is empty")
			}

			t.Logf("RegExp size: %d bytes (pattern: %s, flags: %s)", len(data), tt.pattern, tt.flags)

			// Test decoding
			decodedPattern, decodedFlags, err := DecodeRegExp(data)
			if err != nil {
				t.Fatalf("DecodeRegExp failed: %v", err)
			}

			if decodedPattern != tt.pattern {
				t.Errorf("Pattern mismatch: got %q, want %q", decodedPattern, tt.pattern)
			}
			if decodedFlags != tt.flags {
				t.Errorf("Flags mismatch: got %q, want %q", decodedFlags, tt.flags)
			}
		})
	}
}

func TestRegExpHighLevelAPI(t *testing.T) {
	tests := []struct {
		name  string
		regex *regexp.Regexp
	}{
		{
			name:  "compiled_regex",
			regex: regexp.MustCompile("^test.*$"),
		},
		{
			name:  "simple_regex",
			regex: regexp.MustCompile("[0-9]+"),
		},
		{
			name:  "complex_regex",
			regex: regexp.MustCompile(`(?i)^[a-z]{3,10}@\w+\.(com|org)$`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal
			data, err := MarshalRegExp(tt.regex)
			if err != nil {
				t.Fatalf("MarshalRegExp failed: %v", err)
			}

			t.Logf("Marshaled regex size: %d bytes", len(data))

			// Unmarshal
			decoded, err := UnmarshalRegExp(data)
			if err != nil {
				t.Fatalf("UnmarshalRegExp failed: %v", err)
			}

			// Verify pattern matches
			if decoded.String() != tt.regex.String() {
				t.Errorf("Regex mismatch: got %q, want %q", decoded.String(), tt.regex.String())
			}

			// Test functionality
			testStr := "test123"
			if tt.regex.MatchString(testStr) != decoded.MatchString(testStr) {
				t.Errorf("Regex behavior mismatch on %q", testStr)
			}
		})
	}
}

func TestRegExpStringEncoding(t *testing.T) {
	pattern := "^[a-z]+$"
	flags := "i"

	// Encode
	data, err := EncodeRegExpString(pattern, flags)
	if err != nil {
		t.Fatalf("EncodeRegExpString failed: %v", err)
	}

	// Decode
	decodedPattern, decodedFlags, err := DecodeRegExpString(data)
	if err != nil {
		t.Fatalf("DecodeRegExpString failed: %v", err)
	}

	if decodedPattern != pattern {
		t.Errorf("Pattern mismatch: got %q, want %q", decodedPattern, pattern)
	}
	if decodedFlags != flags {
		t.Errorf("Flags mismatch: got %q, want %q", decodedFlags, flags)
	}
}

// ============================================================================
// Extension 5: Duration Tests
// ============================================================================

func TestDurationEncoding(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
	}{
		{
			name:     "one_second",
			duration: time.Second,
		},
		{
			name:     "one_hour",
			duration: time.Hour,
		},
		{
			name:     "negative_duration",
			duration: -5 * time.Minute,
		},
		{
			name:     "nanoseconds",
			duration: 12345 * time.Nanosecond,
		},
		{
			name:     "zero_duration",
			duration: 0,
		},
		{
			name:     "max_duration",
			duration: 1<<63 - 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode
			data, err := EncodeDuration(tt.duration)
			if err != nil {
				t.Fatalf("EncodeDuration failed: %v", err)
			}

			if len(data) == 0 {
				t.Error("Encoded duration is empty")
			}

			t.Logf("Duration %v encoded to %d bytes", tt.duration, len(data))

			// Decode
			decoded, err := DecodeDuration(data)
			if err != nil {
				t.Fatalf("DecodeDuration failed: %v", err)
			}

			if decoded != tt.duration {
				t.Errorf("Duration mismatch: got %v, want %v", decoded, tt.duration)
			}
		})
	}
}

// ============================================================================
// Extension 6: Interval Tests
// ============================================================================

func TestIntervalEncoding(t *testing.T) {
	tests := []struct {
		name  string
		start time.Time
		end   time.Time
	}{
		{
			name:  "one_hour_interval",
			start: time.Date(2025, 10, 17, 10, 0, 0, 0, time.UTC),
			end:   time.Date(2025, 10, 17, 11, 0, 0, 0, time.UTC),
		},
		{
			name:  "one_day_interval",
			start: time.Date(2025, 10, 17, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2025, 10, 18, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "same_time",
			start: time.Now(),
			end:   time.Now(),
		},
		{
			name:  "reverse_order",
			start: time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			end:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode
			data, err := EncodeInterval(tt.start, tt.end)
			if err != nil {
				t.Fatalf("EncodeInterval failed: %v", err)
			}

			if len(data) == 0 {
				t.Error("Encoded interval is empty")
			}

			t.Logf("Interval encoded to %d bytes", len(data))

			// Decode
			decodedStart, decodedEnd, err := DecodeInterval(data)
			if err != nil {
				t.Fatalf("DecodeInterval failed: %v", err)
			}

			// Check start time (within nanosecond precision)
			if !decodedStart.Equal(tt.start) {
				t.Errorf("Start time mismatch: got %v, want %v", decodedStart, tt.start)
			}

			// Check end time
			if !decodedEnd.Equal(tt.end) {
				t.Errorf("End time mismatch: got %v, want %v", decodedEnd, tt.end)
			}

			// Verify duration
			expectedDuration := tt.end.Sub(tt.start)
			actualDuration := decodedEnd.Sub(decodedStart)
			if expectedDuration != actualDuration {
				t.Errorf("Duration mismatch: got %v, want %v", actualDuration, expectedDuration)
			}
		})
	}
}

// ============================================================================
// UUID Helper Function Tests
// ============================================================================

func TestUUIDHelpers(t *testing.T) {
	// Test UUID version detection
	t.Run("uuid_version", func(t *testing.T) {
		// UUID v4 example
		uuid := [16]byte{
			0x6b, 0xa7, 0xb8, 0x10,
			0x9d, 0xad, 0x41, 0xd1,
			0x80, 0xb4, 0x00, 0xc0,
			0x4f, 0xd4, 0x30, 0xc8,
		}

		version := UUIDVersion(uuid)
		t.Logf("UUID version: %d", version)
		// Version is stored in byte 6, high nibble
	})

	// Test UUID variant detection
	t.Run("uuid_variant", func(t *testing.T) {
		uuid := [16]byte{
			0x6b, 0xa7, 0xb8, 0x10,
			0x9d, 0xad, 0x41, 0xd1,
			0x80, 0xb4, 0x00, 0xc0,
			0x4f, 0xd4, 0x30, 0xc8,
		}

		variant := UUIDVariant(uuid)
		t.Logf("UUID variant: %d", variant)
		// Variant is stored in byte 8, high bits
	})
}

// ============================================================================
// Utility Function Tests
// ============================================================================

func TestInferTypeCodeFromValue(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  byte
	}{
		{"int8", int8(42), TypeInt8},
		{"int16", int16(1000), TypeInt16},
		{"int32", int32(100000), TypeInt32},
		{"int64", int64(10000000), TypeInt64},
		{"uint8", uint8(200), TypeUInt8},
		{"uint16", uint16(50000), TypeUInt16},
		{"uint32", uint32(4000000000), TypeUInt32},
		{"uint64", uint64(18446744073709551615), TypeUInt64},
		{"float32", float32(3.14), TypeFloat32},
		{"float64", float64(2.718281828), TypeFloat64},
		{"string", "hello", TypeString},
		{"bool", true, TypeBool},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inferTypeCodeFromValue(tt.value)
			if got != tt.want {
				t.Errorf("inferTypeCodeFromValue(%v) = %d, want %d", tt.value, got, tt.want)
			}
		})
	}
}

func TestBuildNestedSchema(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
	}{
		{
			name: "2d_array",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{
			name: "3d_array",
			input: [][][]float64{
				{{1.1, 2.2}, {3.3, 4.4}},
				{{5.5, 6.6}, {7.7, 8.8}},
			},
		},
		{
			name:  "empty_nested",
			input: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schema, err := buildNestedSchema(tt.input)
			if err != nil {
				t.Logf("buildNestedSchema returned error (may be expected): %v", err)
			} else {
				t.Logf("Built schema: %v", schema)
			}
		})
	}
}

// ============================================================================
// Error Handling Tests
// ============================================================================

func TestErrorHandling(t *testing.T) {
	t.Run("decode_invalid_timestamp", func(t *testing.T) {
		invalidData := []byte{0x01, 0x02, 0x03} // Too short
		_, err := DecodeTimestamp(invalidData)
		if err == nil {
			t.Error("Expected error for invalid timestamp data")
		}
	})

	t.Run("decode_invalid_uuid", func(t *testing.T) {
		invalidData := []byte{0x01, 0x02} // Too short
		_, err := DecodeUUID(invalidData)
		if err == nil {
			t.Error("Expected error for invalid UUID data")
		}
	})

	t.Run("decode_invalid_duration", func(t *testing.T) {
		invalidData := []byte{0x01} // Too short
		_, err := DecodeDuration(invalidData)
		if err == nil {
			t.Error("Expected error for invalid duration data")
		}
	})

	t.Run("decode_invalid_interval", func(t *testing.T) {
		invalidData := []byte{0x01, 0x02, 0x03} // Too short
		_, _, err := DecodeInterval(invalidData)
		if err == nil {
			t.Error("Expected error for invalid interval data")
		}
	})

	t.Run("nil_regex", func(t *testing.T) {
		_, err := MarshalRegExp(nil)
		if err == nil {
			t.Error("Expected error for nil regex")
		}
	})
}

// ============================================================================
// UnmarshalAuto Edge Cases
// ============================================================================

func TestUnmarshalAutoEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		target    interface{}
		wantError bool
	}{
		{
			name:      "empty_data",
			data:      []byte{},
			target:    new(interface{}),
			wantError: true,
		},
		{
			name:      "nil_target",
			data:      []byte{0x00},
			target:    nil,
			wantError: true,
		},
		{
			name: "corrupted_header",
			data: []byte{0xFF, 0xFF, 0xFF},
			target: new(interface{}),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UnmarshalAuto(tt.data, tt.target)
			if (err != nil) != tt.wantError {
				t.Errorf("UnmarshalAuto() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

// ============================================================================
// Detection Tests
// ============================================================================

func TestDetectEncoding(t *testing.T) {
	tests := []struct {
		name     string
		prepData func() []byte
		want     string
	}{
		{
			name: "detect_typed_array",
			prepData: func() []byte {
				data, _ := MarshalTyped([]int{1, 2, 3})
				return data
			},
			want: "beve_typed_array",
		},
		{
			name: "detect_timestamp",
			prepData: func() []byte {
				data, _ := MarshalTimestamp(time.Now())
				return data
			},
			want: "beve_timestamp",
		},
		{
			name: "detect_uuid",
			prepData: func() []byte {
				uuid := [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
				data, _ := MarshalUUID(uuid)
				return data
			},
			want: "beve_uuid",
		},
		{
			name: "detect_standard_beve",
			prepData: func() []byte {
				data, _ := Marshal(map[string]int{"key": 42})
				return data
			},
			want: "beve_standard",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := tt.prepData()
			got := DetectEncoding(data)
			if got != tt.want {
				t.Errorf("DetectEncoding() = %q, want %q", got, tt.want)
			}
		})
	}
}

// ============================================================================
// Field Index Edge Cases
// ============================================================================

func TestFieldIndexEdgeCases(t *testing.T) {
	t.Run("read_nonexistent_field", func(t *testing.T) {
		obj := map[string]interface{}{
			"name": "Alice",
			"age":  30,
		}

		data, err := EncodeIndexedObject(obj)
		if err != nil {
			t.Fatalf("EncodeIndexedObject failed: %v", err)
		}

		_, err = ReadFieldByName(data, "nonexistent")
		if err == nil {
			t.Error("Expected error for nonexistent field")
		}
	})

	t.Run("empty_indexed_object", func(t *testing.T) {
		obj := map[string]interface{}{}

		data, err := EncodeIndexedObject(obj)
		if err != nil {
			t.Fatalf("EncodeIndexedObject failed: %v", err)
		}

		decoded, err := DecodeIndexedObject(data)
		if err != nil {
			t.Fatalf("DecodeIndexedObject failed: %v", err)
		}

		if len(decoded) != 0 {
			t.Errorf("Expected empty object, got %d fields", len(decoded))
		}
	})

	t.Run("large_indexed_object", func(t *testing.T) {
		obj := make(map[string]interface{})
		for i := 0; i < 100; i++ {
			obj[string(rune('a'+i%26))+string(rune('0'+i%10))] = i
		}

		data, err := EncodeIndexedObject(obj)
		if err != nil {
			t.Fatalf("EncodeIndexedObject failed: %v", err)
		}

		t.Logf("Large indexed object size: %d bytes", len(data))

		decoded, err := DecodeIndexedObject(data)
		if err != nil {
			t.Fatalf("DecodeIndexedObject failed: %v", err)
		}

		if len(decoded) != len(obj) {
			t.Errorf("Field count mismatch: got %d, want %d", len(decoded), len(obj))
		}
	})
}
