// Assembly implementation of varint encoding for AMD64 (x86-64).
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
// Register usage conventions (System V ABI):
//   - AX (RAX): Return value, scratch
//   - BX (RBX): Callee-saved
//   - CX (RCX): Arg 2, scratch
//   - DX (RDX): Arg 3, scratch
//   - SI (RSI): Arg 1 (buf pointer)
//   - DI (RDI): Arg 0 (value)
//   - R8-R15: Various uses
//
// Function signature:
//   func encodeVarintAsm(buf []byte, value uint64) int
//
// Args:
//   buf (SI): Pointer to output buffer (base address)
//   value (DI): uint64 value to encode
//
// Returns:
//   AX: Number of bytes written (1, 2, 3, or 4)
//
// Note: This is a template/stub. Actual assembly optimization requires
// careful benchmarking and may not provide significant gains over
// compiler-optimized Go code on modern CPUs.

#include "textflag.h"

// func encodeVarintAsm(buf []byte, value uint64) int
TEXT ·encodeVarintAsm(SB), NOSPLIT, $0-32
    // Load arguments
    MOVQ buf_base+0(FP), SI   // SI = buf pointer
    MOVQ value+24(FP), DI      // DI = value to encode

    // Fast path: value < 64 (1-byte encoding)
    CMPQ DI, $64
    JGE  two_bytes
    
    // 1-byte encoding: value << 2
    SHLQ $2, DI                // DI = value << 2
    MOVB DI, 0(SI)             // buf[0] = (value << 2)
    MOVQ $1, ret+32(FP)        // return 1
    RET

two_bytes:
    // Check for 2-byte encoding: value < 16384
    CMPQ DI, $16384
    JGE  three_bytes
    
    // 2-byte encoding
    // byte 0: 0x01 | ((value >> 8) << 2)
    // byte 1: value & 0xFF
    MOVQ DI, AX                // AX = value
    SHRQ $8, AX                // AX = value >> 8
    SHLQ $2, AX                // AX = (value >> 8) << 2
    ORQ  $0x01, AX             // AX = 0x01 | ((value >> 8) << 2)
    MOVB AX, 0(SI)             // buf[0] = first byte
    
    MOVB DI, 1(SI)             // buf[1] = value & 0xFF
    MOVQ $2, ret+32(FP)        // return 2
    RET

three_bytes:
    // Check for 3-byte encoding: value < 1073741824
    CMPQ DI, $1073741824
    JGE  four_bytes
    
    // 3-byte encoding
    // byte 0: 0x02 | ((value >> 16) << 2)
    // byte 1: (value >> 8) & 0xFF
    // byte 2: value & 0xFF
    MOVQ DI, AX                // AX = value
    SHRQ $16, AX               // AX = value >> 16
    SHLQ $2, AX                // AX = (value >> 16) << 2
    ORQ  $0x02, AX             // AX = 0x02 | ((value >> 16) << 2)
    MOVB AX, 0(SI)             // buf[0] = first byte
    
    MOVQ DI, AX                // AX = value
    SHRQ $8, AX                // AX = value >> 8
    MOVB AX, 1(SI)             // buf[1] = (value >> 8) & 0xFF
    
    MOVB DI, 2(SI)             // buf[2] = value & 0xFF
    MOVQ $3, ret+32(FP)        // return 3
    RET

four_bytes:
    // 4-byte encoding
    // byte 0: 0x03 | ((value >> 24) << 2)
    // byte 1: (value >> 16) & 0xFF
    // byte 2: (value >> 8) & 0xFF
    // byte 3: value & 0xFF
    MOVQ DI, AX                // AX = value
    SHRQ $24, AX               // AX = value >> 24
    SHLQ $2, AX                // AX = (value >> 24) << 2
    ORQ  $0x03, AX             // AX = 0x03 | ((value >> 24) << 2)
    MOVB AX, 0(SI)             // buf[0] = first byte
    
    MOVQ DI, AX                // AX = value
    SHRQ $16, AX               // AX = value >> 16
    MOVB AX, 1(SI)             // buf[1] = (value >> 16) & 0xFF
    
    MOVQ DI, AX                // AX = value
    SHRQ $8, AX                // AX = value >> 8
    MOVB AX, 2(SI)             // buf[2] = (value >> 8) & 0xFF
    
    MOVB DI, 3(SI)             // buf[3] = value & 0xFF
    MOVQ $4, ret+32(FP)        // return 4
    RET
