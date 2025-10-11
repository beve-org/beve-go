// Package beve implements BEVE (Binary Efficient Versatile Encoding) for Go,
// providing high-performance binary serialization with zero-allocation goals
// and full compatibility with Go's encoding/json interfaces.
package beve

import (
	"bytes"
	"io"
	"reflect"
	"sync"

	"github.com/beve-org/beve-go/core"
)

// Version represents the BEVE specification version
const Version = "1.2.0"

// ZeroCopyBytes represents a leased encoder buffer. Call Release to recycle it.
type ZeroCopyBytes = core.BufferLease

// Marshal encodes v into BEVE binary format.
// It follows the same interface as encoding/json.Marshal.
//
// Fast paths for common types avoid reflection overhead:
//   - Primitives: int, string, bool, float64
//   - Slices: []byte directly encoded
//   - Everything else: Uses reflection
func Marshal(v interface{}) ([]byte, error) {
	// Fast path: avoid reflect.ValueOf for common types
	switch val := v.(type) {
	case int:
		return marshalInt(val)
	case string:
		return marshalString(val)
	case bool:
		return marshalBool(val)
	case float64:
		return marshalFloat64(val)
	case []byte:
		return marshalBytes(val)
	}

	// Slow path: use reflection
	lease, err := MarshalZeroCopy(v)
	if err != nil {
		return nil, err
	}

	data := lease.Bytes()
	result := make([]byte, len(data))
	copy(result, data)
	lease.Release()
	return result, nil
}

// MarshalZeroCopy encodes v into BEVE binary format without copying the result.
//
// The returned lease shares the encoder's internal buffer. Callers must invoke
// Release when done so the buffer can be recycled. The byte slice remains valid
// until Release is called.
func MarshalZeroCopy(v interface{}) (ZeroCopyBytes, error) {
	enc := getEncoderFromPool()
	if enc.Buf != nil {
		enc.Buf.Reset()
	}

	// Encode the value
	rv := reflect.ValueOf(v)
	if err := enc.Encode(rv); err != nil {
		putEncoderToPool(enc)
		return ZeroCopyBytes{}, err
	}

	lease := enc.DetachBytes()
	putEncoderToPool(enc)

	return lease, nil
}

// Unmarshal decodes BEVE binary data into v.
// It follows the same interface as encoding/json.Unmarshal.
func Unmarshal(data []byte, v interface{}) error {
	d := NewDecoder(data)
	return d.Decode(v)
}

// Encoder provides streaming BEVE encoding.
type Encoder struct {
	w io.Writer
}

// NewEncoder creates a new BEVE encoder.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode encodes v to BEVE format and writes to the underlying writer.
// If no writer is set, returns the encoded bytes.
func (e *Encoder) Encode(v interface{}) ([]byte, error) {
	if e.w == nil {
		buf := getBuffer()
		defer putBuffer(buf)
		enc := newEncoder(buf)
		if err := enc.Encode(reflect.ValueOf(v)); err != nil {
			return nil, err
		}
		result := make([]byte, buf.Len())
		copy(result, buf.Bytes())
		return result, nil
	}
	// Streaming encoding
	enc := newEncoder(e.w)
	return nil, enc.Encode(reflect.ValueOf(v))
}

// Decoder provides streaming BEVE decoding.
type Decoder struct {
	r interface{} // can be []byte or io.Reader
}

// NewDecoder creates a new BEVE decoder.
// Accepts either []byte or io.Reader.
func NewDecoder(r interface{}) *Decoder {
	return &Decoder{r: r}
}

// Decode decodes BEVE data into v.
// If r is a []byte, decodes from the slice.
func (d *Decoder) Decode(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{Type: reflect.TypeOf(v)}
	}

	if data, ok := d.r.([]byte); ok {
		dec := newDecoder(data)
		return dec.decode(rv.Elem())
	}

	if reader, ok := d.r.(io.Reader); ok {
		buf := getBuffer()
		defer putBuffer(buf)
		if _, err := buf.ReadFrom(reader); err != nil {
			return err
		}
		data := make([]byte, buf.Len())
		copy(data, buf.Bytes())
		dec := newDecoder(data)
		return dec.decode(rv.Elem())
	}

	return &UnsupportedError{Msg: "unsupported reader type"}
}

