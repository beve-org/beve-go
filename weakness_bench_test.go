package beve

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"
)

// This file benchmarks scenarios where BEVE might be slower than alternatives

// Scenario 1: Very deep nested structures (reflection overhead)
type DeepNested struct {
	Level1 *DeepNested1 `beve:"l1" json:"l1"`
	Value  string       `beve:"v" json:"v"`
}

type DeepNested1 struct {
	Level2 *DeepNested2 `beve:"l2" json:"l2"`
	Value  string       `beve:"v" json:"v"`
}

type DeepNested2 struct {
	Level3 *DeepNested3 `beve:"l3" json:"l3"`
	Value  string       `beve:"v" json:"v"`
}

type DeepNested3 struct {
	Level4 *DeepNested4 `beve:"l4" json:"l4"`
	Value  string       `beve:"v" json:"v"`
}

type DeepNested4 struct {
	Level5 *DeepNested5 `beve:"l5" json:"l5"`
	Value  string       `beve:"v" json:"v"`
}

type DeepNested5 struct {
	Level6 *DeepNested6 `beve:"l6" json:"l6"`
	Value  string       `beve:"v" json:"v"`
}

type DeepNested6 struct {
	Level7 *DeepNested7 `beve:"l7" json:"l7"`
	Value  string       `beve:"v" json:"v"`
}

type DeepNested7 struct {
	Level8 *DeepNested8 `beve:"l8" json:"l8"`
	Value  string       `beve:"v" json:"v"`
}

type DeepNested8 struct {
	Level9 *DeepNested9 `beve:"l9" json:"l9"`
	Value  string       `beve:"v" json:"v"`
}

type DeepNested9 struct {
	Level10 *DeepNested10 `beve:"l10" json:"l10"`
	Value   string        `beve:"v" json:"v"`
}

type DeepNested10 struct {
	Value string `beve:"v" json:"v"`
}

