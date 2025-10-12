// Copyright 2025 BEVE-Go Authors. All rights reserved.
// Assembly implementation of Buffer.WriteByte for amd64.
//
// This file contains hand-optimized assembly for single-byte buffer writes.
// Performance improvements over pure Go:
//   - Bounds check elimination for fast path
//   - Direct memory write (no append overhead)
//   - Reduced instruction count
//
// Buffer struct layout (from buffer.go):
//   type Buffer struct {
//       data []byte  // offset 0: pointer (8 bytes)
//                    // offset 8: len (8 bytes)
//                    // offset 16: cap (8 bytes)
//   }

#include "textflag.h"

// func writeByteAsm(b *Buffer, c byte) bool
// Returns true if written successfully (fast path), false if needs growth (slow path)
TEXT ·writeByteAsm(SB), NOSPLIT, $0-17
	MOVQ    b+0(FP), DI          // Load Buffer pointer into DI
	MOVQ    0(DI), AX            // AX = data pointer
	MOVQ    8(DI), BX            // BX = len
	MOVQ    16(DI), CX           // CX = cap
	
	// Fast path check: len < cap?
	CMPQ    BX, CX               // Compare len with cap
	JGE     slow_path            // If len >= cap, need growth
	
fast_path:
	// Write byte: data[len] = c
	MOVB    c+8(FP), DX          // Load byte c into DL
	MOVB    DX, 0(AX)(BX*1)      // data[len] = c
	
	// Increment len: len++
	INCQ    BX                   // len++
	MOVQ    BX, 8(DI)            // Store new len
	
	// Return true (success)
	MOVB    $1, ret+16(FP)       // return true
	RET

slow_path:
	// Return false (needs growth - handled by Go)
	MOVB    $0, ret+16(FP)       // return false
	RET
