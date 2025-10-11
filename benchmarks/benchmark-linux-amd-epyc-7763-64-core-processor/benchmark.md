# BEVE Benchmark Snapshot

> Generated: 2025-10-11T21:53:22Z
> Hostname: runnervmwhb2z
> OS: Linux
> Kernel: Linux 6.11.0-1018-azure #18~24.04.1-Ubuntu SMP Sat Jun 28 04:46:03 UTC 2025
> Architecture: x86_64
> CPU: AMD EPYC 7763 64-Core Processor
> Go: go version go1.25.1 linux/amd64
> Git: a8e6011

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE | Marshal | 1278 | 1570 | 3 |
| Small Struct | Sonic | Marshal | 1482 | 2037 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 1525 | 289 | 2 |
| Small Struct | CBOR | Marshal | 1937 | 1937 | 2 |
| Small Struct | JSON | Marshal | 3751 | 2192 | 2 |
| Small Struct | MessagePack | Marshal | 3983 | 8322 | 9 |
| Small Struct | BEVE | Unmarshal | 639.5 | 312 | 3 |
| Small Struct | MessagePack | Unmarshal | 1959 | 1024 | 24 |
| Small Struct | CBOR | Unmarshal | 2653 | 1096 | 26 |
| Small Struct | Sonic | Unmarshal | 4180 | 7783 | 10 |
| Small Struct | JSON | Unmarshal | 6706 | 1384 | 30 |
| Medium Payload | Sonic | Marshal | 12401 | 17061 | 4 |
| Medium Payload | BEVE ZeroCopy | Marshal | 12767 | 128 | 2 |
| Medium Payload | BEVE | Marshal | 16204 | 20854 | 3 |
| Medium Payload | CBOR | Marshal | 22923 | 21937 | 2 |
| Medium Payload | MessagePack | Marshal | 35212 | 65839 | 22 |
| Medium Payload | JSON | Marshal | 44650 | 20930 | 9 |
| Medium Payload | BEVE | Unmarshal | 18153 | 13210 | 59 |
| Medium Payload | Sonic | Unmarshal | 36322 | 55672 | 74 |
| Medium Payload | MessagePack | Unmarshal | 58818 | 39455 | 733 |
| Medium Payload | CBOR | Unmarshal | 66762 | 30200 | 621 |
| Medium Payload | JSON | Unmarshal | 195028 | 47832 | 645 |
| Large Payload | BEVE ZeroCopy | Marshal | 121265 | 566 | 2 |
| Large Payload | BEVE | Marshal | 153257 | 197825 | 3 |
| Large Payload | Sonic | Marshal | 165252 | 218088 | 4 |
| Large Payload | CBOR | Marshal | 191663 | 181296 | 2 |
| Large Payload | MessagePack | Marshal | 316452 | 526856 | 115 |
| Large Payload | JSON | Marshal | 432428 | 206182 | 9 |
| Large Payload | BEVE | Unmarshal | 202342 | 157044 | 419 |
| Large Payload | Sonic | Unmarshal | 373085 | 573628 | 593 |
| Large Payload | MessagePack | Unmarshal | 589999 | 382933 | 7052 |
| Large Payload | CBOR | Unmarshal | 715455 | 330683 | 6746 |
| Large Payload | JSON | Unmarshal | 2205067 | 557458 | 7270 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=5000x -run=\^\$ ./...` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=3000x -run=\^\$ ./...` |
