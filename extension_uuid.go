package beve

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// EncodeUUID encodes a UUID (Extension 8)
// Size: 18 bytes (header + version + 16 bytes)
func EncodeUUID(u [16]byte) ([]byte, error) {
	buf := make([]byte, 18)

	// Header
	buf[0] = ExtUUID

	// Version (extracted from UUID byte 6, bits 4-7)
	version := (u[6] >> 4) & 0x0F
	buf[1] = version

	// Copy UUID bytes
	copy(buf[2:], u[:])

	return buf, nil
}

// DecodeUUID decodes Extension 8 UUID
func DecodeUUID(data []byte) ([16]byte, error) {
	if len(data) < 18 || data[0] != ExtUUID {
		return [16]byte{}, fmt.Errorf("invalid UUID header")
	}

	var u [16]byte
	copy(u[:], data[2:18])

	return u, nil
}

// EncodeUUIDString encodes a UUID string (e.g., "550e8400-e29b-41d4-a716-446655440000")
func EncodeUUIDString(s string) ([]byte, error) {
	// Remove hyphens
	s = strings.ReplaceAll(s, "-", "")

	if len(s) != 32 {
		return nil, fmt.Errorf("invalid UUID string length: %d", len(s))
	}

	// Decode hex string
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID string: %w", err)
	}

	var u [16]byte
	copy(u[:], bytes)

	return EncodeUUID(u)
}

// DecodeUUIDString decodes Extension 8 UUID to string
func DecodeUUIDString(data []byte) (string, error) {
	u, err := DecodeUUID(data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		u[0:4],
		u[4:6],
		u[6:8],
		u[8:10],
		u[10:16],
	), nil
}

// UUIDVersion extracts the UUID version
func UUIDVersion(u [16]byte) byte {
	return (u[6] >> 4) & 0x0F
}

// UUIDVariant extracts the UUID variant
func UUIDVariant(u [16]byte) byte {
	return (u[8] >> 6) & 0x03
}

// MarshalUUID marshals a UUID to BEVE bytes
func MarshalUUID(u [16]byte) ([]byte, error) {
	return EncodeUUID(u)
}

// UnmarshalUUID unmarshals BEVE UUID bytes
func UnmarshalUUID(data []byte) ([16]byte, error) {
	return DecodeUUID(data)
}

// MarshalUUIDString marshals a UUID string to BEVE bytes
func MarshalUUIDString(s string) ([]byte, error) {
	return EncodeUUIDString(s)
}

// UnmarshalUUIDString unmarshals BEVE UUID bytes to string
func UnmarshalUUIDString(data []byte) (string, error) {
	return DecodeUUIDString(data)
}
