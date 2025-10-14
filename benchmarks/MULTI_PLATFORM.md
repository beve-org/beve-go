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

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1375701 | 293109 | 419 |
| Large Payload | Sonic | Unmarshal | 1718375 | 543695 | 575 |
| Large Payload | MessagePack | Unmarshal | 1721846 | 364352 | 6660 |
| Large Payload | CBOR | Unmarshal | 1949040 | 328173 | 6688 |
| Large Payload | JSON | Unmarshal | 3927704 | 543802 | 7131 |
| Small Struct | BEVE ZeroCopy | Marshal | 1860 | 288 | 2 |
| Small Struct | CBOR | Marshal | 4960 | 560 | 2 |
| Small Struct | Sonic | Marshal | 6061 | 1354 | 3 |
| Small Struct | JSON | Marshal | 9319 | 656 | 2 |
| Small Struct | BEVE | Marshal | 11226 | 3360 | 3 |
| Small Struct | MessagePack | Marshal | 11455 | 2176 | 7 |
| Large Payload | BEVE ZeroCopy | Marshal | 603653 | 259 | 2 |
| Large Payload | BEVE | Marshal | 680299 | 197333 | 3 |
| Large Payload | Sonic | Marshal | 777660 | 229290 | 4 |
| Large Payload | CBOR | Marshal | 1020495 | 197364 | 2 |
| Large Payload | MessagePack | Marshal | 1580960 | 526792 | 115 |
| Large Payload | JSON | Marshal | 2156489 | 205748 | 9 |
| Medium Payload | BEVE | Unmarshal | 177263 | 33374 | 59 |
| Medium Payload | Sonic | Unmarshal | 235522 | 58985 | 76 |
| Medium Payload | MessagePack | Unmarshal | 322609 | 37551 | 703 |
| Medium Payload | CBOR | Unmarshal | 421019 | 32312 | 665 |
| Medium Payload | JSON | Unmarshal | 1242097 | 54009 | 710 |
| Small Struct | BEVE | Unmarshal | 10086 | 1848 | 4 |
| Small Struct | JSON | Unmarshal | 23379 | 872 | 21 |
| Small Struct | CBOR | Unmarshal | 26271 | 1296 | 30 |
| Small Struct | Sonic | Unmarshal | 28479 | 7782 | 10 |
| Small Struct | MessagePack | Unmarshal | 32742 | 3968 | 84 |
| Medium Payload | BEVE | Marshal | 72175 | 16523 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 83152 | 136 | 2 |
| Medium Payload | Sonic | Marshal | 91840 | 22522 | 4 |
| Medium Payload | CBOR | Marshal | 121238 | 21860 | 2 |
| Medium Payload | MessagePack | Marshal | 140485 | 33060 | 21 |
| Medium Payload | JSON | Marshal | 316318 | 24874 | 9 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 742424 | 267730 | 418 |
| Large Payload | Sonic | Unmarshal | 758286 | 417122 | 213 |
| Large Payload | MessagePack | Unmarshal | 803998 | 340218 | 6179 |
| Large Payload | CBOR | Unmarshal | 1074256 | 322443 | 6568 |
| Large Payload | JSON | Unmarshal | 2550255 | 553475 | 7252 |
| Small Struct | BEVE ZeroCopy | Marshal | 2996 | 288 | 2 |
| Small Struct | BEVE | Marshal | 3454 | 1312 | 3 |
| Small Struct | CBOR | Marshal | 6513 | 2192 | 2 |
| Small Struct | JSON | Marshal | 7179 | 1296 | 2 |
| Small Struct | MessagePack | Marshal | 9688 | 4224 | 8 |
| Small Struct | Sonic | Marshal | 10525 | 2265 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 214027 | 207 | 2 |
| Large Payload | BEVE | Marshal | 393565 | 188990 | 3 |
| Large Payload | CBOR | Marshal | 516309 | 180931 | 2 |
| Large Payload | MessagePack | Marshal | 846074 | 526794 | 115 |
| Large Payload | Sonic | Marshal | 900091 | 220639 | 4 |
| Large Payload | JSON | Marshal | 1063017 | 213988 | 9 |
| Medium Payload | Sonic | Unmarshal | 78176 | 30272 | 33 |
| Medium Payload | BEVE | Unmarshal | 88403 | 30782 | 59 |
| Medium Payload | MessagePack | Unmarshal | 190834 | 43264 | 813 |
| Medium Payload | CBOR | Unmarshal | 252954 | 39096 | 804 |
| Medium Payload | JSON | Unmarshal | 465709 | 39448 | 509 |
| Small Struct | Sonic | Unmarshal | 2825 | 622 | 6 |
| Small Struct | BEVE | Unmarshal | 5835 | 2104 | 4 |
| Small Struct | MessagePack | Unmarshal | 7329 | 1568 | 35 |
| Small Struct | CBOR | Unmarshal | 11703 | 1352 | 31 |
| Small Struct | JSON | Unmarshal | 43486 | 4200 | 68 |
| Medium Payload | BEVE ZeroCopy | Marshal | 25963 | 131 | 2 |
| Medium Payload | BEVE | Marshal | 47418 | 21898 | 3 |
| Medium Payload | CBOR | Marshal | 73575 | 27363 | 2 |
| Medium Payload | Sonic | Marshal | 90093 | 21314 | 4 |
| Medium Payload | JSON | Marshal | 102250 | 18733 | 9 |
| Medium Payload | MessagePack | Marshal | 118936 | 65833 | 22 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

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

