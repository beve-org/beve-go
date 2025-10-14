# BEVE Benchmark Snapshot

> Generated: 2025-10-14T20:22:29Z
> Hostname: runnervmwhb2z
> OS: Linux
> Kernel: Linux 6.11.0-1018-azure #18~24.04.1-Ubuntu SMP Sat Jun 28 04:46:03 UTC 2025
> Architecture: x86_64
> CPU: AMD EPYC 7763 64-Core Processor
> Go: go version go1.25.1 linux/amd64
> Git: e26943b

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1375701 | 293109 | 419 |
| Large Payload | Sonic | Unmarshal | 1718375 | 543695 | 575 |
| Large Payload | MessagePack | Unmarshal | 1721846 | 364352 | 6660 |
| Large Payload | CBOR | Unmarshal | 1949040 | 328173 | 6688 |
| Large Payload | JSON | Unmarshal | 3927704 | 543802 | 7131 |
| Small Struct | BEVE ZeroCopy | Marshal | 1860 | 288 | 2 |
| Small Struct | CBOR | Marshal | 4960 | 560 | 2 |
| Small Struct | Sonic | Marshal | 6061 | 1354 | 3 |
| Small Struct | JSON | Marshal | 9319 | 656 | 2 |
| Small Struct | BEVE | Marshal | 11226 | 3360 | 3 |
| Small Struct | MessagePack | Marshal | 11455 | 2176 | 7 |
| Large Payload | BEVE ZeroCopy | Marshal | 603653 | 259 | 2 |
| Large Payload | BEVE | Marshal | 680299 | 197333 | 3 |
| Large Payload | Sonic | Marshal | 777660 | 229290 | 4 |
| Large Payload | CBOR | Marshal | 1020495 | 197364 | 2 |
| Large Payload | MessagePack | Marshal | 1580960 | 526792 | 115 |
| Large Payload | JSON | Marshal | 2156489 | 205748 | 9 |
| Medium Payload | BEVE | Unmarshal | 177263 | 33374 | 59 |
| Medium Payload | Sonic | Unmarshal | 235522 | 58985 | 76 |
| Medium Payload | MessagePack | Unmarshal | 322609 | 37551 | 703 |
| Medium Payload | CBOR | Unmarshal | 421019 | 32312 | 665 |
| Medium Payload | JSON | Unmarshal | 1242097 | 54009 | 710 |
| Small Struct | BEVE | Unmarshal | 10086 | 1848 | 4 |
| Small Struct | JSON | Unmarshal | 23379 | 872 | 21 |
| Small Struct | CBOR | Unmarshal | 26271 | 1296 | 30 |
| Small Struct | Sonic | Unmarshal | 28479 | 7782 | 10 |
| Small Struct | MessagePack | Unmarshal | 32742 | 3968 | 84 |
| Medium Payload | BEVE | Marshal | 72175 | 16523 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 83152 | 136 | 2 |
| Medium Payload | Sonic | Marshal | 91840 | 22522 | 4 |
| Medium Payload | CBOR | Marshal | 121238 | 21860 | 2 |
| Medium Payload | MessagePack | Marshal | 140485 | 33060 | 21 |
| Medium Payload | JSON | Marshal | 316318 | 24874 | 9 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
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
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
