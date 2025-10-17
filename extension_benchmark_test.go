package beve

import (
	"testing"
	"time"
)

// ============================================================================
// Benchmark: Typed Nested Arrays
// ============================================================================

func BenchmarkTypedNestedArray2D(b *testing.B) {
	data := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	b.Run("marshal", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = EncodeTypedNestedArray(data)
		}
	})

	encoded, _ := EncodeTypedNestedArray(data)
	b.Run("unmarshal", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = DecodeTypedNestedArray(encoded)
		}
	})
}

func BenchmarkTypedNestedArray3D(b *testing.B) {
	data := [][][]float32{
		{{1.1, 2.2}, {3.3, 4.4}},
		{{5.5, 6.6}, {7.7, 8.8}},
		{{9.9, 10.1}, {11.2, 12.3}},
	}

	b.Run("marshal", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = EncodeTypedNestedArray(data)
		}
	})

	encoded, _ := EncodeTypedNestedArray(data)
	b.Run("unmarshal", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = DecodeTypedNestedArray(encoded)
		}
	})
}

// ============================================================================
// Benchmark: RegExp
// ============================================================================

func BenchmarkRegExpMarshal(b *testing.B) {
	tests := []struct {
		name    string
		pattern string
		flags   byte
	}{
		{"simple", "^test$", 0},
		{"complex", "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", 0x01},
		{"unicode", "\\p{L}+\\s*\\p{N}+", 0x10},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = EncodeRegExp(tt.pattern, tt.flags)
			}
		})
	}
}

func BenchmarkRegExpUnmarshal(b *testing.B) {
	pattern := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	data, _ := EncodeRegExp(pattern, 0x01)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecodeRegExp(data)
	}
}

// ============================================================================
// Benchmark: Duration
// ============================================================================

func BenchmarkDurationMarshal(b *testing.B) {
	duration := 5*time.Hour + 30*time.Minute + 15*time.Second

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncodeDuration(duration)
	}
}

func BenchmarkDurationUnmarshal(b *testing.B) {
	duration := 5*time.Hour + 30*time.Minute + 15*time.Second
	data, _ := EncodeDuration(duration)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecodeDuration(data)
	}
}

// ============================================================================
// Benchmark: Interval
// ============================================================================

func BenchmarkIntervalMarshal(b *testing.B) {
	start := time.Date(2025, 10, 17, 10, 0, 0, 0, time.UTC)
	end := time.Date(2025, 10, 17, 11, 30, 0, 0, time.UTC)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncodeInterval(start, end)
	}
}

func BenchmarkIntervalUnmarshal(b *testing.B) {
	start := time.Date(2025, 10, 17, 10, 0, 0, 0, time.UTC)
	end := time.Date(2025, 10, 17, 11, 30, 0, 0, time.UTC)
	data, _ := EncodeInterval(start, end)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = DecodeInterval(data)
	}
}

// ============================================================================
// Benchmark: Field Index with Large Objects
// ============================================================================

func BenchmarkFieldIndexLargeObject(b *testing.B) {
	// Create object with 100 fields
	obj := make(map[string]interface{})
	for i := 0; i < 100; i++ {
		key := string(rune('a'+i%26)) + string(rune('A'+i%26)) + string(rune('0'+i%10))
		obj[key] = i * 1000
	}

	encoded, _ := EncodeIndexedObject(obj)

	b.Run("encode", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = EncodeIndexedObject(obj)
		}
	})

	b.Run("decode", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = DecodeIndexedObject(encoded)
		}
	})

	b.Run("read_first_field", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = ReadFieldByName(encoded, "aA0")
		}
	})

	b.Run("read_middle_field", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = ReadFieldByName(encoded, "zZ9") // Near middle
		}
	})

	b.Run("read_last_field", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = ReadFieldByName(encoded, "dD9") // Near end
		}
	})
}

// ============================================================================
// Benchmark: Typed Array vs Standard (Size Comparison)
// ============================================================================

