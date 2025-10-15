//go:build amd64 && !purego
// +build amd64,!purego

#include "textflag.h"

// func validateUTF8ASM(data []byte) bool
TEXT ·validateUTF8ASM(SB), NOSPLIT, $0-25
    MOVQ    data_base+0(FP), SI    // SI = data pointer
    MOVQ    data_len+8(FP), CX     // CX = length
    
    // Handle empty string
    TESTQ   CX, CX
    JZ      valid
    
    // Check if AVX2 is available (should be checked at init)
    // Assume available for now, fallback handled in Go
    
    // Process 32-byte chunks using AVX2
    CMPQ    CX, $32
    JB      scalar_loop
    
    // Load ASCII detection constant (0x80) into YMM1
    MOVQ    $0x8080808080808080, AX
    MOVQ    AX, X1
    VPBROADCASTQ    X1, Y1         // Y1 = [0x80 × 32]
    
    // Main SIMD loop: process 32 bytes at once
simd_loop:
    // Load 32 bytes into YMM0
    VMOVDQU (SI), Y0
    
    // Find maximum byte in the vector
    // If all bytes < 0x80, max will be < 0x80 (ASCII)
    VPMAXUB Y0, Y1, Y2              // Y2 = max(Y0, Y1)
    VPCMPEQB Y1, Y2, Y3             // Y3 = (Y2 == Y1) ? 0xFF : 0x00
    VPMOVMSKB Y3, AX                // Move mask to AX
    
    // If AX == 0xFFFFFFFF, all bytes are >= 0x80 (need detailed check)
    // If AX == 0, all bytes are < 0x80 (pure ASCII, fast path)
    TESTL   AX, AX
    JZ      advance_simd
    
    // Non-ASCII detected, use scalar validation
    // Fall through to scalar loop
    JMP     scalar_loop

advance_simd:
    ADDQ    $32, SI
    SUBQ    $32, CX
    CMPQ    CX, $32
    JAE     simd_loop
    
    // Handle remaining bytes
    TESTQ   CX, CX
    JZ      valid
    JMP     scalar_loop

scalar_loop:
    // Load one byte
    MOVBQZX (SI), AX
    
    // Check if ASCII (< 0x80)
    CMPB    AL, $0x80
    JB      next_byte
    
    // Non-ASCII: validate multi-byte sequence
    // 2-byte: 0xC2-0xDF
    CMPB    AL, $0xC2
    JB      invalid
    CMPB    AL, $0xDF
    JBE     validate_2byte
    
    // 3-byte: 0xE0-0xEF
    CMPB    AL, $0xEF
    JBE     validate_3byte
    
    // 4-byte: 0xF0-0xF4
    CMPB    AL, $0xF4
    JA      invalid
    
    // Validate 4-byte sequence
    CMPQ    CX, $4
    JB      invalid
    
    // Load remaining 3 bytes
    MOVBQZX 1(SI), BX
    MOVBQZX 2(SI), DX
    MOVBQZX 3(SI), DI
    
    // Check continuation bytes (10xxxxxx)
    MOVQ    BX, R8
    ANDQ    $0xC0, R8
    CMPQ    R8, $0x80
    JNE     invalid
    
    MOVQ    DX, R8
    ANDQ    $0xC0, R8
    CMPQ    R8, $0x80
    JNE     invalid
    
    MOVQ    DI, R8
    ANDQ    $0xC0, R8
    CMPQ    R8, $0x80
    JNE     invalid
    
    // Check for overlong (0xF0 with < 0x90)
    CMPB    AL, $0xF0
    JNE     check_4byte_range
    CMPB    BL, $0x90
    JB      invalid
    
check_4byte_range:
    // Check for out of range (0xF4 with >= 0x90)
    CMPB    AL, $0xF4
    JNE     skip_4byte
    CMPB    BL, $0x90
    JAE     invalid
    
skip_4byte:
    ADDQ    $4, SI
    SUBQ    $4, CX
    JMP     continue_loop

validate_3byte:
    // Validate 3-byte sequence
    CMPQ    CX, $3
    JB      invalid
    
    MOVBQZX 1(SI), BX
    MOVBQZX 2(SI), DX
    
    // Check continuation bytes
    MOVQ    BX, R8
    ANDQ    $0xC0, R8
    CMPQ    R8, $0x80
    JNE     invalid
    
    MOVQ    DX, R8
    ANDQ    $0xC0, R8
    CMPQ    R8, $0x80
    JNE     invalid
    
    // Check for overlong (0xE0 with < 0xA0)
    CMPB    AL, $0xE0
    JNE     check_surrogate
    CMPB    BL, $0xA0
    JB      invalid
    
check_surrogate:
    // Check for surrogate pair (0xED with >= 0xA0)
    CMPB    AL, $0xED
    JNE     skip_3byte
    CMPB    BL, $0xA0
    JAE     invalid
    
skip_3byte:
    ADDQ    $3, SI
    SUBQ    $3, CX
    JMP     continue_loop

validate_2byte:
    // Validate 2-byte sequence
    CMPQ    CX, $2
    JB      invalid
    
    MOVBQZX 1(SI), BX
    
    // Check continuation byte
    MOVQ    BX, R8
    ANDQ    $0xC0, R8
    CMPQ    R8, $0x80
    JNE     invalid
    
    ADDQ    $2, SI
    SUBQ    $2, CX
    JMP     continue_loop

next_byte:
    ADDQ    $1, SI
    SUBQ    $1, CX

continue_loop:
    TESTQ   CX, CX
    JNZ     scalar_loop

valid:
    MOVB    $1, ret+24(FP)
    VZEROUPPER                      // Clean up AVX state
    RET

invalid:
    MOVB    $0, ret+24(FP)
    VZEROUPPER                      // Clean up AVX state
    RET
