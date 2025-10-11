# BEVE Benchmark Snapshot

> Generated: 2025-10-11T21:53:23Z
> Hostname: runnervmonn4v
> OS: Linux
> Kernel: Linux 6.11.0-1018-azure #18~24.04.1-Ubuntu SMP Sat Jun 28 04:41:58 UTC 2025
> Architecture: aarch64
> CPU: Neoverse-N2
> Go: go version go1.25.1 linux/arm64
> Git: a8e6011

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 516.2 | 288 | 2 |
| Small Struct | BEVE | Marshal | 1483 | 2083 | 3 |
| Small Struct | CBOR | Marshal | 1566 | 1553 | 2 |
| Small Struct | JSON | Marshal | 3600 | 2193 | 2 |
| Small Struct | Sonic | Marshal | 3658 | 2905 | 3 |
| Small Struct | MessagePack | Marshal | 3881 | 8322 | 9 |
| Small Struct | BEVE | Unmarshal | 712.4 | 408 | 4 |
| Small Struct | Sonic | Unmarshal | 1411 | 1791 | 6 |
| Small Struct | JSON | Unmarshal | 4339 | 896 | 22 |
| Small Struct | MessagePack | Unmarshal | 5779 | 4641 | 97 |
| Small Struct | CBOR | Unmarshal | 8050 | 5192 | 107 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9124 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 12945 | 19428 | 3 |
| Medium Payload | CBOR | Marshal | 16830 | 18519 | 2 |
| Medium Payload | MessagePack | Marshal | 22428 | 33063 | 21 |
| Medium Payload | Sonic | Marshal | 28603 | 20925 | 4 |
| Medium Payload | JSON | Marshal | 43105 | 24895 | 9 |
| Medium Payload | BEVE | Unmarshal | 19598 | 15484 | 59 |
| Medium Payload | Sonic | Unmarshal | 27125 | 34322 | 33 |
| Medium Payload | MessagePack | Unmarshal | 62430 | 47153 | 883 |
| Medium Payload | CBOR | Unmarshal | 63470 | 31177 | 646 |
| Medium Payload | JSON | Unmarshal | 186478 | 51209 | 673 |
| Large Payload | BEVE ZeroCopy | Marshal | 89188 | 479 | 2 |
| Large Payload | BEVE | Marshal | 141350 | 205295 | 3 |
| Large Payload | CBOR | Marshal | 191275 | 205891 | 2 |
| Large Payload | MessagePack | Marshal | 270398 | 526866 | 115 |
| Large Payload | Sonic | Marshal | 314852 | 224816 | 4 |
| Large Payload | JSON | Marshal | 365578 | 197638 | 9 |
| Large Payload | BEVE | Unmarshal | 181946 | 152987 | 419 |
| Large Payload | Sonic | Unmarshal | 286168 | 384122 | 213 |
| Large Payload | MessagePack | Unmarshal | 504028 | 351766 | 6401 |
| Large Payload | CBOR | Unmarshal | 639972 | 323851 | 6593 |
| Large Payload | JSON | Unmarshal | 1997108 | 546253 | 7164 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
