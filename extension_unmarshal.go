package beve

import (
	"fmt"
	"reflect"
)

// UnmarshalAuto auto-detects the encoding format and unmarshals accordingly
// Supports both generic BEVE v1.0 and all extensions
func UnmarshalAuto(data []byte, v interface{}) error {
	if len(data) == 0 {
		return fmt.Errorf("empty data")
	}

	header := data[0]

	// Check for extension headers (0x86-0xF6)
	if header >= 0x86 && header <= 0xF6 {
		return unmarshalExtension(data, v, header)
	}

	// Generic BEVE v1.0 decoding
	return Unmarshal(data, v)
}

// unmarshalExtension handles extension-specific unmarshaling
func unmarshalExtension(data []byte, v interface{}, header byte) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("unmarshal target must be non-nil pointer")
	}

	switch header {
	case ExtTypedArray:
		// Unmarshal typed object array
		result, err := DecodeTypedArray(data)
		if err != nil {
			return err
		}
		return assignValue(rv.Elem(), reflect.ValueOf(result))

	case ExtTypedNestedArray:
		// Unmarshal nested typed array
		result, err := DecodeTypedNestedArray(data)
		if err != nil {
			return err
		}
		return assignValue(rv.Elem(), reflect.ValueOf(result))

	case ExtFieldIndex:
		// Unmarshal indexed object
		result, err := DecodeIndexedObject(data)
		if err != nil {
			return err
		}
		return assignValue(rv.Elem(), reflect.ValueOf(result))

	case ExtTimestamp:
		// Unmarshal timestamp
		ts, err := DecodeTimestamp(data)
		if err != nil {
			return err
		}
		// Try to convert to time.Time if target is time.Time
		if rv.Elem().Type().String() == "time.Time" {
			return assignValue(rv.Elem(), reflect.ValueOf(ts.ToTime()))
		}
		return assignValue(rv.Elem(), reflect.ValueOf(ts))

	case ExtDuration:
		// Unmarshal duration
		d, err := DecodeDuration(data)
		if err != nil {
			return err
		}
		return assignValue(rv.Elem(), reflect.ValueOf(d))

	case ExtInterval:
		// Unmarshal interval
		start, end, err := DecodeInterval(data)
		if err != nil {
			return err
		}
		// Return as [2]time.Time array
		interval := [2]interface{}{start, end}
		return assignValue(rv.Elem(), reflect.ValueOf(interval))

	case ExtUUID:
		// Unmarshal UUID
		uuid, err := DecodeUUID(data)
		if err != nil {
			return err
		}
		// Check if target wants string
		if rv.Elem().Kind() == reflect.String {
			uuidStr, err := DecodeUUIDString(data)
			if err != nil {
				return err
			}
			return assignValue(rv.Elem(), reflect.ValueOf(uuidStr))
		}
		return assignValue(rv.Elem(), reflect.ValueOf(uuid))

	case ExtRegex:
		// Unmarshal regex
		regexData, err := DecodeRegExp(data)
		if err != nil {
			return err
		}
		// Check if target wants *regexp.Regexp
		if rv.Elem().Type().String() == "*regexp.Regexp" {
			r, err := UnmarshalRegExp(data)
			if err != nil {
				return err
			}
			return assignValue(rv.Elem(), reflect.ValueOf(r))
		}
		return assignValue(rv.Elem(), reflect.ValueOf(regexData))

	default:
		return fmt.Errorf("unsupported extension: 0x%02X", header)
	}
}

// assignValue safely assigns a value to a reflect.Value
func assignValue(dst, src reflect.Value) error {
	if !dst.CanSet() {
		return fmt.Errorf("cannot set value")
	}

	// Handle type conversion
	if src.Type().AssignableTo(dst.Type()) {
		dst.Set(src)
		return nil
	}

	// Try type conversion
	if src.Type().ConvertibleTo(dst.Type()) {
		dst.Set(src.Convert(dst.Type()))
		return nil
	}

	return fmt.Errorf("cannot assign %s to %s", src.Type(), dst.Type())
}

// DetectEncoding returns the encoding type used
func DetectEncoding(data []byte) string {
	if len(data) == 0 {
		return "empty"
	}

	header := data[0]

	switch header {
	case ExtFieldIndex:
		return "field_index"
	case ExtTypedArray:
		return "typed_array"
	case ExtTypedNestedArray:
		return "typed_nested_array"
	case ExtCompressionHint:
		return "compression_hint"
	case ExtTimestamp:
		return "timestamp"
	case ExtDuration:
		return "duration"
	case ExtInterval:
		return "interval"
	case ExtRecurringEvent:
		return "recurring_event"
	case ExtUUID:
		return "uuid"
	case ExtRegex:
		return "regex"
	default:
		// Check if it's generic BEVE
		typeCode := header & 0x07
		switch typeCode {
		case 0:
			return "beve_null_or_bool"
		case 1:
			return "beve_number"
		case 2:
			return "beve_string"
		case 3:
			return "beve_object"
		case 4:
			return "beve_typed_array"
		case 5:
			return "beve_generic_array"
		case 6:
			return "beve_extension"
		default:
			return fmt.Sprintf("unknown_0x%02X", header)
		}
	}
}

// IsExtension checks if data uses BEVE extensions
func IsExtension(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	header := data[0]
	return header >= 0x86 && header <= 0xF6
}

// GetExtensionID extracts the extension ID from header
func GetExtensionID(data []byte) (int, error) {
	if len(data) == 0 {
		return -1, fmt.Errorf("empty data")
	}

	header := data[0]
	if header < 0x86 || header > 0xF6 {
		return -1, fmt.Errorf("not an extension header: 0x%02X", header)
	}

	// Extension ID = (header - 0x86) >> 3
	// 0x86 = Extension 0, 0x8E = Extension 1, etc.
	extID := int(header-0x86) >> 3

	return extID, nil
}
