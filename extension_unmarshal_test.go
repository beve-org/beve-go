package beve

import (
	"reflect"
	"regexp"
	"testing"
	"time"
)

// ============================================================================
// Tests for unmarshalExtension (12% → 50% coverage target)
// ============================================================================

func TestUnmarshalExtension_TypedArray(t *testing.T) {
	users := []map[string]interface{}{
		{"name": "Alice", "age": float64(30)},
		{"name": "Bob", "age": float64(25)},
	}

	data, err := EncodeTypedArray(users)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	var result []map[string]interface{}
	err = UnmarshalAuto(data, &result)
	if err != nil {
		t.Fatalf("UnmarshalAuto failed: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 users, got %d", len(result))
	}
}

func TestUnmarshalExtension_FieldIndex(t *testing.T) {
	obj := map[string]interface{}{
		"id":   123,
		"name": "test",
	}

	data, err := EncodeIndexedObject(obj)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	var result map[string]interface{}
	err = UnmarshalAuto(data, &result)
	if err != nil {
		t.Fatalf("UnmarshalAuto failed: %v", err)
	}

	if result["name"] != "test" {
		t.Errorf("Expected name=test, got %v", result["name"])
	}
}

func TestUnmarshalExtension_Timestamp(t *testing.T) {
	now := time.Now().UTC()
	ts := Timestamp{
		Seconds:     now.Unix(),
		Nanoseconds: uint32(now.Nanosecond()),
	}

	data, err := EncodeTimestamp(ts)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	t.Run("to_time.Time", func(t *testing.T) {
		var result time.Time
		err = UnmarshalAuto(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalAuto failed: %v", err)
		}

		if result.Unix() != ts.Seconds {
			t.Errorf("Expected unix=%d, got %d", ts.Seconds, result.Unix())
		}
	})

	t.Run("to_Timestamp", func(t *testing.T) {
		var result Timestamp
		err = UnmarshalAuto(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalAuto failed: %v", err)
		}

		if result.Seconds != ts.Seconds {
			t.Errorf("Expected seconds=%d, got %d", ts.Seconds, result.Seconds)
		}
	})
}

func TestUnmarshalExtension_Duration(t *testing.T) {
	duration := 1 * time.Hour

	data, err := EncodeDuration(duration)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	var result time.Duration
	err = UnmarshalAuto(data, &result)
	if err != nil {
		t.Fatalf("UnmarshalAuto failed: %v", err)
	}

	if result != duration {
		t.Errorf("Expected %v, got %v", duration, result)
	}
}

func TestUnmarshalExtension_Interval(t *testing.T) {
	start := time.Now().UTC()
	end := start.Add(1 * time.Hour)

	data, err := EncodeInterval(start, end)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	var result [2]interface{}
	err = UnmarshalAuto(data, &result)
	if err != nil {
		t.Fatalf("UnmarshalAuto failed: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected array of 2, got %d", len(result))
	}
}

func TestUnmarshalExtension_UUID(t *testing.T) {
	uuid := [16]byte{0x55, 0x0e, 0x84, 0x00, 0xe2, 0x9b, 0x41, 0xd4, 0xa7, 0x16, 0x44, 0x66, 0x55, 0x44, 0x00, 0x00}

	data, err := MarshalUUID(uuid)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	t.Run("to_binary", func(t *testing.T) {
		var result [16]byte
		err = UnmarshalAuto(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalAuto failed: %v", err)
		}

		if result != uuid {
			t.Errorf("UUID mismatch")
		}
	})

	t.Run("to_string", func(t *testing.T) {
		var result string
		err = UnmarshalAuto(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalAuto failed: %v", err)
		}

		if len(result) == 0 {
			t.Error("Expected non-empty UUID string")
		}
	})
}

