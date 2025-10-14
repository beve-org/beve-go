# BEVE Benchmark Snapshot

> Generated: 2025-10-14T20:22:27Z
> Hostname: sjc22-be107-5bda3f77-9adc-410a-b0b8-72b416ac2973-36C06D29299F.local
> OS: Darwin
> Kernel: Darwin 24.6.0 Darwin Kernel Version 24.6.0: Mon Jul 14 11:30:18 PDT 2025; root:xnu-11417.140.69~1/RELEASE_ARM64_VMAPPLE
> Architecture: arm64
> CPU: Apple M1 (Virtual)
> Go: go version go1.25.1 darwin/arm64
> Git: e26943b

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1491535 | 279431 | 419 |
| Large Payload | Sonic | Unmarshal | 1817693 | 377283 | 207 |
| Large Payload | MessagePack | Unmarshal | 2082285 | 333859 | 6067 |
| Large Payload | CBOR | Unmarshal | 2352011 | 310763 | 6326 |
| Large Payload | JSON | Unmarshal | 3661498 | 509979 | 6720 |
| Small Struct | BEVE ZeroCopy | Marshal | 5394 | 288 | 2 |
| Small Struct | BEVE | Marshal | 15073 | 2080 | 3 |
| Small Struct | CBOR | Marshal | 21857 | 1936 | 2 |
| Small Struct | JSON | Marshal | 29880 | 2449 | 2 |
| Small Struct | MessagePack | Marshal | 46642 | 8321 | 9 |
| Small Struct | Sonic | Marshal | 118670 | 3294 | 3 |
| Large Payload | BEVE | Marshal | 1108661 | 197553 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 1172730 | 233 | 2 |
| Large Payload | CBOR | Marshal | 1231507 | 197641 | 2 |
| Large Payload | MessagePack | Marshal | 1637349 | 526803 | 115 |
| Large Payload | JSON | Marshal | 2215958 | 222654 | 9 |
| Large Payload | Sonic | Marshal | 2282839 | 226490 | 4 |
| Medium Payload | BEVE | Unmarshal | 193306 | 28188 | 59 |
| Medium Payload | Sonic | Unmarshal | 361936 | 44774 | 33 |
| Medium Payload | MessagePack | Unmarshal | 489974 | 35821 | 664 |
| Medium Payload | CBOR | Unmarshal | 578584 | 28296 | 582 |
| Medium Payload | JSON | Unmarshal | 1083101 | 45448 | 584 |
| Small Struct | BEVE | Unmarshal | 25045 | 2616 | 4 |
| Small Struct | MessagePack | Unmarshal | 28771 | 2784 | 59 |
| Small Struct | Sonic | Unmarshal | 37515 | 2446 | 6 |
| Small Struct | CBOR | Unmarshal | 46064 | 3624 | 78 |
| Small Struct | JSON | Unmarshal | 227740 | 4328 | 72 |
| Medium Payload | BEVE | Marshal | 150634 | 20617 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 190822 | 133 | 2 |
| Medium Payload | CBOR | Marshal | 197633 | 24693 | 2 |
| Medium Payload | MessagePack | Marshal | 352739 | 65831 | 22 |
| Medium Payload | JSON | Marshal | 439950 | 24883 | 9 |
| Medium Payload | Sonic | Marshal | 534821 | 21058 | 4 |

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
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
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
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
