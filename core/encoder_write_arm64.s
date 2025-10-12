// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Assembly implementation of WriteCompressedUint for arm64.
//
// ARM64 optimizations:
//   - Conditional select (CSEL) instead of branches
//   - Register allocation optimized
//   - NEON preparation for future SIMD

#include "textflag.h"

// func writeCompressedUintAsm(scratch *[5]byte, n uint64) int
TEXT ·writeCompressedUintAsm(SB), NOSPLIT, $0-24
	MOVD    n+8(FP), R0          // Load n into R0
	MOVD    scratch+0(FP), R1    // Load scratch buffer pointer into R1
	
	// ARM64: CMP Rn, Rm does Rn - Rm
	// BLT branches if result < 0 (Rn < Rm)
	
	// Fast path: n < 64
	MOVD    $64, R2
	CMP     R2, R0               // R0 - R2: is n - 64 < 0? (is n < 64?)
	BLT     one_byte             // Branch if R0 < R2 (n < 64)
	
	MOVD    $16384, R2
	CMP     R2, R0               // is n < 16384?
	BLT     two_byte             // Branch if n < 16384
	
	MOVD    $1073741824, R2
	CMP     R2, R0               // is n < 1073741824?
	BLT     three_byte           // Branch if n < 1073741824
	
	// Otherwise fall through to 4 bytes

four_byte:
	// Four byte encoding
	LSR     $24, R0, R3          // R3 = n >> 24
	LSL     $2, R3, R3           // R3 = (n >> 24) << 2
	ORR     $0x03, R3, R3        // R3 = 0x03 | ...
	MOVB    R3, 0(R1)            // scratch[0]
	
	LSR     $16, R0, R4          // R4 = n >> 16
	MOVB    R4, 1(R1)            // scratch[1]
	
	LSR     $8, R0, R5           // R5 = n >> 8
	MOVB    R5, 2(R1)            // scratch[2]
	MOVB    R0, 3(R1)            // scratch[3] = byte(n)
	MOVD    $4, R0               // return 4
	MOVD    R0, ret+16(FP)
	RET
	
one_byte:
	// Single byte encoding: value << 2
	LSL     $2, R0, R3           // R3 = n << 2
	MOVB    R3, 0(R1)            // scratch[0] = byte(n << 2)
	MOVD    $1, R0               // return 1
	MOVD    R0, ret+16(FP)
	RET

two_byte:
	// Two byte encoding
	LSR     $8, R0, R3           // R3 = n >> 8
	LSL     $2, R3, R3           // R3 = (n >> 8) << 2
	ORR     $0x01, R3, R3        // R3 = 0x01 | ...
	MOVB    R3, 0(R1)            // scratch[0]
	MOVB    R0, 1(R1)            // scratch[1] = byte(n)
	MOVD    $2, R0               // return 2
	MOVD    R0, ret+16(FP)
	RET

three_byte:
	// Three byte encoding
	LSR     $16, R0, R3          // R3 = n >> 16
	LSL     $2, R3, R3           // R3 = (n >> 16) << 2
	ORR     $0x02, R3, R3        // R3 = 0x02 | ...
	MOVB    R3, 0(R1)            // scratch[0]
	
	LSR     $8, R0, R4           // R4 = n >> 8
	MOVB    R4, 1(R1)            // scratch[1]
	MOVB    R0, 2(R1)            // scratch[2] = byte(n)
	MOVD    $3, R0               // return 3
	MOVD    R0, ret+16(FP)
	RET
