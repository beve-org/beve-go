package core

import (
	"reflect"
	"sync"
)

// UnsupportedError is returned when encoding/decoding encounters an unsupported type or value.
type UnsupportedError struct {
	Msg string
}

func (e *UnsupportedError) Error() string {
	return "beve: " + e.Msg
}

// isRawMessageType checks if the type is RawMessage.
func isRawMessageType(t reflect.Type) bool {
	// RawMessage is a []byte type in the main package
	// We check if it's a slice of bytes with name "RawMessage"
	if t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Uint8 {
		return t.Name() == "RawMessage"
	}
	return false
}

// setRawMessageValue sets the RawMessage value.
func setRawMessageValue(v reflect.Value, raw []byte) {
	if raw == nil {
		v.Set(reflect.Zero(v.Type()))
		return
	}
	v.SetBytes(raw)
}

// BinaryUnmarshaler is the interface implemented by types that can
// unmarshal a BEVE description of themselves.
type BinaryUnmarshaler interface {
	UnmarshalBEVE([]byte) error
}

// BinaryMarshaler is the interface implemented by types that can
// marshal themselves into a BEVE description.
type BinaryMarshaler interface {
	MarshalBEVE() ([]byte, error)
}

// encoderFunc is a cached encoder function for a specific type.
type encoderFunc func(*Encoder, reflect.Value) error

// encoderFuncCache caches encoder functions by type for fast dispatch.
//
// Instead of going through the type switch in Encode() every time,
// we cache the appropriate encoder function for each type after the
// first lookup. This provides ~30-40% performance improvement.
//
// The cache is thread-safe using sync.Map.
var encoderFuncCache sync.Map // map[reflect.Type]encoderFunc

// getEncoderFunc returns the cached encoder function for a type,
// or determines and caches it on first use.
func getEncoderFunc(t reflect.Type) encoderFunc {
	// Check cache first
	if fn, ok := encoderFuncCache.Load(t); ok {
		return fn.(encoderFunc)
	}

	// Determine the encoder function based on type
	var build encoderFunc

	// Check for time.Time (encode as int64 Unix nanos)
	if t.PkgPath() == "time" && t.Name() == "Time" {
		build = func(e *Encoder, v reflect.Value) error {
			// Get time.Time value and convert to Unix nanos
			t := v.Interface().(interface{ UnixNano() int64 })
			nanos := t.UnixNano()
			return encodeByKind(e, reflect.ValueOf(nanos))
		}
		return storeEncoderFunc(t, build)
	}

	// Check for RawMessage
	if isRawMessageType(t) {
		build = func(e *Encoder, v reflect.Value) error {
			return e.encodeRawMessage(v.Bytes())
		}
		return storeEncoderFunc(t, build)
	}

	// Check for BinaryMarshaler interface
	typeInfo := getTypeInfo(t)
	if typeInfo.implMarsh {
		build = func(e *Encoder, v reflect.Value) error {
			if v.CanInterface() {
				if bm, ok := v.Interface().(BinaryMarshaler); ok {
					return e.encodeBinaryMarshaler(bm)
				}
			}
			// Fallback to kind-based encoding
			return encodeByKind(e, v)
		}
		return storeEncoderFunc(t, build)
	}

	if typeInfo.ptrImplMarsh {
		build = func(e *Encoder, v reflect.Value) error {
			if v.Kind() != reflect.Ptr && v.CanAddr() {
				addr := v.Addr()
				if addr.CanInterface() {
					if bm, ok := addr.Interface().(BinaryMarshaler); ok {
						return e.encodeBinaryMarshaler(bm)
					}
				}
			}
			// Fallback to kind-based encoding
			return encodeByKind(e, v)
		}
		return storeEncoderFunc(t, build)
	}

	if t.Kind() == reflect.Struct {
		build = buildStructEncoder(t)
		return storeEncoderFunc(t, build)
	}

	// Kind-based encoder
	build = encodeByKind
	return storeEncoderFunc(t, build)
}

func storeEncoderFunc(t reflect.Type, fn encoderFunc) encoderFunc {
	if existing, loaded := encoderFuncCache.LoadOrStore(t, fn); loaded {
		return existing.(encoderFunc)
	}
	return fn
}

// encodeByKind encodes a value based on its kind.
func encodeByKind(e *Encoder, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		return e.EncodeNull()

	case reflect.Bool:
		return e.encodeBool(extractBool(v))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return e.encodeInt(extractInt(v))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return e.encodeUint(extractUint(v))

	case reflect.Float32, reflect.Float64:
		return e.encodeFloat(extractFloat(v), v.Kind())

	case reflect.String:
		return e.EncodeString(extractString(v))

	case reflect.Slice, reflect.Array:
		return e.encodeSlice(v)

	case reflect.Map:
		return e.encodeMapFast(v)

	case reflect.Struct:
		return e.encodeStructFast(v)

	case reflect.Interface:
		if v.IsNil() {
			return e.EncodeNull()
		}
		return e.Encode(v.Elem())

	case reflect.Ptr:
		if v.IsNil() {
			return e.EncodeNull()
		}
		return e.Encode(v.Elem())

	default:
		return &UnsupportedError{"unsupported type: " + v.Type().String()}
	}
}

// ClearEncoderCache clears the encoder function cache.
// This is useful when changing struct tag configuration.
func ClearEncoderCache() {
	encoderFuncCache.Range(func(key, value interface{}) bool {
		encoderFuncCache.Delete(key)
		return true
	})
}

// ClearDecoderCache clears the decoder cache.
// This is useful when changing struct tag configuration.
func ClearDecoderCache() {
	// Decoder doesn't have a separate cache yet, but we include this
	// for future-proofing and API consistency.
}
