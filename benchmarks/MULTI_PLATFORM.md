# Multi-Platform Benchmark Results

| CPU | OS | Artifacts |
|-----|----|-----------|
| Apple M1 (Virtual) | Darwin | [Markdown](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.md) · [JSON](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.json) · [PNG](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.png) |
| AMD EPYC 7763 64-Core Processor | Linux | [Markdown](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md) · [JSON](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.json) · [PNG](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png) |
| Neoverse-N2 | Linux | [Markdown](benchmarks/benchmark-linux-neoverse-n2/benchmark.md) · [JSON](benchmarks/benchmark-linux-neoverse-n2/benchmark.json) · [PNG](benchmarks/benchmark-linux-neoverse-n2/benchmark.png) |
| unknown | MINGW64_NT-10.0-26100 | [Markdown](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.md) · [JSON](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.json) · [PNG](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png) |

## Apple M1 (Virtual) — Darwin

![Benchmark Chart](benchmark-darwin-apple-m1-virtual/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1130630 | 144449 | 419 |
| Large Payload | Sonic | Unmarshal | 2027899 | 358677 | 213 |
| Large Payload | MessagePack | Unmarshal | 2277705 | 328738 | 5943 |
| Large Payload | CBOR | Unmarshal | 2590314 | 291579 | 5945 |
| Large Payload | JSON | Unmarshal | 4066396 | 511501 | 6660 |
| Small Struct | MessagePack | Marshal | 3516 | 4224 | 8 |
| Small Struct | BEVE | Marshal | 9871 | 2977 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 17508 | 288 | 2 |
| Small Struct | CBOR | Marshal | 25803 | 2193 | 2 |
| Small Struct | JSON | Marshal | 42766 | 2192 | 2 |
| Small Struct | Sonic | Marshal | 45897 | 1985 | 3 |
| Large Payload | BEVE | Marshal | 1043429 | 189042 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 1088986 | 233 | 2 |
| Large Payload | CBOR | Marshal | 1316790 | 190018 | 2 |
| Large Payload | MessagePack | Marshal | 1769767 | 526804 | 115 |
| Large Payload | JSON | Marshal | 2080171 | 206217 | 9 |
| Large Payload | Sonic | Marshal | 2838372 | 202846 | 4 |
| Medium Payload | BEVE | Unmarshal | 211149 | 16890 | 59 |
| Medium Payload | Sonic | Unmarshal | 334861 | 43572 | 33 |
| Medium Payload | MessagePack | Unmarshal | 582087 | 41582 | 777 |
| Medium Payload | CBOR | Unmarshal | 600315 | 29560 | 608 |
| Medium Payload | JSON | Unmarshal | 1250425 | 51640 | 707 |
| Small Struct | BEVE | Unmarshal | 10575 | 2104 | 4 |
| Small Struct | Sonic | Unmarshal | 45103 | 2796 | 6 |
| Small Struct | MessagePack | Unmarshal | 59222 | 2304 | 50 |
| Small Struct | CBOR | Unmarshal | 109909 | 4616 | 97 |
| Small Struct | JSON | Unmarshal | 194693 | 3912 | 59 |
| Medium Payload | BEVE ZeroCopy | Marshal | 130090 | 133 | 2 |
| Medium Payload | BEVE | Marshal | 168187 | 19209 | 3 |
| Medium Payload | CBOR | Marshal | 244205 | 21862 | 2 |
| Medium Payload | MessagePack | Marshal | 330678 | 65831 | 22 |
| Medium Payload | JSON | Marshal | 393445 | 22097 | 9 |
| Medium Payload | Sonic | Marshal | 586468 | 20960 | 4 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1040360 | 145481 | 418 |
| Large Payload | MessagePack | Unmarshal | 1760434 | 340984 | 6187 |
| Large Payload | Sonic | Unmarshal | 1789026 | 578100 | 599 |
| Large Payload | CBOR | Unmarshal | 2055200 | 314252 | 6416 |
| Large Payload | JSON | Unmarshal | 3887559 | 561995 | 7285 |
| Small Struct | BEVE | Marshal | 2234 | 480 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 4765 | 288 | 2 |
| Small Struct | Sonic | Marshal | 5873 | 897 | 3 |
| Small Struct | CBOR | Marshal | 8014 | 1937 | 2 |
| Small Struct | JSON | Marshal | 9390 | 848 | 2 |
| Small Struct | MessagePack | Marshal | 10776 | 2176 | 7 |
| Large Payload | BEVE ZeroCopy | Marshal | 587592 | 286 | 2 |
| Large Payload | BEVE | Marshal | 667483 | 188905 | 3 |
| Large Payload | Sonic | Marshal | 792213 | 222150 | 4 |
| Large Payload | CBOR | Marshal | 871858 | 172887 | 2 |
| Large Payload | MessagePack | Marshal | 1364207 | 526787 | 115 |
| Large Payload | JSON | Marshal | 2063768 | 205695 | 9 |
| Medium Payload | BEVE | Unmarshal | 126050 | 15899 | 59 |
| Medium Payload | Sonic | Unmarshal | 196516 | 49058 | 70 |
| Medium Payload | MessagePack | Unmarshal | 329395 | 39520 | 740 |
| Medium Payload | CBOR | Unmarshal | 434364 | 32728 | 669 |
| Medium Payload | JSON | Unmarshal | 1067337 | 45529 | 612 |
| Small Struct | Sonic | Unmarshal | 5422 | 830 | 6 |
| Small Struct | BEVE | Unmarshal | 9134 | 1848 | 4 |
| Small Struct | MessagePack | Unmarshal | 20060 | 2240 | 48 |
| Small Struct | JSON | Unmarshal | 59009 | 2184 | 40 |
| Small Struct | CBOR | Unmarshal | 61641 | 4200 | 88 |
| Medium Payload | BEVE ZeroCopy | Marshal | 73311 | 135 | 2 |
| Medium Payload | Sonic | Marshal | 92539 | 19071 | 4 |
| Medium Payload | BEVE | Marshal | 93101 | 20626 | 3 |
| Medium Payload | CBOR | Marshal | 102183 | 18533 | 2 |
| Medium Payload | MessagePack | Marshal | 156230 | 33060 | 21 |
| Medium Payload | JSON | Marshal | 312703 | 24888 | 9 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 548235 | 148759 | 418 |
| Large Payload | Sonic | Unmarshal | 672239 | 356963 | 211 |
| Large Payload | MessagePack | Unmarshal | 721220 | 339018 | 6148 |
| Large Payload | CBOR | Unmarshal | 1037323 | 332299 | 6768 |
| Large Payload | JSON | Unmarshal | 2412413 | 522180 | 6813 |
| Small Struct | CBOR | Marshal | 1518 | 320 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 2367 | 289 | 2 |
| Small Struct | BEVE | Marshal | 3246 | 1696 | 3 |
| Small Struct | Sonic | Marshal | 7972 | 2554 | 3 |
| Small Struct | JSON | Marshal | 10371 | 2833 | 2 |
| Small Struct | MessagePack | Marshal | 12993 | 8321 | 9 |
| Large Payload | BEVE ZeroCopy | Marshal | 256688 | 260 | 2 |
| Large Payload | BEVE | Marshal | 388346 | 188857 | 3 |
| Large Payload | CBOR | Marshal | 537385 | 197747 | 2 |
| Large Payload | MessagePack | Marshal | 806015 | 526797 | 115 |
| Large Payload | Sonic | Marshal | 839535 | 222423 | 4 |
| Large Payload | JSON | Marshal | 991237 | 205693 | 9 |
| Medium Payload | BEVE | Unmarshal | 45384 | 13498 | 58 |
| Medium Payload | Sonic | Unmarshal | 89197 | 45930 | 33 |
| Medium Payload | MessagePack | Unmarshal | 151418 | 31341 | 575 |
| Medium Payload | CBOR | Unmarshal | 215373 | 33192 | 686 |
| Medium Payload | JSON | Unmarshal | 541173 | 50584 | 652 |
| Small Struct | BEVE | Unmarshal | 5625 | 1848 | 4 |
| Small Struct | MessagePack | Unmarshal | 6702 | 3936 | 83 |
| Small Struct | Sonic | Unmarshal | 7953 | 4559 | 6 |
| Small Struct | CBOR | Unmarshal | 16979 | 3592 | 77 |
| Small Struct | JSON | Unmarshal | 21489 | 2312 | 44 |
| Medium Payload | BEVE ZeroCopy | Marshal | 23077 | 131 | 2 |
| Medium Payload | BEVE | Marshal | 56094 | 19211 | 3 |
| Medium Payload | CBOR | Marshal | 62259 | 20566 | 2 |
| Medium Payload | JSON | Marshal | 99690 | 19371 | 9 |
| Medium Payload | MessagePack | Marshal | 104166 | 65832 | 22 |
| Medium Payload | Sonic | Marshal | 106630 | 25446 | 4 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1107554 | 150944 | 417 |
| Large Payload | MessagePack | Unmarshal | 1325433 | 323706 | 5850 |
| Large Payload | Sonic | Unmarshal | 1613816 | 572431 | 613 |
| Large Payload | CBOR | Unmarshal | 2046596 | 325964 | 6651 |
| Large Payload | JSON | Unmarshal | 3629644 | 523797 | 6759 |
| Small Struct | BEVE ZeroCopy | Marshal | 2812 | 288 | 2 |
| Small Struct | MessagePack | Marshal | 8096 | 2176 | 7 |
| Small Struct | JSON | Marshal | 14060 | 1552 | 2 |
| Small Struct | BEVE | Marshal | 16048 | 2593 | 3 |
| Small Struct | CBOR | Marshal | 17164 | 2833 | 2 |
| Small Struct | Sonic | Marshal | 17353 | 2942 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 243501 | 207 | 2 |
| Large Payload | CBOR | Marshal | 916548 | 197901 | 2 |
| Large Payload | Sonic | Marshal | 950318 | 218731 | 4 |
| Large Payload | BEVE | Marshal | 1033267 | 180562 | 3 |
| Large Payload | MessagePack | Marshal | 1380349 | 526768 | 115 |
| Large Payload | JSON | Marshal | 1504587 | 223697 | 9 |
| Medium Payload | BEVE | Unmarshal | 69895 | 14234 | 59 |
| Medium Payload | MessagePack | Unmarshal | 247550 | 31132 | 565 |
| Medium Payload | CBOR | Unmarshal | 377090 | 34936 | 716 |
| Medium Payload | Sonic | Unmarshal | 440557 | 63945 | 77 |
| Medium Payload | JSON | Unmarshal | 670193 | 44728 | 617 |
| Small Struct | BEVE | Unmarshal | 7362 | 1592 | 4 |
| Small Struct | MessagePack | Unmarshal | 12924 | 3072 | 64 |
| Small Struct | Sonic | Unmarshal | 15836 | 7416 | 10 |
| Small Struct | CBOR | Unmarshal | 23041 | 4776 | 102 |
| Small Struct | JSON | Unmarshal | 28493 | 2280 | 43 |
| Medium Payload | BEVE ZeroCopy | Marshal | 28679 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 48576 | 20620 | 3 |
| Medium Payload | CBOR | Marshal | 83660 | 21843 | 2 |
| Medium Payload | Sonic | Marshal | 96737 | 22298 | 4 |
| Medium Payload | MessagePack | Marshal | 123691 | 65829 | 22 |
| Medium Payload | JSON | Marshal | 176544 | 24900 | 9 |

