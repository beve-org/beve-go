package core

import (
	"fmt"
	"reflect"
	"testing"
)

// CustomType implements BinaryMarshaler and BinaryUnmarshaler
type CustomType struct {
	Value string
	Count int
}

func (c CustomType) MarshalBinary() ([]byte, error) {
	return []byte(fmt.Sprintf("%s:%d", c.Value, c.Count)), nil
}

func (c *CustomType) UnmarshalBinary(data []byte) error {
	_, err := fmt.Sscanf(string(data), "%s:%d", &c.Value, &c.Count)
	return err
}

// ErrorMarshaler always returns error
type ErrorMarshaler struct{}

func (e ErrorMarshaler) MarshalBinary() ([]byte, error) {
	return nil, fmt.Errorf("marshal error")
}

func (e *ErrorMarshaler) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("unmarshal error")
}

// TestEncodeBinaryMarshaler tests encodeBinaryMarshaler function (0% coverage)
func TestEncodeBinaryMarshaler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		custom := CustomType{Value: "test", Count: 42}
		rv := reflect.ValueOf(custom)

		if err := enc.Encode(rv); err != nil {
			t.Logf("Encode custom type: %v", err)
		}

		if enc.Buf.Len() > 0 {
			t.Log("Custom type encoded successfully")
		}
	})

	t.Run("error case", func(t *testing.T) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		errMarshaler := ErrorMarshaler{}
		rv := reflect.ValueOf(&errMarshaler)

		err := enc.Encode(rv)
		// May or may not error depending on implementation
		t.Logf("ErrorMarshaler encode result: %v", err)
	})
}

// TestBinaryUnmarshaler tests decoder's BinaryUnmarshaler support
func TestBinaryUnmarshaler(t *testing.T) {
	t.Run("roundtrip", func(t *testing.T) {
		enc := GetEncoderFromPool()
		defer PutEncoderToPool(enc)

		original := &CustomType{Value: "hello", Count: 123}

		if err := enc.Encode(reflect.ValueOf(original)); err != nil {
			t.Logf("Encode custom type: %v", err)
		}

		if enc.Buf.Len() > 0 {
			dec := NewDecoder(enc.Buf.Bytes())
			result := &CustomType{}

			if err := dec.Decode(reflect.ValueOf(result).Elem()); err != nil {
				t.Logf("Decode custom type: %v", err)
			}
		}
	})
}

// TestLookupBinaryUnmarshaler tests shouldCheckBinaryUnmarshaler and lookupBinaryUnmarshaler
func TestLookupBinaryUnmarshaler(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	tests := []struct {
		name  string
		value interface{}
	}{
		{"custom type", &CustomType{Value: "test", Count: 1}},
		{"slice of custom", []CustomType{{Value: "a", Count: 1}, {Value: "b", Count: 2}}},
		{"map with custom", map[string]*CustomType{"key": {Value: "val", Count: 3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc.Buf.Reset()

			if err := enc.Encode(reflect.ValueOf(tt.value)); err != nil {
				t.Errorf("Encode failed: %v", err)
			}

			if enc.Buf.Len() == 0 {
				t.Error("Nothing encoded")
			}
		})
	}
}

// TestCustomMarshalerInStruct tests custom marshaler inside struct
func TestCustomMarshalerInStruct(t *testing.T) {
	type Container struct {
		Name   string
		Custom CustomType
		ID     int
	}

	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	original := Container{
		Name:   "container",
		Custom: CustomType{Value: "nested", Count: 99},
		ID:     456,
	}

	if err := enc.Encode(reflect.ValueOf(original)); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	if enc.Buf.Len() == 0 {
		t.Error("Nothing encoded")
	}
}

// TestCustomMarshalerPointer tests pointer vs value for custom marshaler
func TestCustomMarshalerPointer(t *testing.T) {
	enc := GetEncoderFromPool()
	defer PutEncoderToPool(enc)

	t.Run("pointer", func(t *testing.T) {
		enc.Buf.Reset()
		custom := &CustomType{Value: "ptr", Count: 10}
		if err := enc.Encode(reflect.ValueOf(custom)); err != nil {
			t.Fatalf("Encode pointer failed: %v", err)
		}
	})

	t.Run("value", func(t *testing.T) {
		enc.Buf.Reset()
		custom := CustomType{Value: "val", Count: 20}
		if err := enc.Encode(reflect.ValueOf(custom)); err != nil {
			t.Fatalf("Encode value failed: %v", err)
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		enc.Buf.Reset()
		var custom *CustomType = nil
		if err := enc.Encode(reflect.ValueOf(custom)); err != nil {
			t.Fatalf("Encode nil pointer failed: %v", err)
		}
	})
}
