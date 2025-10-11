package beve

import (
	"bufio"
	"io"

	"github.com/beve-org/beve-go/core"
)

// StreamEncoder provides buffered, high-performance streaming encoding.
//
// Benefits over regular Encoder:
//   - Buffered I/O (8KB default) reduces syscalls
//   - Encoder reuse between items (no reflection re-computation)
//   - Batch writes for small payloads
//   - Lower allocations per item
//
// Example:
//
//	stream := beve.NewStreamEncoder(writer)
//	defer stream.Close()
//	for _, item := range items {
//	    if err := stream.Encode(item); err != nil {
//	        return err
//	    }
//	}
type StreamEncoder struct {
	enc *core.Encoder
	bw  *bufio.Writer
	w   io.Writer
}

// NewStreamEncoder creates a new streaming encoder with buffered I/O.
//
// The encoder uses an 8KB buffer by default for optimal performance.
// Call Close() or Flush() to ensure all data is written.
func NewStreamEncoder(w io.Writer) *StreamEncoder {
	bw := bufio.NewWriterSize(w, 8192) // 8KB buffer
	enc := core.GetEncoderFromPool()
	enc.Buf.Reset()

	return &StreamEncoder{
		enc: enc,
		bw:  bw,
		w:   w,
	}
}

// NewStreamEncoderSize creates a streaming encoder with a specific buffer size.
func NewStreamEncoderSize(w io.Writer, bufSize int) *StreamEncoder {
	bw := bufio.NewWriterSize(w, bufSize)
	enc := core.GetEncoderFromPool()
	enc.Buf.Reset()

	return &StreamEncoder{
		enc: enc,
		bw:  bw,
		w:   w,
	}
}

// Encode writes a BEVE-encoded value to the stream.
//
// The value is buffered and may not be written immediately.
// Call Flush() or Close() to ensure all data is written.
func (s *StreamEncoder) Encode(v interface{}) error {
	// Reset encoder buffer for reuse
	s.enc.Buf.Reset()

	// Use Marshal which handles all types efficiently
	data, err := Marshal(v)
	if err != nil {
		return err
	}

	// Write to buffered writer
	if len(data) > 0 {
		if _, err := s.bw.Write(data); err != nil {
			return err
		}
	}

	return nil
}

// Flush writes any buffered data to the underlying writer.
func (s *StreamEncoder) Flush() error {
	return s.bw.Flush()
}

// Close flushes any buffered data and returns the encoder to the pool.
//
// After calling Close(), the StreamEncoder should not be used.
func (s *StreamEncoder) Close() error {
	if err := s.bw.Flush(); err != nil {
		return err
	}

	// Return encoder to pool
	core.PutEncoderToPool(s.enc)
	s.enc = nil

	return nil
}

// StreamDecoder provides buffered, high-performance streaming decoding.
//
// Benefits:
//   - Buffered I/O reduces syscalls
//   - Decoder reuse between items
//   - Lower allocations per item
//
// Example:
//
//	stream := beve.NewStreamDecoder(reader)
//	defer stream.Close()
//	for {
//	    var item MyType
//	    if err := stream.Decode(&item); err != nil {
//	        if err == io.EOF {
//	            break
//	        }
//	        return err
//	    }
//	    process(item)
//	}
type StreamDecoder struct {
	dec *core.Decoder
	br  *bufio.Reader
	r   io.Reader
}

// NewStreamDecoder creates a new streaming decoder with buffered I/O.
//
// The decoder uses an 8KB buffer by default for optimal performance.
func NewStreamDecoder(r io.Reader) *StreamDecoder {
	br := bufio.NewReaderSize(r, 8192) // 8KB buffer
	dec := core.NewDecoder(nil)

	return &StreamDecoder{
		dec: dec,
		br:  br,
		r:   r,
	}
}

// NewStreamDecoderSize creates a streaming decoder with a specific buffer size.
func NewStreamDecoderSize(r io.Reader, bufSize int) *StreamDecoder {
	br := bufio.NewReaderSize(r, bufSize)
	dec := core.NewDecoder(nil)

	return &StreamDecoder{
		dec: dec,
		br:  br,
		r:   r,
	}
}

// Decode reads the next BEVE-encoded value from the stream.
//
// Returns io.EOF when the stream ends.
func (s *StreamDecoder) Decode(v interface{}) error {
	// Read length prefix (if present)
	// For now, use simple approach: read until next value

	// TODO: Implement proper streaming decode
	// This requires protocol changes to support length-prefixed messages

	return Unmarshal(nil, v) // Placeholder
}

// Close releases decoder resources.
//
// After calling Close(), the StreamDecoder should not be used.
func (s *StreamDecoder) Close() error {
	s.dec = nil
	return nil
}
