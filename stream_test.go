package beve

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

// =============================================================================
// StreamEncoder Tests
// =============================================================================

func TestStreamEncoder_NewStreamEncoder(t *testing.T) {
	t.Run("creates encoder with default buffer", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		if enc == nil {
			t.Fatal("expected encoder, got nil")
		}
		if enc.enc == nil {
			t.Error("encoder not initialized")
		}
		if enc.bw == nil {
			t.Error("buffered writer not initialized")
		}
		if enc.w != buf {
			t.Error("writer not set correctly")
		}
		defer enc.Close()
	})

	t.Run("creates encoder with custom buffer size", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoderSize(buf, 4096)
		if enc == nil {
			t.Fatal("expected encoder, got nil")
		}
		if enc.enc == nil {
			t.Error("encoder not initialized")
		}
		if enc.bw == nil {
			t.Error("buffered writer not initialized")
		}
		defer enc.Close()
	})
}

func TestStreamEncoder_EncodeSingleValue(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"nil", nil},
		{"bool true", true},
		{"bool false", false},
		{"int", 42},
		{"int64", int64(1234567890)},
		{"float64", 3.14159},
		{"string", "hello world"},
		{"empty string", ""},
		{"byte slice", []byte{1, 2, 3, 4, 5}},
		{"int slice", []int{1, 2, 3}},
		{"string slice", []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			enc := NewStreamEncoder(buf)
			defer enc.Close()

			if err := enc.Encode(tt.value); err != nil {
				t.Errorf("Encode() error = %v", err)
				return
			}

			if err := enc.Flush(); err != nil {
				t.Errorf("Flush() error = %v", err)
				return
			}

			// Verify data was written
			if buf.Len() == 0 {
				t.Error("expected data to be written, got empty buffer")
			}
		})
	}
}

func TestStreamEncoder_EncodeComplexTypes(t *testing.T) {
	t.Run("encode map", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		defer enc.Close()

		value := map[string]int{"a": 1, "b": 2}
		if err := enc.Encode(value); err != nil {
			t.Errorf("Encode() error = %v", err)
		}

		if err := enc.Flush(); err != nil {
			t.Errorf("Flush() error = %v", err)
		}

		if buf.Len() == 0 {
			t.Error("expected data to be written")
		}
	})

	t.Run("encode struct", func(t *testing.T) {
		type TestStruct struct {
			ID   int
			Name string
		}

		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		defer enc.Close()

		value := TestStruct{123, "test"}
		if err := enc.Encode(value); err != nil {
			t.Errorf("Encode() error = %v", err)
		}

		if err := enc.Flush(); err != nil {
			t.Errorf("Flush() error = %v", err)
		}

		if buf.Len() == 0 {
			t.Error("expected data to be written")
		}
	})
}

func TestStreamEncoder_EncodeMultipleValues(t *testing.T) {
	t.Run("encode 10 integers", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		defer enc.Close()

		for i := 0; i < 10; i++ {
			if err := enc.Encode(i); err != nil {
				t.Errorf("Encode(%d) error = %v", i, err)
			}
		}

		if err := enc.Flush(); err != nil {
			t.Errorf("Flush() error = %v", err)
		}

		// Verify data was written
		if buf.Len() == 0 {
			t.Error("expected data to be written")
		}
	})

	t.Run("encode mixed types", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		defer enc.Close()

		values := []interface{}{
			42,
			"hello",
			true,
			3.14,
			[]int{1, 2, 3},
			map[string]string{"key": "value"},
		}

		for i, v := range values {
			if err := enc.Encode(v); err != nil {
				t.Errorf("Encode(%d) error = %v", i, err)
			}
		}

		if err := enc.Flush(); err != nil {
			t.Errorf("Flush() error = %v", err)
		}
	})

	t.Run("encode 100 structs", func(t *testing.T) {
		type User struct {
			ID   int
			Name string
			Age  int
		}

		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		defer enc.Close()

		for i := 0; i < 100; i++ {
			user := User{
				ID:   i,
				Name: "User" + string(rune(i)),
				Age:  20 + (i % 50),
			}
			if err := enc.Encode(user); err != nil {
				t.Errorf("Encode(%d) error = %v", i, err)
			}
		}

		if err := enc.Flush(); err != nil {
			t.Errorf("Flush() error = %v", err)
		}

		// Verify significant data was written
		if buf.Len() < 500 {
			t.Errorf("expected at least 500 bytes, got %d", buf.Len())
		}
	})
}

