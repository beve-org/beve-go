// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Assembly implementation of Buffer.WriteByte for arm64.
//
// Buffer struct layout:
//   type Buffer struct {
//       data []byte  // offset 0: pointer (8 bytes)
//                    // offset 8: len (8 bytes)
//                    // offset 16: cap (8 bytes)
//   }

#include "textflag.h"

// func writeByteAsm(b *Buffer, c byte) bool
TEXT ·writeByteAsm(SB), NOSPLIT, $0-17
	MOVD    b+0(FP), R0          // Load Buffer pointer into R0
	MOVD    0(R0), R1            // R1 = data pointer
	MOVD    8(R0), R2            // R2 = len
	MOVD    16(R0), R3           // R3 = cap
	
	// Fast path check: len < cap?
	CMP     R3, R2               // Compare len with cap (R2 - R3)
	BGE     slow_path            // If len >= cap, need growth
	
fast_path:
	// Write byte: data[len] = c
	MOVBU   c+8(FP), R4          // Load byte c into R4
	MOVBU   R4, (R1)(R2)         // data[len] = c
	
	// Increment len: len++
	ADD     $1, R2, R2           // R2 = len + 1
	MOVD    R2, 8(R0)            // Store new len
	
	// Return true (success)
	MOVD    $1, R0               // return true
	MOVB    R0, ret+16(FP)
	RET

slow_path:
	// Return false (needs growth)
	MOVD    $0, R0               // return false
	MOVB    R0, ret+16(FP)
	RET
