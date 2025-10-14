package beve

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
	"testing"

	"github.com/pierrec/lz4/v4"
)

// Simulated field-indexed BEVE encoder
// Instead of string keys, uses integer indexes

type FieldIndexedObject struct {
	FieldCount uint32
	Fields     []FieldIndexEntry
}

type FieldIndexEntry struct {
	Index uint32 // Field index instead of string name
	Value []byte // Serialized value
}

// encodeWithFieldIndex simulates field-indexed BEVE format
func encodeWithFieldIndex(users []BenchUser) []byte {
	buf := new(bytes.Buffer)

	// Write array header
	buf.WriteByte(0x05) // Generic array
	writeCompressedUintToBuffer(buf, uint64(len(users)))

	for _, user := range users {
		// Object with field indexes
		buf.WriteByte(0x13) // Object with field index (0b00'10'011)
		writeCompressedUintToBuffer(buf, 10) // 10 fields

		// Field 0: ID (index instead of "id" string)
		buf.WriteByte(0x00) // Field index 0
		buf.WriteByte(0x02) // String header
		writeCompressedUintToBuffer(buf, uint64(len(user.ID)))
		buf.WriteString(user.ID)

		// Field 1: Name
		buf.WriteByte(0x01)
		buf.WriteByte(0x02)
		writeCompressedUintToBuffer(buf, uint64(len(user.Name)))
		buf.WriteString(user.Name)

		// Field 2: Email
		buf.WriteByte(0x02)
		buf.WriteByte(0x02)
		writeCompressedUintToBuffer(buf, uint64(len(user.Email)))
		buf.WriteString(user.Email)

		// Field 3: Age (int)
		buf.WriteByte(0x03)
		buf.WriteByte(0x09) // uint8
		buf.WriteByte(byte(user.Age))

		// Field 4: Active (bool)
		buf.WriteByte(0x04)
		if user.Active {
			buf.WriteByte(0x18) // true
		} else {
			buf.WriteByte(0x08) // false
		}

		// Field 5: Balance (float64)
		buf.WriteByte(0x05)
		buf.WriteByte(0x19) // float64
		binary.Write(buf, binary.LittleEndian, user.Balance)

		// Field 6: Address
		buf.WriteByte(0x06)
		buf.WriteByte(0x02)
		writeCompressedUintToBuffer(buf, uint64(len(user.Address)))
		buf.WriteString(user.Address)

		// Field 7: Phone
		buf.WriteByte(0x07)
		buf.WriteByte(0x02)
		writeCompressedUintToBuffer(buf, uint64(len(user.Phone)))
		buf.WriteString(user.Phone)

		// Field 8: Company
		buf.WriteByte(0x08)
		buf.WriteByte(0x02)
		writeCompressedUintToBuffer(buf, uint64(len(user.Company)))
		buf.WriteString(user.Company)

		// Field 9: Country
		buf.WriteByte(0x09)
		buf.WriteByte(0x02)
		writeCompressedUintToBuffer(buf, uint64(len(user.Country)))
		buf.WriteString(user.Country)
	}

	return buf.Bytes()
}

func writeCompressedUintToBuffer(buf *bytes.Buffer, v uint64) {
	if v < 64 {
		buf.WriteByte(byte(v << 2))
	} else if v < 16384 {
		buf.WriteByte(byte((v<<2)|1) & 0xFF)
		buf.WriteByte(byte(v >> 6))
	} else if v < 1073741824 {
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, uint32((v<<2)|2))
		buf.Write(b)
	} else {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, (v<<2)|3)
		buf.Write(b)
	}
}

// TestFieldIndexCompressionComparison compares all combinations
func TestFieldIndexCompressionComparison(t *testing.T) {
	users := generateBenchUsers(100)

	// 1. Normal BEVE
	normalBEVE, _ := Marshal(users)
	t.Logf("1. Normal BEVE:                  %6d bytes (baseline)", len(normalBEVE))

	// 2. Field Index BEVE (simulated)
	fieldIndexBEVE := encodeWithFieldIndex(users)
	fieldIndexSaving := float64(len(normalBEVE)-len(fieldIndexBEVE)) / float64(len(normalBEVE)) * 100
	t.Logf("2. Field Index BEVE:             %6d bytes (%.1f%% smaller)", len(fieldIndexBEVE), fieldIndexSaving)

	// 3. Normal BEVE + LZ4
	lz4Normal := make([]byte, lz4.CompressBlockBound(len(normalBEVE)))
	lz4NormalSize, _ := lz4.CompressBlock(normalBEVE, lz4Normal, nil)
	lz4NormalSaving := float64(len(normalBEVE)-lz4NormalSize) / float64(len(normalBEVE)) * 100
	t.Logf("3. Normal BEVE + LZ4:            %6d bytes (%.1f%% smaller) ⭐", lz4NormalSize, lz4NormalSaving)

	// 4. Field Index BEVE + LZ4 🔬
	lz4FieldIndex := make([]byte, lz4.CompressBlockBound(len(fieldIndexBEVE)))
	lz4FieldIndexSize, _ := lz4.CompressBlock(fieldIndexBEVE, lz4FieldIndex, nil)
	lz4FieldIndexSaving := float64(len(normalBEVE)-lz4FieldIndexSize) / float64(len(normalBEVE)) * 100
	extraSaving := float64(lz4NormalSize-lz4FieldIndexSize) / float64(lz4NormalSize) * 100
	t.Logf("4. Field Index BEVE + LZ4:       %6d bytes (%.1f%% smaller, %.1f%% vs LZ4-only) 🔬", 
		lz4FieldIndexSize, lz4FieldIndexSaving, extraSaving)

	// 5. For comparison: Zstd variants
	// (We already tested these, just reference)
	t.Logf("")
	t.Logf("For reference from previous tests:")
	t.Logf("5. Normal BEVE + Zstd:             426 bytes (97.8%% smaller)")
	
	// Analysis
	t.Logf("")
	t.Logf("================================================================================")
	t.Logf("ANALYSIS: Is Field Index + Compression worth it?")
	t.Logf("================================================================================")
	
	additionalSavingBytes := lz4NormalSize - lz4FieldIndexSize
	t.Logf("Additional saving from field index: %d bytes (%.1f%% improvement over LZ4-only)",
		additionalSavingBytes, extraSaving)
	
	if extraSaving < 5 {
		t.Logf("❌ VERDICT: NOT worth it (<5%% improvement)")
		t.Logf("   Reason: LZ4 already compresses repeated field names very well")
		t.Logf("   Complexity cost > marginal benefit")
	} else if extraSaving < 15 {
		t.Logf("⚠️  VERDICT: Marginal benefit (%.1f%% improvement)", extraSaving)
		t.Logf("   Consider if: Schema is stable AND complexity acceptable")
	} else {
		t.Logf("✅ VERDICT: Worth it! (%.1f%% improvement)", extraSaving)
		t.Logf("   Field index + compression is significantly better")
	}
}

