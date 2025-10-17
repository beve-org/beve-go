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
//   - arena: Optional arena allocator for temporary buffers (reduces GC pressure)
//
// Thread Safety:
//   - Decoders are NOT thread-safe
//   - Each goroutine should use its own decoder
//
// Arena Support:
//   - When arena is non-nil, temporary allocations (slices, maps, raw captures)
//     use arena memory instead of heap allocation
//   - Reduces GC pressure by 10-100× for high-allocation workloads
//   - All arena memory is freed in one operation when arena.Free() is called
//
// Example with arena:
//
//	arena := core.NewArena(64 * 1024) // 64KB arena
//	defer arena.Free()
//
//	dec := core.NewDecoderWithArena(data, arena)
//	var result MyStruct
//	dec.Decode(reflect.ValueOf(&result).Elem())
//	// All temporary allocations freed when arena.Free() is called
type Decoder struct {
	Data  []byte // Input data - exported for access
	Pos   int    // Current position - exported for access
	arena *Arena // Optional arena allocator (nil = standard heap allocation)
}

// decoderPool reuses decoder instances to avoid repeated allocations.
var decoderPool = sync.Pool{
	New: func() interface{} {
		return &Decoder{}
	},
}

// NewDecoder creates a new decoder for the given data.
func NewDecoder(data []byte) *Decoder {
	dec := decoderPool.Get().(*Decoder)
	dec.Data = data
	dec.Pos = 0
	dec.arena = nil // Standard heap allocation
	return dec
}

// NewDecoderWithArena creates a new decoder that uses arena allocation.
//
// The arena will be used for temporary allocations during decoding:
//   - Raw value captures (BinaryUnmarshaler, RawMessage)
//   - Typed array slices (int, uint, float, bool, string)
//   - Generic array/map allocations (when size known upfront)
//
// Benefits:
//   - Reduces GC pressure by 10-100× for high-allocation workloads
//   - Faster allocation (bump allocator vs heap)
//   - Bulk deallocation (one arena.Free() vs many GC cycles)
//
// Arena must outlive the decoder. Typical pattern:
//
//	arena := NewArena(64 * 1024)
//	defer arena.Free()
//
//	dec := NewDecoderWithArena(data, arena)
//	// ... decode ...
//
// Performance: ~2ns allocation overhead vs ~20ns heap allocation
func NewDecoderWithArena(data []byte, arena *Arena) *Decoder {
	dec := decoderPool.Get().(*Decoder)
	dec.Data = data
	dec.Pos = 0
	dec.arena = arena
	return dec
}

// PutDecoderToPool returns a decoder instance to the global pool for reuse.
//
// Callers must ensure the decoder is no longer in use. The decoder's state is
// cleared before pooling to avoid retaining references to large buffers.
func PutDecoderToPool(dec *Decoder) {
	dec.Data = nil
	dec.Pos = 0
	dec.arena = nil // Clear arena reference
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
// captureRawValue captures raw BEVE bytes for current value.
//
// If decoder has an arena, uses arena allocation for the raw buffer,
// otherwise falls back to standard heap allocation.
//
// Performance with arena: ~2ns allocation vs ~20ns heap
func (d *Decoder) captureRawValue() ([]byte, error) {
	start := d.Pos
	if err := d.SkipValue(); err != nil {
		if ue, ok := err.(*UnsupportedError); ok && ue.Msg == "unexpected end of data" {
			size := len(d.Data) - start
			var raw []byte
			if d.arena != nil {
				// Arena allocation (fast bump allocator)
				raw = d.arena.AllocBytes(size)
			} else {
				// Fallback to heap allocation
				raw = make([]byte, size)
			}
			copy(raw, d.Data[start:])
			d.Pos = len(d.Data)
			return raw, nil
		}
		return nil, err
	}

	size := d.Pos - start
	var raw []byte
	if d.arena != nil {
		// Arena allocation (fast bump allocator)
		raw = d.arena.AllocBytes(size)
	} else {
		// Fallback to heap allocation
		raw = make([]byte, size)
	}
	copy(raw, d.Data[start:d.Pos])
	return raw, nil
}
