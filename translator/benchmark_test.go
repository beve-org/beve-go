package translator

import (
	"encoding/json"
	"testing"

	"github.com/beve-org/beve-go"
)

// Benchmark data structures
var (
	smallJSON = []byte(`{"id":123,"name":"John","active":true}`)

	mediumJSON = []byte(`{
		"id": 12345,
		"username": "john_doe",
		"email": "john@example.com",
		"profile": {
			"firstName": "John",
			"lastName": "Doe",
			"age": 30,
			"address": {
				"street": "123 Main St",
				"city": "New York",
				"country": "USA"
			}
		},
		"preferences": {
			"theme": "dark",
			"notifications": true,
			"language": "en"
		},
		"tags": ["developer", "golang", "beve"]
	}`)

	// Pre-encoded BEVE versions
	smallBEVE, _  = FromJSON(smallJSON)
	mediumBEVE, _ = FromJSON(mediumJSON)
)

// BenchmarkFromJSON_Small tests JSON→BEVE conversion for small payloads
func BenchmarkFromJSON_Small(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(smallJSON)))

	for i := 0; i < b.N; i++ {
		_, err := FromJSON(smallJSON)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFromJSON_Medium tests JSON→BEVE conversion for medium payloads
func BenchmarkFromJSON_Medium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(mediumJSON)))

	for i := 0; i < b.N; i++ {
		_, err := FromJSON(mediumJSON)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkToJSON_Small tests BEVE→JSON conversion for small payloads
func BenchmarkToJSON_Small(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(smallBEVE)))

	for i := 0; i < b.N; i++ {
		_, err := ToJSON(smallBEVE)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkToJSON_Medium tests BEVE→JSON conversion for medium payloads
func BenchmarkToJSON_Medium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(mediumBEVE)))

	for i := 0; i < b.N; i++ {
		_, err := ToJSON(mediumBEVE)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFromJSONString tests string wrapper overhead
func BenchmarkFromJSONString(b *testing.B) {
	jsonStr := string(smallJSON)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := FromJSONString(jsonStr)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkToJSONString tests string wrapper overhead
func BenchmarkToJSONString(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := ToJSONString(smallBEVE)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkRoundTrip_JSONtoBEVEtoJSON tests full round-trip performance
func BenchmarkRoundTrip_JSONtoBEVEtoJSON(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		beveData, err := FromJSON(mediumJSON)
		if err != nil {
			b.Fatal(err)
		}

		_, err = ToJSON(beveData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkValidateJSON tests JSON validation performance
func BenchmarkValidateJSON(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = ValidateJSON(mediumJSON)
	}
}

// BenchmarkValidateBEVE tests BEVE validation performance
func BenchmarkValidateBEVE(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = ValidateBEVE(mediumBEVE)
	}
}

// BenchmarkFromJSONWithStats tests stats collection overhead
func BenchmarkFromJSONWithStats(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _, err := FromJSONWithStats(mediumJSON)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Comparison benchmarks against standard library

// BenchmarkStdlib_JSONUnmarshalMarshal compares against json.Unmarshal + json.Marshal
func BenchmarkStdlib_JSONUnmarshalMarshal(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var data interface{}
		if err := json.Unmarshal(mediumJSON, &data); err != nil {
			b.Fatal(err)
		}

		_, err := json.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkStdlib_BEVEUnmarshalMarshal compares against beve.Unmarshal + beve.Marshal
func BenchmarkStdlib_BEVEUnmarshalMarshal(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var data interface{}
		if err := beve.Unmarshal(mediumBEVE, &data); err != nil {
			b.Fatal(err)
		}

		_, err := beve.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkTranslator_vs_Stdlib compares translator performance
func BenchmarkTranslator_vs_Stdlib(b *testing.B) {
	b.Run("Translator_FromJSON", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, err := FromJSON(mediumJSON)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Stdlib_Unmarshal+Marshal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var data interface{}
			if err := json.Unmarshal(mediumJSON, &data); err != nil {
				b.Fatal(err)
			}
			_, err := beve.Marshal(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkArrayConversion tests array-specific performance
func BenchmarkArrayConversion(b *testing.B) {
	arrays := [][]byte{
		[]byte(`[1,2,3,4,5]`),
		[]byte(`["a","b","c","d","e"]`),
		[]byte(`[true,false,true,false,true]`),
		[]byte(`[1,"two",3.0,true,null]`),
	}

	for i, arr := range arrays {
		b.Run(string(rune('A'+i)), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				beveData, err := FromJSON(arr)
				if err != nil {
					b.Fatal(err)
				}
				_, err = ToJSON(beveData)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkObjectConversion tests object-specific performance
func BenchmarkObjectConversion(b *testing.B) {
	objects := [][]byte{
		[]byte(`{}`),
		[]byte(`{"a":1}`),
		[]byte(`{"a":1,"b":2,"c":3}`),
		[]byte(`{"nested":{"deep":{"value":42}}}`),
	}

	for i, obj := range objects {
		b.Run(string(rune('A'+i)), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				beveData, err := FromJSON(obj)
				if err != nil {
					b.Fatal(err)
				}
				_, err = ToJSON(beveData)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkLargePayload tests performance with larger data
func BenchmarkLargePayload(b *testing.B) {
	// Generate a large JSON payload
	data := make(map[string]interface{})
	for i := 0; i < 1000; i++ {
		data[string(rune('a'+i%26))] = map[string]interface{}{
			"id":    i,
			"value": float64(i) * 3.14,
			"text":  "Lorem ipsum dolor sit amet consectetur adipiscing elit",
		}
	}

	largeJSON, _ := json.Marshal(data)

	b.Run("FromJSON", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(largeJSON)))

		for i := 0; i < b.N; i++ {
			_, err := FromJSON(largeJSON)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	largeBEVE, _ := FromJSON(largeJSON)

	b.Run("ToJSON", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(largeBEVE)))

		for i := 0; i < b.N; i++ {
			_, err := ToJSON(largeBEVE)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkMemoryEfficiency tests allocation patterns
func BenchmarkMemoryEfficiency(b *testing.B) {
	b.Run("Small_Allocations", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			beveData, _ := FromJSON(smallJSON)
			ToJSON(beveData)
		}
	})

	b.Run("Medium_Allocations", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			beveData, _ := FromJSON(mediumJSON)
			ToJSON(beveData)
		}
	})
}
