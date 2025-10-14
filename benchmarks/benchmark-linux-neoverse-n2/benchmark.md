# BEVE Benchmark Snapshot

> Generated: 2025-10-14T16:59:30Z
> Hostname: runnervmonn4v
> OS: Linux
> Kernel: Linux 6.11.0-1018-azure #18~24.04.1-Ubuntu SMP Sat Jun 28 04:41:58 UTC 2025
> Architecture: aarch64
> CPU: Neoverse-N2
> Go: go version go1.25.1 linux/arm64
> Git: b63fdae

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 542801 | 147401 | 419 |
| Large Payload | Sonic | Unmarshal | 693451 | 375292 | 211 |
| Large Payload | MessagePack | Unmarshal | 717536 | 372065 | 6820 |
| Large Payload | CBOR | Unmarshal | 1051033 | 318909 | 6495 |
| Large Payload | JSON | Unmarshal | 2452704 | 527085 | 6880 |
| Small Struct | BEVE ZeroCopy | Marshal | 1389 | 288 | 2 |
| Small Struct | MessagePack | Marshal | 2968 | 1152 | 6 |
| Small Struct | BEVE | Marshal | 4820 | 2977 | 3 |
| Small Struct | JSON | Marshal | 7358 | 1936 | 2 |
| Small Struct | CBOR | Marshal | 8574 | 2833 | 2 |
| Small Struct | Sonic | Marshal | 10722 | 2970 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 121527 | 286 | 2 |
| Large Payload | BEVE | Marshal | 423928 | 205823 | 3 |
| Large Payload | CBOR | Marshal | 555848 | 198369 | 2 |
| Large Payload | Sonic | Marshal | 854680 | 220380 | 4 |
| Large Payload | MessagePack | Marshal | 886904 | 526796 | 115 |
| Large Payload | JSON | Marshal | 983670 | 189621 | 9 |
| Medium Payload | BEVE | Unmarshal | 55899 | 17274 | 59 |
| Medium Payload | Sonic | Unmarshal | 81973 | 31833 | 33 |
| Medium Payload | MessagePack | Unmarshal | 160000 | 33726 | 621 |
| Medium Payload | CBOR | Unmarshal | 220440 | 31784 | 657 |
| Medium Payload | JSON | Unmarshal | 652995 | 66872 | 884 |
| Small Struct | BEVE | Unmarshal | 1796 | 888 | 4 |
| Small Struct | Sonic | Unmarshal | 7614 | 2744 | 6 |
| Small Struct | CBOR | Unmarshal | 9795 | 1288 | 30 |
| Small Struct | MessagePack | Unmarshal | 14596 | 4000 | 85 |
| Small Struct | JSON | Unmarshal | 83706 | 7880 | 112 |
| Medium Payload | BEVE ZeroCopy | Marshal | 23581 | 133 | 2 |
| Medium Payload | BEVE | Marshal | 38880 | 21903 | 3 |
| Medium Payload | CBOR | Marshal | 66850 | 21855 | 2 |
| Medium Payload | Sonic | Marshal | 78764 | 19168 | 4 |
| Medium Payload | MessagePack | Marshal | 83895 | 33061 | 21 |
| Medium Payload | JSON | Marshal | 120075 | 24874 | 9 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
