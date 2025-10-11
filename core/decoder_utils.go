package core

import (
	"reflect"
	"time"
)

// decoder_utils.go - Utility functions for decoder

// setIntValue writes a signed integer into the destination reflect.Value.
func setIntValue(v reflect.Value, value int64) error {
	// Special case: time.Time (stored as int64 Unix nanos)
	if v.Type().PkgPath() == "time" && v.Type().Name() == "Time" {
		t := time.Unix(0, value)
		v.Set(reflect.ValueOf(t))
		return nil
	}

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(value)
		return nil
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
		return nil
	default:
		return &UnsupportedError{"cannot set value of type " + v.Type().String()}
	}
}

// setUintValue writes an unsigned integer into the destination reflect.Value.
func setUintValue(v reflect.Value, value uint64) error {
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v.SetUint(value)
		return nil
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
		return nil
	default:
		return &UnsupportedError{"cannot set value of type " + v.Type().String()}
	}
}

// setFloatValue writes a floating-point number into the destination reflect.Value.
func setFloatValue(v reflect.Value, value float64) error {
	switch v.Kind() {
	case reflect.Float32:
		v.SetFloat(float64(float32(value)))
		return nil
	case reflect.Float64:
		v.SetFloat(value)
		return nil
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
		return nil
	default:
		return &UnsupportedError{"cannot set value of type " + v.Type().String()}
	}
}

// setStringValue writes a string into the destination reflect.Value.
func setStringValue(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
		return nil
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
		return nil
	default:
		return &UnsupportedError{"cannot set value of type " + v.Type().String()}
	}
}

// SetBool sets a boolean value.
//
// Used for decoding boolean values to bool-typed variables.
func (d *Decoder) SetBool(v reflect.Value, b bool) error {
	if v.Kind() == reflect.Bool {
		v.SetBool(b)
		return nil
	}
	if v.Kind() == reflect.Interface {
		v.Set(reflect.ValueOf(b))
		return nil
	}
	return &UnsupportedError{"expected bool"}
}

// SetNil sets a nil/zero value.
//
// Used for decoding BEVE null values.
func (d *Decoder) SetNil(v reflect.Value) error {
	if v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		v.Set(reflect.Zero(v.Type()))
		return nil
	}
	// For other types, set zero value
	v.Set(reflect.Zero(v.Type()))
	return nil
}

// EnsureSliceLength ensures a slice has the specified length.
//
// This is used when decoding arrays to ensure the destination slice
// is large enough to hold all elements.
func EnsureSliceLength(v reflect.Value, length int) error {
	if v.Kind() != reflect.Slice {
		return &UnsupportedError{"expected slice"}
	}

	if v.IsNil() {
		// Create new slice
		v.Set(reflect.MakeSlice(v.Type(), length, length))
	} else if v.Len() < length {
		// Grow existing slice
		newSlice := reflect.MakeSlice(v.Type(), length, length)
		reflect.Copy(newSlice, v)
		v.Set(newSlice)
	} else if v.Len() > length {
		// Truncate slice
		v.SetLen(length)
	}

	return nil
}

// SetBoolElement sets a boolean element in a slice or array.
func SetBoolElement(elem reflect.Value, value bool) error {
	switch elem.Kind() {
	case reflect.Bool:
		elem.SetBool(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if value {
			elem.SetInt(1)
		} else {
			elem.SetInt(0)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if value {
			elem.SetUint(1)
		} else {
			elem.SetUint(0)
		}
	case reflect.Interface:
		elem.Set(reflect.ValueOf(value))
	default:
		return &UnsupportedError{"cannot set bool to " + elem.Type().String()}
	}
	return nil
}

// SetSignedElement sets a signed integer element with proper sizing.
func SetSignedElement(elem reflect.Value, value int64, byteCount int) error {
	switch elem.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		elem.SetInt(value)
	case reflect.Interface:
		// Choose appropriate int type based on byte count
		switch byteCount {
		case 1:
			elem.Set(reflect.ValueOf(int8(value)))
		case 2:
			elem.Set(reflect.ValueOf(int16(value)))
		case 4:
			elem.Set(reflect.ValueOf(int32(value)))
		default:
			elem.Set(reflect.ValueOf(value))
		}
	default:
		return &UnsupportedError{"cannot set int to " + elem.Type().String()}
	}
	return nil
}

// SetUnsignedElement sets an unsigned integer element with proper sizing.
func SetUnsignedElement(elem reflect.Value, value uint64, byteCount int) error {
	switch elem.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		elem.SetUint(value)
	case reflect.Interface:
		// Choose appropriate uint type based on byte count
		switch byteCount {
		case 1:
			elem.Set(reflect.ValueOf(uint8(value)))
		case 2:
			elem.Set(reflect.ValueOf(uint16(value)))
		case 4:
			elem.Set(reflect.ValueOf(uint32(value)))
		default:
			elem.Set(reflect.ValueOf(value))
		}
	default:
		return &UnsupportedError{"cannot set uint to " + elem.Type().String()}
	}
	return nil
}

// SetFloatElement sets a floating-point element with proper sizing.
func SetFloatElement(elem reflect.Value, value float64, byteCount int) error {
	switch elem.Kind() {
	case reflect.Float32, reflect.Float64:
		elem.SetFloat(value)
	case reflect.Interface:
		if byteCount == 4 {
			elem.Set(reflect.ValueOf(float32(value)))
		} else {
			elem.Set(reflect.ValueOf(value))
		}
	default:
		return &UnsupportedError{"cannot set float to " + elem.Type().String()}
	}
	return nil
}

// CheckedLength validates a length value and returns it as int.
//
// This prevents integer overflow attacks.
func CheckedLength(size uint64) (int, error) {
	if size > 0x7fffffff {
		return 0, &UnsupportedError{"array size too large"}
	}
	return int(size), nil
}
