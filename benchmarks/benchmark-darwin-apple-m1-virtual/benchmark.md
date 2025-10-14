# BEVE Benchmark Snapshot

> Generated: 2025-10-14T20:11:39Z
> Hostname: sat12-bq161-f63eeef8-5436-479e-aecf-c9a7c095d6fb-A2F2C9259122.local
> OS: Darwin
> Kernel: Darwin 24.6.0 Darwin Kernel Version 24.6.0: Mon Jul 14 11:30:18 PDT 2025; root:xnu-11417.140.69~1/RELEASE_ARM64_VMAPPLE
> Architecture: arm64
> CPU: Apple M1 (Virtual)
> Go: go version go1.25.1 darwin/arm64
> Git: b814044

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

## Summary

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1554466 | 275141 | 417 |
| Large Payload | Sonic | Unmarshal | 2081949 | 366434 | 213 |
| Large Payload | MessagePack | Unmarshal | 2426497 | 354854 | 6473 |
| Large Payload | CBOR | Unmarshal | 2460638 | 313211 | 6382 |
| Large Payload | JSON | Unmarshal | 4411674 | 559739 | 7310 |
| Small Struct | MessagePack | Marshal | 1664 | 2176 | 7 |
| Small Struct | BEVE | Marshal | 8599 | 1824 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 10145 | 288 | 2 |
| Small Struct | CBOR | Marshal | 28447 | 1680 | 2 |
| Small Struct | JSON | Marshal | 110782 | 1552 | 2 |
| Small Struct | Sonic | Marshal | 157889 | 3303 | 3 |
| Large Payload | BEVE | Marshal | 1025215 | 205429 | 3 |
| Large Payload | CBOR | Marshal | 1229549 | 197538 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 1410277 | 207 | 2 |
| Large Payload | MessagePack | Marshal | 1663249 | 526797 | 115 |
| Large Payload | JSON | Marshal | 1972992 | 214726 | 9 |
| Large Payload | Sonic | Marshal | 2404221 | 218558 | 4 |
| Medium Payload | BEVE | Unmarshal | 306982 | 32605 | 59 |
| Medium Payload | Sonic | Unmarshal | 372559 | 38604 | 33 |
| Medium Payload | MessagePack | Unmarshal | 525078 | 43086 | 808 |
| Medium Payload | CBOR | Unmarshal | 652926 | 30216 | 626 |
| Medium Payload | JSON | Unmarshal | 1210829 | 62504 | 834 |
| Small Struct | BEVE | Unmarshal | 10032 | 1848 | 4 |
| Small Struct | Sonic | Unmarshal | 53954 | 5060 | 6 |
| Small Struct | CBOR | Unmarshal | 105884 | 4384 | 94 |
| Small Struct | MessagePack | Unmarshal | 133939 | 4384 | 93 |
| Small Struct | JSON | Unmarshal | 342851 | 7208 | 91 |
| Medium Payload | BEVE | Marshal | 175018 | 19221 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 224622 | 133 | 2 |
| Medium Payload | CBOR | Marshal | 310188 | 20570 | 2 |
| Medium Payload | MessagePack | Marshal | 355349 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 412286 | 20784 | 9 |
| Medium Payload | Sonic | Marshal | 557273 | 22350 | 4 |

## Commands

| Scenario | Codec | Operation | Command |
|----------|-------|-----------|---------|
| Large Payload | BEVE | Unmarshal | `go test -bench=\^BenchmarkLarge_BEVE_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Unmarshal | `go test -bench=\^BenchmarkLarge_Sonic_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Unmarshal | `go test -bench=\^BenchmarkLarge_MessagePack_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Unmarshal | `go test -bench=\^BenchmarkLarge_CBOR_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Unmarshal | `go test -bench=\^BenchmarkLarge_JSON_Unmarshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Small Struct | MessagePack | Marshal | `go test -bench=\^BenchmarkSmallStruct_MessagePack_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkSmallStruct_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | CBOR | Marshal | `go test -bench=\^BenchmarkSmallStruct_CBOR_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | JSON | Marshal | `go test -bench=\^BenchmarkSmallStruct_JSON_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Small Struct | Sonic | Marshal | `go test -bench=\^BenchmarkSmallStruct_Sonic_Marshal\$ -benchmem -benchtime=30000x -run=\^\$ .` |
| Large Payload | BEVE | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | CBOR | Marshal | `go test -bench=\^BenchmarkLarge_CBOR_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkLarge_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkLarge_MessagePack_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | JSON | Marshal | `go test -bench=\^BenchmarkLarge_JSON_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
| Large Payload | Sonic | Marshal | `go test -bench=\^BenchmarkLarge_Sonic_Marshal\$ -benchmem -benchtime=10000x -run=\^\$ .` |
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
| Medium Payload | BEVE | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | BEVE ZeroCopy | Marshal | `go test -bench=\^BenchmarkMedium_BEVE_MarshalZeroCopy\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | CBOR | Marshal | `go test -bench=\^BenchmarkMedium_CBOR_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | MessagePack | Marshal | `go test -bench=\^BenchmarkMedium_MessagePack_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | JSON | Marshal | `go test -bench=\^BenchmarkMedium_JSON_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
| Medium Payload | Sonic | Marshal | `go test -bench=\^BenchmarkMedium_Sonic_Marshal\$ -benchmem -benchtime=20000x -run=\^\$ .` |
