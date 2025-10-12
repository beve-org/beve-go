# BEVE Benchmark Snapshot

> Generated: 2025-10-12T20:16:59Z
> Hostname: sat12-dp150-77b6d7e7-244f-49f7-95b1-522dcf14b9bc-8ABD71683F5F.local
> OS: Darwin
> Kernel: Darwin 24.6.0 Darwin Kernel Version 24.6.0: Mon Jul 14 11:30:18 PDT 2025; root:xnu-11417.140.69~1/RELEASE_ARM64_VMAPPLE
> Architecture: arm64
> CPU: Apple M1 (Virtual)
> Go: go version go1.25.1 darwin/arm64
> Git: 386d2ed

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1111419 | 152481 | 418 |
| Large Payload | Sonic | Unmarshal | 1759607 | 346360 | 211 |
| Large Payload | MessagePack | Unmarshal | 2063338 | 347429 | 6333 |
| Large Payload | CBOR | Unmarshal | 2238993 | 335930 | 6848 |
| Large Payload | JSON | Unmarshal | 3578360 | 496957 | 6545 |
| Small Struct | CBOR | Marshal | 520.4 | 432 | 2 |
| Small Struct | BEVE | Marshal | 9788 | 496 | 3 |
| Small Struct | MessagePack | Marshal | 14240 | 1152 | 6 |
| Small Struct | BEVE ZeroCopy | Marshal | 30885 | 288 | 2 |
| Small Struct | Sonic | Marshal | 37264 | 1731 | 3 |
| Small Struct | JSON | Marshal | 46846 | 1936 | 2 |
| Large Payload | BEVE | Marshal | 1000146 | 197420 | 3 |
| Large Payload | CBOR | Marshal | 1011417 | 189393 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 1163388 | 180 | 2 |
| Large Payload | MessagePack | Marshal | 1432618 | 526800 | 115 |
| Large Payload | JSON | Marshal | 1893788 | 214445 | 9 |
| Large Payload | Sonic | Marshal | 2090595 | 234651 | 4 |
| Medium Payload | BEVE | Unmarshal | 141852 | 16106 | 59 |
| Medium Payload | Sonic | Unmarshal | 224419 | 33121 | 33 |
| Medium Payload | MessagePack | Unmarshal | 361742 | 35661 | 656 |
| Medium Payload | CBOR | Unmarshal | 414926 | 33160 | 679 |
| Medium Payload | JSON | Unmarshal | 1226807 | 56056 | 750 |
| Small Struct | BEVE | Unmarshal | 4777 | 552 | 4 |
| Small Struct | MessagePack | Unmarshal | 33550 | 2080 | 45 |
| Small Struct | Sonic | Unmarshal | 33786 | 2753 | 6 |
| Small Struct | CBOR | Unmarshal | 46226 | 2472 | 54 |
| Small Struct | JSON | Unmarshal | 85233 | 1256 | 26 |
| Medium Payload | BEVE | Marshal | 171407 | 16523 | 3 |
| Medium Payload | CBOR | Marshal | 186146 | 21855 | 2 |
| Medium Payload | BEVE ZeroCopy | Marshal | 236037 | 132 | 2 |
| Medium Payload | MessagePack | Marshal | 282811 | 65831 | 22 |
| Medium Payload | JSON | Marshal | 321932 | 22081 | 9 |
| Medium Payload | Sonic | Marshal | 349749 | 19578 | 4 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=100000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=50000x -run=\^\$ ./...` |
