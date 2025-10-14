# BEVE Benchmark Snapshot

> Generated: 2025-10-14T19:59:30Z
> Hostname: runnervmd3hz3
> OS: MINGW64_NT-10.0-26100
> Kernel: MINGW64_NT-10.0-26100 3.6.4-b9f03e96.x86_64 2025-07-16 18:17 UTC
> Architecture: x86_64
> CPU: unknown
> Go: go version go1.25.1 windows/amd64
> Git: b7f817d

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1157606 | 225754 | 418 |
| Large Payload | MessagePack | Unmarshal | 1377340 | 375800 | 6898 |
| Large Payload | Sonic | Unmarshal | 1440202 | 542780 | 562 |
| Large Payload | CBOR | Unmarshal | 1997668 | 330188 | 6728 |
| Large Payload | JSON | Unmarshal | 3539630 | 530085 | 6960 |
| Small Struct | BEVE ZeroCopy | Marshal | 2409 | 288 | 2 |
| Small Struct | JSON | Marshal | 5374 | 656 | 2 |
| Small Struct | BEVE | Marshal | 8759 | 2080 | 3 |
| Small Struct | Sonic | Marshal | 10929 | 1995 | 3 |
| Small Struct | CBOR | Marshal | 20776 | 3217 | 2 |
| Small Struct | MessagePack | Marshal | 28196 | 4224 | 8 |
| Large Payload | BEVE ZeroCopy | Marshal | 229896 | 259 | 2 |
| Large Payload | Sonic | Marshal | 632861 | 217564 | 4 |
| Large Payload | BEVE | Marshal | 655694 | 180814 | 3 |
| Large Payload | CBOR | Marshal | 799748 | 198056 | 2 |
| Large Payload | MessagePack | Marshal | 1087105 | 526763 | 115 |
| Large Payload | JSON | Marshal | 1385034 | 215245 | 9 |
| Medium Payload | BEVE | Unmarshal | 78669 | 21819 | 59 |
| Medium Payload | Sonic | Unmarshal | 167361 | 48582 | 69 |
| Medium Payload | MessagePack | Unmarshal | 261510 | 30925 | 566 |
| Medium Payload | CBOR | Unmarshal | 409771 | 42504 | 877 |
| Medium Payload | JSON | Unmarshal | 554760 | 41816 | 529 |
| Small Struct | MessagePack | Unmarshal | 7332 | 1088 | 25 |
| Small Struct | Sonic | Unmarshal | 11219 | 3632 | 9 |
| Small Struct | BEVE | Unmarshal | 17098 | 3512 | 4 |
| Small Struct | CBOR | Unmarshal | 39188 | 5088 | 104 |
| Small Struct | JSON | Unmarshal | 52844 | 4264 | 70 |
| Medium Payload | BEVE ZeroCopy | Marshal | 23547 | 133 | 2 |
| Medium Payload | BEVE | Marshal | 52054 | 20625 | 3 |
| Medium Payload | Sonic | Marshal | 81818 | 25254 | 4 |
| Medium Payload | CBOR | Marshal | 108336 | 24675 | 2 |
| Medium Payload | MessagePack | Marshal | 128358 | 65829 | 22 |
| Medium Payload | JSON | Marshal | 159174 | 24910 | 9 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
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
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
