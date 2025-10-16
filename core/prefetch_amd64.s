//go:build amd64 && !purego
// +build amd64,!purego

#include "textflag.h"

// func prefetchRead(addr unsafe.Pointer, len int)
// AMD64 PREFETCHT0 instruction for software prefetch hint
TEXT ·prefetchRead(SB), NOSPLIT, $0-16
    MOVQ addr+0(FP), AX     // Load address into AX
    // PREFETCHT0 (AX)        // Prefetch to L1 cache (temporal data)
    BYTE $0x0F; BYTE $0x18; BYTE $0x00  // PREFETCHT0 [RAX]
    RET