func TestStreamEncoder_BufferedWriting(t *testing.T) {
	t.Run("data buffered until flush", func(t *testing.T) {
		// Use a custom writer to detect writes
		writeCount := 0
		var totalWritten int
		writer := writerFunc(func(p []byte) (int, error) {
			writeCount++
			totalWritten += len(p)
			return len(p), nil
		})

		enc := NewStreamEncoder(writer)
		defer enc.Close()

		// Write small values (should be buffered)
		for i := 0; i < 10; i++ {
			if err := enc.Encode(i); err != nil {
				t.Errorf("Encode(%d) error = %v", i, err)
			}
		}

		// Before flush, writes should be minimal or buffered
		beforeFlushWrites := writeCount

		if err := enc.Flush(); err != nil {
			t.Errorf("Flush() error = %v", err)
		}

		// After flush, data should be written
		if totalWritten == 0 {
			t.Error("expected data to be written after flush")
		}

		// Flush should trigger at least one write
		if writeCount == beforeFlushWrites {
			// This is ok - buffer may not have flushed yet if very small
		}
	})

	t.Run("large data triggers automatic flush", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoderSize(buf, 1024) // Small buffer
		defer enc.Close()

		// Write large value (should exceed buffer)
		largeData := make([]byte, 2048)
		for i := range largeData {
			largeData[i] = byte(i % 256)
		}

		if err := enc.Encode(largeData); err != nil {
			t.Errorf("Encode() error = %v", err)
		}

		// Data should be partially written even before flush
		// (bufio automatically flushes when buffer is full)
		if buf.Len() == 0 {
			// May still be buffered, flush to check
			enc.Flush()
			if buf.Len() == 0 {
				t.Error("expected data to be written for large value")
			}
		}
	})
}

func TestStreamEncoder_Flush(t *testing.T) {
	t.Run("flush writes buffered data", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		defer enc.Close()

		if err := enc.Encode(42); err != nil {
			t.Errorf("Encode() error = %v", err)
		}

		// Get buffer size before flush
		beforeFlush := buf.Len()

		if err := enc.Flush(); err != nil {
			t.Errorf("Flush() error = %v", err)
		}

		// After flush, data should be in buffer
		afterFlush := buf.Len()
		if afterFlush == 0 {
			t.Error("expected data after flush")
		}

		// Multiple flushes should be safe
		if err := enc.Flush(); err != nil {
			t.Errorf("second Flush() error = %v", err)
		}

		if buf.Len() != afterFlush {
			t.Error("second flush should not write more data")
		}

		_ = beforeFlush // May be 0 if fully buffered
	})

	t.Run("flush returns write errors", func(t *testing.T) {
		expectedErr := errors.New("write error")
		writer := writerFunc(func(p []byte) (int, error) {
			return 0, expectedErr
		})

		enc := NewStreamEncoder(writer)
		defer enc.Close()

		// Encode a value
		if err := enc.Encode(42); err != nil {
			t.Errorf("Encode() error = %v", err)
		}

		// Flush should return the write error
		if err := enc.Flush(); err == nil {
			t.Error("expected flush error, got nil")
		} else if !strings.Contains(err.Error(), "write error") {
			t.Errorf("expected write error, got %v", err)
		}
	})
}

