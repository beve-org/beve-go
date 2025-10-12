package beve

// RawMessage is a raw encoded BEVE value.
// It can be used with Marshal and Unmarshal to delay decoding or
// precompute BEVE payloads.
type RawMessage []byte

// MarshalBEVE returns m as the raw BEVE payload.
func (m RawMessage) MarshalBEVE() ([]byte, error) {
	if m == nil {
		return []byte{0x00}, nil // encode as null when nil
	}
	buf := make([]byte, len(m))
	copy(buf, m)
	return buf, nil
}

// UnmarshalBEVE stores the raw BEVE payload.
func (m *RawMessage) UnmarshalBEVE(data []byte) error {
	if data == nil {
		*m = nil
		return nil
	}
	buf := make([]byte, len(data))
	copy(buf, data)
	*m = RawMessage(buf)
	return nil
}
