package beve

import (
	"reflect"

	"github.com/beve-org/beve-go/core"
)

// decoder wraps core.Decoder for backward compatibility
type decoder struct {
	core *core.Decoder
}

// newDecoder creates a new decoder
func newDecoder(data []byte) decoder {
	return decoder{
		core: core.NewDecoder(data),
	}
}

// decode decodes BEVE data into a reflect.Value
func (d decoder) decode(v reflect.Value) error {
	dec := d.core
	defer core.PutDecoderToPool(dec)
	return dec.Decode(v)
}
