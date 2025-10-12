// Package beve implements BEVE (Binary Efficient Versatile Encoding) for Go,
// providing high-performance binary serialization with zero-allocation goals
// and full compatibility with Go's encoding/json interfaces.
package beve

import (
	"bytes"
	"io"
	"reflect"
	"sync"
	"time"

	"github.com/beve-org/beve-go/core"
)

// Version represents the BEVE specification version
const Version = "1.2.0"

// StructTag is the default struct tag name to use for field configuration.
// Default is "beve", but can be changed to "json", "msgpack", or any other tag
// for compatibility with existing codebases.
//
// Example:
//
//	beve.SetStructTag("json") // Use json tags instead of beve tags
//
// The library always falls back to "json" tags if the configured tag is not found.
var StructTag = "beve"

var structTagMu sync.RWMutex

func init() {
	// Set up the function pointer for core package to access StructTag
	core.GetStructTag = GetStructTag
}

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
	case time.Time:
		return marshalTime(val)
	}

	// Slow path: use reflection with buffer pooling
	enc := getEncoderFromPool()
	if enc.Buf != nil {
		enc.Buf.Reset()
	}

	rv := reflect.ValueOf(v)
	if err := enc.Encode(rv); err != nil {
		putEncoderToPool(enc)
		return nil, err
	}

	// Copy encoded data - make+copy is consistently faster than append for small sizes
	encoded := enc.Buf.Bytes()
	n := len(encoded)
	result := make([]byte, n)
	copy(result, encoded)

	// Return encoder to pool (buffer is retained and reset for next use)
	putEncoderToPool(enc)

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
	w   io.Writer
	enc *core.Encoder // Reusable encoder
	buf *bytes.Buffer // Reusable buffer
}

// NewEncoder creates a new BEVE encoder.
func NewEncoder(w io.Writer) *Encoder {
	buf := getBuffer()
	enc := core.GetEncoderFromPool()
	if enc.Buf == nil {
		enc.Buf = &core.Buffer{}
	}
	enc.Buf.Reset()
	return &Encoder{
		w:   w,
		enc: enc,
		buf: buf,
	}
}

// Encode encodes v to BEVE format and writes to the underlying writer.
// If no writer is set, returns the encoded bytes.
func (e *Encoder) Encode(v interface{}) ([]byte, error) {
	// Reset encoder state
	if e.enc != nil && e.enc.Buf != nil {
		e.enc.Buf.Reset()
	}

	if e.w == nil {
		// Non-streaming mode
		if err := e.enc.Encode(reflect.ValueOf(v)); err != nil {
			return nil, err
		}
		data := e.enc.Buf.Bytes()
		result := make([]byte, len(data))
		copy(result, data)
		return result, nil
	}

	// Streaming encoding - write to buffer first then to writer
	if err := e.enc.Encode(reflect.ValueOf(v)); err != nil {
		return nil, err
	}

	data := e.enc.Buf.Bytes()
	if _, err := e.w.Write(data); err != nil {
		return nil, err
	}

	return nil, nil
}

// Close releases resources. Call when done with the encoder.
func (e *Encoder) Close() error {
	if e.enc != nil {
		core.PutEncoderToPool(e.enc)
		e.enc = nil
	}
	if e.buf != nil {
		putBuffer(e.buf)
		e.buf = nil
	}
	return nil
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
	// Copy data - make+copy for consistent small allocations
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
	// Use pooled byte slice
	result := getByteSlice()
	*result = growSlice(result, len(data))
	copy(*result, data)
	return *result, nil
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
	// Use pooled byte slice
	result := getByteSlice()
	*result = growSlice(result, len(data))
	copy(*result, data)
	return *result, nil
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
	// Use pooled byte slice
	result := getByteSlice()
	*result = growSlice(result, len(data))
	copy(*result, data)
	return *result, nil
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
	// Use pooled byte slice
	result := getByteSlice()
	*result = growSlice(result, len(data))
	copy(*result, data)
	return *result, nil
}

// marshalTime is a fast-path encoder for time.Time (avoids reflection).
//
// Performance: ~10-15ns for time encoding.
// Format: Encodes as Unix nanoseconds (int64) for maximum precision.
func marshalTime(t time.Time) ([]byte, error) {
	// Encode as Unix nanoseconds (int64)
	// This preserves full nanosecond precision
	nanos := t.UnixNano()
	return marshalInt64(nanos)
}

// marshalInt64 is a helper for encoding int64 values.
func marshalInt64(n int64) ([]byte, error) {
	enc := getEncoderFromPool()
	defer putEncoderToPool(enc)
	if enc.Buf != nil {
		enc.Buf.Reset()
	}
	// Use reflection to encode int64
	rv := reflect.ValueOf(n)
	if err := enc.Encode(rv); err != nil {
		return nil, err
	}
	data := enc.Buf.Bytes()
	// Use pooled byte slice
	result := getByteSlice()
	*result = growSlice(result, len(data))
	copy(*result, data)
	return *result, nil
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

// SetStructTag changes the struct tag name used for field configuration.
// Default is "beve". Common alternatives are "json", "msgpack", etc.
//
// This is useful for compatibility with existing codebases that already use
// json tags. The library always falls back to "json" tags if the configured
// tag is not found.
//
// Example:
//
//	beve.SetStructTag("json") // Use json tags
//	data, _ := beve.Marshal(myStruct) // Will read json:"..." tags
//
// Note: Changing the tag name clears the encoder/decoder cache, so it's
// recommended to set this once at application startup.
func SetStructTag(tag string) {
	if tag == "" {
		tag = "beve" // Default
	}
	structTagMu.Lock()
	StructTag = tag
	structTagMu.Unlock()

	// Clear caches to force rebuilding with new tag
	core.ClearEncoderCache()
	core.ClearDecoderCache()
}

// GetStructTag returns the current struct tag name being used.
func GetStructTag() string {
	structTagMu.RLock()
	defer structTagMu.RUnlock()
	return StructTag
}

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
