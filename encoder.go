package beve

import (
	"github.com/beve-org/beve-go/core"
)

// encoder is an alias for core.Encoder
// This provides backward compatibility during the migration to modular architecture
type encoder = core.Encoder

// Buffer is an alias for core.Buffer
// This provides backward compatibility during the migration to modular architecture
type Buffer = core.Buffer

// Wrapper functions for backward compatibility
// These delegate directly to core package functions

// getEncoderFromPool acquires a pooled encoder
func getEncoderFromPool() *encoder {
	return core.GetEncoderFromPool()
}

// putEncoderToPool returns an encoder to the pool for reuse
func putEncoderToPool(enc *encoder) {
	core.PutEncoderToPool(enc)
}
