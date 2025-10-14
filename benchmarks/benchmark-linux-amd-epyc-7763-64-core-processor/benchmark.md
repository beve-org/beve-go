# BEVE Benchmark Snapshot

> Generated: 2025-10-14T16:59:18Z
> Hostname: runnervmwhb2z
> OS: Linux
> Kernel: Linux 6.11.0-1018-azure #18~24.04.1-Ubuntu SMP Sat Jun 28 04:46:03 UTC 2025
> Architecture: x86_64
> CPU: AMD EPYC 7763 64-Core Processor
> Go: go version go1.25.1 linux/amd64
> Git: b63fdae

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1045426 | 152922 | 418 |
| Large Payload | Sonic | Unmarshal | 1676737 | 495222 | 548 |
| Large Payload | MessagePack | Unmarshal | 1677074 | 350141 | 6380 |
| Large Payload | CBOR | Unmarshal | 1964356 | 324491 | 6599 |
| Large Payload | JSON | Unmarshal | 3944807 | 562586 | 7348 |
| Small Struct | BEVE ZeroCopy | Marshal | 7820 | 289 | 2 |
| Small Struct | BEVE | Marshal | 9188 | 2977 | 3 |
| Small Struct | CBOR | Marshal | 10075 | 1936 | 2 |
| Small Struct | Sonic | Marshal | 11076 | 2956 | 3 |
| Small Struct | MessagePack | Marshal | 26600 | 8321 | 9 |
| Small Struct | JSON | Marshal | 33421 | 2833 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 536215 | 286 | 2 |
| Large Payload | BEVE | Marshal | 665717 | 188772 | 3 |
| Large Payload | Sonic | Marshal | 769631 | 221328 | 4 |
| Large Payload | CBOR | Marshal | 1013375 | 198100 | 2 |
| Large Payload | MessagePack | Marshal | 1443824 | 526782 | 115 |
| Large Payload | JSON | Marshal | 2039922 | 205904 | 9 |
| Medium Payload | BEVE | Unmarshal | 113443 | 15723 | 59 |
| Medium Payload | Sonic | Unmarshal | 216689 | 55827 | 68 |
| Medium Payload | MessagePack | Unmarshal | 324616 | 36783 | 684 |
| Medium Payload | CBOR | Unmarshal | 398394 | 31544 | 647 |
| Medium Payload | JSON | Unmarshal | 1096385 | 50776 | 644 |
| Small Struct | BEVE | Unmarshal | 9336 | 1848 | 4 |
| Small Struct | MessagePack | Unmarshal | 15246 | 2056 | 44 |
| Small Struct | Sonic | Unmarshal | 22308 | 7025 | 10 |
| Small Struct | CBOR | Unmarshal | 46973 | 4640 | 98 |
| Small Struct | JSON | Unmarshal | 151803 | 7792 | 109 |
| Medium Payload | BEVE ZeroCopy | Marshal | 58998 | 136 | 2 |
| Medium Payload | BEVE | Marshal | 85008 | 20622 | 3 |
| Medium Payload | CBOR | Marshal | 95062 | 14412 | 2 |
| Medium Payload | Sonic | Marshal | 106629 | 25475 | 4 |
| Medium Payload | MessagePack | Marshal | 205933 | 65833 | 22 |
| Medium Payload | JSON | Marshal | 278643 | 24888 | 9 |

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
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
