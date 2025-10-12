package beve

import (
	"sync"
)

// byteSlicePool manages reusable byte slices to reduce allocations
var byteSlicePool = sync.Pool{
	New: func() interface{} {
		// Start with a reasonable default size (256 bytes)
		// This covers most small structs without excessive memory use
		b := make([]byte, 0, 256)
		return &b
	},
}

// getByteSlice retrieves a byte slice from the pool
func getByteSlice() *[]byte {
	return byteSlicePool.Get().(*[]byte)
}

// growSlice ensures the slice has enough capacity and returns a slice of the desired length
func growSlice(b *[]byte, n int) []byte {
	if cap(*b) < n {
		// Grow by 2x or to exactly n, whichever is larger
		newCap := cap(*b) * 2
		if newCap < n {
			newCap = n
		}
		newSlice := make([]byte, n, newCap)
		copy(newSlice, *b)
		*b = newSlice
		return newSlice
	}
	*b = (*b)[:n]
	return *b
}
