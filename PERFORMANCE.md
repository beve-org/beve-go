# BEVE Performance Snapshot – October 2025

This document captures the latest benchmark runs comparing **BEVE** against popular Go codecs on an Apple M2 Max using Go 1.25.1. Benchmarks follow the curated suite in [`comparison_advanced_test.go`](comparison_advanced_test.go) with size-specific durations: small payloadlar `-benchtime=10000x`, medium payloadlar `-benchtime=5000x`, large payloadlar `-benchtime=3000x`.

## Environment

- macOS (Apple M2 Max)
- Go 1.25.1
- Command template: `go test -bench=<NAME> -benchmem -benchtime=<ITER>x ./...`

## Small Struct (User profile)

| Operation | Codec | ns/op | B/op | allocs/op | Delta vs BEVE |
|-----------|-------|-------|------|-----------|----------------|
| Marshal | **BEVE** | **404.5** | **793** | **2** | — |
|          | CBOR | 2,258 | 1,681 | 2 | BEVE is **5.6× faster**, **53% less heap** |
|          | JSON | 2,732 | 1,939 | 2 | BEVE is **6.8× faster**; JSON uses **2.4×** more heap |
| Unmarshal | **BEVE** | **744.6** | **1,209** | **4** | — |
|            | CBOR | 4,781 | 4,237 | 89 | BEVE is **6.4× faster**, **71% less heap**, **22× fewer allocs** |
|            | JSON | 2,683 | 680 | 18 | BEVE is **2.4× faster**; JSON uses **44% less heap** |

## Medium Payload (10 users + 20 orders)

| Operation | Codec | ns/op | B/op | allocs/op | Delta vs BEVE |
|-----------|-------|-------|------|-----------|----------------|
| Marshal | **BEVE** | **10,166** | 21,356 | **2** | — |
|          | CBOR | 12,982 | 16,665 | 2 | BEVE is **27% faster**; CBOR saves **22% heap** |
|          | JSON | 30,646 | 22,092 | 9 | BEVE is **3.0× faster**, **4.5× fewer allocs** |

## Large Payload (100 users + 200 orders)

| Operation | Codec | ns/op | B/op | allocs/op | Delta vs BEVE |
|-----------|-------|-------|------|-----------|----------------|
| Marshal | **BEVE** | **85,795** | 209,872 | **2** | — |
|          | CBOR | 123,768 | 191,431 | 3 | BEVE is **1.4× faster**, similar heap |
|          | JSON | 307,572 | 224,098 | 9 | BEVE is **3.6× faster**, **4.5× fewer allocs** |

## Key Takeaways

- **Throughput**: BEVE now leads every measured marshal/unmarshal workload, topping CBOR by up to **5.6×** and JSON by **6.8×** on hot marshal paths.
- **Allocations**: Decoder work holds at **4 allocs/op**, delivering **22× fewer** allocations than CBOR and **4.5× fewer** than JSON for small structs.
- **Memory**: Small-struct decode stays near **1.2 KB/op**, while marshal paths use **53% less heap** than CBOR and stay below **0.8 KB/op**.
- **Consistency**: Encoder pooling and slice fast paths keep allocations flat (2–4 per op) even as payload size grows.
- **Zero-copy varyantları**: `BEVE ZeroCopy` satırları artık raporlarda yer alıyor; özellikle küçük payload'larda kopyasız yolun CPU/bellek avantajı belirgin.
- **Decode kapsamı**: Orta ve büyük payload senaryoları için unmarshal (decode) benchmark'ları da otomasyona eklendi; BEVE’nin parse performansı ve allocation profili doğrudan görülebiliyor.

For additional scenarios (MessagePack, Sonic/JSON round trips, etc.) re-run the relevant benchmarks in `comparison_advanced_test.go` with the same parameters.

## Automation & Visualization

- Run `./scripts/bench.sh` to reproduce the curated suite locally. It now emits both `benchmarks/latest.md` and a structured `benchmarks/latest.json` snapshot with raw metrics.
- Generate comparison charts with `python scripts/plot_benchmarks.py benchmarks/latest.json benchmarks/latest.png` (requires `matplotlib`). The tool auto-selects readable units and groups codecs per scenario.
- GitHub Actions workflow [`benchmarks.yml`](.github/workflows/benchmarks.yml) executes the same suite on macOS, Linux, and Windows runners, publishes the Markdown, JSON, and PNG artefacts, and makes them downloadable from each run.
