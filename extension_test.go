package beve

import (
	"testing"
	"time"
)

func TestTypedArrayEncoding(t *testing.T) {
	type User struct {
		Name  string `beve:"name"`
		Email string `beve:"email"`
		Age   int    `beve:"age"`
	}

	tests := []struct {
		name  string
		input []User
	}{
		{
			name: "small array",
			input: []User{
				{"Alice", "alice@example.com", 30},
				{"Bob", "bob@example.com", 25},
			},
		},
		{
			name:  "empty array",
			input: []User{},
		},
		{
			name: "single user",
			input: []User{
				{"Charlie", "charlie@example.com", 35},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode
			data, err := MarshalTyped(tt.input)
			if err != nil {
				t.Fatalf("MarshalTyped failed: %v", err)
			}

			// Decode
			var decoded []map[string]interface{}
			err = UnmarshalAuto(data, &decoded)
			if err != nil {
				t.Fatalf("UnmarshalAuto failed: %v", err)
			}

			// Verify length
			if len(decoded) != len(tt.input) {
				t.Errorf("Expected %d users, got %d", len(tt.input), len(decoded))
			}

			// Verify fields
			for i, user := range tt.input {
				if i >= len(decoded) {
					break
				}

				if decoded[i]["name"] != user.Name {
					t.Errorf("User %d: expected name %s, got %v", i, user.Name, decoded[i]["name"])
				}
				if decoded[i]["email"] != user.Email {
					t.Errorf("User %d: expected email %s, got %v", i, user.Email, decoded[i]["email"])
				}
				if int(decoded[i]["age"].(int64)) != user.Age {
					t.Errorf("User %d: expected age %d, got %v", i, user.Age, decoded[i]["age"])
				}
			}
		})
	}
}

func TestTimestampEncoding(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
	}{
		{
			name: "current time",
			time: time.Now(),
		},
		{
			name: "unix epoch",
			time: time.Unix(0, 0).UTC(),
		},
		{
			name: "with nanoseconds",
			time: time.Unix(1234567890, 123456789).UTC(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode
			data, err := MarshalTimestamp(tt.time)
			if err != nil {
				t.Fatalf("MarshalTimestamp failed: %v", err)
			}

			// Decode
			decoded, err := UnmarshalTimestamp(data)
			if err != nil {
				t.Fatalf("UnmarshalTimestamp failed: %v", err)
			}

			// Verify (nanosecond precision)
			if !tt.time.Equal(decoded) {
				t.Errorf("Expected %v, got %v", tt.time, decoded)
			}
		})
	}
}

func TestUUIDEncoding(t *testing.T) {
	tests := []struct {
		name string
		uuid [16]byte
		str  string
	}{
		{
			name: "uuid v4",
			uuid: [16]byte{0x55, 0x0e, 0x84, 0x00, 0xe2, 0x9b, 0x41, 0xd4,
				0xa7, 0x16, 0x44, 0x66, 0x55, 0x44, 0x00, 0x00},
			str: "550e8400-e29b-41d4-a716-446655440000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode binary
			data, err := MarshalUUID(tt.uuid)
			if err != nil {
				t.Fatalf("MarshalUUID failed: %v", err)
			}

			// Should be 18 bytes (header + version + 16 bytes)
			if len(data) != 18 {
				t.Errorf("Expected 18 bytes, got %d", len(data))
			}

			// Decode
			decoded, err := UnmarshalUUID(data)
			if err != nil {
				t.Fatalf("UnmarshalUUID failed: %v", err)
			}

			// Verify
			if decoded != tt.uuid {
				t.Errorf("Expected %v, got %v", tt.uuid, decoded)
			}

			// Test string encoding
			strData, err := MarshalUUIDString(tt.str)
			if err != nil {
				t.Fatalf("MarshalUUIDString failed: %v", err)
			}

			decodedStr, err := UnmarshalUUIDString(strData)
			if err != nil {
				t.Fatalf("UnmarshalUUIDString failed: %v", err)
			}

			if decodedStr != tt.str {
				t.Errorf("Expected %s, got %s", tt.str, decodedStr)
			}
		})
	}
}