func TestUnmarshalExtension_RegExp(t *testing.T) {
	pattern := "^test$"

	data, err := EncodeRegExp(pattern, 0)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	t.Run("to_regexp.Regexp", func(t *testing.T) {
		var result *regexp.Regexp
		err = UnmarshalAuto(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalAuto failed: %v", err)
		}

		if result == nil {
			t.Error("Expected non-nil regexp")
		}
	})

	t.Run("to_RegExpData", func(t *testing.T) {
		var result RegExpData
		err = UnmarshalAuto(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalAuto failed: %v", err)
		}

		if result.Pattern != pattern {
			t.Errorf("Expected pattern=%s, got %s", pattern, result.Pattern)
		}
	})
}

func TestUnmarshalExtension_Errors(t *testing.T) {
	t.Run("nil_pointer", func(t *testing.T) {
		data := []byte{ExtTypedArray, 0x00}
		var result *map[string]interface{}
		err := UnmarshalAuto(data, result) // nil pointer
		if err == nil {
			t.Error("Expected error for nil pointer")
		}
	})

	t.Run("non_pointer", func(t *testing.T) {
		data := []byte{ExtTypedArray, 0x00}
		var result map[string]interface{}
		err := UnmarshalAuto(data, result) // not a pointer
		if err == nil {
			t.Error("Expected error for non-pointer")
		}
	})

	t.Run("unsupported_extension", func(t *testing.T) {
		data := []byte{0xFF, 0x00} // Invalid extension
		var result interface{}
		err := UnmarshalAuto(data, &result)
		if err == nil {
			t.Error("Expected error for unsupported extension")
		}
	})

	t.Run("invalid_typed_array_data", func(t *testing.T) {
		data := []byte{ExtTypedArray, 0xFF, 0xFF} // Corrupted data
		var result []map[string]interface{}
		err := UnmarshalAuto(data, &result)
		if err == nil {
			t.Error("Expected error for invalid data")
		}
	})

	t.Run("invalid_field_index_data", func(t *testing.T) {
		data := []byte{ExtFieldIndex, 0xFF, 0xFF} // Corrupted data
		var result map[string]interface{}
		err := UnmarshalAuto(data, &result)
		if err == nil {
			t.Error("Expected error for invalid data")
		}
	})

	t.Run("invalid_timestamp_data", func(t *testing.T) {
		data := []byte{ExtTimestamp, 0x00} // Too short
		var result time.Time
		err := UnmarshalAuto(data, &result)
		if err == nil {
			t.Error("Expected error for invalid timestamp")
		}
	})

	t.Run("invalid_duration_data", func(t *testing.T) {
		data := []byte{ExtDuration, 0x00} // Too short
		var result time.Duration
		err := UnmarshalAuto(data, &result)
		if err == nil {
			t.Error("Expected error for invalid duration")
		}
	})

	t.Run("invalid_uuid_data", func(t *testing.T) {
		data := []byte{ExtUUID, 0x00} // Too short
		var result [16]byte
		err := UnmarshalAuto(data, &result)
		if err == nil {
			t.Error("Expected error for invalid UUID")
		}
	})

	t.Run("invalid_regexp_data", func(t *testing.T) {
		data := []byte{ExtRegex, 0x00} // Too short
		var result RegExpData
		err := UnmarshalAuto(data, &result)
		if err == nil {
			t.Error("Expected error for invalid regexp")
		}
	})
}

// ============================================================================
// Tests for assignValue (44% → 70% coverage target)
// ============================================================================

func TestAssignValue_IntTypes(t *testing.T) {
	tests := []struct {
		name   string
		source interface{}
		target interface{}
	}{
		{"int_to_int", 42, new(int)},
		{"int_to_int8", 42, new(int8)},
		{"int_to_int16", 42, new(int16)},
		{"int_to_int32", 42, new(int32)},
		{"int_to_int64", 42, new(int64)},
		{"int_to_uint", 42, new(uint)},
		{"int_to_uint8", 42, new(uint8)},
		{"int_to_uint16", 42, new(uint16)},
		{"int_to_uint32", 42, new(uint32)},
		{"int_to_uint64", 42, new(uint64)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := reflect.ValueOf(tt.target).Elem()
			err := assignValue(rv, reflect.ValueOf(tt.source))
			if err != nil {
				t.Errorf("assignValue failed: %v", err)
			}
		})
	}
}

