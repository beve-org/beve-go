# BEVE Benchmark Snapshot

> Generated: 2025-10-14T20:12:00Z
> Hostname: runnervmrcw8b
> OS: Linux
> Kernel: Linux 6.11.0-1018-azure #18~24.04.1-Ubuntu SMP Sat Jun 28 04:41:58 UTC 2025
> Architecture: aarch64
> CPU: Neoverse-N2
> Go: go version go1.25.1 linux/arm64
> Git: b814044

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | MessagePack | Unmarshal | 708161 | 355902 | 6494 |
| Large Payload | BEVE | Unmarshal | 758134 | 285493 | 417 |
| Large Payload | Sonic | Unmarshal | 762498 | 372324 | 209 |
| Large Payload | CBOR | Unmarshal | 1046717 | 324556 | 6608 |
| Large Payload | JSON | Unmarshal | 2647401 | 573884 | 7450 |
| Small Struct | Sonic | Marshal | 1110 | 420 | 3 |
| Small Struct | BEVE | Marshal | 1622 | 2594 | 3 |
| Small Struct | CBOR | Marshal | 2384 | 1296 | 2 |
| Small Struct | MessagePack | Marshal | 3255 | 1152 | 6 |
| Small Struct | BEVE ZeroCopy | Marshal | 4307 | 289 | 2 |
| Small Struct | JSON | Marshal | 8403 | 1168 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 199601 | 259 | 2 |
| Large Payload | BEVE | Marshal | 376586 | 188907 | 3 |
| Large Payload | CBOR | Marshal | 553098 | 197952 | 2 |
| Large Payload | MessagePack | Marshal | 842551 | 526798 | 115 |
| Large Payload | Sonic | Marshal | 935884 | 230658 | 4 |
| Large Payload | JSON | Marshal | 1054072 | 205745 | 9 |
| Medium Payload | Sonic | Unmarshal | 87100 | 42709 | 33 |
| Medium Payload | BEVE | Unmarshal | 87449 | 31933 | 59 |
| Medium Payload | MessagePack | Unmarshal | 145992 | 31822 | 583 |
| Medium Payload | CBOR | Unmarshal | 224384 | 35928 | 738 |
| Medium Payload | JSON | Unmarshal | 572304 | 48889 | 652 |
| Small Struct | BEVE | Unmarshal | 1444 | 600 | 4 |
| Small Struct | Sonic | Unmarshal | 8452 | 3075 | 6 |
| Small Struct | CBOR | Unmarshal | 9707 | 4360 | 93 |
| Small Struct | MessagePack | Unmarshal | 19796 | 5153 | 105 |
| Small Struct | JSON | Unmarshal | 66356 | 7304 | 94 |
| Medium Payload | BEVE ZeroCopy | Marshal | 13513 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 50021 | 24721 | 3 |
| Medium Payload | CBOR | Marshal | 76166 | 24661 | 2 |
| Medium Payload | JSON | Marshal | 97949 | 19376 | 9 |
| Medium Payload | MessagePack | Marshal | 107752 | 65832 | 22 |
| Medium Payload | Sonic | Marshal | 124409 | 28402 | 4 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
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
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
