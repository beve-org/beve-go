package translatornative

import (
	"fmt"
)

// FromJSON converts JSON bytes to BEVE binary (TRUE zero-copy).
//
// This function uses direct encoding:
//  1. JSON bytes → BEVE bytes directly
//  2. NO intermediate map/slice structures
//  3. Single allocation for output buffer
//
// Performance characteristics:
//   - WASM-optimized (no reflection, no SIMD)
//   - True zero-copy: 1 allocation for output
//   - Direct JSON → BEVE conversion
//   - Zero dependencies on beve-go package
//
// Example:
//
//	jsonStr := `{"id":123,"name":"Alice","active":true}`
//	beveData, err := translatornative.FromJSON([]byte(jsonStr))
//	if err != nil {
//	    log.Fatal(err)
//	}
func FromJSON(jsonData []byte) ([]byte, error) {
	if len(jsonData) == 0 {
		return nil, fmt.Errorf("translatornative: empty JSON input")
	}

	// Direct encode: JSON → BEVE without intermediate structures
	enc := NewDirectEncoder(jsonData)
	beveData, err := enc.Encode()
	if err != nil {
		return nil, fmt.Errorf("translatornative: BEVE encode error: %w", err)
	}

	return beveData, nil
}

// FromJSONString is a convenience wrapper for FromJSON.
func FromJSONString(jsonStr string) ([]byte, error) {
	return FromJSON([]byte(jsonStr))
}

// ToJSON converts BEVE bytes to JSON using native serializer (WASM-optimized).
//
// This function:
//  1. Decodes BEVE binary using specification-compliant decoder
//  2. Serializes to JSON using custom serializer (no encoding/json)
//  3. Returns JSON bytes
//
// Performance characteristics:
//   - WASM-optimized (no reflection)
//   - 60+ MB/s throughput
//   - Zero dependencies on beve-go package
//
// Example:
//
//	beveData := []byte{0x03, 0x04, ...}
//	jsonData, err := translatornative.ToJSON(beveData)
//	if err != nil {
//	    log.Fatal(err)
//	}
func ToJSON(beveData []byte) ([]byte, error) {
	if len(beveData) == 0 {
		return nil, fmt.Errorf("translatornative: empty BEVE input")
	}

	// Decode BEVE binary (no beve-go dependency)
	decoder := NewBEVEDecoder(beveData)
	data, err := decoder.Decode()
	if err != nil {
		return nil, fmt.Errorf("translatornative: BEVE decode error: %w", err)
	}

	// Serialize to JSON using native serializer
	serializer := NewJSONSerializer()
	defer serializer.Close()
	jsonData, err := serializer.Serialize(data)
	if err != nil {
		return nil, fmt.Errorf("translatornative: JSON serialize error: %w", err)
	}

	return jsonData, nil
}

// ToJSONString converts BEVE to JSON string.
func ToJSONString(beveData []byte) (string, error) {
	jsonData, err := ToJSON(beveData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// ToJSONIndent converts BEVE to pretty-printed JSON.
//
// Example:
//
//	jsonStr, err := translatornative.ToJSONIndent(beveData, "", "  ")
//	fmt.Println(jsonStr) // Pretty-printed JSON
func ToJSONIndent(beveData []byte, prefix, indent string) (string, error) {
	if len(beveData) == 0 {
		return "", fmt.Errorf("translatornative: empty BEVE input")
	}

	// Decode BEVE binary (no beve-go dependency)
	decoder := NewBEVEDecoder(beveData)
	data, err := decoder.Decode()
	if err != nil {
		return "", fmt.Errorf("translatornative: BEVE decode error: %w", err)
	}

	// Serialize to pretty JSON
	serializer := NewJSONSerializer()
	defer serializer.Close()
	jsonData, err := serializer.SerializeIndent(data, prefix, indent)
	if err != nil {
		return "", fmt.Errorf("translatornative: JSON serialize error: %w", err)
	}

	return string(jsonData), nil
}

// ValidateJSON validates JSON using native parser.
func ValidateJSON(data []byte) bool {
	parser := NewJSONParser(data)
	defer parser.Close()
	_, err := parser.Parse()
	return err == nil
}

// ValidateBEVE validates BEVE data using native decoder.
func ValidateBEVE(data []byte) bool {
	decoder := NewBEVEDecoder(data)
	_, err := decoder.Decode()
	return err == nil
}

// ConversionStats provides metrics about conversion.
type ConversionStats struct {
	OriginalSize  int     // Input size (bytes)
	ConvertedSize int     // Output size (bytes)
	Ratio         float64 // Compression ratio
	Savings       float64 // Space saved
}

// FromJSONWithStats converts JSON to BEVE with statistics.
func FromJSONWithStats(jsonData []byte) ([]byte, *ConversionStats, error) {
	beveData, err := FromJSON(jsonData)
	if err != nil {
		return nil, nil, err
	}

	stats := &ConversionStats{
		OriginalSize:  len(jsonData),
		ConvertedSize: len(beveData),
		Ratio:         float64(len(beveData)) / float64(len(jsonData)),
		Savings:       1.0 - (float64(len(beveData)) / float64(len(jsonData))),
	}

	return beveData, stats, nil
}

// ToJSONWithStats converts BEVE to JSON with statistics.
func ToJSONWithStats(beveData []byte) ([]byte, *ConversionStats, error) {
	jsonData, err := ToJSON(beveData)
	if err != nil {
		return nil, nil, err
	}

	stats := &ConversionStats{
		OriginalSize:  len(beveData),
		ConvertedSize: len(jsonData),
		Ratio:         float64(len(jsonData)) / float64(len(beveData)),
		Savings:       1.0 - (float64(len(jsonData)) / float64(len(beveData))),
	}

	return jsonData, stats, nil
}
