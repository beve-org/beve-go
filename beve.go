// Package beve implements BEVE (Binary Efficient Versatile Encoding) for Go,
// providing high-performance binary serialization with zero-allocation goals
// and full compatibility with Go's encoding/json interfaces.
package beve

import (
	"bytes"
	"io"
	"reflect"
	"sync"
)

// Version represents the BEVE specification version
const Version = "1.0"

// Marshal encodes v into BEVE binary format.
// It follows the same interface as encoding/json.Marshal.
func Marshal(v interface{}) ([]byte, error) {
	e := NewEncoder(nil)
	return e.Encode(v)
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
		if err := enc.encode(reflect.ValueOf(v)); err != nil {
			return nil, err
		}
		result := make([]byte, buf.Len())
		copy(result, buf.Bytes())
		return result, nil
	}
	// Streaming encoding
	enc := newEncoder(e.w)
	return nil, enc.encode(reflect.ValueOf(v))
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

	return &UnsupportedError{"unsupported reader type"}
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

type UnsupportedError struct {
	msg string
}

func (e *UnsupportedError) Error() string {
	return "beve: " + e.msg
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
