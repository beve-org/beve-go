package core

import (
	"testing"
)

func TestEncodeUintAsmDirect(t *testing.T) {
	tests := []struct {
		value      uint64
		wantLen    int
		wantHeader byte
	}{
		{0, 2, 0x11},          // 1 byte + header
		{255, 2, 0x11},        // 1 byte + header
		{256, 3, 0x31},        // 2 bytes + header
		{65535, 3, 0x31},      // 2 bytes + header
		{65536, 5, 0x51},      // 4 bytes + header
		{4294967295, 5, 0x51}, // 4 bytes + header
		{4294967296, 9, 0x71}, // 8 bytes + header
		{^uint64(0), 9, 0x71}, // MaxUint64: 8 bytes + header
	}

	for _, tt := range tests {
		var scratch [9]byte
		n := encodeUintAsm(&scratch, tt.value)

		if n != tt.wantLen {
			t.Errorf("encodeUintAsm(%d / 0x%x) = %d bytes, want %d bytes [got header=0x%02x]",
				tt.value, tt.value, n, tt.wantLen, scratch[0])
		}
		if scratch[0] != tt.wantHeader {
			t.Errorf("encodeUintAsm(%d / 0x%x) header = 0x%02x, want 0x%02x",
				tt.value, tt.value, scratch[0], tt.wantHeader)
		}
	}
}