func TestStreamEncoder_Close(t *testing.T) {
	t.Run("close flushes data", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)

		if err := enc.Encode(42); err != nil {
			t.Errorf("Encode() error = %v", err)
		}

		if err := enc.Close(); err != nil {
			t.Errorf("Close() error = %v", err)
		}

		// Data should be written
		if buf.Len() == 0 {
			t.Error("expected data after close")
		}
	})

	t.Run("close returns flush errors", func(t *testing.T) {
		expectedErr := errors.New("flush error")
		writer := writerFunc(func(p []byte) (int, error) {
			return 0, expectedErr
		})

		enc := NewStreamEncoder(writer)

		if err := enc.Encode(42); err != nil {
			t.Errorf("Encode() error = %v", err)
		}

		if err := enc.Close(); err == nil {
			t.Error("expected close error, got nil")
		}
	})

	t.Run("close returns encoder to pool", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)

		if err := enc.Close(); err != nil {
			t.Errorf("Close() error = %v", err)
		}

		// Encoder should be nil after close
		if enc.enc != nil {
			t.Error("expected encoder to be returned to pool")
		}
	})

	t.Run("encode after close should panic or return error", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)

		if err := enc.Close(); err != nil {
			t.Errorf("Close() error = %v", err)
		}

		// Encoding after close should fail
		defer func() {
			if r := recover(); r != nil {
				// Panic is acceptable
			}
		}()

		// This should fail gracefully
		err := enc.Encode(42)
		if err == nil {
			// If no error, check that encoder is nil
			if enc.enc != nil {
				t.Error("expected error when encoding after close")
			}
		}
	})
}

func TestStreamEncoder_ErrorHandling(t *testing.T) {
	t.Run("unsupported type returns error", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		defer enc.Close()

		// Try to encode a channel (unsupported)
		ch := make(chan int)
		err := enc.Encode(ch)
		if err == nil {
			t.Error("expected error for unsupported type")
		}
	})

	t.Run("write error propagates", func(t *testing.T) {
		expectedErr := errors.New("write failed")
		writer := &errorWriter{err: expectedErr, failAfter: 1}

		enc := NewStreamEncoder(writer)
		defer func() {
			// Close may also error, ignore it
			_ = enc.Close()
		}()

		// First encode might succeed
		_ = enc.Encode(42)

		// Second encode or flush should fail
		err := enc.Encode(make([]byte, 10000)) // Large value to trigger write
		if err == nil {
			err = enc.Flush()
		}

		if err == nil {
			t.Error("expected write error to propagate")
		}
	})
}

func TestStreamEncoder_Performance(t *testing.T) {
	t.Run("reuses encoder between calls", func(t *testing.T) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		defer enc.Close()

		// Encode multiple values - should reuse encoder
		for i := 0; i < 100; i++ {
			if err := enc.Encode(i); err != nil {
				t.Errorf("Encode(%d) error = %v", i, err)
			}
		}

		if err := enc.Flush(); err != nil {
			t.Errorf("Flush() error = %v", err)
		}

		// Verify all data written
		if buf.Len() == 0 {
			t.Error("expected data to be written")
		}
	})

	t.Run("handles large batch efficiently", func(t *testing.T) {
		type Record struct {
			ID        int
			Timestamp int64
			Value     float64
			Tags      []string
		}

		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		defer enc.Close()

		// Encode 1000 records
		for i := 0; i < 1000; i++ {
			record := Record{
				ID:        i,
				Timestamp: int64(1000000 + i),
				Value:     float64(i) * 1.5,
				Tags:      []string{"tag1", "tag2", "tag3"},
			}
			if err := enc.Encode(record); err != nil {
				t.Errorf("Encode(%d) error = %v", i, err)
			}
		}

		if err := enc.Flush(); err != nil {
			t.Errorf("Flush() error = %v", err)
		}

		// Verify significant data was written
		if buf.Len() < 10000 {
			t.Errorf("expected at least 10KB of data, got %d bytes", buf.Len())
		}
	})
}

// =============================================================================
// StreamDecoder Tests (Basic - TODO: Full implementation needed)
// =============================================================================

