package beve

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"
)

// Test data structures for I/O performance comparison
type IOTestData struct {
	ID          int               `json:"id" msgpack:"id" cbor:"id"`
	Name        string            `json:"name" msgpack:"name" cbor:"name"`
	Email       string            `json:"email" msgpack:"email" cbor:"email"`
	Age         int               `json:"age" msgpack:"age" cbor:"age"`
	Active      bool              `json:"active" msgpack:"active" cbor:"active"`
	Score       float64           `json:"score" msgpack:"score" cbor:"score"`
	Tags        []string          `json:"tags" msgpack:"tags" cbor:"tags"`
	Metadata    map[string]string `json:"metadata" msgpack:"metadata" cbor:"metadata"`
	Description string            `json:"description" msgpack:"description" cbor:"description"`
}

type LargeIOTestData struct {
	Users    []IOTestData      `json:"users" msgpack:"users" cbor:"users"`
	Settings map[string]string `json:"settings" msgpack:"settings" cbor:"settings"`
	Count    int               `json:"count" msgpack:"count" cbor:"count"`
	Total    float64           `json:"total" msgpack:"total" cbor:"total"`
}

// Helper functions to create test data
func createIOTestData() IOTestData {
	return IOTestData{
		ID:          12345,
		Name:        "John Doe",
		Email:       "john.doe@example.com",
		Age:         30,
		Active:      true,
		Score:       95.5,
		Tags:        []string{"developer", "golang", "backend", "senior"},
		Metadata:    map[string]string{"city": "San Francisco", "country": "USA", "timezone": "PST"},
		Description: "Experienced software engineer with 10+ years in backend development",
	}
}

func createLargeIOTestData(userCount int) LargeIOTestData {
	users := make([]IOTestData, userCount)
	for i := 0; i < userCount; i++ {
		users[i] = IOTestData{
			ID:          i,
			Name:        "User " + string(rune(i)),
			Email:       "user" + string(rune(i)) + "@example.com",
			Age:         20 + (i % 50),
			Active:      i%2 == 0,
			Score:       float64(50 + (i % 50)),
			Tags:        []string{"tag1", "tag2", "tag3"},
			Metadata:    map[string]string{"key1": "value1", "key2": "value2"},
			Description: "User description " + string(rune(i)),
		}
	}

	return LargeIOTestData{
		Users: users,
		Settings: map[string]string{
			"theme":    "dark",
			"language": "en",
			"timezone": "UTC",
		},
		Count: userCount,
		Total: float64(userCount * 100),
	}
}

// ==============================================================================
// WRITE Performance Tests (Encoding to Writer)
// ==============================================================================

