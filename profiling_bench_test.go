package beve

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"
)

// This file contains CPU profiling benchmarks to identify bottlenecks

// =============================================================================
// PROFILING BENCHMARK 1: Wide Struct (Current Bottleneck)
// =============================================================================

type WideStructProfile struct {
	F1, F2, F3, F4, F5      int `beve:"f1" json:"f1"`
	F6, F7, F8, F9, F10     int `beve:"f6" json:"f6"`
	F11, F12, F13, F14, F15 int `beve:"f11" json:"f11"`
	F16, F17, F18, F19, F20 int `beve:"f16" json:"f16"`
	F21, F22, F23, F24, F25 int `beve:"f21" json:"f21"`
	F26, F27, F28, F29, F30 int `beve:"f26" json:"f26"`
	F31, F32, F33, F34, F35 int `beve:"f31" json:"f31"`
	F36, F37, F38, F39, F40 int `beve:"f36" json:"f36"`
	F41, F42, F43, F44, F45 int `beve:"f41" json:"f41"`
	F46, F47, F48, F49, F50 int `beve:"f46" json:"f46"`
}

func BenchmarkProfile_WideStruct_BEVE(b *testing.B) {
	data := WideStructProfile{
		F1: 1, F2: 2, F3: 3, F4: 4, F5: 5,
		F6: 6, F7: 7, F8: 8, F9: 9, F10: 10,
		F11: 11, F12: 12, F13: 13, F14: 14, F15: 15,
		F16: 16, F17: 17, F18: 18, F19: 19, F20: 20,
		F21: 21, F22: 22, F23: 23, F24: 24, F25: 25,
		F26: 26, F27: 27, F28: 28, F29: 29, F30: 30,
		F31: 31, F32: 32, F33: 33, F34: 34, F35: 35,
		F36: 36, F37: 37, F38: 38, F39: 39, F40: 40,
		F41: 41, F42: 42, F43: 43, F44: 44, F45: 45,
		F46: 46, F47: 47, F48: 48, F49: 49, F50: 50,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_WideStruct_JSON(b *testing.B) {
	data := WideStructProfile{
		F1: 1, F2: 2, F3: 3, F4: 4, F5: 5,
		F6: 6, F7: 7, F8: 8, F9: 9, F10: 10,
		F11: 11, F12: 12, F13: 13, F14: 14, F15: 15,
		F16: 16, F17: 17, F18: 18, F19: 19, F20: 20,
		F21: 21, F22: 22, F23: 23, F24: 24, F25: 25,
		F26: 26, F27: 27, F28: 28, F29: 29, F30: 30,
		F31: 31, F32: 32, F33: 33, F34: 34, F35: 35,
		F36: 36, F37: 37, F38: 38, F39: 39, F40: 40,
		F41: 41, F42: 42, F43: 43, F44: 44, F45: 45,
		F46: 46, F47: 47, F48: 48, F49: 49, F50: 50,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

// =============================================================================
// PROFILING BENCHMARK 2: Large Map (High Allocation)
// =============================================================================

func createProfileMap(size int) map[string]int {
	m := make(map[string]int, size)
	for i := 0; i < size; i++ {
		key := string(rune('a'+(i%26))) + string(rune('0'+(i/26)%10))
		m[key] = i
	}
	return m
}

func BenchmarkProfile_LargeMap_BEVE(b *testing.B) {
	data := createProfileMap(1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_LargeMap_MessagePack(b *testing.B) {
	data := createProfileMap(1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(data)
	}
}

// =============================================================================
// PROFILING BENCHMARK 3: Deep Nesting (Reflection Overhead)
// =============================================================================

type ProfileNested struct {
	Value string         `beve:"v" json:"v"`
	Next  *ProfileNested `beve:"n" json:"n"`
}

func createProfileNested(depth int) ProfileNested {
	root := ProfileNested{Value: "root"}
	current := &root
	for i := 1; i < depth; i++ {
		current.Next = &ProfileNested{Value: string(rune('0' + i))}
		current = current.Next
	}
	return root
}

func BenchmarkProfile_DeepNested_BEVE(b *testing.B) {
	data := createProfileNested(10)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_DeepNested_CBOR(b *testing.B) {
	data := createProfileNested(10)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cbor.Marshal(data)
	}
}

// =============================================================================
// PROFILING BENCHMARK 4: Interface Slice (Type Switching)
// =============================================================================

func createProfileInterfaceSlice(size int) []interface{} {
	result := make([]interface{}, size)
	for i := 0; i < size; i++ {
		switch i % 4 {
		case 0:
			result[i] = i
		case 1:
			result[i] = float64(i) * 1.5
		case 2:
			result[i] = string(rune('a' + (i % 26)))
		case 3:
			result[i] = i%2 == 0
		}
	}
	return result
}

func BenchmarkProfile_InterfaceSlice_BEVE(b *testing.B) {
	data := createProfileInterfaceSlice(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_InterfaceSlice_CBOR(b *testing.B) {
	data := createProfileInterfaceSlice(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cbor.Marshal(data)
	}
}

// =============================================================================
// PROFILING BENCHMARK 5: Struct Field Access Pattern Analysis
// =============================================================================

// Test different struct layouts to understand cache effects
type CacheFriendlyStruct struct {
	// Hot fields first (most accessed)
	ID     int64 `beve:"id" json:"id"`
	Status byte  `beve:"status" json:"status"`
	Count  int32 `beve:"count" json:"count"`
	// Cold fields last
	Metadata string `beve:"metadata" json:"metadata"`
	Extra    []byte `beve:"extra" json:"extra"`
}

type CacheUnfriendlyStruct struct {
	// Mixed hot and cold fields
	Metadata string `beve:"metadata" json:"metadata"`
	ID       int64  `beve:"id" json:"id"`
	Extra    []byte `beve:"extra" json:"extra"`
	Status   byte   `beve:"status" json:"status"`
	Count    int32  `beve:"count" json:"count"`
}

func BenchmarkProfile_CacheFriendly_BEVE(b *testing.B) {
	data := CacheFriendlyStruct{
		ID:       123,
		Status:   1,
		Count:    100,
		Metadata: "test",
		Extra:    []byte("extra"),
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_CacheUnfriendly_BEVE(b *testing.B) {
	data := CacheUnfriendlyStruct{
		ID:       123,
		Status:   1,
		Count:    100,
		Metadata: "test",
		Extra:    []byte("extra"),
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

// =============================================================================
// PROFILING BENCHMARK 6: String vs []byte encoding
// =============================================================================

type StringStruct struct {
	Data1 string `beve:"d1" json:"d1"`
	Data2 string `beve:"d2" json:"d2"`
	Data3 string `beve:"d3" json:"d3"`
}

type ByteSliceStruct struct {
	Data1 []byte `beve:"d1" json:"d1"`
	Data2 []byte `beve:"d2" json:"d2"`
	Data3 []byte `beve:"d3" json:"d3"`
}

func BenchmarkProfile_String_BEVE(b *testing.B) {
	data := StringStruct{
		Data1: "Lorem ipsum dolor sit amet",
		Data2: "consectetur adipiscing elit",
		Data3: "sed do eiusmod tempor",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_ByteSlice_BEVE(b *testing.B) {
	data := ByteSliceStruct{
		Data1: []byte("Lorem ipsum dolor sit amet"),
		Data2: []byte("consectetur adipiscing elit"),
		Data3: []byte("sed do eiusmod tempor"),
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

// =============================================================================
// PROFILING BENCHMARK 7: Map Key Type Performance
// =============================================================================

func BenchmarkProfile_MapStringKey_BEVE(b *testing.B) {
	data := map[string]int{
		"key1": 1, "key2": 2, "key3": 3, "key4": 4, "key5": 5,
		"key6": 6, "key7": 7, "key8": 8, "key9": 9, "key10": 10,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_MapIntKey_BEVE(b *testing.B) {
	data := map[int]int{
		1: 1, 2: 2, 3: 3, 4: 4, 5: 5,
		6: 6, 7: 7, 8: 8, 9: 9, 10: 10,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

// =============================================================================
// PROFILING BENCHMARK 8: Small vs Large Allocations
// =============================================================================

type TinyStruct struct {
	A int `beve:"a" json:"a"`
	B int `beve:"b" json:"b"`
}

type MediumStruct struct {
	A, B, C, D, E int `beve:"a" json:"a"`
	F, G, H, I, J int `beve:"f" json:"f"`
}

type LargeStruct struct {
	A, B, C, D, E    int    `beve:"a" json:"a"`
	F, G, H, I, J    int    `beve:"f" json:"f"`
	K, L, M, N, O    int    `beve:"k" json:"k"`
	P, Q, R, S, T    int    `beve:"p" json:"p"`
	U, V, W, X, Y, Z int    `beve:"u" json:"u"`
	Data             string `beve:"data" json:"data"`
}

func BenchmarkProfile_TinyStruct_BEVE(b *testing.B) {
	data := TinyStruct{A: 1, B: 2}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_MediumStruct_BEVE(b *testing.B) {
	data := MediumStruct{
		A: 1, B: 2, C: 3, D: 4, E: 5,
		F: 6, G: 7, H: 8, I: 9, J: 10,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_LargeStruct_BEVE(b *testing.B) {
	data := LargeStruct{
		A: 1, B: 2, C: 3, D: 4, E: 5,
		F: 6, G: 7, H: 8, I: 9, J: 10,
		K: 11, L: 12, M: 13, N: 14, O: 15,
		P: 16, Q: 17, R: 18, S: 19, T: 20,
		U: 21, V: 22, W: 23, X: 24, Y: 25, Z: 26,
		Data: "Some data here",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

// =============================================================================
// PROFILING BENCHMARK 9: Slice Size Impact
// =============================================================================

func BenchmarkProfile_SmallSlice_BEVE(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_MediumSlice_BEVE(b *testing.B) {
	data := make([]int, 100)
	for i := range data {
		data[i] = i
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

func BenchmarkProfile_LargeSlice_BEVE(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

// =============================================================================
// PROFILING BENCHMARK 10: Encoder Reuse vs New
// =============================================================================

func BenchmarkProfile_EncoderReuse_BEVE(b *testing.B) {
	data := MediumStruct{
		A: 1, B: 2, C: 3, D: 4, E: 5,
		F: 6, G: 7, H: 8, I: 9, J: 10,
	}

	lease, _ := MarshalZeroCopy(data)
	lease.Release()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lease, _ := MarshalZeroCopy(data)
		lease.Release()
	}
}

func BenchmarkProfile_EncoderNew_BEVE(b *testing.B) {
	data := MediumStruct{
		A: 1, B: 2, C: 3, D: 4, E: 5,
		F: 6, G: 7, H: 8, I: 9, J: 10,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}

// =============================================================================
// PROFILING BENCHMARK 11: Comparison with Sonic (fastest JSON)
// =============================================================================

func BenchmarkProfile_Sonic_SmallStruct(b *testing.B) {
	data := TinyStruct{A: 1, B: 2}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonic.Marshal(data)
	}
}

func BenchmarkProfile_BEVE_SmallStruct(b *testing.B) {
	data := TinyStruct{A: 1, B: 2}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(data)
	}
}
