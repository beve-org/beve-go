// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Tests and benchmarks for Buffer.WriteByte assembly implementation.

package core

import (
	"testing"
)

// TestBufferWriteByte_Assembly tests correctness of assembly implementation
func TestBufferWriteByte_Assembly(t *testing.T) {
	t.Run("empty buffer", func(t *testing.T) {
		buf := AcquireBuffer(0)
		defer ReleaseBuffer(buf)

		if err := buf.WriteByte(0x42); err != nil {
			t.Fatalf("WriteByte failed: %v", err)
		}

		if buf.Len() != 1 {
			t.Errorf("Length mismatch: got %d, want 1", buf.Len())
		}

		if buf.Bytes()[0] != 0x42 {
			t.Errorf("Byte mismatch: got 0x%02X, want 0x42", buf.Bytes()[0])
		}
	})

	t.Run("multiple writes", func(t *testing.T) {
		buf := AcquireBuffer(0)
		defer ReleaseBuffer(buf)

		expected := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
		for _, b := range expected {
			if err := buf.WriteByte(b); err != nil {
				t.Fatalf("WriteByte failed: %v", err)
			}
		}

		if buf.Len() != len(expected) {
			t.Errorf("Length mismatch: got %d, want %d", buf.Len(), len(expected))
		}

		result := buf.Bytes()
		for i, want := range expected {
			if result[i] != want {
				t.Errorf("Byte %d mismatch: got 0x%02X, want 0x%02X", i, result[i], want)
			}
		}
	})

	t.Run("force growth", func(t *testing.T) {
		// Start with small buffer to force growth
		buf := AcquireBuffer(2)
		defer ReleaseBuffer(buf)

		initialCap := buf.Cap()

		// Write beyond initial capacity
		for i := 0; i < initialCap+10; i++ {
			if err := buf.WriteByte(byte(i)); err != nil {
				t.Fatalf("WriteByte failed at index %d: %v", i, err)
			}
		}

		if buf.Len() != initialCap+10 {
			t.Errorf("Length mismatch after growth: got %d, want %d", buf.Len(), initialCap+10)
		}

		// Verify all bytes
		result := buf.Bytes()
		for i := 0; i < initialCap+10; i++ {
			if result[i] != byte(i) {
				t.Errorf("Byte %d mismatch after growth: got 0x%02X, want 0x%02X", i, result[i], byte(i))
			}
		}
	})

	t.Run("stress test", func(t *testing.T) {
		buf := AcquireBuffer(0)
		defer ReleaseBuffer(buf)

		// Write 10,000 bytes
		for i := 0; i < 10000; i++ {
			if err := buf.WriteByte(byte(i % 256)); err != nil {
				t.Fatalf("WriteByte failed at index %d: %v", i, err)
			}
		}

		if buf.Len() != 10000 {
			t.Errorf("Length mismatch: got %d, want 10000", buf.Len())
		}

		// Spot check
		result := buf.Bytes()
		for i := 0; i < 10000; i += 100 {
			expected := byte(i % 256)
			if result[i] != expected {
				t.Errorf("Byte %d mismatch: got 0x%02X, want 0x%02X", i, result[i], expected)
			}
		}
	})
}

// BenchmarkBufferWriteByte_Assembly benchmarks assembly implementation
func BenchmarkBufferWriteByte_Assembly(b *testing.B) {
	buf := AcquireBuffer(1024) // Pre-allocate to test fast path
	defer ReleaseBuffer(buf)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		for j := 0; j < 100; j++ {
			_ = buf.WriteByte(byte(j))
		}
	}
}

// BenchmarkBufferWriteByte_FastPath benchmarks pure fast path (no growth)
func BenchmarkBufferWriteByte_FastPath(b *testing.B) {
	buf := AcquireBuffer(10000) // Large buffer, never grows
	defer ReleaseBuffer(buf)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = buf.WriteByte(0x42)
	}
}

// BenchmarkBufferWriteByte_SlowPath benchmarks with growth
func BenchmarkBufferWriteByte_SlowPath(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf := AcquireBuffer(0) // Start empty, will grow
		for j := 0; j < 10; j++ {
			_ = buf.WriteByte(byte(j))
		}
		ReleaseBuffer(buf)
	}
}

// BenchmarkBufferWriteByte_SingleWrite compares single WriteByte call
func BenchmarkBufferWriteByte_SingleWrite(b *testing.B) {
	b.Run("Assembly", func(b *testing.B) {
		buf := AcquireBuffer(1024)
		defer ReleaseBuffer(buf)

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			buf.Reset()
			_ = buf.WriteByte(0x42)
		}
	})

	b.Run("PureGo_Struct", func(b *testing.B) {
		type PureGoBuffer struct {
			data []byte
		}
		buf := &PureGoBuffer{data: make([]byte, 0, 1024)}

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			buf.data = buf.data[:0]
			buf.data = append(buf.data, 0x42)
		}
	})
}
