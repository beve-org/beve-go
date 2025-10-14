package core

import (
	"reflect"
	"testing"
)

// TestEncoderPoolBufferReset validates that encoder pool properly resets buffers.
//
// This test prevents regression of the critical bug where PutEncoderToPool()
// was not resetting enc.Buf.data, causing stale data contamination.
//
// Bug history: Fixed in commit e26943b (14 Oct 2025)
func TestEncoderPoolBufferReset(t *testing.T) {
	// Step 1: Encode string array
	enc1 := GetEncoderFromPool()
	strings := []string{"hello", "world", "test"}
	enc1.Encode(reflect.ValueOf(strings))
	
	// Capture first encoding
	firstData := make([]byte, len(enc1.Buf.Bytes()))
	copy(firstData, enc1.Buf.Bytes())
	
	// Return to pool (MUST reset buffer)
	PutEncoderToPool(enc1)
	
	// Step 2: Get encoder from pool again
	enc2 := GetEncoderFromPool()
	defer PutEncoderToPool(enc2)
	
	// CRITICAL CHECK: Buffer should be empty (length 0)
	if len(enc2.Buf.data) != 0 {
		t.Errorf("CRITICAL BUG: Encoder buffer not reset! Found %d bytes of stale data", len(enc2.Buf.data))
		t.Errorf("Stale data: %v", enc2.Buf.data[:min(20, len(enc2.Buf.data))])
		t.Fatal("Encoder pool contamination detected - this causes data corruption!")
	}
	
	// Step 3: Encode different data (int32 array)
	ints := []int32{1, 2, 3, 4, 5}
	enc2.Encode(reflect.ValueOf(ints))
	
	secondData := enc2.Buf.Bytes()
	
	// Step 4: Verify no contamination from first encoding
	// The second encoding should NOT contain any bytes from the first
	if len(secondData) > len(firstData) {
		// Check if first data is present in second
		for i := 0; i <= len(secondData)-len(firstData); i++ {
			match := true
			for j := 0; j < len(firstData); j++ {
				if secondData[i+j] != firstData[j] {
					match = false
					break
				}
			}
			if match {
				t.Fatal("CRITICAL BUG: First encoding data found in second encoding - buffer not reset!")
			}
		}
	}
	
	t.Log("✓ Encoder pool buffer reset working correctly")
}

// TestEncoderPoolCrossContamination tests the original bug scenario.
//
// This reproduces the exact failure mode that caused:
// panic: reflect: call of reflect.Value.SetString on int32 Value
func TestEncoderPoolCrossContamination(t *testing.T) {
	// Simulate benchmark scenario that discovered the bug
	
	// Encode string array
	enc1 := GetEncoderFromPool()
	strings := []string{"a", "b", "c", "d", "e"}
	enc1.Encode(reflect.ValueOf(strings))
	encoded1 := make([]byte, len(enc1.Buf.Bytes()))
	copy(encoded1, enc1.Buf.Bytes())
	PutEncoderToPool(enc1)
	
	// Get same encoder from pool
	enc2 := GetEncoderFromPool()
	
	// CRITICAL: Buffer must be empty before encoding
	bufLenBeforeEncode := len(enc2.Buf.data)
	if bufLenBeforeEncode != 0 {
		t.Fatalf("CRITICAL: Encoder buffer has %d bytes before encoding (expected 0)", bufLenBeforeEncode)
	}
	
	// Encode int32 array
	ints := []int32{1, 2, 3, 4, 5}
	enc2.Encode(reflect.ValueOf(ints))
	encoded2 := enc2.Buf.Bytes()
	PutEncoderToPool(enc2)
	
	// Decode both to verify correctness
	dec1 := NewDecoder(encoded1)
	var result1 []string
	if err := dec1.Decode(reflect.ValueOf(&result1).Elem()); err != nil {
		t.Fatalf("Failed to decode strings: %v", err)
	}
	
	dec2 := NewDecoder(encoded2)
	var result2 []int32
	if err := dec2.Decode(reflect.ValueOf(&result2).Elem()); err != nil {
		t.Fatalf("Failed to decode ints (CROSS-CONTAMINATION BUG?): %v", err)
	}
	
	// Verify decoded values
	if len(result1) != 5 || result1[0] != "a" {
		t.Errorf("String decode corrupted: got %v", result1)
	}
	if len(result2) != 5 || result2[0] != 1 {
		t.Errorf("Int32 decode corrupted: got %v", result2)
	}
	
	t.Log("✓ No cross-contamination between different data types")
}

// TestEncoderPoolConcurrentUsage tests pool safety under concurrent access.
func TestEncoderPoolConcurrentUsage(t *testing.T) {
	const goroutines = 10
	const iterations = 100
	
	done := make(chan bool, goroutines)
	
	for g := 0; g < goroutines; g++ {
		go func(id int) {
			for i := 0; i < iterations; i++ {
				enc := GetEncoderFromPool()
				
				// Check buffer is empty
				if len(enc.Buf.data) != 0 {
					t.Errorf("Goroutine %d iteration %d: buffer not empty (%d bytes)", id, i, len(enc.Buf.data))
				}
				
				// Encode some data
				data := []int32{int32(id), int32(i)}
				enc.Encode(reflect.ValueOf(data))
				
				// Return to pool
				PutEncoderToPool(enc)
			}
			done <- true
		}(g)
	}
	
	// Wait for all goroutines
	for g := 0; g < goroutines; g++ {
		<-done
	}
	
	t.Logf("✓ %d goroutines × %d iterations = %d pool operations (no contamination)", 
		goroutines, iterations, goroutines*iterations)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