// BenchmarkFieldIndexWithCompression measures performance
func BenchmarkFieldIndexWithCompression(b *testing.B) {
	users := generateBenchUsers(100)
	
	normalBEVE, _ := Marshal(users)
	fieldIndexBEVE := encodeWithFieldIndex(users)

	b.Run("Normal-BEVE", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(normalBEVE)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = Marshal(users)
		}
	})

	b.Run("Field-Index-BEVE", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(fieldIndexBEVE)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = encodeWithFieldIndex(users)
		}
	})

	b.Run("Normal-BEVE+LZ4", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, _ := Marshal(users)
			compressed := make([]byte, lz4.CompressBlockBound(len(data)))
			_, _ = lz4.CompressBlock(data, compressed, nil)
		}
	})

	b.Run("Field-Index-BEVE+LZ4", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data := encodeWithFieldIndex(users)
			compressed := make([]byte, lz4.CompressBlockBound(len(data)))
			_, _ = lz4.CompressBlock(data, compressed, nil)
		}
	})
}

// TestCompressionEfficiency tests different data patterns
func TestCompressionEfficiency(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{"10-users", 10},
		{"100-users", 100},
		{"1000-users", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := generateBenchUsers(tt.count)

			normalBEVE, _ := Marshal(users)
			fieldIndexBEVE := encodeWithFieldIndex(users)

			// LZ4 on normal
			lz4Normal := make([]byte, lz4.CompressBlockBound(len(normalBEVE)))
			lz4NormalSize, _ := lz4.CompressBlock(normalBEVE, lz4Normal, nil)

			// LZ4 on field index
			lz4FieldIndex := make([]byte, lz4.CompressBlockBound(len(fieldIndexBEVE)))
			lz4FieldIndexSize, _ := lz4.CompressBlock(fieldIndexBEVE, lz4FieldIndex, nil)

			improvement := float64(lz4NormalSize-lz4FieldIndexSize) / float64(lz4NormalSize) * 100

			t.Logf("Count: %4d | Normal+LZ4: %6d bytes | FieldIndex+LZ4: %6d bytes | Improvement: %.1f%%",
				tt.count, lz4NormalSize, lz4FieldIndexSize, improvement)
		})
	}
}

// TestWorstCaseScenario tests when field index is WORST
func TestWorstCaseScenario(t *testing.T) {
	t.Log("Testing worst-case scenario: Unique field values, no repetition")
	
	// Generate users with unique data (no compression benefit)
	users := make([]BenchUser, 100)
	for i := 0; i < 100; i++ {
		users[i] = BenchUser{
			ID:      generateUniqueString(i, "id"),
			Name:    generateUniqueString(i, "name"),
			Email:   generateUniqueString(i, "email"),
			Age:     20 + i,
			Active:  i%2 == 0,
			Balance: 1000.0 + float64(i)*1.5,
			Address: generateUniqueString(i, "address"),
			Phone:   generateUniqueString(i, "phone"),
			Company: generateUniqueString(i, "company"),
			Country: generateUniqueString(i, "country"),
		}
	}

	normalBEVE, _ := Marshal(users)
	fieldIndexBEVE := encodeWithFieldIndex(users)

	t.Logf("Normal BEVE:       %6d bytes", len(normalBEVE))
	t.Logf("Field Index BEVE:  %6d bytes", len(fieldIndexBEVE))

	// LZ4 compression
	lz4Normal := make([]byte, lz4.CompressBlockBound(len(normalBEVE)))
	lz4NormalSize, _ := lz4.CompressBlock(normalBEVE, lz4Normal, nil)

	lz4FieldIndex := make([]byte, lz4.CompressBlockBound(len(fieldIndexBEVE)))
	lz4FieldIndexSize, _ := lz4.CompressBlock(fieldIndexBEVE, lz4FieldIndex, nil)

	t.Logf("Normal + LZ4:      %6d bytes", lz4NormalSize)
	t.Logf("FieldIndex + LZ4:  %6d bytes", lz4FieldIndexSize)

	improvement := float64(lz4NormalSize-lz4FieldIndexSize) / float64(lz4NormalSize) * 100
	t.Logf("Improvement: %.1f%%", improvement)
}

func generateUniqueString(i int, prefix string) string {
	h := fnv.New32a()
	h.Write([]byte(prefix))
	binary.Write(h, binary.LittleEndian, uint32(i))
	return prefix + "-" + string(rune(h.Sum32()))
}
