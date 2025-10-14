package core

import "testing"

func TestEncodeMapStringBoolHeaders(t *testing.T) {
	enc := &Encoder{Buf: AcquireBuffer(0)}
	if enc.Buf == nil {
		t.Fatal("expected buffer from AcquireBuffer")
	}
	defer ReleaseBuffer(enc.Buf)

	m := map[string]bool{"t": true, "f": false}
	if err := enc.encodeMapStringBool(m, len(m)); err != nil {
		t.Fatalf("encodeMapStringBool error: %v", err)
	}

	data := append([]byte(nil), enc.Buf.Bytes()...)

	dec := &Decoder{Data: data}
	header, err := dec.ReadByte()
	if err != nil {
		t.Fatalf("read header: %v", err)
	}
	if header != 0x03 {
		t.Fatalf("unexpected map header 0x%02X", header)
	}

	size, err := dec.ReadCompressedUint()
	if err != nil {
		t.Fatalf("read map size: %v", err)
	}
	if size != uint64(len(m)) {
		t.Fatalf("unexpected map size %d", size)
	}

	expected := map[string]byte{"t": 0x18, "f": 0x08}
	seen := make(map[string]byte)

	for i := uint64(0); i < size; i++ {
		keyLen, err := dec.ReadCompressedUint()
		if err != nil {
			t.Fatalf("read key length: %v", err)
		}

		keyBytes, err := dec.ReadBytes(int(keyLen))
		if err != nil {
			t.Fatalf("read key bytes: %v", err)
		}

		valueHeader, err := dec.ReadByte()
		if err != nil {
			t.Fatalf("read value header: %v", err)
		}

		seen[string(keyBytes)] = valueHeader
	}

	if len(seen) != len(expected) {
		t.Fatalf("unexpected number of entries: %d", len(seen))
	}

	for k, want := range expected {
		got, ok := seen[k]
		if !ok {
			t.Fatalf("missing key %q", k)
		}
		if got != want {
			t.Fatalf("key %q expected header 0x%02X, got 0x%02X", k, want, got)
		}
	}

	if dec.Pos != len(data) {
		t.Fatalf("decoder did not consume all bytes: pos=%d len=%d", dec.Pos, len(data))
	}
}
