# BEVE Benchmark Snapshot

> Generated: 2025-10-14T17:00:41Z
> Hostname: sat12-bq147-513501cd-1e5c-48f3-b5a8-b98bd3df4298-BEE2502D1074.local
> OS: Darwin
> Kernel: Darwin 24.6.0 Darwin Kernel Version 24.6.0: Mon Jul 14 11:30:18 PDT 2025; root:xnu-11417.140.69~1/RELEASE_ARM64_VMAPPLE
> Architecture: arm64
> CPU: Apple M1 (Virtual)
> Go: go version go1.25.1 darwin/arm64
> Git: b63fdae

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1061839 | 151600 | 419 |
| Large Payload | Sonic | Unmarshal | 1707611 | 366316 | 213 |
| Large Payload | MessagePack | Unmarshal | 2029736 | 345716 | 6291 |
| Large Payload | CBOR | Unmarshal | 2203998 | 319515 | 6510 |
| Large Payload | JSON | Unmarshal | 3535936 | 539092 | 6972 |
| Small Struct | CBOR | Marshal | 365.5 | 288 | 2 |
| Small Struct | BEVE | Marshal | 814.4 | 1696 | 3 |
| Small Struct | Sonic | Marshal | 846.2 | 382 | 3 |
| Small Struct | MessagePack | Marshal | 19027 | 8321 | 9 |
| Small Struct | BEVE ZeroCopy | Marshal | 21717 | 288 | 2 |
| Small Struct | JSON | Marshal | 41366 | 1296 | 2 |
| Large Payload | BEVE | Marshal | 858904 | 189042 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 889719 | 233 | 2 |
| Large Payload | CBOR | Marshal | 1277962 | 189823 | 2 |
| Large Payload | MessagePack | Marshal | 1408925 | 526802 | 115 |
| Large Payload | JSON | Marshal | 1845362 | 213728 | 9 |
| Large Payload | Sonic | Marshal | 2089784 | 216949 | 4 |
| Medium Payload | BEVE | Unmarshal | 140941 | 15418 | 59 |
| Medium Payload | Sonic | Unmarshal | 263512 | 35374 | 33 |
| Medium Payload | MessagePack | Unmarshal | 508849 | 37982 | 710 |
| Medium Payload | CBOR | Unmarshal | 555993 | 30552 | 625 |
| Medium Payload | JSON | Unmarshal | 1075877 | 54616 | 691 |
| Small Struct | BEVE | Unmarshal | 1614 | 1592 | 4 |
| Small Struct | Sonic | Unmarshal | 22522 | 2481 | 6 |
| Small Struct | MessagePack | Unmarshal | 44084 | 784 | 19 |
| Small Struct | CBOR | Unmarshal | 78265 | 2792 | 60 |
| Small Struct | JSON | Unmarshal | 371677 | 7816 | 110 |
| Medium Payload | BEVE | Marshal | 122710 | 20619 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 126351 | 134 | 2 |
| Medium Payload | CBOR | Marshal | 165435 | 24690 | 2 |
| Medium Payload | MessagePack | Marshal | 273059 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 297798 | 18749 | 9 |
| Medium Payload | Sonic | Marshal | 490775 | 25262 | 4 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
