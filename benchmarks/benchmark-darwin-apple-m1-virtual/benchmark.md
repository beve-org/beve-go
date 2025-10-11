# BEVE Benchmark Snapshot

> Generated: 2025-10-11T21:53:11Z
> Hostname: sat12-bq166-0c068cd5-cd23-4577-a829-38877c50c1b0-CABA16475D91.local
> OS: Darwin
> Kernel: Darwin 24.6.0 Darwin Kernel Version 24.6.0: Mon Jul 14 11:30:18 PDT 2025; root:xnu-11417.140.69~1/RELEASE_ARM64_VMAPPLE
> Architecture: arm64
> CPU: Apple M1 (Virtual)
> Go: go version go1.25.1 darwin/arm64
> Git: a8e6011

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE | Marshal | 497.8 | 930 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 917.3 | 289 | 2 |
| Small Struct | CBOR | Marshal | 1016 | 1425 | 2 |
| Small Struct | MessagePack | Marshal | 1656 | 4224 | 8 |
| Small Struct | JSON | Marshal | 1837 | 1296 | 2 |
| Small Struct | Sonic | Marshal | 2990 | 1967 | 3 |
| Small Struct | BEVE | Unmarshal | 773.7 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 1513 | 1722 | 6 |
| Small Struct | MessagePack | Unmarshal | 1538 | 1512 | 34 |
| Small Struct | CBOR | Unmarshal | 4958 | 4328 | 92 |
| Small Struct | JSON | Unmarshal | 12117 | 4168 | 67 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5935 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 7557 | 14576 | 3 |
| Medium Payload | CBOR | Marshal | 16317 | 24664 | 2 |
| Medium Payload | MessagePack | Marshal | 21822 | 65834 | 22 |
| Medium Payload | JSON | Marshal | 29399 | 20797 | 9 |
| Medium Payload | Sonic | Marshal | 47340 | 27588 | 4 |
| Medium Payload | BEVE | Unmarshal | 13445 | 16250 | 59 |
| Medium Payload | Sonic | Unmarshal | 27490 | 34840 | 33 |
| Medium Payload | MessagePack | Unmarshal | 39418 | 40670 | 765 |
| Medium Payload | CBOR | Unmarshal | 53636 | 33721 | 695 |
| Medium Payload | JSON | Unmarshal | 153905 | 47145 | 643 |
| Large Payload | BEVE ZeroCopy | Marshal | 61688 | 303 | 2 |
| Large Payload | BEVE | Marshal | 104991 | 200873 | 3 |
| Large Payload | CBOR | Marshal | 145670 | 197668 | 2 |
| Large Payload | MessagePack | Marshal | 187330 | 526834 | 115 |
| Large Payload | JSON | Marshal | 351637 | 230409 | 9 |
| Large Payload | Sonic | Marshal | 380269 | 213905 | 4 |
| Large Payload | BEVE | Unmarshal | 124112 | 156051 | 419 |
| Large Payload | Sonic | Unmarshal | 234004 | 329671 | 209 |
| Large Payload | MessagePack | Unmarshal | 361269 | 352778 | 6434 |
| Large Payload | CBOR | Unmarshal | 416311 | 317641 | 6477 |
| Large Payload | JSON | Unmarshal | 1496116 | 518393 | 6812 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
