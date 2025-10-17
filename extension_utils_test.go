package beve

import (
	"fmt"
	"testing"
)

// ============================================================================
// Tests for readCompressedSize (48% → 60% coverage target)
// ============================================================================

func TestReadCompressedSize_RoundTrip(t *testing.T) {
	// Test various sizes for proper encoding/decoding
	testSizes := []int{
		0, 1, 5, 63,           // 1-byte encoding
		64, 100, 16383,         // 2-byte encoding
		16384, 1000000,         // 4-byte encoding
		1073741823,             // Max 4-byte (30-bit)
	}

	for _, size := range testSizes {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			// Write
			buf := make([]byte, 8)
			written := writeCompressedSize(buf, size)
			if written == 0 {
				t.Fatalf("writeCompressedSize failed for %d", size)
			}

			// Read
			decoded, bytesRead, err := readCompressedSize(buf, 0)
			if err != nil {
				t.Fatalf("readCompressedSize failed: %v", err)
			}

			if decoded != size {
				t.Errorf("Size mismatch: expected %d, got %d", size, decoded)
			}

			if bytesRead != written {
				t.Errorf("Bytes read (%d) != bytes written (%d)", bytesRead, written)
			}
		})
	}
}

func TestReadCompressedSize_Errors(t *testing.T) {
	t.Run("empty_data", func(t *testing.T) {
		_, _, err := readCompressedSize([]byte{}, 0)
		if err == nil {
			t.Error("Expected error for empty data")
		}
	})

	t.Run("insufficient_data_2byte", func(t *testing.T) {
		_, _, err := readCompressedSize([]byte{0x01}, 0) // Needs 2 bytes
		if err == nil {
			t.Error("Expected error for insufficient data")
		}
	})

	t.Run("insufficient_data_4byte", func(t *testing.T) {
		_, _, err := readCompressedSize([]byte{0x02, 0x00, 0x00}, 0) // Needs 4 bytes
		if err == nil {
			t.Error("Expected error for insufficient data")
		}
	})

	t.Run("insufficient_data_8byte", func(t *testing.T) {
		_, _, err := readCompressedSize([]byte{0x03, 0x00, 0x00, 0x00, 0x00}, 0) // Needs 8 bytes
		if err == nil {
			t.Error("Expected error for insufficient data")
		}
	})

	t.Run("offset_out_of_bounds", func(t *testing.T) {
		_, _, err := readCompressedSize([]byte{0x00}, 10) // Offset beyond data
		if err == nil {
			t.Error("Expected error for offset out of bounds")
		}
	})
}

func TestReadCompressedSize_WithOffset(t *testing.T) {
	// Create buffer with size=5 encoded at offset 2
	buf := make([]byte, 10)
	buf[2] = byte((5 << 2) | 0) // 1-byte encoding: size=5

	size, bytesRead, err := readCompressedSize(buf, 2)
	if err != nil {
		t.Fatalf("readCompressedSize failed: %v", err)
	}

	if size != 5 {
		t.Errorf("Expected size 5, got %d", size)
	}

	if bytesRead != 1 {
		t.Errorf("Expected 1 byte read, got %d", bytesRead)
	}
}

// ============================================================================
// Tests for writeCompressedSize (improve existing coverage)
// ============================================================================

func TestWriteCompressedSize_AllSizes(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int // Expected byte count
	}{
		{"zero", 0, 1},
		{"1byte_max", 63, 1},
		{"2byte_min", 64, 2},
		{"2byte_max", 16383, 2},
		{"4byte_min", 16384, 4},
		{"4byte_max", 1073741823, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]byte, 8)
			written := writeCompressedSize(buf, tt.size)

			if written != tt.expected {
				t.Errorf("Expected %d bytes written, got %d", tt.expected, written)
			}

			// Verify round-trip
			size, _, err := readCompressedSize(buf, 0)
			if err != nil {
				t.Fatalf("readCompressedSize failed: %v", err)
			}

			if size != tt.size {
				t.Errorf("Round-trip failed: expected %d, got %d", tt.size, size)
			}
		})
	}
}

func TestWriteCompressedSize_RoundTrip(t *testing.T) {
	// Test various sizes for round-trip correctness
	sizes := []int{0, 1, 63, 64, 100, 16383, 16384, 1000000, 1073741823}

	for _, size := range sizes {
		buf := make([]byte, 8)
		written := writeCompressedSize(buf, size)

		if written == 0 {
			t.Fatalf("writeCompressedSize(%d) wrote 0 bytes", size)
		}

		decoded, _, err := readCompressedSize(buf, 0)
		if err != nil {
			t.Fatalf("readCompressedSize failed for size %d: %v", size, err)
		}

		if decoded != size {
			t.Errorf("Round-trip failed for %d: got %d", size, decoded)
		}
	}
}
