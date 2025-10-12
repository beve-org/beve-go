# BEVE Benchmark Snapshot

> Generated: 2025-10-12T20:17:15Z
> Hostname: runnervmonn4v
> OS: Linux
> Kernel: Linux 6.11.0-1018-azure #18~24.04.1-Ubuntu SMP Sat Jun 28 04:41:58 UTC 2025
> Architecture: aarch64
> CPU: Neoverse-N2
> Go: go version go1.25.1 linux/arm64
> Git: 386d2ed

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 598554 | 151608 | 418 |
| Large Payload | Sonic | Unmarshal | 656135 | 405046 | 213 |
| Large Payload | MessagePack | Unmarshal | 803322 | 364094 | 6661 |
| Large Payload | CBOR | Unmarshal | 1105018 | 323323 | 6600 |
| Large Payload | JSON | Unmarshal | 2499003 | 512891 | 6746 |
| Small Struct | MessagePack | Marshal | 5458 | 2176 | 7 |
| Small Struct | BEVE ZeroCopy | Marshal | 6288 | 289 | 2 |
| Small Struct | CBOR | Marshal | 8641 | 2833 | 2 |
| Small Struct | Sonic | Marshal | 9595 | 2278 | 3 |
| Small Struct | BEVE | Marshal | 10509 | 2336 | 3 |
| Small Struct | JSON | Marshal | 12542 | 1552 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 290295 | 171 | 2 |
| Large Payload | CBOR | Marshal | 531310 | 189620 | 2 |
| Large Payload | BEVE | Marshal | 610748 | 197573 | 3 |
| Large Payload | MessagePack | Marshal | 868233 | 526795 | 115 |
| Large Payload | Sonic | Marshal | 870737 | 219473 | 4 |
| Large Payload | JSON | Marshal | 1080446 | 230600 | 9 |
| Medium Payload | BEVE | Unmarshal | 76298 | 18539 | 59 |
| Medium Payload | Sonic | Unmarshal | 102842 | 44867 | 33 |
| Medium Payload | MessagePack | Unmarshal | 167966 | 36302 | 673 |
| Medium Payload | CBOR | Unmarshal | 198931 | 29336 | 601 |
| Medium Payload | JSON | Unmarshal | 638076 | 55208 | 711 |
| Small Struct | BEVE | Unmarshal | 3897 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 12741 | 5157 | 6 |
| Small Struct | CBOR | Unmarshal | 14847 | 1928 | 43 |
| Small Struct | MessagePack | Unmarshal | 16533 | 3928 | 83 |
| Small Struct | JSON | Unmarshal | 20786 | 1320 | 28 |
| Medium Payload | BEVE ZeroCopy | Marshal | 42861 | 131 | 2 |
| Medium Payload | CBOR | Marshal | 62502 | 18527 | 2 |
| Medium Payload | BEVE | Marshal | 66518 | 20619 | 3 |
| Medium Payload | Sonic | Marshal | 103861 | 22624 | 4 |
| Medium Payload | MessagePack | Marshal | 114724 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 132704 | 24882 | 9 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