// BinaryMarshaler is the interface implemented by an object that can
// marshal itself into a binary form.
type BinaryMarshaler interface {
	MarshalBEVE() ([]byte, error)
}

// BinaryUnmarshaler is the interface implemented by an object that can
// unmarshal a binary representation of itself.
type BinaryUnmarshaler interface {
	UnmarshalBEVE(data []byte) error
}

// Fast path helpers - pre-allocate and avoid full reflection overhead

func marshalInt(v int) ([]byte, error) {
	enc := getEncoderFromPool()
	defer putEncoderToPool(enc)
	if enc.Buf != nil {
		enc.Buf.Reset()
		enc.Buf.Grow(16) // int needs max 10 bytes
	}
	rv := reflect.ValueOf(v)
	if err := enc.Encode(rv); err != nil {
		return nil, err
	}
	data := enc.Buf.Bytes()
	result := make([]byte, len(data))
	copy(result, data)
	return result, nil
}

func marshalString(v string) ([]byte, error) {
	enc := getEncoderFromPool()
	defer putEncoderToPool(enc)
	if enc.Buf != nil {
		enc.Buf.Reset()
		enc.Buf.Grow(len(v) + 8) // string + length encoding
	}
	if err := enc.EncodeString(v); err != nil {
		return nil, err
	}
	data := enc.Buf.Bytes()
	result := make([]byte, len(data))
	copy(result, data)
	return result, nil
}

func marshalBool(v bool) ([]byte, error) {
	enc := getEncoderFromPool()
	defer putEncoderToPool(enc)
	if enc.Buf != nil {
		enc.Buf.Reset()
		enc.Buf.Grow(4) // bool needs 1 byte
	}
	rv := reflect.ValueOf(v)
	if err := enc.Encode(rv); err != nil {
		return nil, err
	}
	data := enc.Buf.Bytes()
	result := make([]byte, len(data))
	copy(result, data)
	return result, nil
}

func marshalFloat64(v float64) ([]byte, error) {
	enc := getEncoderFromPool()
	defer putEncoderToPool(enc)
	if enc.Buf != nil {
		enc.Buf.Reset()
		enc.Buf.Grow(16) // float64 needs 9 bytes
	}
	rv := reflect.ValueOf(v)
	if err := enc.Encode(rv); err != nil {
		return nil, err
	}
	data := enc.Buf.Bytes()
	result := make([]byte, len(data))
	copy(result, data)
	return result, nil
}

func marshalBytes(v []byte) ([]byte, error) {
	enc := getEncoderFromPool()
	defer putEncoderToPool(enc)
	if enc.Buf != nil {
		enc.Buf.Reset()
		enc.Buf.Grow(len(v) + 8) // bytes + length encoding
	}
	rv := reflect.ValueOf(v)
	if err := enc.Encode(rv); err != nil {
		return nil, err
	}
	data := enc.Buf.Bytes()
	result := make([]byte, len(data))
	copy(result, data)
	return result, nil
}

// Errors
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "beve: Unmarshal(nil)"
	}
	if e.Type.Kind() != reflect.Ptr {
		return "beve: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "beve: Unmarshal(nil " + e.Type.String() + ")"
}

// UnsupportedError is an alias for core.UnsupportedError
type UnsupportedError = core.UnsupportedError

// Indent sets the indentation for pretty-printed output.
// Currently not implemented as BEVE is binary.
func (e *Encoder) Indent(prefix, indent string) {
	// BEVE is binary, no indentation
}

// SetEscapeHTML specifies whether problematic HTML characters should be escaped.
// Not applicable for binary format.
func (e *Encoder) SetEscapeHTML(on bool) {
	// Not applicable
}

// buffer pool for zero-allocation
var bufferPool = &sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func getBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func putBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}
