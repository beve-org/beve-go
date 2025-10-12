# BEVE Benchmark Snapshot

> Generated: 2025-10-12T19:14:05Z
> Hostname: sat12-bq147-bca79e1f-d7d0-4a1c-8b49-93ec099e4970-1E918F871B13.local
> OS: Darwin
> Kernel: Darwin 24.6.0 Darwin Kernel Version 24.6.0: Mon Jul 14 11:30:18 PDT 2025; root:xnu-11417.140.69~1/RELEASE_ARM64_VMAPPLE
> Architecture: arm64
> CPU: Apple M1 (Virtual)
> Go: go version go1.25.1 darwin/arm64
> Git: 458517e

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 235 | 288 | 2 |
| Small Struct | JSON | Marshal | 583.1 | 464 | 2 |
| Small Struct | CBOR | Marshal | 607.6 | 848 | 2 |
| Small Struct | MessagePack | Marshal | 698.8 | 1152 | 6 |
| Small Struct | BEVE | Marshal | 1747 | 2977 | 3 |
| Small Struct | Sonic | Marshal | 2899 | 1469 | 3 |
| Small Struct | BEVE | Unmarshal | 897 | 1464 | 4 |
| Small Struct | CBOR | Unmarshal | 1588 | 952 | 23 |
| Small Struct | Sonic | Unmarshal | 1796 | 1722 | 6 |
| Small Struct | MessagePack | Unmarshal | 3881 | 3880 | 81 |
| Small Struct | JSON | Unmarshal | 5276 | 1448 | 32 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7324 | 128 | 2 |
| Medium Payload | BEVE | Marshal | 10676 | 16514 | 3 |
| Medium Payload | CBOR | Marshal | 20412 | 21847 | 2 |
| Medium Payload | MessagePack | Marshal | 32239 | 65833 | 22 |
| Medium Payload | Sonic | Marshal | 43515 | 22008 | 4 |
| Medium Payload | JSON | Marshal | 50740 | 27569 | 9 |
| Medium Payload | BEVE | Unmarshal | 18767 | 15850 | 59 |
| Medium Payload | Sonic | Unmarshal | 34280 | 35554 | 33 |
| Medium Payload | MessagePack | Unmarshal | 47895 | 32094 | 589 |
| Medium Payload | CBOR | Unmarshal | 56783 | 30536 | 621 |
| Medium Payload | JSON | Unmarshal | 205944 | 50776 | 666 |
| Large Payload | BEVE ZeroCopy | Marshal | 116291 | 478 | 2 |
| Large Payload | BEVE | Marshal | 172184 | 188927 | 3 |
| Large Payload | MessagePack | Marshal | 223009 | 526832 | 115 |
| Large Payload | CBOR | Marshal | 238545 | 189800 | 2 |
| Large Payload | JSON | Marshal | 444829 | 214018 | 9 |
| Large Payload | Sonic | Marshal | 461790 | 206869 | 4 |
| Large Payload | BEVE | Unmarshal | 160292 | 154482 | 418 |
| Large Payload | Sonic | Unmarshal | 313669 | 335231 | 209 |
| Large Payload | MessagePack | Unmarshal | 556732 | 347881 | 6342 |
| Large Payload | CBOR | Unmarshal | 665491 | 336462 | 6868 |
| Large Payload | JSON | Unmarshal | 1933466 | 545308 | 7076 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
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
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
