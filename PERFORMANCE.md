# BEVE Performance Snapshot – October 2025

This document captures the latest benchmark runs comparing **BEVE** against popular Go codecs on an Apple M2 Max using Go 1.22 and `-benchtime=1000x` (large payloads use `-benchtime=50x`). All benchmarks live in [`comparison_advanced_test.go`](comparison_advanced_test.go).

## Environment

- macOS (Apple M2 Max)
- Go 1.22
- Command template: `go test -bench=<NAME> -benchmem -benchtime=<ITER>x ./...`

## Small Struct (User profile)

| Operation | Codec | ns/op | B/op | allocs/op | Delta vs BEVE |
|-----------|-------|-------|------|-----------|----------------|
| Marshal | **BEVE** | **294.1** | **393** | **2** | — |
|          | CBOR | 1,618 | 2,841 | 2 | BEVE is **5.5× faster**, **6.7× less heap** |
|          | JSON | 2,369 | 1,942 | 2 | BEVE is **8.1× faster**, **5× less heap** |
| Unmarshal | **BEVE** | **665.5** | **889** | **4** | — |
|            | CBOR | 3,059 | 2,440 | 53 | BEVE is **4.6× faster**, **2.7× less heap**, **13× fewer allocs** |
|            | JSON | 6,495 | 2,312 | 44 | BEVE is **9.8× faster**, **2.6× less heap** |

## Medium Payload (10 users + 20 orders)

| Operation | Codec | ns/op | B/op | allocs/op | Delta vs BEVE |
|-----------|-------|-------|------|-----------|----------------|
| Marshal | **BEVE** | **11,636** | 22,388 | **2** | — |
|          | CBOR | 14,087 | 18,577 | 2 | BEVE is **21% faster**, a modest heap trade-off |

## Large Payload (100 users + 200 orders)

| Operation | Codec | ns/op | B/op | allocs/op | Delta vs BEVE |
|-----------|-------|-------|------|-----------|----------------|
| Marshal | **BEVE** | **101,629** | 223,402 | **2** | — |
|          | CBOR | 131,509 | 199,622 | 3 | BEVE is **29% faster** while staying near-constant allocations |

## Key Takeaways

- **Throughput**: BEVE now leads every measured marshal/unmarshal workload, topping CBOR by up to **5.5×** on hot paths.
- **Allocations**: Decoder work was reduced to **4 allocs/op**, over **13× fewer** than CBOR in the small-struct scenario.
- **Memory**: Heap usage fell below **1 KB/op** for small-struct decode and stays competitive on larger payloads.
- **Consistency**: Encoder pooling and typed-array fast paths keep allocations flat (2–4 per op) regardless of payload size.

For additional scenarios (MessagePack, Sonic/JSON round trips, etc.) re-run the relevant benchmarks in `comparison_advanced_test.go` with the same parameters.
