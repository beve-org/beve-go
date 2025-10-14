# BEVE Benchmark Snapshot

> Generated: 2025-10-14T20:12:01Z
> Hostname: runnervmd3hz3
> OS: MINGW64_NT-10.0-26100
> Kernel: MINGW64_NT-10.0-26100 3.6.4-b9f03e96.x86_64 2025-07-16 18:17 UTC
> Architecture: x86_64
> CPU: unknown
> Go: go version go1.25.1 windows/amd64
> Git: b814044

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1334329 | 262017 | 418 |
| Large Payload | Sonic | Unmarshal | 1412887 | 523827 | 568 |
| Large Payload | MessagePack | Unmarshal | 1479432 | 342544 | 6230 |
| Large Payload | CBOR | Unmarshal | 1848957 | 308635 | 6285 |
| Large Payload | JSON | Unmarshal | 3851570 | 541838 | 7137 |
| Small Struct | BEVE ZeroCopy | Marshal | 3743 | 289 | 2 |
| Small Struct | JSON | Marshal | 10236 | 1296 | 2 |
| Small Struct | BEVE | Marshal | 10292 | 1568 | 3 |
| Small Struct | CBOR | Marshal | 12309 | 3217 | 2 |
| Small Struct | Sonic | Marshal | 12889 | 2539 | 3 |
| Small Struct | MessagePack | Marshal | 16894 | 4224 | 8 |
| Large Payload | BEVE ZeroCopy | Marshal | 228394 | 259 | 2 |
| Large Payload | Sonic | Marshal | 577903 | 209714 | 4 |
| Large Payload | CBOR | Marshal | 772214 | 181553 | 2 |
| Large Payload | BEVE | Marshal | 985823 | 205073 | 3 |
| Large Payload | MessagePack | Marshal | 1130270 | 526759 | 115 |
| Large Payload | JSON | Marshal | 1433619 | 215191 | 9 |
| Medium Payload | BEVE | Unmarshal | 90366 | 29212 | 59 |
| Medium Payload | MessagePack | Unmarshal | 215159 | 30619 | 555 |
| Medium Payload | Sonic | Unmarshal | 229713 | 47756 | 68 |
| Medium Payload | CBOR | Unmarshal | 308817 | 33432 | 682 |
| Medium Payload | JSON | Unmarshal | 751920 | 57961 | 789 |
| Small Struct | BEVE | Unmarshal | 13315 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 13987 | 2432 | 52 |
| Small Struct | JSON | Unmarshal | 20999 | 2184 | 40 |
| Small Struct | CBOR | Unmarshal | 23640 | 3944 | 84 |
| Small Struct | Sonic | Unmarshal | 27687 | 4173 | 9 |
| Medium Payload | BEVE ZeroCopy | Marshal | 24631 | 136 | 2 |
| Medium Payload | Sonic | Marshal | 81018 | 22272 | 4 |
| Medium Payload | CBOR | Marshal | 81969 | 18517 | 2 |
| Medium Payload | BEVE | Marshal | 104659 | 21908 | 3 |
| Medium Payload | MessagePack | Marshal | 151562 | 65828 | 22 |
| Medium Payload | JSON | Marshal | 169458 | 22074 | 9 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
