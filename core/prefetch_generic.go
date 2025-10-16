//go:build !arm64 && !amd64 && !purego
// +build !arm64,!amd64,!purego

package core

import "unsafe"

// prefetchRead is a no-op on unsupported platforms
func prefetchRead(addr unsafe.Pointer, len int) {
	// No-op: Generic fallback, no prefetching
	// Compiler may still optimize sequential access patterns
}