func createDeepNested() DeepNested {
	return DeepNested{
		Value: "root",
		Level1: &DeepNested1{
			Value: "l1",
			Level2: &DeepNested2{
				Value: "l2",
				Level3: &DeepNested3{
					Value: "l3",
					Level4: &DeepNested4{
						Value: "l4",
						Level5: &DeepNested5{
							Value: "l5",
							Level6: &DeepNested6{
								Value: "l6",
								Level7: &DeepNested7{
									Value: "l7",
									Level8: &DeepNested8{
										Value: "l8",
										Level9: &DeepNested9{
											Value: "l9",
											Level10: &DeepNested10{
												Value: "l10",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Scenario 2: Many small strings (varint overhead vs fixed-size)
type ManySmallStrings struct {
	S1  string `beve:"s1" json:"s1"`
	S2  string `beve:"s2" json:"s2"`
	S3  string `beve:"s3" json:"s3"`
	S4  string `beve:"s4" json:"s4"`
	S5  string `beve:"s5" json:"s5"`
	S6  string `beve:"s6" json:"s6"`
	S7  string `beve:"s7" json:"s7"`
	S8  string `beve:"s8" json:"s8"`
	S9  string `beve:"s9" json:"s9"`
	S10 string `beve:"s10" json:"s10"`
	S11 string `beve:"s11" json:"s11"`
	S12 string `beve:"s12" json:"s12"`
	S13 string `beve:"s13" json:"s13"`
	S14 string `beve:"s14" json:"s14"`
	S15 string `beve:"s15" json:"s15"`
	S16 string `beve:"s16" json:"s16"`
	S17 string `beve:"s17" json:"s17"`
	S18 string `beve:"s18" json:"s18"`
	S19 string `beve:"s19" json:"s19"`
	S20 string `beve:"s20" json:"s20"`
}

// Scenario 3: Map with many entries (O(n) map iteration)
func createLargeMap(size int) map[string]interface{} {
	m := make(map[string]interface{}, size)
	for i := 0; i < size; i++ {
		m[string(rune('a'+(i%26)))+string(rune('a'+((i/26)%26)))] = i
	}
	return m
}

// Scenario 4: Slice of interfaces (type checking overhead)
func createInterfaceSlice(size int) []interface{} {
	result := make([]interface{}, size)
	for i := 0; i < size; i++ {
		switch i % 5 {
		case 0:
			result[i] = i
		case 1:
			result[i] = float64(i) * 1.5
		case 2:
			result[i] = string(rune('a' + (i % 26)))
		case 3:
			result[i] = i%2 == 0
		case 4:
			result[i] = []int{i, i + 1, i + 2}
		}
	}
	return result
}

// Scenario 5: Struct with many fields (struct field enumeration overhead)
type WideStruct struct {
	F1, F2, F3, F4, F5, F6, F7, F8, F9, F10           int `beve:"f1,f2,f3,f4,f5,f6,f7,f8,f9,f10" json:"f1,f2,f3,f4,f5,f6,f7,f8,f9,f10"`
	F11, F12, F13, F14, F15, F16, F17, F18, F19, F20  int `beve:"f11,f12,f13,f14,f15,f16,f17,f18,f19,f20" json:"f11,f12,f13,f14,f15,f16,f17,f18,f19,f20"`
	F21, F22, F23, F24, F25, F26, F27, F28, F29, F30  int `beve:"f21,f22,f23,f24,f25,f26,f27,f28,f29,f30" json:"f21,f22,f23,f24,f25,f26,f27,f28,f29,f30"`
	F31, F32, F33, F34, F35, F36, F37, F38, F39, F40  int `beve:"f31,f32,f33,f34,f35,f36,f37,f38,f39,f40" json:"f31,f32,f33,f34,f35,f36,f37,f38,f39,f40"`
	F41, F42, F43, F44, F45, F46, F47, F48, F49, F50  int `beve:"f41,f42,f43,f44,f45,f46,f47,f48,f49,f50" json:"f41,f42,f43,f44,f45,f46,f47,f48,f49,f50"`
}

// =============================================================================
// BENCHMARK 1: Deep Nested Structures
// =============================================================================

func BenchmarkDeepNested_BEVE_Marshal(b *testing.B) {
	data := createDeepNested()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkDeepNested_JSON_Marshal(b *testing.B) {
	data := createDeepNested()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkDeepNested_Sonic_Marshal(b *testing.B) {
	data := createDeepNested()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonic.Marshal(data)
	}
}

func BenchmarkDeepNested_MessagePack_Marshal(b *testing.B) {
	data := createDeepNested()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(data)
	}
}

func BenchmarkDeepNested_CBOR_Marshal(b *testing.B) {
	data := createDeepNested()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cbor.Marshal(data)
	}
}

// =============================================================================
// BENCHMARK 2: Many Small Strings
// =============================================================================

func BenchmarkManySmallStrings_BEVE_Marshal(b *testing.B) {
	data := ManySmallStrings{
		S1: "a", S2: "b", S3: "c", S4: "d", S5: "e",
		S6: "f", S7: "g", S8: "h", S9: "i", S10: "j",
		S11: "k", S12: "l", S13: "m", S14: "n", S15: "o",
		S16: "p", S17: "q", S18: "r", S19: "s", S20: "t",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkManySmallStrings_JSON_Marshal(b *testing.B) {
	data := ManySmallStrings{
		S1: "a", S2: "b", S3: "c", S4: "d", S5: "e",
		S6: "f", S7: "g", S8: "h", S9: "i", S10: "j",
		S11: "k", S12: "l", S13: "m", S14: "n", S15: "o",
		S16: "p", S17: "q", S18: "r", S19: "s", S20: "t",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkManySmallStrings_Sonic_Marshal(b *testing.B) {
	data := ManySmallStrings{
		S1: "a", S2: "b", S3: "c", S4: "d", S5: "e",
		S6: "f", S7: "g", S8: "h", S9: "i", S10: "j",
		S11: "k", S12: "l", S13: "m", S14: "n", S15: "o",
		S16: "p", S17: "q", S18: "r", S19: "s", S20: "t",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonic.Marshal(data)
	}
}

func BenchmarkManySmallStrings_MessagePack_Marshal(b *testing.B) {
	data := ManySmallStrings{
		S1: "a", S2: "b", S3: "c", S4: "d", S5: "e",
		S6: "f", S7: "g", S8: "h", S9: "i", S10: "j",
		S11: "k", S12: "l", S13: "m", S14: "n", S15: "o",
		S16: "p", S17: "q", S18: "r", S19: "s", S20: "t",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(data)
	}
}

func BenchmarkManySmallStrings_CBOR_Marshal(b *testing.B) {
	data := ManySmallStrings{
		S1: "a", S2: "b", S3: "c", S4: "d", S5: "e",
		S6: "f", S7: "g", S8: "h", S9: "i", S10: "j",
		S11: "k", S12: "l", S13: "m", S14: "n", S15: "o",
		S16: "p", S17: "q", S18: "r", S19: "s", S20: "t",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cbor.Marshal(data)
	}
}

// =============================================================================
// BENCHMARK 3: Large Map
// =============================================================================

func BenchmarkLargeMap_BEVE_Marshal(b *testing.B) {
	data := createLargeMap(1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkLargeMap_JSON_Marshal(b *testing.B) {
	data := createLargeMap(1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkLargeMap_Sonic_Marshal(b *testing.B) {
	data := createLargeMap(1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonic.Marshal(data)
	}
}

func BenchmarkLargeMap_MessagePack_Marshal(b *testing.B) {
	data := createLargeMap(1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(data)
	}
}

func BenchmarkLargeMap_CBOR_Marshal(b *testing.B) {
	data := createLargeMap(1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cbor.Marshal(data)
	}
}

// =============================================================================
// BENCHMARK 4: Interface Slice (Type Switching Overhead)
// =============================================================================

func BenchmarkInterfaceSlice_BEVE_Marshal(b *testing.B) {
	data := createInterfaceSlice(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkInterfaceSlice_JSON_Marshal(b *testing.B) {
	data := createInterfaceSlice(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkInterfaceSlice_Sonic_Marshal(b *testing.B) {
	data := createInterfaceSlice(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonic.Marshal(data)
	}
}

func BenchmarkInterfaceSlice_MessagePack_Marshal(b *testing.B) {
	data := createInterfaceSlice(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(data)
	}
}

func BenchmarkInterfaceSlice_CBOR_Marshal(b *testing.B) {
	data := createInterfaceSlice(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cbor.Marshal(data)
	}
}

// =============================================================================
// BENCHMARK 5: Wide Struct (Many Fields)
// =============================================================================

func BenchmarkWideStruct_BEVE_Marshal(b *testing.B) {
	data := WideStruct{
		F1: 1, F2: 2, F3: 3, F4: 4, F5: 5, F6: 6, F7: 7, F8: 8, F9: 9, F10: 10,
		F11: 11, F12: 12, F13: 13, F14: 14, F15: 15, F16: 16, F17: 17, F18: 18, F19: 19, F20: 20,
		F21: 21, F22: 22, F23: 23, F24: 24, F25: 25, F26: 26, F27: 27, F28: 28, F29: 29, F30: 30,
		F31: 31, F32: 32, F33: 33, F34: 34, F35: 35, F36: 36, F37: 37, F38: 38, F39: 39, F40: 40,
		F41: 41, F42: 42, F43: 43, F44: 44, F45: 45, F46: 46, F47: 47, F48: 48, F49: 49, F50: 50,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkWideStruct_JSON_Marshal(b *testing.B) {
	data := WideStruct{
		F1: 1, F2: 2, F3: 3, F4: 4, F5: 5, F6: 6, F7: 7, F8: 8, F9: 9, F10: 10,
		F11: 11, F12: 12, F13: 13, F14: 14, F15: 15, F16: 16, F17: 17, F18: 18, F19: 19, F20: 20,
		F21: 21, F22: 22, F23: 23, F24: 24, F25: 25, F26: 26, F27: 27, F28: 28, F29: 29, F30: 30,
		F31: 31, F32: 32, F33: 33, F34: 34, F35: 35, F36: 36, F37: 37, F38: 38, F39: 39, F40: 40,
		F41: 41, F42: 42, F43: 43, F44: 44, F45: 45, F46: 46, F47: 47, F48: 48, F49: 49, F50: 50,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkWideStruct_Sonic_Marshal(b *testing.B) {
	data := WideStruct{
		F1: 1, F2: 2, F3: 3, F4: 4, F5: 5, F6: 6, F7: 7, F8: 8, F9: 9, F10: 10,
		F11: 11, F12: 12, F13: 13, F14: 14, F15: 15, F16: 16, F17: 17, F18: 18, F19: 19, F20: 20,
		F21: 21, F22: 22, F23: 23, F24: 24, F25: 25, F26: 26, F27: 27, F28: 28, F29: 29, F30: 30,
		F31: 31, F32: 32, F33: 33, F34: 34, F35: 35, F36: 36, F37: 37, F38: 38, F39: 39, F40: 40,
		F41: 41, F42: 42, F43: 43, F44: 44, F45: 45, F46: 46, F47: 47, F48: 48, F49: 49, F50: 50,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonic.Marshal(data)
	}
}

func BenchmarkWideStruct_MessagePack_Marshal(b *testing.B) {
	data := WideStruct{
		F1: 1, F2: 2, F3: 3, F4: 4, F5: 5, F6: 6, F7: 7, F8: 8, F9: 9, F10: 10,
		F11: 11, F12: 12, F13: 13, F14: 14, F15: 15, F16: 16, F17: 17, F18: 18, F19: 19, F20: 20,
		F21: 21, F22: 22, F23: 23, F24: 24, F25: 25, F26: 26, F27: 27, F28: 28, F29: 29, F30: 30,
		F31: 31, F32: 32, F33: 33, F34: 34, F35: 35, F36: 36, F37: 37, F38: 38, F39: 39, F40: 40,
		F41: 41, F42: 42, F43: 43, F44: 44, F45: 45, F46: 46, F47: 47, F48: 48, F49: 49, F50: 50,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(data)
	}
}

func BenchmarkWideStruct_CBOR_Marshal(b *testing.B) {
	data := WideStruct{
		F1: 1, F2: 2, F3: 3, F4: 4, F5: 5, F6: 6, F7: 7, F8: 8, F9: 9, F10: 10,
		F11: 11, F12: 12, F13: 13, F14: 14, F15: 15, F16: 16, F17: 17, F18: 18, F19: 19, F20: 20,
		F21: 21, F22: 22, F23: 23, F24: 24, F25: 25, F26: 26, F27: 27, F28: 28, F29: 29, F30: 30,
		F31: 31, F32: 32, F33: 33, F34: 34, F35: 35, F36: 36, F37: 37, F38: 38, F39: 39, F40: 40,
		F41: 41, F42: 42, F43: 43, F44: 44, F45: 45, F46: 46, F47: 47, F48: 48, F49: 49, F50: 50,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cbor.Marshal(data)
	}
}

// =============================================================================
// BENCHMARK 6: String-Heavy Data (Text vs Binary Trade-off)
// =============================================================================

type StringHeavyData struct {
	Description string   `beve:"desc" json:"desc"`
	Tags        []string `beve:"tags" json:"tags"`
	Content     string   `beve:"content" json:"content"`
	Metadata    string   `beve:"meta" json:"meta"`
}

func BenchmarkStringHeavy_BEVE_Marshal(b *testing.B) {
	data := StringHeavyData{
		Description: "This is a very long description that contains a lot of text to simulate real-world string-heavy data structures",
		Tags:        []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10"},
		Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.",
		Metadata:    "{\"author\": \"John Doe\", \"date\": \"2025-10-12\", \"version\": \"1.0.0\", \"category\": \"benchmark\"}",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkStringHeavy_JSON_Marshal(b *testing.B) {
	data := StringHeavyData{
		Description: "This is a very long description that contains a lot of text to simulate real-world string-heavy data structures",
		Tags:        []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10"},
		Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.",
		Metadata:    "{\"author\": \"John Doe\", \"date\": \"2025-10-12\", \"version\": \"1.0.0\", \"category\": \"benchmark\"}",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkStringHeavy_Sonic_Marshal(b *testing.B) {
	data := StringHeavyData{
		Description: "This is a very long description that contains a lot of text to simulate real-world string-heavy data structures",
		Tags:        []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10"},
		Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.",
		Metadata:    "{\"author\": \"John Doe\", \"date\": \"2025-10-12\", \"version\": \"1.0.0\", \"category\": \"benchmark\"}",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonic.Marshal(data)
	}
}

func BenchmarkStringHeavy_MessagePack_Marshal(b *testing.B) {
	data := StringHeavyData{
		Description: "This is a very long description that contains a lot of text to simulate real-world string-heavy data structures",
		Tags:        []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10"},
		Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.",
		Metadata:    "{\"author\": \"John Doe\", \"date\": \"2025-10-12\", \"version\": \"1.0.0\", \"category\": \"benchmark\"}",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(data)
	}
}

func BenchmarkStringHeavy_CBOR_Marshal(b *testing.B) {
	data := StringHeavyData{
		Description: "This is a very long description that contains a lot of text to simulate real-world string-heavy data structures",
		Tags:        []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10"},
		Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.",
		Metadata:    "{\"author\": \"John Doe\", \"date\": \"2025-10-12\", \"version\": \"1.0.0\", \"category\": \"benchmark\"}",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cbor.Marshal(data)
	}
}
