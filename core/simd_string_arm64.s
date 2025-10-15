//go:build arm64 && !purego
// +build arm64,!purego

#include "textflag.h"

// func validateUTF8ASM(data []byte) bool
TEXT ·validateUTF8ASM(SB), NOSPLIT, $0-25
    MOVD    data_base+0(FP), R0    // R0 = data pointer
    MOVD    data_len+8(FP), R1     // R1 = length
    
    // Handle empty string
    CBZ     R1, valid
    
    // Process 16-byte chunks using NEON
    CMP     $16, R1
    BLT     scalar_loop
    
    // Prepare NEON registers
    MOVD    $0x80, R2
    VMOV    R2, V1.D[0]            // V1 = 0x80 for ASCII check
    VMOV    R2, V1.D[1]
    
    // Main SIMD loop: process 16 bytes at once
simd_loop:
    VLD1    (R0), [V0.B16]         // Load 16 bytes into V0
    
    // Check if all bytes are ASCII (< 0x80)
    // Use UMAXV to find the maximum byte in the vector
    // If max < 0x80, all bytes are ASCII
    WORD    $0x6e30a800             // UMAXV B0, V0.16B (max byte -> V0[0])
    VMOV    V0.B[0], R3
    CMP     $0x80, R3
    BGE     non_ascii_chunk
    
    // All ASCII, advance quickly
    ADD     $16, R0
    SUB     $16, R1
    CMP     $16, R1
    BGE     simd_loop
    
    // Handle remaining bytes
    CBZ     R1, valid
    B       scalar_loop

non_ascii_chunk:
    // Found non-ASCII, validate this chunk carefully
    // For now, fall back to scalar validation
    B       scalar_loop

scalar_loop:
    // Load one byte
    MOVBU   (R0), R2
    
    // Check if ASCII (< 0x80)
    CMP     $0x80, R2
    BLT     next_byte
    
    // Non-ASCII: validate multi-byte sequence
    // 2-byte: 0xC2-0xDF
    CMP     $0xC2, R2
    BLT     invalid
    CMP     $0xDF, R2
    BLE     validate_2byte
    
    // 3-byte: 0xE0-0xEF
    CMP     $0xEF, R2
    BLE     validate_3byte
    
    // 4-byte: 0xF0-0xF4
    CMP     $0xF4, R2
    BGT     invalid
    
    // Validate 4-byte sequence
    CMP     $4, R1
    BLT     invalid
    
    // Load remaining 3 bytes
    MOVBU   1(R0), R3
    MOVBU   2(R0), R4
    MOVBU   3(R0), R5
    
    // Check continuation bytes (10xxxxxx)
    AND     $0xC0, R3, R6
    CMP     $0x80, R6
    BNE     invalid
    AND     $0xC0, R4, R6
    CMP     $0x80, R6
    BNE     invalid
    AND     $0xC0, R5, R6
    CMP     $0x80, R6
    BNE     invalid
    
    // Check for overlong (0xF0 with < 0x90)
    CMP     $0xF0, R2
    BNE     check_4byte_range
    CMP     $0x90, R3
    BLT     invalid
    
check_4byte_range:
    // Check for out of range (0xF4 with >= 0x90)
    CMP     $0xF4, R2
    BNE     skip_4byte
    CMP     $0x90, R3
    BGE     invalid
    
skip_4byte:
    ADD     $4, R0
    SUB     $4, R1
    B       continue_loop

validate_3byte:
    // Validate 3-byte sequence
    CMP     $3, R1
    BLT     invalid
    
    MOVBU   1(R0), R3
    MOVBU   2(R0), R4
    
    // Check continuation bytes
    AND     $0xC0, R3, R5
    CMP     $0x80, R5
    BNE     invalid
    AND     $0xC0, R4, R5
    CMP     $0x80, R5
    BNE     invalid
    
    // Check for overlong (0xE0 with < 0xA0)
    CMP     $0xE0, R2
    BNE     check_surrogate
    CMP     $0xA0, R3
    BLT     invalid
    
check_surrogate:
    // Check for surrogate pair (0xED with >= 0xA0)
    CMP     $0xED, R2
    BNE     skip_3byte
    CMP     $0xA0, R3
    BGE     invalid
    
skip_3byte:
    ADD     $3, R0
    SUB     $3, R1
    B       continue_loop

validate_2byte:
    // Validate 2-byte sequence
    CMP     $2, R1
    BLT     invalid
    
    MOVBU   1(R0), R3
    
    // Check continuation byte
    AND     $0xC0, R3, R4
    CMP     $0x80, R4
    BNE     invalid
    
    ADD     $2, R0
    SUB     $2, R1
    B       continue_loop

next_byte:
    ADD     $1, R0
    SUB     $1, R1

continue_loop:
    CBNZ    R1, scalar_loop

valid:
    MOVD    $1, R0
    MOVB    R0, ret+24(FP)
    RET

invalid:
    MOVB    ZR, ret+24(FP)
    RET
