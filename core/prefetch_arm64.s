//go:build arm64 && !purego
// +build arm64,!purego

#include "textflag.h"

// func prefetchRead(addr unsafe.Pointer, len int)
// ARM64 NEON PRFM instruction for software prefetch hint
TEXT ·prefetchRead(SB), NOSPLIT, $0-16
    MOVD addr+0(FP), R0     // Load address into R0
    // PRFM PLDL1KEEP, [R0]   // Prefetch for read, L1 cache, keep
    // ARM64 assembly: Prefetch to L1 data cache
    WORD $0xf9800000        // PRFM PLDL1KEEP, [X0]
    RET