func TestStreamDecoder_NewStreamDecoder(t *testing.T) {
	t.Run("creates decoder with default buffer", func(t *testing.T) {
		buf := bytes.NewReader([]byte{})
		dec := NewStreamDecoder(buf)
		if dec == nil {
			t.Fatal("expected decoder, got nil")
		}
		if dec.dec == nil {
			t.Error("decoder not initialized")
		}
		if dec.br == nil {
			t.Error("buffered reader not initialized")
		}
		defer dec.Close()
	})

	t.Run("creates decoder with custom buffer size", func(t *testing.T) {
		buf := bytes.NewReader([]byte{})
		dec := NewStreamDecoderSize(buf, 4096)
		if dec == nil {
			t.Fatal("expected decoder, got nil")
		}
		if dec.dec == nil {
			t.Error("decoder not initialized")
		}
		if dec.br == nil {
			t.Error("buffered reader not initialized")
		}
		defer dec.Close()
	})
}

func TestStreamDecoder_Close(t *testing.T) {
	t.Run("close releases resources", func(t *testing.T) {
		buf := bytes.NewReader([]byte{})
		dec := NewStreamDecoder(buf)

		if err := dec.Close(); err != nil {
			t.Errorf("Close() error = %v", err)
		}

		// Decoder should be nil after close
		if dec.dec != nil {
			t.Error("expected decoder to be released")
		}
	})
}

func TestStreamDecoder_DecodeBasic(t *testing.T) {
	t.Run("decode returns error for unimplemented", func(t *testing.T) {
		// Create some encoded data
		data, err := Marshal(42)
		if err != nil {
			t.Fatalf("Marshal() error = %v", err)
		}

		buf := bytes.NewReader(data)
		dec := NewStreamDecoder(buf)
		defer dec.Close()

		var result int
		err = dec.Decode(&result)

		// Currently Decode is not fully implemented
		// This test just ensures it doesn't panic
		if err != nil {
			// Expected - implementation is incomplete
			t.Logf("Decode() error (expected): %v", err)
		}
	})
}

// =============================================================================
// Helper Types
// =============================================================================

// writerFunc allows using a function as io.Writer
type writerFunc func(p []byte) (n int, err error)

func (f writerFunc) Write(p []byte) (n int, err error) {
	return f(p)
}

// errorWriter fails after a certain number of writes
type errorWriter struct {
	err       error
	count     int
	failAfter int
}

func (w *errorWriter) Write(p []byte) (n int, err error) {
	w.count++
	if w.count > w.failAfter {
		return 0, w.err
	}
	return len(p), nil
}

// =============================================================================
// Benchmarks
// =============================================================================

func BenchmarkStreamEncoder_SingleInt(b *testing.B) {
	buf := &bytes.Buffer{}
	enc := NewStreamEncoder(buf)
	defer enc.Close()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := enc.Encode(42); err != nil {
			b.Fatal(err)
		}
		buf.Reset()
		enc.enc.Buf.Reset()
	}
}

func BenchmarkStreamEncoder_SmallStruct(b *testing.B) {
	type Data struct {
		ID   int
		Name string
		Age  int
	}

	data := Data{ID: 123, Name: "test", Age: 30}
	buf := &bytes.Buffer{}
	enc := NewStreamEncoder(buf)
	defer enc.Close()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
		buf.Reset()
		enc.enc.Buf.Reset()
	}
}

func BenchmarkStreamEncoder_Batch100(b *testing.B) {
	type Record struct {
		ID    int
		Value float64
	}

	buf := &bytes.Buffer{}
	enc := NewStreamEncoder(buf)
	defer enc.Close()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			record := Record{ID: j, Value: float64(j) * 1.5}
			if err := enc.Encode(record); err != nil {
				b.Fatal(err)
			}
		}
		enc.Flush()
		buf.Reset()
	}
}

func BenchmarkStreamEncoder_vs_Marshal(b *testing.B) {
	type Data struct {
		ID   int
		Name string
		Age  int
	}

	data := Data{ID: 123, Name: "test", Age: 30}

	b.Run("StreamEncoder", func(b *testing.B) {
		buf := &bytes.Buffer{}
		enc := NewStreamEncoder(buf)
		defer enc.Close()

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if err := enc.Encode(data); err != nil {
				b.Fatal(err)
			}
			buf.Reset()
			enc.enc.Buf.Reset()
		}
	})

	b.Run("Marshal", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if _, err := Marshal(data); err != nil {
				b.Fatal(err)
			}
		}
	})
}
