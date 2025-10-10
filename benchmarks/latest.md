# BEVE Benchmark Snapshot

> Generated: 2025-10-10T21:11:29Z
> Hostname: Mac.home
> OS: Darwin
> Kernel: Darwin 24.6.0 Darwin Kernel Version 24.6.0: Mon Jul 14 11:30:30 PDT 2025; root:xnu-11417.140.69~1/RELEASE_ARM64_T6020
> Architecture: arm64
> CPU: Apple M2 Max
> Go: go version go1.25.1 darwin/arm64
> Git: 36d9454

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 716.4 | 145 | 1 |
| Small Struct | BEVE | Marshal | 919.1 | 2197 | 2 |
| Small Struct | CBOR | Marshal | 1483 | 2835 | 2 |
| Small Struct | MessagePack | Marshal | 1504 | 4227 | 8 |
| Small Struct | JSON | Marshal | 1829 | 1426 | 2 |
| Small Struct | Sonic | Marshal | 2120 | 1485 | 3 |
| Small Struct | BEVE | Unmarshal | 548.9 | 312 | 3 |
| Small Struct | Sonic | Unmarshal | 1846 | 3619 | 6 |
| Small Struct | MessagePack | Unmarshal | 2521 | 3226 | 69 |
| Small Struct | CBOR | Unmarshal | 7448 | 3656 | 79 |
| Small Struct | JSON | Unmarshal | 8198 | 3689 | 52 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6126 | 64 | 1 |
| Medium Payload | BEVE | Marshal | 9106 | 21028 | 2 |
| Medium Payload | CBOR | Marshal | 13674 | 20597 | 2 |
| Medium Payload | MessagePack | Marshal | 18097 | 65871 | 22 |
| Medium Payload | JSON | Marshal | 35762 | 24893 | 9 |
| Medium Payload | Sonic | Marshal | 40835 | 24960 | 4 |
| Medium Payload | BEVE | Unmarshal | 14268 | 20072 | 59 |
| Medium Payload | Sonic | Unmarshal | 26564 | 39305 | 31 |
| Medium Payload | MessagePack | Unmarshal | 39127 | 45047 | 849 |
| Medium Payload | CBOR | Unmarshal | 50450 | 38378 | 793 |
| Medium Payload | JSON | Unmarshal | 157302 | 57692 | 764 |
| Large Payload | BEVE ZeroCopy | Marshal | 62132 | 506 | 1 |
| Large Payload | BEVE | Marshal | 85983 | 210656 | 2 |
| Large Payload | CBOR | Marshal | 131275 | 208528 | 2 |
| Large Payload | MessagePack | Marshal | 163574 | 527172 | 115 |
| Large Payload | JSON | Marshal | 330399 | 230551 | 9 |
| Large Payload | Sonic | Marshal | 341281 | 207358 | 4 |
| Large Payload | BEVE | Unmarshal | 121944 | 161728 | 418 |
| Large Payload | Sonic | Unmarshal | 242516 | 362193 | 213 |
| Large Payload | MessagePack | Unmarshal | 302349 | 320241 | 5768 |
| Large Payload | CBOR | Unmarshal | 426994 | 309279 | 6301 |
| Large Payload | JSON | Unmarshal | 1428735 | 493271 | 6432 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
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
