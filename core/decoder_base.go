// Package core provides the internal implementation of BEVE encoding/decoding.
//
// This file contains the base decoder structure and main dispatch logic.
package core

import (
	"reflect"
	"sync"
	"time"
)

// timeFromUnixNano converts Unix nanoseconds to time.Time.
//
//go:inline
func timeFromUnixNano(nanos int64) time.Time {
	return time.Unix(0, nanos)
}

// Decoder handles the decoding of BEVE format to values.
//
// The decoder maintains state during decoding:
//   - data: The input byte slice
//   - pos: Current read position
//
// Thread Safety:
//   - Decoders are NOT thread-safe
//   - Each goroutine should use its own decoder
type Decoder struct {
	Data []byte // Input data - exported for access
	Pos  int    // Current position - exported for access
}

// decoderPool reuses decoder instances to avoid repeated allocations.
var decoderPool = sync.Pool{
	New: func() interface{} {
		return &Decoder{}
	},
}

// stringSlicePool reuses string slices to reduce allocations in decodeStringTypedArray.
// Slices are pooled by size buckets: 16, 64, 256, 1024, 4096.
var stringSlicePools = [5]sync.Pool{
	{New: func() interface{} { s := make([]string, 0, 16); return &s }},   // 0-16
	{New: func() interface{} { s := make([]string, 0, 64); return &s }},   // 17-64
	{New: func() interface{} { s := make([]string, 0, 256); return &s }},  // 65-256
	{New: func() interface{} { s := make([]string, 0, 1024); return &s }}, // 257-1024
	{New: func() interface{} { s := make([]string, 0, 4096); return &s }}, // 1025-4096
}

// getStringSlice returns a string slice from the pool with at least the requested capacity.
// Only pools slices >= 256 elements to avoid overhead on small arrays.
func getStringSlice(length int) []string {
	// OPTIMIZATION: Don't pool small slices - allocation is faster
	if length < 256 {
		return make([]string, length)
	}

	var poolIdx int
	switch {
	case length <= 256:
		poolIdx = 2
	case length <= 1024:
		poolIdx = 3
	case length <= 4096:
		poolIdx = 4
	default:
		// Too large for pooling, allocate directly
		return make([]string, length)
	}

	slicePtr := stringSlicePools[poolIdx].Get().(*[]string)
	slice := *slicePtr
	if cap(slice) < length {
		slice = make([]string, length)
	} else {
		slice = slice[:length]
		// No need to clear - will be overwritten during decode
	}
	return slice
}

// putStringSlice returns a string slice to the pool.
func putStringSlice(slice []string) {
	if cap(slice) > 4096 {
		return // Too large, let GC handle it
	}

	// Clear strings to allow GC
	for i := range slice {
		slice[i] = ""
	}
	slice = slice[:0]

	var poolIdx int
	switch cap(slice) {
	case 16:
		poolIdx = 0
	case 64:
		poolIdx = 1
	case 256:
		poolIdx = 2
	case 1024:
		poolIdx = 3
	case 4096:
		poolIdx = 4
	default:
		return // Not from our pool
	}

	stringSlicePools[poolIdx].Put(&slice)
}

// NewDecoder creates a new decoder for the given data.
func NewDecoder(data []byte) *Decoder {
	dec := decoderPool.Get().(*Decoder)
	dec.Data = data
	dec.Pos = 0
	return dec
}

// PutDecoderToPool returns a decoder instance to the global pool for reuse.
//
// Callers must ensure the decoder is no longer in use. The decoder's state is
// cleared before pooling to avoid retaining references to large buffers.
func PutDecoderToPool(dec *Decoder) {
	dec.Data = nil
	dec.Pos = 0
	decoderPool.Put(dec)
}

