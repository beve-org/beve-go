package translatornative

import (
	"strconv"
	"testing"
)

// BenchmarkDirectEncoder_Small: Small JSON object
func BenchmarkDirectEncoder_Small(b *testing.B) {
	data := []byte(`{"id":123,"name":"Alice","age":30,"active":true}`)
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc := NewDirectEncoder(data)
		_, err := enc.Encode()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkDirectEncoder_Medium: Medium JSON array
func BenchmarkDirectEncoder_Medium(b *testing.B) {
	data := []byte(`{
		"users": [
			{"id":1,"name":"Alice","age":30,"email":"alice@example.com","active":true},
			{"id":2,"name":"Bob","age":25,"email":"bob@example.com","active":false},
			{"id":3,"name":"Charlie","age":35,"email":"charlie@example.com","active":true}
		],
		"total": 3,
		"page": 1
	}`)
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc := NewDirectEncoder(data)
		_, err := enc.Encode()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkDirectEncoder_Large: Large JSON array (100 objects)
func BenchmarkDirectEncoder_Large(b *testing.B) {
	// Generate 100 user objects
	json := `{"users":[`
	for i := 0; i < 100; i++ {
		if i > 0 {
			json += `,`
		}
		json += `{"id":` + strconv.Itoa(i) + `,"name":"User` + strconv.Itoa(i) + `","age":` + strconv.Itoa(i%50) + `,"email":"user` + strconv.Itoa(i) + `@example.com","active":true,"score":` + strconv.Itoa(i%100) + `}`
	}
	json += `],"total":100,"page":1}`
	data := []byte(json)

	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc := NewDirectEncoder(data)
		_, err := enc.Encode()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFromJSON_Small: Full API with small payload
func BenchmarkFromJSON_Small(b *testing.B) {
	data := []byte(`{"id":123,"name":"Alice","age":30,"active":true}`)
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := FromJSON(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFromJSON_Medium: Full API with medium payload
func BenchmarkFromJSON_Medium(b *testing.B) {
	data := []byte(`{
		"users": [
			{"id":1,"name":"Alice","age":30,"email":"alice@example.com","active":true},
			{"id":2,"name":"Bob","age":25,"email":"bob@example.com","active":false},
			{"id":3,"name":"Charlie","age":35,"email":"charlie@example.com","active":true}
		],
		"total": 3,
		"page": 1
	}`)
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := FromJSON(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFromJSON_Large: Full API with large payload
func BenchmarkFromJSON_Large(b *testing.B) {
	// Generate 100 user objects
	json := `{"users":[`
	for i := 0; i < 100; i++ {
		if i > 0 {
			json += `,`
		}
		json += `{"id":` + strconv.Itoa(i) + `,"name":"User` + strconv.Itoa(i) + `","age":` + strconv.Itoa(i%50) + `,"email":"user` + strconv.Itoa(i) + `@example.com","active":true,"score":` + strconv.Itoa(i%100) + `}`
	}
	json += `],"total":100,"page":1}`
	data := []byte(json)

	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := FromJSON(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}
