//go:build !go1.21
// +build !go1.21

package core

// UseLockFreePool is always false for Go versions < 1.21
// Lock-free pooling requires runtime.procPin/Unpin which is available in Go 1.21+
const UseLockFreePool = false

// GetLockFreePoolStats returns zeros for unsupported Go versions
func GetLockFreePoolStats() (hits, misses, puts, discards, overflows uint64) {
	return 0, 0, 0, 0, 0
}

// ResetLockFreePoolStats is a no-op for unsupported Go versions
func ResetLockFreePoolStats() {
	// No-op
}

// getEncoderFromLockFreePool is a no-op stub for unsupported Go versions
// This should never be called due to UseLockFreePool=false
func getEncoderFromLockFreePool() *Encoder {
	panic("lock-free pool requires Go 1.21+")
}

// putEncoderToLockFreePool is a no-op stub for unsupported Go versions
// This should never be called due to UseLockFreePool=false
func putEncoderToLockFreePool(enc *Encoder) {
	panic("lock-free pool requires Go 1.21+")
}
