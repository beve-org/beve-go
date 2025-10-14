// Assembly implementation of varint encoding for ARM64.
//
// This file provides hand-optimized assembly for the most critical
// hot-path operations in BEVE encoding: variable-length integer encoding.
//
// Performance targets:
//   - encodeVarintAsm: <5ns for small values (<64)
//   - encodeVarintAsm: <10ns for medium values (64-16383)
//   - Compare to pure Go: ~8ns for small, ~15ns for medium
//   - Target speedup: 40-50% faster than pure Go
//
// BEVE varint encoding format:
//   [0, 63]:              1 byte  (value << 2)
//   [64, 16383]:          2 bytes (0x01 | (value << 2 for high bits))
//   [16384, 1073741823]:  3 bytes (0x02 | ...)
//   [1073741824, ...]:    4 bytes (0x03 | ...)
//
// Register usage conventions (ARM64 ABI):
//   - R0-R7: Arguments and return values
//   - R8-R15: Temporary registers
//   - R16-R17: Intra-procedure-call temporary (IP0, IP1)
//   - R18: Platform register (reserved)
//   - R19-R28: Callee-saved
//   - R29: Frame pointer (FP)
//   - R30: Link register (LR)
//   - SP: Stack pointer
//
// Function signature:
//   func encodeVarintAsm(buf []byte, value uint64) int
//
// Args:
//   R0: buf pointer (base address)
//   R2: value to encode (uint64)
//
// Returns:
//   R0: Number of bytes written (1, 2, 3, or 4)
//
// Note: This is a template/stub. Actual assembly optimization requires
// careful benchmarking and may not provide significant gains over
// compiler-optimized Go code on modern ARM64 CPUs.

#include "textflag.h"

// func encodeVarintAsm(buf []byte, value uint64) int
TEXT ·encodeVarintAsm(SB), NOSPLIT, $0-32
    // Load arguments
    MOVD buf_base+0(FP), R0    // R0 = buf pointer
    MOVD value+24(FP), R1      // R1 = value to encode

    // Fast path: value < 64 (1-byte encoding)
    MOVD $64, R2
    CMP  R2, R1
    BGE  two_bytes
    
    // 1-byte encoding: value << 2
    LSL  $2, R1, R1            // R1 = value << 2
    MOVB R1, 0(R0)             // buf[0] = (value << 2)
    MOVD $1, R0                // return 1
    MOVD R0, ret+32(FP)
    RET

two_bytes:
    // Check for 2-byte encoding: value < 16384
    MOVD $16384, R2
    CMP  R2, R1
    BGE  three_bytes
    
    // 2-byte encoding
    // byte 0: 0x01 | ((value >> 8) << 2)
    // byte 1: value & 0xFF
    LSR  $8, R1, R3            // R3 = value >> 8
    LSL  $2, R3, R3            // R3 = (value >> 8) << 2
    ORR  $0x01, R3, R3         // R3 = 0x01 | ((value >> 8) << 2)
    MOVB R3, 0(R0)             // buf[0] = first byte
    
    MOVB R1, 1(R0)             // buf[1] = value & 0xFF
    MOVD $2, R0                // return 2
    MOVD R0, ret+32(FP)
    RET

three_bytes:
    // Check for 3-byte encoding: value < 1073741824
    MOVD $1073741824, R2
    CMP  R2, R1
    BGE  four_bytes
    
    // 3-byte encoding
    // byte 0: 0x02 | ((value >> 16) << 2)
    // byte 1: (value >> 8) & 0xFF
    // byte 2: value & 0xFF
    LSR  $16, R1, R3           // R3 = value >> 16
    LSL  $2, R3, R3            // R3 = (value >> 16) << 2
    ORR  $0x02, R3, R3         // R3 = 0x02 | ((value >> 16) << 2)
    MOVB R3, 0(R0)             // buf[0] = first byte
    
    LSR  $8, R1, R3            // R3 = value >> 8
    MOVB R3, 1(R0)             // buf[1] = (value >> 8) & 0xFF
    
    MOVB R1, 2(R0)             // buf[2] = value & 0xFF
    MOVD $3, R0                // return 3
    MOVD R0, ret+32(FP)
    RET

four_bytes:
    // 4-byte encoding
    // byte 0: 0x03 | ((value >> 24) << 2)
    // byte 1: (value >> 16) & 0xFF
    // byte 2: (value >> 8) & 0xFF
    // byte 3: value & 0xFF
    LSR  $24, R1, R3           // R3 = value >> 24
    LSL  $2, R3, R3            // R3 = (value >> 24) << 2
    ORR  $0x03, R3, R3         // R3 = 0x03 | ((value >> 24) << 2)
    MOVB R3, 0(R0)             // buf[0] = first byte
    
    LSR  $16, R1, R3           // R3 = value >> 16
    MOVB R3, 1(R0)             // buf[1] = (value >> 16) & 0xFF
    
    LSR  $8, R1, R3            // R3 = value >> 8
    MOVB R3, 2(R0)             // buf[2] = (value >> 8) & 0xFF
    
    MOVB R1, 3(R0)             // buf[3] = value & 0xFF
    MOVD $4, R0                // return 4
    MOVD R0, ret+32(FP)
    RET
