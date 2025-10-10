package beve

import (
	"unsafe"
)

// stringToBytes converts string to []byte without allocation.
// SAFETY: The returned slice must not be modified and should not outlive the string.
// This is safe in our use case because we immediately write the data.
//
//go:inline
func stringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// bytesToString converts []byte to string without allocation.
// SAFETY: The input slice must not be modified after conversion.
//
//go:inline
func bytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b), len(b))
}
