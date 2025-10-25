package translatornative

import (
	"sync"
)

// ByteBuffer is a pooled buffer for zero-allocation parsing and serialization.
type ByteBuffer struct {
	data []byte
	pos  int
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &ByteBuffer{
			data: make([]byte, 0, 256*1024), // 4KB default capacity
		}
	},
}

// GetBuffer gets a buffer from the pool.
func GetBuffer() *ByteBuffer {
	buf := bufferPool.Get().(*ByteBuffer)
	buf.Reset()
	return buf
}

// PutBuffer returns a buffer to the pool.
func PutBuffer(buf *ByteBuffer) {
	if cap(buf.data) > 1024*1024 { // Don't pool buffers > 1MB
		return
	}
	bufferPool.Put(buf)
}

// Reset resets the buffer for reuse.
func (b *ByteBuffer) Reset() {
	b.data = b.data[:0]
	b.pos = 0
}

// WriteByte appends a byte.
func (b *ByteBuffer) WriteByte(c byte) error {
	b.data = append(b.data, c)
	return nil
}

// WriteString appends a string.
func (b *ByteBuffer) WriteString(s string) {
	b.data = append(b.data, s...)
}

// WriteBytes appends bytes.
func (b *ByteBuffer) WriteBytes(p []byte) {
	b.data = append(b.data, p...)
}

// Bytes returns the buffer data.
func (b *ByteBuffer) Bytes() []byte {
	return b.data
}

// Len returns the buffer length.
func (b *ByteBuffer) Len() int {
	return len(b.data)
}

// Grow grows the buffer to guarantee space for n more bytes.
func (b *ByteBuffer) Grow(n int) {
	if cap(b.data)-len(b.data) < n {
		newCap := cap(b.data) * 2
		if newCap < len(b.data)+n {
			newCap = len(b.data) + n
		}
		newData := make([]byte, len(b.data), newCap)
		copy(newData, b.data)
		b.data = newData
	}
}

// String pool for common strings to reduce allocations
var stringPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 256)
	},
}

// GetStringBuffer gets a string buffer from pool.
func GetStringBuffer() []byte {
	buf := stringPool.Get().([]byte)
	return buf[:0]
}

// PutStringBuffer returns a string buffer to pool.
func PutStringBuffer(buf []byte) {
	if cap(buf) > 4096 { // Don't pool large buffers
		return
	}
	stringPool.Put(buf)
}
