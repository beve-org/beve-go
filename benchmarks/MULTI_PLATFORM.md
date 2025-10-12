# Multi-Platform Benchmark Results

| CPU | OS | Artifacts |
|-----|----|-----------|
| Apple M1 (Virtual) | Darwin | [Markdown](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.md) · [JSON](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.json) · [PNG](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.png) |
| AMD EPYC 7763 64-Core Processor | Linux | [Markdown](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md) · [JSON](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.json) · [PNG](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png) |
| Neoverse-N2 | Linux | [Markdown](benchmarks/benchmark-linux-neoverse-n2/benchmark.md) · [JSON](benchmarks/benchmark-linux-neoverse-n2/benchmark.json) · [PNG](benchmarks/benchmark-linux-neoverse-n2/benchmark.png) |
| unknown | MINGW64_NT-10.0-26100 | [Markdown](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.md) · [JSON](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.json) · [PNG](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png) |

## Apple M1 (Virtual) — Darwin

![Benchmark Chart](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | CBOR | Marshal | 582.80 | 400 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 819.90 | 288 | 2 |
| Small Struct | BEVE | Marshal | 1068 | 866 | 3 |
| Small Struct | Sonic | Marshal | 1476 | 590 | 3 |
| Small Struct | MessagePack | Marshal | 6387 | 8321 | 9 |
| Small Struct | JSON | Marshal | 7795 | 2192 | 2 |
| Small Struct | BEVE | Unmarshal | 995.80 | 1336 | 4 |
| Small Struct | CBOR | Unmarshal | 2968 | 1480 | 34 |
| Small Struct | Sonic | Unmarshal | 4482 | 4350 | 6 |
| Small Struct | MessagePack | Unmarshal | 5858 | 4385 | 93 |
| Small Struct | JSON | Unmarshal | 13701 | 3784 | 55 |
| Medium Payload | BEVE | Marshal | 12102 | 18652 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 12115 | 134 | 2 |
| Medium Payload | CBOR | Marshal | 19153 | 21847 | 2 |
| Medium Payload | MessagePack | Marshal | 32444 | 33062 | 21 |
| Medium Payload | JSON | Marshal | 48002 | 20797 | 9 |
| Medium Payload | Sonic | Marshal | 62442 | 25124 | 4 |
| Medium Payload | BEVE | Unmarshal | 21989 | 13674 | 59 |
| Medium Payload | Sonic | Unmarshal | 41982 | 36520 | 33 |
| Medium Payload | MessagePack | Unmarshal | 65045 | 29037 | 528 |
| Medium Payload | CBOR | Unmarshal | 75972 | 30633 | 631 |
| Medium Payload | JSON | Unmarshal | 192421 | 47097 | 596 |
| Large Payload | BEVE ZeroCopy | Marshal | 87814 | 303 | 2 |
| Large Payload | BEVE | Marshal | 129777 | 201487 | 3 |
| Large Payload | CBOR | Marshal | 272266 | 189288 | 2 |
| Large Payload | MessagePack | Marshal | 331496 | 526822 | 115 |
| Large Payload | JSON | Marshal | 483981 | 222215 | 9 |
| Large Payload | Sonic | Marshal | 553298 | 230727 | 4 |
| Large Payload | BEVE | Unmarshal | 263776 | 154003 | 417 |
| Large Payload | Sonic | Unmarshal | 329678 | 351097 | 211 |
| Large Payload | MessagePack | Unmarshal | 503645 | 358269 | 6554 |
| Large Payload | CBOR | Unmarshal | 567521 | 318460 | 6500 |
| Large Payload | JSON | Unmarshal | 2114734 | 514038 | 6710 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 437.00 | 288 | 2 |
| Small Struct | BEVE | Marshal | 814.40 | 931 | 3 |
| Small Struct | JSON | Marshal | 920.10 | 432 | 2 |
| Small Struct | Sonic | Marshal | 1171 | 1511 | 3 |
| Small Struct | CBOR | Marshal | 2403 | 2452 | 2 |
| Small Struct | MessagePack | Marshal | 2489 | 4225 | 8 |
| Small Struct | Sonic | Unmarshal | 695.50 | 453 | 4 |
| Small Struct | BEVE | Unmarshal | 1166 | 1208 | 4 |
| Small Struct | MessagePack | Unmarshal | 3209 | 2144 | 47 |
| Small Struct | CBOR | Unmarshal | 4297 | 2024 | 44 |
| Small Struct | JSON | Unmarshal | 14013 | 3976 | 61 |
| Medium Payload | BEVE ZeroCopy | Marshal | 10770 | 141 | 2 |
| Medium Payload | CBOR | Marshal | 15107 | 13687 | 2 |
| Medium Payload | BEVE | Marshal | 15391 | 19356 | 3 |
| Medium Payload | Sonic | Marshal | 22791 | 34222 | 4 |
| Medium Payload | MessagePack | Marshal | 34784 | 65838 | 22 |
| Medium Payload | JSON | Marshal | 44093 | 22145 | 9 |
| Medium Payload | BEVE | Unmarshal | 19053 | 15595 | 59 |
| Medium Payload | Sonic | Unmarshal | 39630 | 63503 | 76 |
| Medium Payload | MessagePack | Unmarshal | 59997 | 40128 | 750 |
| Medium Payload | CBOR | Unmarshal | 78386 | 33704 | 698 |
| Medium Payload | JSON | Unmarshal | 258037 | 68280 | 893 |
| Large Payload | BEVE ZeroCopy | Marshal | 103663 | 479 | 2 |
| Large Payload | Sonic | Marshal | 153258 | 202181 | 4 |
| Large Payload | BEVE | Marshal | 160545 | 204005 | 3 |
| Large Payload | CBOR | Marshal | 213150 | 205710 | 2 |
| Large Payload | MessagePack | Marshal | 305344 | 526857 | 115 |
| Large Payload | JSON | Marshal | 432833 | 206357 | 9 |
| Large Payload | BEVE | Unmarshal | 184250 | 152753 | 419 |
| Large Payload | Sonic | Unmarshal | 377302 | 568272 | 596 |
| Large Payload | MessagePack | Unmarshal | 539271 | 340721 | 6196 |
| Large Payload | CBOR | Unmarshal | 721181 | 306664 | 6249 |
| Large Payload | JSON | Unmarshal | 2324238 | 575524 | 7468 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmarks/benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 751.50 | 288 | 2 |
| Small Struct | JSON | Marshal | 1117 | 624 | 2 |
| Small Struct | BEVE | Marshal | 1776 | 2597 | 3 |
| Small Struct | CBOR | Marshal | 2525 | 2835 | 2 |
| Small Struct | Sonic | Marshal | 3291 | 2507 | 3 |
| Small Struct | MessagePack | Marshal | 3648 | 8321 | 9 |
| Small Struct | BEVE | Unmarshal | 1372 | 1720 | 4 |
| Small Struct | CBOR | Unmarshal | 2800 | 1288 | 30 |
| Small Struct | MessagePack | Unmarshal | 3269 | 2304 | 50 |
| Small Struct | Sonic | Unmarshal | 3303 | 5491 | 6 |
| Small Struct | JSON | Unmarshal | 16212 | 4424 | 75 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9157 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 12733 | 18903 | 3 |
| Medium Payload | CBOR | Marshal | 19048 | 20557 | 2 |
| Medium Payload | Sonic | Marshal | 28364 | 20831 | 4 |
| Medium Payload | MessagePack | Marshal | 32165 | 65838 | 22 |
| Medium Payload | JSON | Marshal | 43800 | 24907 | 9 |
| Medium Payload | BEVE | Unmarshal | 20089 | 16891 | 59 |
| Medium Payload | Sonic | Unmarshal | 28264 | 34426 | 33 |
| Medium Payload | MessagePack | Unmarshal | 47264 | 31230 | 571 |
| Medium Payload | CBOR | Unmarshal | 57508 | 26537 | 550 |
| Medium Payload | JSON | Unmarshal | 209117 | 58665 | 772 |
| Large Payload | BEVE ZeroCopy | Marshal | 89744 | 829 | 2 |
| Large Payload | BEVE | Marshal | 145901 | 213315 | 3 |
| Large Payload | CBOR | Marshal | 186595 | 197520 | 2 |
| Large Payload | MessagePack | Marshal | 268741 | 526869 | 115 |
| Large Payload | Sonic | Marshal | 292120 | 208305 | 4 |
| Large Payload | JSON | Marshal | 354872 | 197988 | 9 |
| Large Payload | BEVE | Unmarshal | 180253 | 158894 | 418 |
| Large Payload | Sonic | Unmarshal | 288116 | 386243 | 211 |
| Large Payload | MessagePack | Unmarshal | 503043 | 351832 | 6421 |
| Large Payload | CBOR | Unmarshal | 654744 | 331435 | 6746 |
| Large Payload | JSON | Unmarshal | 2030758 | 557342 | 7314 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 1055 | 288 | 2 |
| Small Struct | JSON | Marshal | 1203 | 496 | 2 |
| Small Struct | CBOR | Marshal | 1429 | 1040 | 2 |
| Small Struct | BEVE | Marshal | 1603 | 1186 | 3 |
| Small Struct | Sonic | Marshal | 2741 | 3361 | 3 |
| Small Struct | MessagePack | Marshal | 4084 | 4224 | 8 |
| Small Struct | BEVE | Unmarshal | 2605 | 2104 | 4 |
| Small Struct | CBOR | Unmarshal | 2661 | 856 | 21 |
| Small Struct | MessagePack | Unmarshal | 4461 | 1288 | 29 |
| Small Struct | Sonic | Unmarshal | 5717 | 7409 | 10 |
| Small Struct | JSON | Unmarshal | 21355 | 4616 | 81 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9637 | 148 | 2 |
| Medium Payload | Sonic | Marshal | 20708 | 22234 | 4 |
| Medium Payload | CBOR | Marshal | 22558 | 16462 | 2 |
| Medium Payload | BEVE | Marshal | 36107 | 25045 | 3 |
| Medium Payload | MessagePack | Marshal | 39985 | 65828 | 22 |
| Medium Payload | JSON | Marshal | 64490 | 27636 | 9 |
| Medium Payload | BEVE | Unmarshal | 26214 | 18042 | 59 |
| Medium Payload | Sonic | Unmarshal | 35136 | 39652 | 57 |
| Medium Payload | MessagePack | Unmarshal | 67810 | 34508 | 632 |
| Medium Payload | CBOR | Unmarshal | 90323 | 32873 | 675 |
| Medium Payload | JSON | Unmarshal | 270655 | 60600 | 795 |
| Large Payload | BEVE ZeroCopy | Marshal | 120376 | 479 | 2 |
| Large Payload | Sonic | Marshal | 159979 | 218873 | 4 |
| Large Payload | BEVE | Marshal | 191885 | 213920 | 3 |
| Large Payload | CBOR | Marshal | 223791 | 199890 | 2 |
| Large Payload | MessagePack | Marshal | 292611 | 526770 | 115 |
| Large Payload | JSON | Marshal | 493897 | 225349 | 9 |
| Large Payload | BEVE | Unmarshal | 224178 | 149155 | 418 |
| Large Payload | Sonic | Unmarshal | 424793 | 524661 | 579 |
| Large Payload | MessagePack | Unmarshal | 696740 | 362046 | 6614 |
| Large Payload | CBOR | Unmarshal | 862958 | 316316 | 6449 |
| Large Payload | JSON | Unmarshal | 2430977 | 519614 | 6799 |