// BenchmarkIOWrite_BEVE_Small tests BEVE write performance with small data
func BenchmarkIOWrite_BEVE_Small(b *testing.B) {
	data := createIOTestData()
	buf := &bytes.Buffer{}
	enc := NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		_, err := enc.Encode(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

// BenchmarkIOWrite_JSON_Small tests JSON write performance with small data
func BenchmarkIOWrite_JSON_Small(b *testing.B) {
	data := createIOTestData()
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

// BenchmarkIOWrite_Sonic_Small tests Sonic write performance with small data
func BenchmarkIOWrite_Sonic_Small(b *testing.B) {
	data := createIOTestData()
	buf := &bytes.Buffer{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc := sonic.ConfigDefault.NewEncoder(buf)
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

// BenchmarkIOWrite_MessagePack_Small tests MessagePack write performance
func BenchmarkIOWrite_MessagePack_Small(b *testing.B) {
	data := createIOTestData()
	buf := &bytes.Buffer{}
	enc := msgpack.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

// BenchmarkIOWrite_CBOR_Small tests CBOR write performance
func BenchmarkIOWrite_CBOR_Small(b *testing.B) {
	data := createIOTestData()
	buf := &bytes.Buffer{}
	enc := cbor.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

// ==============================================================================
// WRITE Performance Tests - Medium Data (100 users)
// ==============================================================================

func BenchmarkIOWrite_BEVE_Medium(b *testing.B) {
	data := createLargeIOTestData(100)
	buf := &bytes.Buffer{}
	enc := NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		_, err := enc.Encode(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

func BenchmarkIOWrite_JSON_Medium(b *testing.B) {
	data := createLargeIOTestData(100)
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

func BenchmarkIOWrite_Sonic_Medium(b *testing.B) {
	data := createLargeIOTestData(100)
	buf := &bytes.Buffer{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc := sonic.ConfigDefault.NewEncoder(buf)
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

func BenchmarkIOWrite_MessagePack_Medium(b *testing.B) {
	data := createLargeIOTestData(100)
	buf := &bytes.Buffer{}
	enc := msgpack.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

func BenchmarkIOWrite_CBOR_Medium(b *testing.B) {
	data := createLargeIOTestData(100)
	buf := &bytes.Buffer{}
	enc := cbor.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

// ==============================================================================
// WRITE Performance Tests - Large Data (1000 users)
// ==============================================================================

func BenchmarkIOWrite_BEVE_Large(b *testing.B) {
	data := createLargeIOTestData(1000)
	buf := &bytes.Buffer{}
	enc := NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		_, err := enc.Encode(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

func BenchmarkIOWrite_JSON_Large(b *testing.B) {
	data := createLargeIOTestData(1000)
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

func BenchmarkIOWrite_Sonic_Large(b *testing.B) {
	data := createLargeIOTestData(1000)
	buf := &bytes.Buffer{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc := sonic.ConfigDefault.NewEncoder(buf)
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

func BenchmarkIOWrite_MessagePack_Large(b *testing.B) {
	data := createLargeIOTestData(1000)
	buf := &bytes.Buffer{}
	enc := msgpack.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

func BenchmarkIOWrite_CBOR_Large(b *testing.B) {
	data := createLargeIOTestData(1000)
	buf := &bytes.Buffer{}
	enc := cbor.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(buf.Len()))
}

// ==============================================================================
// READ Performance Tests (Decoding from Reader)
// ==============================================================================

// BenchmarkIORead_BEVE_Small tests BEVE read performance with small data
func BenchmarkIORead_BEVE_Small(b *testing.B) {
	original := createIOTestData()
	data, _ := Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := NewDecoder(data)
		var result IOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

// BenchmarkIORead_JSON_Small tests JSON read performance with small data
func BenchmarkIORead_JSON_Small(b *testing.B) {
	original := createIOTestData()
	data, _ := json.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := json.NewDecoder(reader)
		var result IOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

// BenchmarkIORead_Sonic_Small tests Sonic read performance with small data
func BenchmarkIORead_Sonic_Small(b *testing.B) {
	original := createIOTestData()
	data, _ := sonic.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := sonic.ConfigDefault.NewDecoder(reader)
		var result IOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

// BenchmarkIORead_MessagePack_Small tests MessagePack read performance
func BenchmarkIORead_MessagePack_Small(b *testing.B) {
	original := createIOTestData()
	data, _ := msgpack.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := msgpack.NewDecoder(reader)
		var result IOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

// BenchmarkIORead_CBOR_Small tests CBOR read performance
func BenchmarkIORead_CBOR_Small(b *testing.B) {
	original := createIOTestData()
	data, _ := cbor.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := cbor.NewDecoder(reader)
		var result IOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

// ==============================================================================
// READ Performance Tests - Medium Data (100 users)
// ==============================================================================

func BenchmarkIORead_BEVE_Medium(b *testing.B) {
	original := createLargeIOTestData(100)
	data, _ := Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := NewDecoder(data)
		var result LargeIOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

func BenchmarkIORead_JSON_Medium(b *testing.B) {
	original := createLargeIOTestData(100)
	data, _ := json.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := json.NewDecoder(reader)
		var result LargeIOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

func BenchmarkIORead_Sonic_Medium(b *testing.B) {
	original := createLargeIOTestData(100)
	data, _ := sonic.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := sonic.ConfigDefault.NewDecoder(reader)
		var result LargeIOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

func BenchmarkIORead_MessagePack_Medium(b *testing.B) {
	original := createLargeIOTestData(100)
	data, _ := msgpack.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := msgpack.NewDecoder(reader)
		var result LargeIOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

func BenchmarkIORead_CBOR_Medium(b *testing.B) {
	original := createLargeIOTestData(100)
	data, _ := cbor.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := cbor.NewDecoder(reader)
		var result LargeIOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

// ==============================================================================
// READ Performance Tests - Large Data (1000 users)
// ==============================================================================

func BenchmarkIORead_BEVE_Large(b *testing.B) {
	original := createLargeIOTestData(1000)
	data, _ := Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := NewDecoder(data)
		var result LargeIOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

func BenchmarkIORead_JSON_Large(b *testing.B) {
	original := createLargeIOTestData(1000)
	data, _ := json.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := json.NewDecoder(reader)
		var result LargeIOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

func BenchmarkIORead_Sonic_Large(b *testing.B) {
	original := createLargeIOTestData(1000)
	data, _ := sonic.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := sonic.ConfigDefault.NewDecoder(reader)
		var result LargeIOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

func BenchmarkIORead_MessagePack_Large(b *testing.B) {
	original := createLargeIOTestData(1000)
	data, _ := msgpack.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := msgpack.NewDecoder(reader)
		var result LargeIOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

func BenchmarkIORead_CBOR_Large(b *testing.B) {
	original := createLargeIOTestData(1000)
	data, _ := cbor.Marshal(original)
	reader := bytes.NewReader(data)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		dec := cbor.NewDecoder(reader)
		var result LargeIOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(data)))
}

// ==============================================================================
// ROUND TRIP Performance Tests (Write + Read)
// ==============================================================================

func BenchmarkIORoundTrip_BEVE_Small(b *testing.B) {
	data := createIOTestData()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Write
		buf := &bytes.Buffer{}
		enc := NewEncoder(buf)
		_, err := enc.Encode(data)
		if err != nil {
			b.Fatal(err)
		}

		// Read
		dec := NewDecoder(buf.Bytes())
		var result IOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIORoundTrip_JSON_Small(b *testing.B) {
	data := createIOTestData()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Write
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}

		// Read
		dec := json.NewDecoder(buf)
		var result IOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIORoundTrip_MessagePack_Small(b *testing.B) {
	data := createIOTestData()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Write
		buf := &bytes.Buffer{}
		enc := msgpack.NewEncoder(buf)
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}

		// Read
		dec := msgpack.NewDecoder(buf)
		var result IOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIORoundTrip_CBOR_Small(b *testing.B) {
	data := createIOTestData()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Write
		buf := &bytes.Buffer{}
		enc := cbor.NewEncoder(buf)
		if err := enc.Encode(data); err != nil {
			b.Fatal(err)
		}

		// Read
		dec := cbor.NewDecoder(buf)
		var result IOTestData
		if err := dec.Decode(&result); err != nil {
			b.Fatal(err)
		}
	}
}

// ==============================================================================
// Multiple Sequential Writes Test (Batch Processing)
// ==============================================================================

func BenchmarkIOSequentialWrites_BEVE(b *testing.B) {
	data := createIOTestData()
	buf := &bytes.Buffer{}
	enc := NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			_, err := enc.Encode(data)
			if err != nil {
				b.Fatal(err)
			}
		}
		buf.Reset()
	}
}

func BenchmarkIOSequentialWrites_JSON(b *testing.B) {
	data := createIOTestData()
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			if err := enc.Encode(data); err != nil {
				b.Fatal(err)
			}
		}
		buf.Reset()
	}
}

func BenchmarkIOSequentialWrites_MessagePack(b *testing.B) {
	data := createIOTestData()
	buf := &bytes.Buffer{}
	enc := msgpack.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			if err := enc.Encode(data); err != nil {
				b.Fatal(err)
			}
		}
		buf.Reset()
	}
}

func BenchmarkIOSequentialWrites_CBOR(b *testing.B) {
	data := createIOTestData()
	buf := &bytes.Buffer{}
	enc := cbor.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			if err := enc.Encode(data); err != nil {
				b.Fatal(err)
			}
		}
		buf.Reset()
	}
}
