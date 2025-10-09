package beve

import (
	"bytes"
	"errors"
	"io"
	"math"
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

func TestFloat32RoundTrip(t *testing.T) {
	var input float32 = 3.5
	data, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal float32 failed: %v", err)
	}

	var decoded float32
	if err := Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal float32 failed: %v", err)
	}
	if decoded != input {
		t.Fatalf("float32 round trip mismatch: got %v want %v", decoded, input)
	}

	var any interface{}
	if err := Unmarshal(data, &any); err != nil {
		t.Fatalf("Unmarshal float32 into interface failed: %v", err)
	}
	f64, ok := any.(float64)
	if !ok {
		t.Fatalf("expected interface to hold float64, got %T", any)
	}
	if math.Abs(float64(float32(f64)-input)) > 1e-6 {
		t.Fatalf("interface float mismatch: got %v want %v", f64, input)
	}
}

func TestMapStringInt(t *testing.T) {
	input := map[string]int{"alpha": 1, "beta": 2, "gamma": -3}
	data, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal map failed: %v", err)
	}

	var decoded map[string]int
	if err := Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal map failed: %v", err)
	}

	if !reflect.DeepEqual(decoded, input) {
		t.Fatalf("map round trip mismatch: got %v want %v", decoded, input)
	}
}

func TestTypedStringArray(t *testing.T) {
	input := []string{"alpha", "", "gamma"}
	data, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal string slice failed: %v", err)
	}

	var decoded []string
	if err := Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal string slice failed: %v", err)
	}
	if !reflect.DeepEqual(decoded, input) {
		t.Fatalf("decoded mismatch: got %v want %v", decoded, input)
	}

	var any interface{}
	if err := Unmarshal(data, &any); err != nil {
		t.Fatalf("Unmarshal string slice into interface failed: %v", err)
	}
	values, ok := any.([]string)
	if !ok {
		t.Fatalf("expected []string, got %T", any)
	}
	if !reflect.DeepEqual(values, input) {
		t.Fatalf("interface decoded mismatch: got %v want %v", values, input)
	}
}

func TestMapIntKeys(t *testing.T) {
	input := map[int]string{-2: "neg", 0: "zero", 42: "answer"}
	data, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal map[int]string failed: %v", err)
	}

	var decoded map[int]string
	if err := Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal map[int]string failed: %v", err)
	}

	if !reflect.DeepEqual(decoded, input) {
		t.Fatalf("map round trip mismatch: got %v want %v", decoded, input)
	}
}

func TestStructOmitEmpty(t *testing.T) {
	type sample struct {
		Name string `beve:"name,omitempty"`
		Age  int    `beve:"age"`
		Note string `beve:"-"`
	}

	value := sample{Age: 27, Note: "secret"}
	data, err := Marshal(value)
	if err != nil {
		t.Fatalf("Marshal struct with omitempty failed: %v", err)
	}

	var decoded map[string]interface{}
	if err := Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal into map failed: %v", err)
	}

	if _, ok := decoded["name"]; ok {
		t.Fatalf("expected omitted field 'name' to be absent")
	}
	age, ok := decoded["age"].(int64)
	if !ok || age != 27 {
		t.Fatalf("expected age 27, got %v (%T)", decoded["age"], decoded["age"])
	}
}

func TestStreamingEncodeDecode(t *testing.T) {
	input := map[string]int{"one": 1, "two": 2}
	var buf bytes.Buffer
	enc := NewEncoder(&buf)
	if _, err := enc.Encode(input); err != nil {
		t.Fatalf("streaming encode failed: %v", err)
	}

	dec := NewDecoder(bytes.NewReader(buf.Bytes()))
	var output map[string]int
	if err := dec.Decode(&output); err != nil {
		t.Fatalf("streaming decode failed: %v", err)
	}

	if !reflect.DeepEqual(output, input) {
		t.Fatalf("streaming round trip mismatch: got %v want %v", output, input)
	}
}

func TestUnsupportedMapKeyType(t *testing.T) {
	data := map[float64]string{3.14: "pi"}
	_, err := Marshal(data)
	if err == nil {
		t.Fatalf("expected error when marshaling map with float keys")
	}
	var unsupported *UnsupportedError
	if !errors.As(err, &unsupported) {
		t.Fatalf("expected UnsupportedError, got %T", err)
	}
}

type badReader struct{}

func (badReader) Read(p []byte) (int, error) { return 0, io.EOF }

func TestDecoderUnsupportedReader(t *testing.T) {
	dec := NewDecoder(struct{ badReader }{})
	var out interface{}
	err := dec.Decode(&out)
	if err == nil {
		t.Fatalf("expected error for unsupported reader type")
	}
	var unsupported *UnsupportedError
	if !errors.As(err, &unsupported) {
		t.Fatalf("expected UnsupportedError, got %T", err)
	}
}

type testMarshaler struct{}

func (testMarshaler) MarshalBEVE() ([]byte, error) {
	return []byte{0x18}, nil // true
}

type testUnmarshaler struct {
	raw []byte
}

func (t *testUnmarshaler) UnmarshalBEVE(data []byte) error {
	t.raw = append(t.raw[:0], data...)
	return nil
}

func TestRawMessageAndBinaryMarshaler(t *testing.T) {
	data := []byte{0x08} // false

	var raw RawMessage
	if err := Unmarshal(data, &raw); err != nil {
		t.Fatalf("Unmarshal RawMessage failed: %v", err)
	}
	if !reflect.DeepEqual([]byte(raw), data) {
		t.Fatalf("RawMessage mismatch: got %x want %x", []byte(raw), data)
	}

	encoded, err := Marshal(raw)
	if err != nil {
		t.Fatalf("Marshal RawMessage failed: %v", err)
	}
	if !reflect.DeepEqual(encoded, data) {
		t.Fatalf("Marshal RawMessage mismatch: got %x want %x", encoded, data)
	}

	var rawPtr *RawMessage
	if err := Unmarshal([]byte{0x18}, &rawPtr); err != nil {
		t.Fatalf("Unmarshal *RawMessage failed: %v", err)
	}
	if rawPtr == nil || !reflect.DeepEqual([]byte(*rawPtr), []byte{0x18}) {
		t.Fatalf("pointer RawMessage mismatch")
	}

	marshaled, err := Marshal(testMarshaler{})
	if err != nil {
		t.Fatalf("Marshal BinaryMarshaler failed: %v", err)
	}
	if !reflect.DeepEqual(marshaled, []byte{0x18}) {
		t.Fatalf("Marshal BinaryMarshaler mismatch: got %x", marshaled)
	}

	var u testUnmarshaler
	if err := Unmarshal([]byte{0x0c}, &u); err != nil {
		t.Fatalf("Unmarshal BinaryUnmarshaler failed: %v", err)
	}
	if !reflect.DeepEqual(u.raw, []byte{0x0c}) {
		t.Fatalf("BinaryUnmarshaler raw mismatch: got %x", u.raw)
	}
}
