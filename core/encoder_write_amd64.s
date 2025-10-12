// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Assembly implementation of WriteCompressedUint for amd64.
//
// This file contains hand-optimized assembly for variable-length integer encoding.
// Performance improvements over pure Go:
//   - Branch prediction optimized
//   - Register allocation optimized
//   - Reduced instruction count
//
// BEVE varint encoding:
//   [0, 63]:              1 byte  (value << 2)
//   [64, 16383]:          2 bytes (0x01 | (value << 2 for high bits))
//   [16384, 1073741823]:  3 bytes (0x02 | ...)
//   [1073741824, ...]:    4 bytes (0x03 | ...)

#include "textflag.h"

// func writeCompressedUintAsm(scratch *[5]byte, n uint64) int
TEXT ·writeCompressedUintAsm(SB), NOSPLIT, $0-24
	MOVQ    n+8(FP), AX          // Load n into AX
	MOVQ    scratch+0(FP), DI    // Load scratch buffer pointer into DI
	
	// Plan 9 Assembly: CMPQ arg1, arg2 does arg2 - arg1
	// JLT jumps if result < 0 (i.e. arg2 < arg1)
	
	// Check if n < 64
	CMPQ    AX, $64              // Compare: is AX - 64 < 0? (is AX < 64?)
	JLT     one_byte             // Jump if AX < 64
	
	// Check if n < 16384
	CMPQ    AX, $16384           // Compare: is AX < 16384?
	JLT     two_byte             // Jump if AX < 16384
	
	// Check if n < 1073741824
	CMPQ    AX, $1073741824      // Compare: is AX < 1073741824?
	JLT     three_byte           // Jump if AX < 1073741824
	
	// Otherwise fall through to 4 bytes
	
four_byte:
	// Four byte encoding (32-bit values)
	// scratch[0] = 0x03 | ((n >> 24) << 2)
	// scratch[1] = byte(n >> 16)
	// scratch[2] = byte(n >> 8)
	// scratch[3] = byte(n)
	MOVQ    AX, BX               // Copy n
	SHRQ    $24, BX              // n >> 24
	SHLQ    $2, BX               // (n >> 24) << 2
	ORB     $0x03, BL            // 0x03 | ...
	MOVB    BL, 0(DI)            // scratch[0]
	
	MOVQ    AX, CX               // Copy n
	SHRQ    $16, CX              // n >> 16
	MOVB    CL, 1(DI)            // scratch[1]
	
	MOVQ    AX, DX               // Copy n
	SHRQ    $8, DX               // n >> 8
	MOVB    DL, 2(DI)            // scratch[2]
	MOVB    AL, 3(DI)            // scratch[3] = byte(n)
	MOVQ    $4, ret+16(FP)       // return 4
	RET
	
one_byte:
	// Single byte encoding: value << 2
	SHLQ    $2, AX               // n << 2
	MOVB    AL, 0(DI)            // scratch[0] = byte(n << 2)
	MOVQ    $1, ret+16(FP)       // return 1
	RET

two_byte:
	// Two byte encoding
	// scratch[0] = 0x01 | ((n >> 8) << 2)
	// scratch[1] = byte(n)
	MOVQ    AX, BX               // Copy n to BX
	SHRQ    $8, BX               // n >> 8
	SHLQ    $2, BX               // (n >> 8) << 2
	ORB     $0x01, BL            // 0x01 | ...
	MOVB    BL, 0(DI)            // scratch[0]
	MOVB    AL, 1(DI)            // scratch[1] = byte(n)
	MOVQ    $2, ret+16(FP)       // return 2
	RET

three_byte:
	// Three byte encoding
	// scratch[0] = 0x02 | ((n >> 16) << 2)
	// scratch[1] = byte(n >> 8)
	// scratch[2] = byte(n)
	MOVQ    AX, BX               // Copy n
	SHRQ    $16, BX              // n >> 16
	SHLQ    $2, BX               // (n >> 16) << 2
	ORB     $0x02, BL            // 0x02 | ...
	MOVB    BL, 0(DI)            // scratch[0]
	
	MOVQ    AX, CX               // Copy n again
	SHRQ    $8, CX               // n >> 8
	MOVB    CL, 1(DI)            // scratch[1]
	MOVB    AL, 2(DI)            // scratch[2] = byte(n)
	MOVQ    $3, ret+16(FP)       // return 3
	RET
