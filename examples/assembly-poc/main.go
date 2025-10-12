package main

// Assembly Optimization Proof-of-Concept Example
// This demonstrates how WriteCompressedUint could be implemented in Assembly

import (
	"fmt"
	"testing"
)

// Current Pure Go implementation (simplified)
func writeCompressedUintGo(n uint64, buf []byte) int {
	if n < 64 {
		buf[0] = byte(n << 2)
		return 1
	}

	if n < 16384 {
		buf[0] = byte(0x01 | ((n >> 8) << 2))
		buf[1] = byte(n)
		return 2
	}

	if n < 1073741824 {
		buf[0] = byte(0x02 | ((n >> 16) << 2))
		buf[1] = byte(n >> 8)
		buf[2] = byte(n)
		return 3
	}

	buf[0] = byte(0x03 | ((n >> 24) << 2))
	buf[1] = byte(n >> 16)
	buf[2] = byte(n >> 8)
	buf[3] = byte(n)
	return 4
}

// Assembly stub (would be in .s file)
// This is pseudo-assembly showing the optimization strategy
/*
TEXT ·writeCompressedUintAsm(SB), NOSPLIT, $0-32
    // Arguments:
    //   n uint64    @ 8(SP)
    //   buf []byte  @ 16(SP)

    MOVQ    n+0(FP), AX        // Load n into AX
    MOVQ    buf+8(FP), DI      // Load buf pointer into DI

    // Fast path: n < 64 (no branches with CMOV)
    CMPQ    AX, $64
    JGE     two_byte

one_byte:
    SHLQ    $2, AX             // n << 2
    MOVB    AL, 0(DI)          // buf[0] = byte(n << 2)
    MOVQ    $1, ret+24(FP)     // return 1
    RET

two_byte:
    CMPQ    AX, $16384
    JGE     three_byte

    MOVQ    AX, BX             // Copy n
    SHRQ    $8, BX             // n >> 8
    SHLQ    $2, BX             // (n >> 8) << 2
    ORB     $0x01, BL          // 0x01 | ...
    MOVB    BL, 0(DI)          // buf[0]
    MOVB    AL, 1(DI)          // buf[1] = byte(n)
    MOVQ    $2, ret+24(FP)     // return 2
    RET

three_byte:
    CMPQ    AX, $1073741824
    JGE     four_byte

    MOVQ    AX, BX
    SHRQ    $16, BX
    SHLQ    $2, BX
    ORB     $0x02, BL
    MOVB    BL, 0(DI)          // buf[0]
    MOVQ    AX, CX
    SHRQ    $8, CX
    MOVB    CL, 1(DI)          // buf[1]
    MOVB    AL, 2(DI)          // buf[2]
    MOVQ    $3, ret+24(FP)     // return 3
    RET

four_byte:
    MOVQ    AX, BX
    SHRQ    $24, BX
    SHLQ    $2, BX
    ORB     $0x03, BL
    MOVB    BL, 0(DI)          // buf[0]
    MOVQ    AX, CX
    SHRQ    $16, CX
    MOVB    CL, 1(DI)          // buf[1]
    MOVQ    AX, DX
    SHRQ    $8, DX
    MOVB    DL, 2(DI)          // buf[2]
    MOVB    AL, 3(DI)          // buf[3]
    MOVQ    $4, ret+24(FP)     // return 4
    RET
*/

// Benchmark comparison
func BenchmarkWriteCompressedUint_Go_Small(b *testing.B) {
	buf := make([]byte, 8)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = writeCompressedUintGo(uint64(i%64), buf)
	}
}

func BenchmarkWriteCompressedUint_Go_Medium(b *testing.B) {
	buf := make([]byte, 8)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = writeCompressedUintGo(uint64(i%16384), buf)
	}
}

func BenchmarkWriteCompressedUint_Go_Large(b *testing.B) {
	buf := make([]byte, 8)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = writeCompressedUintGo(uint64(i), buf)
	}
}

// Expected benchmark results:
//
// Pure Go (current):
//   Small:  ~2.5 ns/op    0 B/op   0 allocs/op
//   Medium: ~3.5 ns/op    0 B/op   0 allocs/op
//   Large:  ~5.0 ns/op    0 B/op   0 allocs/op
//
// Assembly (estimated):
//   Small:  ~1.8 ns/op    0 B/op   0 allocs/op  (28% faster)
//   Medium: ~2.5 ns/op    0 B/op   0 allocs/op  (29% faster)
//   Large:  ~3.5 ns/op    0 B/op   0 allocs/op  (30% faster)
//
// Overall impact on SmallStruct Marshal (719 ns/op):
//   If WriteCompressedUint is called 5-10 times per struct:
//   Savings: (3.5 - 2.5) * 7 = ~7 ns
//   New total: 719 - 7 = ~712 ns/op
//   Improvement: ~1% (minimal!)
//
// Conclusion: WriteCompressedUint is already very fast.
// Assembly gains would be <2% on overall Marshal performance.
// NOT WORTH THE EFFORT.

func main() {
	buf := make([]byte, 8)

	testCases := []uint64{0, 10, 63, 64, 100, 16383, 16384, 1073741823, 1073741824, 4294967295}

	for _, n := range testCases {
		len := writeCompressedUintGo(n, buf)
		fmt.Printf("n=%d → %d bytes: %v\n", n, len, buf[:len])
	}
}
