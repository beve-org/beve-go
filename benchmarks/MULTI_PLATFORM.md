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
| Small Struct | BEVE | Marshal | 497.80 | 930 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 917.30 | 289 | 2 |
| Small Struct | CBOR | Marshal | 1016 | 1425 | 2 |
| Small Struct | MessagePack | Marshal | 1656 | 4224 | 8 |
| Small Struct | JSON | Marshal | 1837 | 1296 | 2 |
| Small Struct | Sonic | Marshal | 2990 | 1967 | 3 |
| Small Struct | BEVE | Unmarshal | 773.70 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 1513 | 1722 | 6 |
| Small Struct | MessagePack | Unmarshal | 1538 | 1512 | 34 |
| Small Struct | CBOR | Unmarshal | 4958 | 4328 | 92 |
| Small Struct | JSON | Unmarshal | 12117 | 4168 | 67 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5935 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 7557 | 14576 | 3 |
| Medium Payload | CBOR | Marshal | 16317 | 24664 | 2 |
| Medium Payload | MessagePack | Marshal | 21822 | 65834 | 22 |
| Medium Payload | JSON | Marshal | 29399 | 20797 | 9 |
| Medium Payload | Sonic | Marshal | 47340 | 27588 | 4 |
| Medium Payload | BEVE | Unmarshal | 13445 | 16250 | 59 |
| Medium Payload | Sonic | Unmarshal | 27490 | 34840 | 33 |
| Medium Payload | MessagePack | Unmarshal | 39418 | 40670 | 765 |
| Medium Payload | CBOR | Unmarshal | 53636 | 33721 | 695 |
| Medium Payload | JSON | Unmarshal | 153905 | 47145 | 643 |
| Large Payload | BEVE ZeroCopy | Marshal | 61688 | 303 | 2 |
| Large Payload | BEVE | Marshal | 104991 | 200873 | 3 |
| Large Payload | CBOR | Marshal | 145670 | 197668 | 2 |
| Large Payload | MessagePack | Marshal | 187330 | 526834 | 115 |
| Large Payload | JSON | Marshal | 351637 | 230409 | 9 |
| Large Payload | Sonic | Marshal | 380269 | 213905 | 4 |
| Large Payload | BEVE | Unmarshal | 124112 | 156051 | 419 |
| Large Payload | Sonic | Unmarshal | 234004 | 329671 | 209 |
| Large Payload | MessagePack | Unmarshal | 361269 | 352778 | 6434 |
| Large Payload | CBOR | Unmarshal | 416311 | 317641 | 6477 |
| Large Payload | JSON | Unmarshal | 1496116 | 518393 | 6812 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE | Marshal | 1278 | 1570 | 3 |
| Small Struct | Sonic | Marshal | 1482 | 2037 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 1525 | 289 | 2 |
| Small Struct | CBOR | Marshal | 1937 | 1937 | 2 |
| Small Struct | JSON | Marshal | 3751 | 2192 | 2 |
| Small Struct | MessagePack | Marshal | 3983 | 8322 | 9 |
| Small Struct | BEVE | Unmarshal | 639.50 | 312 | 3 |
| Small Struct | MessagePack | Unmarshal | 1959 | 1024 | 24 |
| Small Struct | CBOR | Unmarshal | 2653 | 1096 | 26 |
| Small Struct | Sonic | Unmarshal | 4180 | 7783 | 10 |
| Small Struct | JSON | Unmarshal | 6706 | 1384 | 30 |
| Medium Payload | Sonic | Marshal | 12401 | 17061 | 4 |
| Medium Payload | BEVE ZeroCopy | Marshal | 12767 | 128 | 2 |
| Medium Payload | BEVE | Marshal | 16204 | 20854 | 3 |
| Medium Payload | CBOR | Marshal | 22923 | 21937 | 2 |
| Medium Payload | MessagePack | Marshal | 35212 | 65839 | 22 |
| Medium Payload | JSON | Marshal | 44650 | 20930 | 9 |
| Medium Payload | BEVE | Unmarshal | 18153 | 13210 | 59 |
| Medium Payload | Sonic | Unmarshal | 36322 | 55672 | 74 |
| Medium Payload | MessagePack | Unmarshal | 58818 | 39455 | 733 |
| Medium Payload | CBOR | Unmarshal | 66762 | 30200 | 621 |
| Medium Payload | JSON | Unmarshal | 195028 | 47832 | 645 |
| Large Payload | BEVE ZeroCopy | Marshal | 121265 | 566 | 2 |
| Large Payload | BEVE | Marshal | 153257 | 197825 | 3 |
| Large Payload | Sonic | Marshal | 165252 | 218088 | 4 |
| Large Payload | CBOR | Marshal | 191663 | 181296 | 2 |
| Large Payload | MessagePack | Marshal | 316452 | 526856 | 115 |
| Large Payload | JSON | Marshal | 432428 | 206182 | 9 |
| Large Payload | BEVE | Unmarshal | 202342 | 157044 | 419 |
| Large Payload | Sonic | Unmarshal | 373085 | 573628 | 593 |
| Large Payload | MessagePack | Unmarshal | 589999 | 382933 | 7052 |
| Large Payload | CBOR | Unmarshal | 715455 | 330683 | 6746 |
| Large Payload | JSON | Unmarshal | 2205067 | 557458 | 7270 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmarks/benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 516.20 | 288 | 2 |
| Small Struct | BEVE | Marshal | 1483 | 2083 | 3 |
| Small Struct | CBOR | Marshal | 1566 | 1553 | 2 |
| Small Struct | JSON | Marshal | 3600 | 2193 | 2 |
| Small Struct | Sonic | Marshal | 3658 | 2905 | 3 |
| Small Struct | MessagePack | Marshal | 3881 | 8322 | 9 |
| Small Struct | BEVE | Unmarshal | 712.40 | 408 | 4 |
| Small Struct | Sonic | Unmarshal | 1411 | 1791 | 6 |
| Small Struct | JSON | Unmarshal | 4339 | 896 | 22 |
| Small Struct | MessagePack | Unmarshal | 5779 | 4641 | 97 |
| Small Struct | CBOR | Unmarshal | 8050 | 5192 | 107 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9124 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 12945 | 19428 | 3 |
| Medium Payload | CBOR | Marshal | 16830 | 18519 | 2 |
| Medium Payload | MessagePack | Marshal | 22428 | 33063 | 21 |
| Medium Payload | Sonic | Marshal | 28603 | 20925 | 4 |
| Medium Payload | JSON | Marshal | 43105 | 24895 | 9 |
| Medium Payload | BEVE | Unmarshal | 19598 | 15484 | 59 |
| Medium Payload | Sonic | Unmarshal | 27125 | 34322 | 33 |
| Medium Payload | MessagePack | Unmarshal | 62430 | 47153 | 883 |
| Medium Payload | CBOR | Unmarshal | 63470 | 31177 | 646 |
| Medium Payload | JSON | Unmarshal | 186478 | 51209 | 673 |
| Large Payload | BEVE ZeroCopy | Marshal | 89188 | 479 | 2 |
| Large Payload | BEVE | Marshal | 141350 | 205295 | 3 |
| Large Payload | CBOR | Marshal | 191275 | 205891 | 2 |
| Large Payload | MessagePack | Marshal | 270398 | 526866 | 115 |
| Large Payload | Sonic | Marshal | 314852 | 224816 | 4 |
| Large Payload | JSON | Marshal | 365578 | 197638 | 9 |
| Large Payload | BEVE | Unmarshal | 181946 | 152987 | 419 |
| Large Payload | Sonic | Unmarshal | 286168 | 384122 | 213 |
| Large Payload | MessagePack | Unmarshal | 504028 | 351766 | 6401 |
| Large Payload | CBOR | Unmarshal | 639972 | 323851 | 6593 |
| Large Payload | JSON | Unmarshal | 1997108 | 546253 | 7164 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | Sonic | Marshal | 659.50 | 664 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 816.70 | 288 | 2 |
| Small Struct | BEVE | Marshal | 1504 | 992 | 3 |
| Small Struct | CBOR | Marshal | 2554 | 2192 | 2 |
| Small Struct | MessagePack | Marshal | 3584 | 4224 | 8 |
| Small Struct | JSON | Marshal | 3774 | 1680 | 2 |
| Small Struct | Sonic | Unmarshal | 897.50 | 592 | 5 |
| Small Struct | BEVE | Unmarshal | 1300 | 1080 | 4 |
| Small Struct | CBOR | Unmarshal | 6375 | 2792 | 60 |
| Small Struct | MessagePack | Unmarshal | 7283 | 4032 | 86 |
| Small Struct | JSON | Unmarshal | 14451 | 3784 | 55 |
| Medium Payload | BEVE ZeroCopy | Marshal | 12453 | 148 | 2 |
| Medium Payload | CBOR | Marshal | 19891 | 16472 | 2 |
| Medium Payload | Sonic | Marshal | 20551 | 22344 | 4 |
| Medium Payload | BEVE | Marshal | 20594 | 19381 | 3 |
| Medium Payload | MessagePack | Marshal | 45748 | 65830 | 22 |
| Medium Payload | JSON | Marshal | 53837 | 24881 | 9 |
| Medium Payload | BEVE | Unmarshal | 24530 | 14827 | 59 |
| Medium Payload | Sonic | Unmarshal | 47169 | 53406 | 75 |
| Medium Payload | CBOR | Unmarshal | 61828 | 21272 | 443 |
| Medium Payload | MessagePack | Unmarshal | 71828 | 36782 | 681 |
| Medium Payload | JSON | Unmarshal | 258626 | 56681 | 759 |
| Large Payload | BEVE ZeroCopy | Marshal | 115918 | 479 | 2 |
| Large Payload | Sonic | Marshal | 177427 | 229920 | 4 |
| Large Payload | BEVE | Marshal | 180134 | 196035 | 3 |
| Large Payload | CBOR | Marshal | 227054 | 192218 | 2 |
| Large Payload | MessagePack | Marshal | 307669 | 526770 | 115 |
| Large Payload | JSON | Marshal | 497878 | 215062 | 9 |
| Large Payload | BEVE | Unmarshal | 246187 | 164553 | 419 |
| Large Payload | Sonic | Unmarshal | 462263 | 572507 | 599 |
| Large Payload | MessagePack | Unmarshal | 670365 | 356586 | 6518 |
| Large Payload | CBOR | Unmarshal | 794306 | 298411 | 6086 |
| Large Payload | JSON | Unmarshal | 2432572 | 521944 | 6832 |

