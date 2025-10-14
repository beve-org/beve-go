# BEVE Benchmark Snapshot

> Generated: 2025-10-14T19:59:08Z
> Hostname: runnervmwhb2z
> OS: Linux
> Kernel: Linux 6.11.0-1018-azure #18~24.04.1-Ubuntu SMP Sat Jun 28 04:46:03 UTC 2025
> Architecture: x86_64
> CPU: AMD EPYC 7763 64-Core Processor
> Go: go version go1.25.1 linux/amd64
> Git: b7f817d

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1350640 | 240044 | 418 |
| Large Payload | Sonic | Unmarshal | 1762269 | 554427 | 576 |
| Large Payload | MessagePack | Unmarshal | 1923763 | 357852 | 6534 |
| Large Payload | CBOR | Unmarshal | 2108138 | 308410 | 6286 |
| Large Payload | JSON | Unmarshal | 4004613 | 526754 | 6984 |
| Small Struct | Sonic | Marshal | 3054 | 714 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 3937 | 288 | 2 |
| Small Struct | CBOR | Marshal | 11526 | 3218 | 2 |
| Small Struct | BEVE | Marshal | 14308 | 2593 | 3 |
| Small Struct | MessagePack | Marshal | 15276 | 4224 | 8 |
| Small Struct | JSON | Marshal | 23142 | 1936 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 590386 | 260 | 2 |
| Large Payload | Sonic | Marshal | 709007 | 211700 | 4 |
| Large Payload | BEVE | Marshal | 709758 | 205395 | 3 |
| Large Payload | CBOR | Marshal | 972728 | 189014 | 2 |
| Large Payload | MessagePack | Marshal | 1441579 | 526786 | 115 |
| Large Payload | JSON | Marshal | 2301548 | 222133 | 9 |
| Medium Payload | BEVE | Unmarshal | 140089 | 26108 | 59 |
| Medium Payload | Sonic | Unmarshal | 209747 | 49688 | 66 |
| Medium Payload | MessagePack | Unmarshal | 387466 | 46513 | 879 |
| Medium Payload | CBOR | Unmarshal | 400762 | 31208 | 643 |
| Medium Payload | JSON | Unmarshal | 1351999 | 63224 | 842 |
| Small Struct | Sonic | Unmarshal | 2352 | 478 | 4 |
| Small Struct | BEVE | Unmarshal | 9041 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 14150 | 1312 | 30 |
| Small Struct | CBOR | Unmarshal | 20700 | 1448 | 33 |
| Small Struct | JSON | Unmarshal | 61310 | 2312 | 44 |
| Medium Payload | BEVE ZeroCopy | Marshal | 69556 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 78514 | 18571 | 3 |
| Medium Payload | Sonic | Marshal | 78836 | 22652 | 4 |
| Medium Payload | CBOR | Marshal | 98199 | 19153 | 2 |
| Medium Payload | MessagePack | Marshal | 210489 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 245626 | 20785 | 9 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
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
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
