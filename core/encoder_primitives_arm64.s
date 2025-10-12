// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Assembly implementation of encodeUint for arm64.

#include "textflag.h"

// func encodeUintAsm(scratch *[9]byte, u uint64) int
TEXT ·encodeUintAsm(SB), NOSPLIT, $0-24
	MOVD    u+8(FP), R0          // Load u into R0
	MOVD    scratch+0(FP), R1    // Load scratch buffer pointer
	
	// Strategy: Check high bits directly (avoids comparison issues)
	// This is slightly slower but guarantees correctness
	
	// Check if upper 32 bits are non-zero
	LSR     $32, R0, R2          // R2 = u >> 32
	CBNZ    R2, eight_bytes      // If non-zero, need 8 bytes
	
	// Upper 32 bits are zero, check bits [31:16]
	LSR     $16, R0, R2          // R2 = u >> 16
	CBNZ    R2, four_bytes       // If non-zero, need 4 bytes
	
	// Bits [31:16] are zero, check bits [15:8]
	LSR     $8, R0, R2           // R2 = u >> 8
	CBNZ    R2, two_bytes        // If non-zero, need 2 bytes
	
	// All high bits are zero, only 1 byte needed
	B       one_byte

one_byte:
	// header = 0x11
	MOVD    u+8(FP), R0          // Reload u (in case it was modified)
	MOVD    $0x11, R2
	MOVB    R2, 0(R1)            // scratch[0] = header
	MOVB    R0, 1(R1)            // scratch[1] = byte(u)
	MOVD    $2, R0               // return 2
	MOVD    R0, ret+16(FP)
	RET

two_bytes:
	// header = 0x31
	MOVD    $0x31, R2
	MOVB    R2, 0(R1)            // scratch[0] = header
	MOVB    R0, 1(R1)            // scratch[1] = byte(u)
	LSR     $8, R0, R3
	MOVB    R3, 2(R1)            // scratch[2] = byte(u >> 8)
	MOVD    $3, R0               // return 3
	MOVD    R0, ret+16(FP)
	RET

four_bytes:
	// header = 0x51
	MOVD    $0x51, R2
	MOVB    R2, 0(R1)            // scratch[0] = header
	
	// Write 4 bytes
	MOVD    u+8(FP), R0          // Reload u
	MOVB    R0, 1(R1)            // scratch[1]
	LSR     $8, R0, R3
	MOVB    R3, 2(R1)            // scratch[2]
	LSR     $16, R0, R3
	MOVB    R3, 3(R1)            // scratch[3]
	LSR     $24, R0, R3
	MOVB    R3, 4(R1)            // scratch[4]
	MOVD    $5, R0               // return 5
	MOVD    R0, ret+16(FP)
	RET

eight_bytes:
	// header = 0x71
	MOVD    $0x71, R2
	MOVB    R2, 0(R1)            // scratch[0] = header
	
	// Write 8 bytes
	MOVD    u+8(FP), R0          // Reload u
	MOVB    R0, 1(R1)            // scratch[1]
	LSR     $8, R0, R3
	MOVB    R3, 2(R1)            // scratch[2]
	LSR     $16, R0, R3
	MOVB    R3, 3(R1)            // scratch[3]
	LSR     $24, R0, R3
	MOVB    R3, 4(R1)            // scratch[4]
	LSR     $32, R0, R3
	MOVB    R3, 5(R1)            // scratch[5]
	LSR     $40, R0, R3
	MOVB    R3, 6(R1)            // scratch[6]
	LSR     $48, R0, R3
	MOVB    R3, 7(R1)            // scratch[7]
	LSR     $56, R0, R3
	MOVB    R3, 8(R1)            // scratch[8]
	MOVD    $9, R0               // return 9
	MOVD    R0, ret+16(FP)
	RET
