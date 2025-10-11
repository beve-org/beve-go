# BEVE Benchmark Snapshot

> Generated: 2025-10-11T21:53:46Z
> Hostname: runnervmd3hz3
> OS: MINGW64_NT-10.0-26100
> Kernel: MINGW64_NT-10.0-26100 3.6.4-b9f03e96.x86_64 2025-07-16 18:17 UTC
> Architecture: x86_64
> CPU: unknown
> Go: go version go1.25.1 windows/amd64
> Git: a8e6011

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | Sonic | Marshal | 659.5 | 664 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 816.7 | 288 | 2 |
| Small Struct | BEVE | Marshal | 1504 | 992 | 3 |
| Small Struct | CBOR | Marshal | 2554 | 2192 | 2 |
| Small Struct | MessagePack | Marshal | 3584 | 4224 | 8 |
| Small Struct | JSON | Marshal | 3774 | 1680 | 2 |
| Small Struct | Sonic | Unmarshal | 897.5 | 592 | 5 |
| Small Struct | BEVE | Unmarshal | 1300 | 1080 | 4 |
| Small Struct | CBOR | Unmarshal | 6375 | 2792 | 60 |
| Small Struct | MessagePack | Unmarshal | 7283 | 4032 | 86 |
| Small Struct | JSON | Unmarshal | 14451 | 3784 | 55 |
| Medium Payload | BEVE ZeroCopy | Marshal | 12453 | 148 | 2 |
| Medium Payload | CBOR | Marshal | 19891 | 16472 | 2 |
| Medium Payload | Sonic | Marshal | 20551 | 22344 | 4 |
| Medium Payload | BEVE | Marshal | 20594 | 19381 | 3 |
| Medium Payload | MessagePack | Marshal | 45748 | 65830 | 22 |
| Medium Payload | JSON | Marshal | 53837 | 24881 | 9 |
| Medium Payload | BEVE | Unmarshal | 24530 | 14827 | 59 |
| Medium Payload | Sonic | Unmarshal | 47169 | 53406 | 75 |
| Medium Payload | CBOR | Unmarshal | 61828 | 21272 | 443 |
| Medium Payload | MessagePack | Unmarshal | 71828 | 36782 | 681 |
| Medium Payload | JSON | Unmarshal | 258626 | 56681 | 759 |
| Large Payload | BEVE ZeroCopy | Marshal | 115918 | 479 | 2 |
| Large Payload | Sonic | Marshal | 177427 | 229920 | 4 |
| Large Payload | BEVE | Marshal | 180134 | 196035 | 3 |
| Large Payload | CBOR | Marshal | 227054 | 192218 | 2 |
| Large Payload | MessagePack | Marshal | 307669 | 526770 | 115 |
| Large Payload | JSON | Marshal | 497878 | 215062 | 9 |
| Large Payload | BEVE | Unmarshal | 246187 | 164553 | 419 |
| Large Payload | Sonic | Unmarshal | 462263 | 572507 | 599 |
| Large Payload | MessagePack | Unmarshal | 670365 | 356586 | 6518 |
| Large Payload | CBOR | Unmarshal | 794306 | 298411 | 6086 |
| Large Payload | JSON | Unmarshal | 2432572 | 521944 | 6832 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
