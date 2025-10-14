# BEVE Benchmark Snapshot

> Generated: 2025-10-14T20:23:16Z
> Hostname: runnervmd3hz3
> OS: MINGW64_NT-10.0-26100
> Kernel: MINGW64_NT-10.0-26100 3.6.4-b9f03e96.x86_64 2025-07-16 18:17 UTC
> Architecture: x86_64
> CPU: unknown
> Go: go version go1.25.1 windows/amd64
> Git: e26943b

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | MessagePack | Unmarshal | 1143926 | 348309 | 6343 |
| Large Payload | BEVE | Unmarshal | 1406824 | 280229 | 418 |
| Large Payload | Sonic | Unmarshal | 1759460 | 580866 | 593 |
| Large Payload | CBOR | Unmarshal | 1815833 | 286106 | 5828 |
| Large Payload | JSON | Unmarshal | 3706456 | 525534 | 6773 |
| Small Struct | CBOR | Marshal | 1876 | 352 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 2159 | 288 | 2 |
| Small Struct | Sonic | Marshal | 7696 | 1624 | 3 |
| Small Struct | BEVE | Marshal | 9077 | 1824 | 3 |
| Small Struct | JSON | Marshal | 10903 | 1936 | 2 |
| Small Struct | MessagePack | Marshal | 17352 | 8321 | 9 |
| Large Payload | BEVE ZeroCopy | Marshal | 221291 | 233 | 2 |
| Large Payload | BEVE | Marshal | 673169 | 188874 | 3 |
| Large Payload | CBOR | Marshal | 762112 | 214692 | 2 |
| Large Payload | Sonic | Marshal | 785486 | 201551 | 4 |
| Large Payload | MessagePack | Marshal | 1089680 | 526756 | 115 |
| Large Payload | JSON | Marshal | 1341927 | 214979 | 9 |
| Medium Payload | BEVE | Unmarshal | 97796 | 23355 | 59 |
| Medium Payload | MessagePack | Unmarshal | 267193 | 34956 | 647 |
| Medium Payload | CBOR | Unmarshal | 281812 | 27448 | 568 |
| Medium Payload | Sonic | Unmarshal | 344378 | 57534 | 76 |
| Medium Payload | JSON | Unmarshal | 572076 | 41544 | 548 |
| Small Struct | MessagePack | Unmarshal | 4670 | 968 | 23 |
| Small Struct | Sonic | Unmarshal | 9829 | 4663 | 9 |
| Small Struct | BEVE | Unmarshal | 11546 | 3384 | 4 |
| Small Struct | CBOR | Unmarshal | 15472 | 2216 | 48 |
| Small Struct | JSON | Unmarshal | 39643 | 4136 | 66 |
| Medium Payload | BEVE ZeroCopy | Marshal | 20370 | 129 | 2 |
| Medium Payload | BEVE | Marshal | 62188 | 18570 | 3 |
| Medium Payload | Sonic | Marshal | 71559 | 22157 | 4 |
| Medium Payload | CBOR | Marshal | 104794 | 24691 | 2 |
| Medium Payload | MessagePack | Marshal | 106417 | 65829 | 22 |
| Medium Payload | JSON | Marshal | 158782 | 22068 | 9 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Medium Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkMedium_BEVE_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkMedium_MessagePack_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkMedium_CBOR_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkMedium_Sonic_Unmarshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
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
