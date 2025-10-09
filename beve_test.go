package beve

import (
	"reflect"
	"testing"
)

func TestBasicTypes(t *testing.T) {
	// Test null
	var n interface{}
	data, err := Marshal(nil)
	if err != nil {
		t.Errorf("Marshal nil failed: %v", err)
	}
	err = Unmarshal(data, &n)
	if err != nil {
		t.Errorf("Unmarshal nil failed: %v", err)
	}
	if n != nil {
		t.Errorf("Expected nil, got %v", n)
	}

	// Test bool
	b := true
	data, err = Marshal(b)
	if err != nil {
		t.Errorf("Marshal bool failed: %v", err)
	}
	var decodedBool bool
	err = Unmarshal(data, &decodedBool)
	if err != nil {
		t.Errorf("Unmarshal bool failed: %v", err)
	}
	if decodedBool != true {
		t.Errorf("Expected true, got %v", decodedBool)
	}

	// Test int
	i := 42
	data, err = Marshal(i)
	if err != nil {
		t.Errorf("Marshal int failed: %v", err)
	}
	var decodedInt int
	err = Unmarshal(data, &decodedInt)
	if err != nil {
		t.Errorf("Unmarshal int failed: %v", err)
	}
	if decodedInt != 42 {
		t.Errorf("Expected 42, got %v", decodedInt)
	}

	// Test string
	s := "hello"
	data, err = Marshal(s)
	if err != nil {
		t.Errorf("Marshal string failed: %v", err)
	}
	var decodedString string
	err = Unmarshal(data, &decodedString)
	if err != nil {
		t.Errorf("Unmarshal string failed: %v", err)
	}
	if decodedString != "hello" {
		t.Errorf("Expected 'hello', got %v", decodedString)
	}
}

func TestStruct(t *testing.T) {
	type Person struct {
		Name string `beve:"name"`
		Age  int    `beve:"age"`
	}

	p := Person{Name: "Alice", Age: 30}
	data, err := Marshal(p)
	if err != nil {
		t.Errorf("Marshal struct failed: %v", err)
	}

	var decoded Person
	err = Unmarshal(data, &decoded)
	if err != nil {
		t.Errorf("Unmarshal struct failed: %v", err)
	}

	if decoded.Name != "Alice" || decoded.Age != 30 {
		t.Errorf("Expected {Alice 30}, got %+v", decoded)
	}
}

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	data, err := Marshal(slice)
	if err != nil {
		t.Errorf("Marshal slice failed: %v", err)
	}

	var decoded []int
	err = Unmarshal(data, &decoded)
	if err != nil {
		t.Errorf("Unmarshal slice failed: %v", err)
	}

	if len(decoded) != 5 || decoded[0] != 1 || decoded[4] != 5 {
		t.Errorf("Expected [1 2 3 4 5], got %v", decoded)
	}
}

func TestTypedArrays(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		input := []int32{1, -2, 3, -4, 5}
		data, err := Marshal(input)
		if err != nil {
			t.Fatalf("Marshal int32 slice failed: %v", err)
		}

		if len(data) < 2 {
			t.Fatalf("encoded data too short: %x", data)
		}
		// Typed array header for signed ints (group 1, byte count indicator for 4 bytes)
		expectedHeader := byte(0x04 | (1 << 3) | (2 << 5))
		if data[0] != expectedHeader {
			t.Fatalf("unexpected header: got 0x%02x want 0x%02x", data[0], expectedHeader)
		}

		var decoded []int32
		if err := Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal int32 slice failed: %v", err)
		}
		if !reflect.DeepEqual(decoded, input) {
			t.Fatalf("decoded mismatch: got %v want %v", decoded, input)
		}

		var any interface{}
		if err := Unmarshal(data, &any); err != nil {
			t.Fatalf("Unmarshal interface failed: %v", err)
		}
		values, ok := any.([]int32)
		if !ok {
			t.Fatalf("expected []int32, got %T", any)
		}
		if !reflect.DeepEqual(values, input) {
			t.Fatalf("interface decoded mismatch: got %v want %v", values, input)
		}
	})

	t.Run("uint16", func(t *testing.T) {
		input := []uint16{0, 1, 255, 1024, 65535}
		data, err := Marshal(input)
		if err != nil {
			t.Fatalf("Marshal uint16 slice failed: %v", err)
		}

		var decoded []uint16
		if err := Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal uint16 slice failed: %v", err)
		}
		if !reflect.DeepEqual(decoded, input) {
			t.Fatalf("decoded mismatch: got %v want %v", decoded, input)
		}
	})

	t.Run("float32", func(t *testing.T) {
		input := []float32{0, 1.5, -2.75, 3.125}
		data, err := Marshal(input)
		if err != nil {
			t.Fatalf("Marshal float32 slice failed: %v", err)
		}

		var decoded []float32
		if err := Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal float32 slice failed: %v", err)
		}
		if len(decoded) != len(input) {
			t.Fatalf("decoded length mismatch: got %d want %d", len(decoded), len(input))
		}
		for i := range input {
			if decoded[i] != input[i] {
				t.Fatalf("decoded[%d]=%v want %v", i, decoded[i], input[i])
			}
		}
	})

	t.Run("bool", func(t *testing.T) {
		input := []bool{true, false, true, true, false, false, true, false, true}
		data, err := Marshal(input)
		if err != nil {
			t.Fatalf("Marshal bool slice failed: %v", err)
		}

		if len(data) != 1+1+2 { // header + size + payload (ceil(9/8)=2)
			t.Fatalf("unexpected encoded length: got %d", len(data))
		}
		if data[0] != 0x1C {
			t.Fatalf("unexpected bool header: got 0x%02x", data[0])
		}
		if data[1] != byte(len(input)<<2) {
			t.Fatalf("unexpected size encoding: got 0x%02x want 0x%02x", data[1], byte(len(input)<<2))
		}

		var decoded []bool
		if err := Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal bool slice failed: %v", err)
		}
		if !reflect.DeepEqual(decoded, input) {
			t.Fatalf("decoded mismatch: got %v want %v", decoded, input)
		}
	})
}