func BenchmarkTypedVsStandardSize(b *testing.B) {
	type User struct {
		Name  string  `beve:"name"`
		Age   int     `beve:"age"`
		Email string  `beve:"email"`
		Score float64 `beve:"score"`
	}

	users := make([]User, 100)
	for i := 0; i < 100; i++ {
		users[i] = User{
			Name:  "User" + string(rune('0'+i%10)),
			Age:   20 + i%50,
			Email: "user" + string(rune('0'+i%10)) + "@example.com",
			Score: float64(i) * 1.5,
		}
	}

	b.Run("typed_marshal", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = MarshalTyped(users)
		}
	})

	b.Run("standard_marshal", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = Marshal(users)
		}
	})

	typedData, _ := MarshalTyped(users)
	standardData, _ := Marshal(users)

	b.Logf("Typed size: %d bytes", len(typedData))
	b.Logf("Standard size: %d bytes", len(standardData))
	savings := float64(len(standardData)-len(typedData)) / float64(len(standardData)) * 100
	b.Logf("Savings: %.1f%%", savings)
}

// ============================================================================
// Benchmark: MarshalAuto Detection
// ============================================================================

func BenchmarkMarshalAutoDetection(b *testing.B) {
	type User struct {
		Name string `beve:"name"`
		Age  int    `beve:"age"`
	}

	users := []User{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	b.Run("auto_small", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = MarshalAuto(users[:1])
		}
	})

	b.Run("auto_medium", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = MarshalAuto(users)
		}
	})

	// Create large array
	largeUsers := make([]User, 100)
	for i := 0; i < 100; i++ {
		largeUsers[i] = User{Name: "User", Age: 20 + i}
	}

	b.Run("auto_large", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = MarshalAuto(largeUsers)
		}
	})
}

// ============================================================================
// Benchmark: UnmarshalAuto Detection
// ============================================================================

func BenchmarkUnmarshalAutoDetection(b *testing.B) {
	type User struct {
		Name string `beve:"name"`
		Age  int    `beve:"age"`
	}

	users := []User{{"Alice", 30}, {"Bob", 25}}

	// Prepare different encoded formats
	typedData, _ := MarshalTyped(users)
	standardData, _ := Marshal(users)

	var result []map[string]interface{}

	b.Run("auto_typed", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = UnmarshalAuto(typedData, &result)
		}
	})

	b.Run("auto_standard", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = UnmarshalAuto(standardData, &result)
		}
	})
}

// ============================================================================
// Benchmark: Extension Detection
// ============================================================================

func BenchmarkExtensionDetection(b *testing.B) {
	// Prepare various extension data
	typedData, _ := MarshalTyped([]int{1, 2, 3})
	timestampData, _ := MarshalTimestamp(time.Now())
	uuid := [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	uuidData, _ := MarshalUUID(uuid)
	standardData, _ := Marshal(map[string]int{"key": 42})

	b.Run("detect_typed_array", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = DetectEncoding(typedData)
		}
	})

	b.Run("detect_timestamp", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = DetectEncoding(timestampData)
		}
	})

	b.Run("detect_uuid", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = DetectEncoding(uuidData)
		}
	})

	b.Run("detect_standard", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = DetectEncoding(standardData)
		}
	})
}

// ============================================================================
// Benchmark: Timestamp Precision
// ============================================================================

func BenchmarkTimestampPrecision(b *testing.B) {
	now := time.Now()
	ts := TimestampFromTime(now)

	b.Run("encode_current_time", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = EncodeTimestamp(ts)
		}
	})

	b.Run("encode_epoch", func(b *testing.B) {
		epoch := TimestampFromTime(time.Unix(0, 0))
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = EncodeTimestamp(epoch)
		}
	})

	b.Run("encode_with_nanoseconds", func(b *testing.B) {
		t := TimestampFromTime(time.Unix(1697545200, 123456789))
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = EncodeTimestamp(t)
		}
	})
}

// ============================================================================
// Benchmark: UUID Operations
// ============================================================================

func BenchmarkUUIDOperations(b *testing.B) {
	uuid := [16]byte{
		0x6b, 0xa7, 0xb8, 0x10,
		0x9d, 0xad, 0x41, 0xd1,
		0x80, 0xb4, 0x00, 0xc0,
		0x4f, 0xd4, 0x30, 0xc8,
	}

	b.Run("marshal_binary", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = MarshalUUID(uuid)
		}
	})

	uuidStr := "6ba7b810-9dad-41d1-80b4-00c04fd430c8"
	b.Run("marshal_string", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = MarshalUUIDString(uuidStr)
		}
	})

	binaryData, _ := MarshalUUID(uuid)
	b.Run("unmarshal_binary", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = UnmarshalUUID(binaryData)
		}
	})

	stringData, _ := MarshalUUIDString(uuidStr)
	b.Run("unmarshal_string", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = UnmarshalUUIDString(stringData)
		}
	})
}
