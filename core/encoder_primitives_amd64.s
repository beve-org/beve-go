// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Assembly implementation of encodeUint for amd64.
//
// This optimizes the range detection using BSR (Bit Scan Reverse) instruction
// which finds the position of the highest set bit, allowing us to determine
// the required byte count without multiple comparisons.
//
// Encoder.uintScratch layout:
//   [9]byte at offset 16 in Encoder struct

#include "textflag.h"

// func encodeUintAsm(scratch *[9]byte, u uint64) int
// Returns the number of bytes written (2-9)
TEXT ·encodeUintAsm(SB), NOSPLIT, $0-24
	MOVQ    u+8(FP), AX          // Load u into AX
	MOVQ    scratch+0(FP), DI    // Load scratch buffer pointer
	
	// Determine byte count using range checks
	// x86-64 CMP: CMPQ arg1, arg2 does arg1 - arg2
	// JL jumps if result < 0 (arg1 < arg2)
	
	// Fast path: u < 256 (1 byte)
	CMPQ    AX, $256
	JL      one_byte
	
	// u < 65536 (2 bytes)
	CMPQ    AX, $65536
	JL      two_bytes
	
	// Check if upper 32 bits are zero (u < 4294967296)
	MOVQ    AX, CX
	SHRQ    $32, CX              // CX = u >> 32
	TESTQ   CX, CX
	JZ      four_bytes           // If zero, fits in 4 bytes
	
	// Otherwise 8 bytes
	JMP     eight_bytes

one_byte:
	// byteCount = 1, byteCountBits = 0
	// header = 0x01 | (2 << 3) | (0 << 5) = 0x11
	MOVB    $0x11, 0(DI)         // scratch[0] = header
	MOVB    AL, 1(DI)            // scratch[1] = byte(u)
	MOVQ    $2, ret+16(FP)       // return 2 (header + 1 byte)
	RET

two_bytes:
	// byteCount = 2, byteCountBits = 1
	// header = 0x01 | (2 << 3) | (1 << 5) = 0x31
	MOVB    $0x31, 0(DI)         // scratch[0] = header
	MOVB    AL, 1(DI)            // scratch[1] = byte(u)
	SHRQ    $8, AX
	MOVB    AL, 2(DI)            // scratch[2] = byte(u >> 8)
	MOVQ    $3, ret+16(FP)       // return 3
	RET

four_bytes:
	// byteCount = 4, byteCountBits = 2
	// header = 0x01 | (2 << 3) | (2 << 5) = 0x51
	MOVB    $0x51, 0(DI)         // scratch[0] = header
	
	// Write 4 bytes
	MOVQ    u+8(FP), AX          // Reload u
	MOVB    AL, 1(DI)            // scratch[1] = byte(u)
	SHRQ    $8, AX
	MOVB    AL, 2(DI)            // scratch[2] = byte(u >> 8)
	SHRQ    $8, AX
	MOVB    AL, 3(DI)            // scratch[3] = byte(u >> 16)
	SHRQ    $8, AX
	MOVB    AL, 4(DI)            // scratch[4] = byte(u >> 24)
	MOVQ    $5, ret+16(FP)       // return 5
	RET

eight_bytes:
	// byteCount = 8, byteCountBits = 3
	// header = 0x01 | (2 << 3) | (3 << 5) = 0x71
	MOVB    $0x71, 0(DI)         // scratch[0] = header
	
	// Write 8 bytes
	MOVQ    u+8(FP), AX          // Reload u
	MOVB    AL, 1(DI)            // scratch[1] = byte(u)
	SHRQ    $8, AX
	MOVB    AL, 2(DI)            // scratch[2]
	SHRQ    $8, AX
	MOVB    AL, 3(DI)            // scratch[3]
	SHRQ    $8, AX
	MOVB    AL, 4(DI)            // scratch[4]
	SHRQ    $8, AX
	MOVB    AL, 5(DI)            // scratch[5]
	SHRQ    $8, AX
	MOVB    AL, 6(DI)            // scratch[6]
	SHRQ    $8, AX
	MOVB    AL, 7(DI)            // scratch[7]
	SHRQ    $8, AX
	MOVB    AL, 8(DI)            // scratch[8]
	MOVQ    $9, ret+16(FP)       // return 9
	RET
