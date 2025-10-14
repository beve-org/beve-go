# BEVE Benchmark Snapshot

> Generated: 2025-10-14T16:59:53Z
> Hostname: runnervmd3hz3
> OS: MINGW64_NT-10.0-26100
> Kernel: MINGW64_NT-10.0-26100 3.6.4-b9f03e96.x86_64 2025-07-16 18:17 UTC
> Architecture: x86_64
> CPU: unknown
> Go: go version go1.25.1 windows/amd64
> Git: b63fdae

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1015118 | 155151 | 416 |
| Large Payload | MessagePack | Unmarshal | 1194394 | 339206 | 6154 |
| Large Payload | Sonic | Unmarshal | 1576098 | 538430 | 574 |
| Large Payload | CBOR | Unmarshal | 1681541 | 327611 | 6695 |
| Large Payload | JSON | Unmarshal | 3708514 | 528093 | 6854 |
| Small Struct | BEVE ZeroCopy | Marshal | 2848 | 288 | 2 |
| Small Struct | Sonic | Marshal | 4694 | 521 | 3 |
| Small Struct | BEVE | Marshal | 10933 | 2336 | 3 |
| Small Struct | CBOR | Marshal | 13165 | 2833 | 2 |
| Small Struct | MessagePack | Marshal | 14331 | 4224 | 8 |
| Small Struct | JSON | Marshal | 17944 | 2834 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 276709 | 207 | 2 |
| Large Payload | Sonic | Marshal | 707293 | 209777 | 4 |
| Large Payload | CBOR | Marshal | 807108 | 198084 | 2 |
| Large Payload | BEVE | Marshal | 842610 | 188978 | 3 |
| Large Payload | MessagePack | Marshal | 1073250 | 526759 | 115 |
| Large Payload | JSON | Marshal | 1498046 | 215611 | 9 |
| Medium Payload | BEVE | Unmarshal | 200945 | 17466 | 59 |
| Medium Payload | MessagePack | Unmarshal | 344865 | 38635 | 719 |
| Medium Payload | CBOR | Unmarshal | 357647 | 27672 | 566 |
| Medium Payload | Sonic | Unmarshal | 391507 | 58509 | 76 |
| Medium Payload | JSON | Unmarshal | 727769 | 47992 | 641 |
| Small Struct | BEVE | Unmarshal | 3843 | 424 | 4 |
| Small Struct | MessagePack | Unmarshal | 16371 | 3904 | 82 |
| Small Struct | Sonic | Unmarshal | 17067 | 7771 | 10 |
| Small Struct | JSON | Unmarshal | 19822 | 1992 | 34 |
| Small Struct | CBOR | Unmarshal | 20423 | 3104 | 66 |
| Medium Payload | BEVE ZeroCopy | Marshal | 27809 | 133 | 2 |
| Medium Payload | Sonic | Marshal | 69580 | 27789 | 4 |
| Medium Payload | CBOR | Marshal | 76650 | 19160 | 2 |
| Medium Payload | BEVE | Marshal | 116709 | 20624 | 3 |
| Medium Payload | JSON | Marshal | 147694 | 19375 | 9 |
| Medium Payload | MessagePack | Marshal | 255013 | 65831 | 22 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkMedium_JSON_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Small Struct | BEVE | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Unmarshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Unmarshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
