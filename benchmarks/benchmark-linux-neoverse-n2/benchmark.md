# BEVE Benchmark Snapshot

> Generated: 2025-10-14T20:22:38Z
> Hostname: runnervmrcw8b
> OS: Linux
> Kernel: Linux 6.11.0-1018-azure #18~24.04.1-Ubuntu SMP Sat Jun 28 04:41:58 UTC 2025
> Architecture: aarch64
> CPU: Neoverse-N2
> Go: go version go1.25.1 linux/arm64
> Git: e26943b

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 742424 | 267730 | 418 |
| Large Payload | Sonic | Unmarshal | 758286 | 417122 | 213 |
| Large Payload | MessagePack | Unmarshal | 803998 | 340218 | 6179 |
| Large Payload | CBOR | Unmarshal | 1074256 | 322443 | 6568 |
| Large Payload | JSON | Unmarshal | 2550255 | 553475 | 7252 |
| Small Struct | BEVE ZeroCopy | Marshal | 2996 | 288 | 2 |
| Small Struct | BEVE | Marshal | 3454 | 1312 | 3 |
| Small Struct | CBOR | Marshal | 6513 | 2192 | 2 |
| Small Struct | JSON | Marshal | 7179 | 1296 | 2 |
| Small Struct | MessagePack | Marshal | 9688 | 4224 | 8 |
| Small Struct | Sonic | Marshal | 10525 | 2265 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 214027 | 207 | 2 |
| Large Payload | BEVE | Marshal | 393565 | 188990 | 3 |
| Large Payload | CBOR | Marshal | 516309 | 180931 | 2 |
| Large Payload | MessagePack | Marshal | 846074 | 526794 | 115 |
| Large Payload | Sonic | Marshal | 900091 | 220639 | 4 |
| Large Payload | JSON | Marshal | 1063017 | 213988 | 9 |
| Medium Payload | Sonic | Unmarshal | 78176 | 30272 | 33 |
| Medium Payload | BEVE | Unmarshal | 88403 | 30782 | 59 |
| Medium Payload | MessagePack | Unmarshal | 190834 | 43264 | 813 |
| Medium Payload | CBOR | Unmarshal | 252954 | 39096 | 804 |
| Medium Payload | JSON | Unmarshal | 465709 | 39448 | 509 |
| Small Struct | Sonic | Unmarshal | 2825 | 622 | 6 |
| Small Struct | BEVE | Unmarshal | 5835 | 2104 | 4 |
| Small Struct | MessagePack | Unmarshal | 7329 | 1568 | 35 |
| Small Struct | CBOR | Unmarshal | 11703 | 1352 | 31 |
| Small Struct | JSON | Unmarshal | 43486 | 4200 | 68 |
| Medium Payload | BEVE ZeroCopy | Marshal | 25963 | 131 | 2 |
| Medium Payload | BEVE | Marshal | 47418 | 21898 | 3 |
| Medium Payload | CBOR | Marshal | 73575 | 27363 | 2 |
| Medium Payload | Sonic | Marshal | 90093 | 21314 | 4 |
| Medium Payload | JSON | Marshal | 102250 | 18733 | 9 |
| Medium Payload | MessagePack | Marshal | 118936 | 65833 | 22 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
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
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