func TestAssignValue_FloatTypes(t *testing.T) {
	tests := []struct {
		name   string
		source interface{}
		target interface{}
	}{
		{"float_to_float32", 3.14, new(float32)},
		{"float_to_float64", 3.14, new(float64)},
		{"float_to_int", 3.14, new(int)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := reflect.ValueOf(tt.target).Elem()
			err := assignValue(rv, reflect.ValueOf(tt.source))
			if err != nil {
				t.Errorf("assignValue failed: %v", err)
			}
		})
	}
}

func TestAssignValue_StringAndBool(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		var target string
		err := assignValue(reflect.ValueOf(&target).Elem(), reflect.ValueOf("hello"))
		if err != nil {
			t.Errorf("assignValue failed: %v", err)
		}
		if target != "hello" {
			t.Errorf("Expected 'hello', got %s", target)
		}
	})

	t.Run("bool", func(t *testing.T) {
		var target bool
		err := assignValue(reflect.ValueOf(&target).Elem(), reflect.ValueOf(true))
		if err != nil {
			t.Errorf("assignValue failed: %v", err)
		}
		if !target {
			t.Error("Expected true")
		}
	})
}

func TestAssignValue_SliceAndMap(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		source := []int{1, 2, 3}
		var target []int
		err := assignValue(reflect.ValueOf(&target).Elem(), reflect.ValueOf(source))
		if err != nil {
			t.Errorf("assignValue failed: %v", err)
		}
		if len(target) != 3 {
			t.Errorf("Expected length 3, got %d", len(target))
		}
	})

	t.Run("map", func(t *testing.T) {
		source := map[string]int{"a": 1}
		var target map[string]int
		err := assignValue(reflect.ValueOf(&target).Elem(), reflect.ValueOf(source))
		if err != nil {
			t.Errorf("assignValue failed: %v", err)
		}
		if target["a"] != 1 {
			t.Error("Map assignment failed")
		}
	})
}

func TestAssignValue_Errors(t *testing.T) {
	t.Run("incompatible_types", func(t *testing.T) {
		var target int
		err := assignValue(reflect.ValueOf(&target).Elem(), reflect.ValueOf("string"))
		if err == nil {
			t.Error("Expected error for incompatible types")
		}
	})
}

// ============================================================================
// Tests for DetectEncoding (38% → 60% coverage target)
// ============================================================================

func TestDetectEncoding_AllExtensions(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected string
	}{
		{"typed_array", []byte{ExtTypedArray, 0x00}, "typed_array"},
		{"typed_nested", []byte{ExtTypedNestedArray, 0x00}, "typed_nested_array"},
		{"field_index", []byte{ExtFieldIndex, TypeObject, 0x00}, "field_index"},
		{"timestamp", []byte{ExtTimestamp, 0x00}, "timestamp"},
		{"duration", []byte{ExtDuration, 0x00}, "duration"},
		{"interval", []byte{ExtInterval, 0x00}, "interval"},
		{"uuid", []byte{ExtUUID, 0x00}, "uuid"},
		{"regex", []byte{ExtRegex, 0x00}, "regex"},
		{"standard_null", []byte{0x00}, "standard"},        // null
		{"standard_bool", []byte{0x08}, "standard"},        // boolean
		{"standard_number", []byte{0x01}, "standard"},      // number
		{"standard_string", []byte{0x02}, "standard"},      // string
		{"standard_object", []byte{0x03}, "standard"},      // object
		{"standard_array", []byte{0x05}, "standard"},       // array
		{"unknown", []byte{0xFF}, "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DetectEncoding(tt.data)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestDetectEncoding_EmptyData(t *testing.T) {
	result := DetectEncoding([]byte{})
	if result != "unknown" {
		t.Errorf("Expected 'unknown' for empty data, got %s", result)
	}
}