// Decode is the main entry point for decoding a BEVE value.
//
// This method dispatches to type-specific decoders based on the header byte.
// It handles:
//   - Null values
//   - RawMessage types
//   - BinaryUnmarshaler interface
//   - All Go primitive types
//   - Collections (slice, map, struct, array)
//
// BEVE Format Header (1 byte):
//   - Bits 0-2: Type (null/bool=0, number=1, string=2, object=3, typed-array=4, generic-array=5, extension=6)
//   - Bits 3+: Type-specific information
func (d *Decoder) Decode(v reflect.Value) error {
	if !v.IsValid() {
		return &UnsupportedError{"invalid destination"}
	}

	header, err := d.ReadByte()
	if err != nil {
		return err
	}
	start := d.Pos - 1

	// Check if destination is RawMessage
	if isRawMessageType(v.Type()) {
		d.Pos = start
		raw, err := d.captureRawValue()
		if err != nil {
			return err
		}
		setRawMessageValue(v, raw)
		return nil
	}

	// Check if destination implements BinaryUnmarshaler
	if shouldCheckBinaryUnmarshaler(v) {
		if um, err := d.lookupBinaryUnmarshaler(v); err != nil {
			return err
		} else if um != nil {
			d.Pos = start
			raw, err := d.captureRawValue()
			if err != nil {
				return err
			}
			return um.UnmarshalBEVE(raw)
		}
	}

	// Special case: time.Time (decode from int64 Unix nanos)
	if v.Type().PkgPath() == "time" && v.Type().Name() == "Time" {
		// Decode as int64 (Unix nanoseconds)
		var nanos int64
		nanosVal := reflect.ValueOf(&nanos).Elem()
		d.Pos = start // Reset position to read header again
		header, _ = d.ReadByte()
		if err := d.DecodeNumber(nanosVal, header); err != nil {
			return err
		}
		// Convert int64 to time.Time
		t := timeFromUnixNano(nanos)
		v.Set(reflect.ValueOf(t))
		return nil
	}

	// Dispatch based on type bits (bits 0-2)
	switch header & 0x07 {
	case 0: // null or bool
		if header&0x08 != 0 { // bool bit
			b := header&0x10 != 0 // true/false bit
			return d.SetBool(v, b)
		}
		return d.SetNil(v)
	case 1: // number
		return d.DecodeNumber(v, header)
	case 2: // string
		return d.DecodeString(v)
	case 3: // object
		return d.DecodeObject(v, header)
	case 4: // typed array
		return d.DecodeTypedArray(v, header)
	case 5: // generic array
		return d.DecodeGenericArray(v)
	case 6: // extension
		return d.DecodeExtension(v, header)
	default:
		return &UnsupportedError{"unknown type"}
	}
}

// shouldCheckBinaryUnmarshaler indicates whether a type could implement
// BinaryUnmarshaler and therefore requires an interface lookup.
func shouldCheckBinaryUnmarshaler(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		return true
	}
	return v.Type().PkgPath() != ""
}

// Common types and functions are in common.go

// lookupBinaryUnmarshaler checks if the value implements BinaryUnmarshaler.
func (d *Decoder) lookupBinaryUnmarshaler(v reflect.Value) (BinaryUnmarshaler, error) {
	if !v.IsValid() {
		return nil, nil
	}

	if v.Kind() != reflect.Interface && v.Kind() != reflect.Ptr {
		t := v.Type()
		if t.PkgPath() == "" {
			switch t.Kind() {
			case reflect.Bool,
				reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
				reflect.Float32, reflect.Float64,
				reflect.String,
				reflect.Slice, reflect.Array, reflect.Map, reflect.Struct:
				return nil, nil
			}
		}
	}

	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil, nil
		}
		return d.lookupBinaryUnmarshaler(v.Elem())
	}

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		if v.CanInterface() {
			if um, ok := v.Interface().(BinaryUnmarshaler); ok {
				return um, nil
			}
		}
		return d.lookupBinaryUnmarshaler(v.Elem())
	}

	if v.CanInterface() {
		if um, ok := v.Interface().(BinaryUnmarshaler); ok {
			return um, nil
		}
	}

	if v.CanAddr() {
		addr := v.Addr()
		if addr.CanInterface() {
			if um, ok := addr.Interface().(BinaryUnmarshaler); ok {
				return um, nil
			}
		}
	}

	return nil, nil
}

// captureRawValue captures raw BEVE bytes for current value.
func (d *Decoder) captureRawValue() ([]byte, error) {
	start := d.Pos
	if err := d.SkipValue(); err != nil {
		if ue, ok := err.(*UnsupportedError); ok && ue.Msg == "unexpected end of data" {
			raw := make([]byte, len(d.Data)-start)
			copy(raw, d.Data[start:])
			d.Pos = len(d.Data)
			return raw, nil
		}
		return nil, err
	}
	raw := make([]byte, d.Pos-start)
	copy(raw, d.Data[start:d.Pos])
	return raw, nil
}