func TestFieldIndexEncoding(t *testing.T) {
	tests := []struct {
		name string
		obj  map[string]interface{}
	}{
		{
			name: "simple object",
			obj: map[string]interface{}{
				"name":  "Alice",
				"age":   30,
				"email": "alice@example.com",
			},
		},
		{
			name: "empty object",
			obj:  map[string]interface{}{},
		},
		{
			name: "nested object",
			obj: map[string]interface{}{
				"id":   123,
				"user": "Alice",
				"meta": map[string]interface{}{
					"created": "2025-10-17",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode
			data, err := EncodeIndexedObject(tt.obj)
			if err != nil {
				t.Fatalf("EncodeIndexedObject failed: %v", err)
			}

			// Decode
			decoded, err := DecodeIndexedObject(data)
			if err != nil {
				t.Fatalf("DecodeIndexedObject failed: %v", err)
			}

			// Verify length
			if len(decoded) != len(tt.obj) {
				t.Errorf("Expected %d fields, got %d", len(tt.obj), len(decoded))
			}

			// Test field access
			for key, expectedValue := range tt.obj {
				value, err := ReadFieldByName(data, key)
				if err != nil {
					t.Errorf("ReadFieldByName(%s) failed: %v", key, err)
					continue
				}

				// Basic type check (not deep comparison)
				if value == nil && expectedValue != nil {
					t.Errorf("Expected non-nil value for %s", key)
				}
			}
		})
	}
}

func TestAutoDetection(t *testing.T) {
	type User struct {
		Name string `beve:"name"`
		Age  int    `beve:"age"`
	}

	users := []User{
		{"Alice", 30},
		{"Bob", 25},
	}

	// Test MarshalAuto
	data, err := MarshalAuto(users)
	if err != nil {
		t.Fatalf("MarshalAuto failed: %v", err)
	}

	// Detect encoding
	encoding := DetectEncoding(data)
	t.Logf("Detected encoding: %s", encoding)

	// Test global Unmarshal (should auto-detect)
	var decoded []map[string]interface{}
	err = Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(decoded) != len(users) {
		t.Errorf("Expected %d users, got %d", len(users), len(decoded))
	}
}

func TestExtensionDetection(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		isExt     bool
		extID     int
		expectErr bool
	}{
		{
			name:  "typed array",
			data:  []byte{ExtTypedArray, 0x00},
			isExt: true,
			extID: 1,
		},
		{
			name:  "timestamp",
			data:  []byte{ExtTimestamp, 0x00},
			isExt: true,
			extID: 4,
		},
		{
			name:  "uuid",
			data:  []byte{ExtUUID, 0x00},
			isExt: true,
			extID: 8,
		},
		{
			name:      "standard beve",
			data:      []byte{0x00},
			isExt:     false,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isExt := IsExtension(tt.data)
			if isExt != tt.isExt {
				t.Errorf("IsExtension: expected %v, got %v", tt.isExt, isExt)
			}

			if tt.isExt {
				extID, err := GetExtensionID(tt.data)
				if err != nil {
					t.Fatalf("GetExtensionID failed: %v", err)
				}
				if extID != tt.extID {
					t.Errorf("Expected extension ID %d, got %d", tt.extID, extID)
				}
			} else if !tt.expectErr {
				_, err := GetExtensionID(tt.data)
				if err == nil {
					t.Error("Expected error for non-extension data")
				}
			}
		})
	}
}

// Benchmark tests
func BenchmarkTypedArrayMarshal(b *testing.B) {
	type User struct {
		Name  string `beve:"name"`
		Email string `beve:"email"`
		Age   int    `beve:"age"`
	}

	users := make([]User, 100)
	for i := 0; i < 100; i++ {
		users[i] = User{
			Name:  "User" + string(rune(i)),
			Email: "user" + string(rune(i)) + "@example.com",
			Age:   20 + i,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MarshalTyped(users)
	}
}

func BenchmarkStandardMarshal(b *testing.B) {
	type User struct {
		Name  string `beve:"name"`
		Email string `beve:"email"`
		Age   int    `beve:"age"`
	}

	users := make([]User, 100)
	for i := 0; i < 100; i++ {
		users[i] = User{
			Name:  "User" + string(rune(i)),
			Email: "user" + string(rune(i)) + "@example.com",
			Age:   20 + i,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(users)
	}
}

func BenchmarkFieldIndexRead(b *testing.B) {
	obj := map[string]interface{}{
		"field1":  "value1",
		"field2":  "value2",
		"field3":  "value3",
		"field4":  "value4",
		"field5":  "value5",
		"field6":  "value6",
		"field7":  "value7",
		"field8":  "value8",
		"field9":  "value9",
		"field10": "value10",
	}

	data, _ := EncodeIndexedObject(obj)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ReadFieldByName(data, "field5")
	}
}
