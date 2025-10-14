# BEVE Benchmark Snapshot

> Generated: 2025-10-14T20:11:54Z
> Hostname: runnervmwhb2z
> OS: Linux
> Kernel: Linux 6.11.0-1018-azure #18~24.04.1-Ubuntu SMP Sat Jun 28 04:46:03 UTC 2025
> Architecture: x86_64
> CPU: AMD EPYC 7763 64-Core Processor
> Go: go version go1.25.1 linux/amd64
> Git: b814044

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1382391 | 282356 | 418 |
| Large Payload | Sonic | Unmarshal | 1785293 | 540912 | 564 |
| Large Payload | MessagePack | Unmarshal | 1960935 | 358303 | 6521 |
| Large Payload | CBOR | Unmarshal | 2137647 | 306299 | 6241 |
| Large Payload | JSON | Unmarshal | 4116227 | 566794 | 7444 |
| Small Struct | BEVE ZeroCopy | Marshal | 1962 | 289 | 2 |
| Small Struct | Sonic | Marshal | 2357 | 641 | 3 |
| Small Struct | JSON | Marshal | 5727 | 560 | 2 |
| Small Struct | BEVE | Marshal | 10888 | 2080 | 3 |
| Small Struct | CBOR | Marshal | 14369 | 2833 | 2 |
| Small Struct | MessagePack | Marshal | 15690 | 4224 | 8 |
| Large Payload | BEVE ZeroCopy | Marshal | 587562 | 259 | 2 |
| Large Payload | BEVE | Marshal | 604587 | 188772 | 3 |
| Large Payload | Sonic | Marshal | 708687 | 220765 | 4 |
| Large Payload | CBOR | Marshal | 1022516 | 205716 | 2 |
| Large Payload | MessagePack | Marshal | 1444610 | 526783 | 115 |
| Large Payload | JSON | Marshal | 2388444 | 221871 | 9 |
| Medium Payload | BEVE | Unmarshal | 153848 | 28445 | 59 |
| Medium Payload | Sonic | Unmarshal | 205493 | 47602 | 71 |
| Medium Payload | MessagePack | Unmarshal | 351165 | 42608 | 803 |
| Medium Payload | CBOR | Unmarshal | 458104 | 34752 | 716 |
| Medium Payload | JSON | Unmarshal | 1287432 | 55576 | 759 |
| Small Struct | Sonic | Unmarshal | 9428 | 2370 | 8 |
| Small Struct | BEVE | Unmarshal | 13115 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 13703 | 1664 | 37 |
| Small Struct | CBOR | Unmarshal | 17350 | 1256 | 29 |
| Small Struct | JSON | Unmarshal | 158065 | 7368 | 96 |
| Medium Payload | BEVE ZeroCopy | Marshal | 65699 | 136 | 2 |
| Medium Payload | BEVE | Marshal | 68894 | 18570 | 3 |
| Medium Payload | Sonic | Marshal | 100977 | 22646 | 4 |
| Medium Payload | CBOR | Marshal | 131632 | 20570 | 2 |
| Medium Payload | JSON | Marshal | 227312 | 18743 | 9 |
| Medium Payload | MessagePack | Marshal | 235382 | 65832 | 22 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
