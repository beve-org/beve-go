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
| Large Payload | BEVE | Unmarshal | 1260586 | 230620 | 418 |
| Large Payload | Sonic | Unmarshal | 1504718 | 348511 | 213 |
| Large Payload | MessagePack | Unmarshal | 1959816 | 341060 | 6194 |
| Large Payload | CBOR | Unmarshal | 2001510 | 307403 | 6283 |
| Large Payload | JSON | Unmarshal | 3608094 | 526405 | 6835 |
| Small Struct | BEVE ZeroCopy | Marshal | 751.20 | 288 | 2 |
| Small Struct | CBOR | Marshal | 1507 | 1424 | 2 |
| Small Struct | BEVE | Marshal | 7059 | 3361 | 3 |
| Small Struct | JSON | Marshal | 16112 | 2448 | 2 |
| Small Struct | MessagePack | Marshal | 22686 | 8321 | 9 |
| Small Struct | Sonic | Marshal | 46383 | 1596 | 3 |
| Large Payload | BEVE | Marshal | 806335 | 188936 | 3 |
| Large Payload | CBOR | Marshal | 1053357 | 197746 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 1055640 | 233 | 2 |
| Large Payload | MessagePack | Marshal | 1234993 | 526800 | 115 |
| Large Payload | Sonic | Marshal | 1895620 | 218256 | 4 |
| Large Payload | JSON | Marshal | 1956906 | 214304 | 9 |
| Medium Payload | BEVE | Unmarshal | 178901 | 26108 | 59 |
| Medium Payload | Sonic | Unmarshal | 261817 | 33198 | 33 |
| Medium Payload | CBOR | Unmarshal | 312051 | 28200 | 579 |
| Medium Payload | MessagePack | Unmarshal | 321578 | 33165 | 611 |
| Medium Payload | JSON | Unmarshal | 936377 | 44280 | 598 |
| Small Struct | BEVE | Unmarshal | 25452 | 3512 | 4 |
| Small Struct | Sonic | Unmarshal | 28790 | 691 | 6 |
| Small Struct | CBOR | Unmarshal | 35648 | 4776 | 102 |
| Small Struct | MessagePack | Unmarshal | 48168 | 3512 | 74 |
| Small Struct | JSON | Unmarshal | 297455 | 7880 | 112 |
| Medium Payload | BEVE | Marshal | 177910 | 24717 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 183181 | 133 | 2 |
| Medium Payload | CBOR | Marshal | 251204 | 21859 | 2 |
| Medium Payload | MessagePack | Marshal | 310870 | 65830 | 22 |
| Medium Payload | Sonic | Marshal | 369196 | 19472 | 4 |
| Medium Payload | JSON | Marshal | 478762 | 24906 | 9 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1350640 | 240044 | 418 |
| Large Payload | Sonic | Unmarshal | 1762269 | 554427 | 576 |
| Large Payload | MessagePack | Unmarshal | 1923763 | 357852 | 6534 |
| Large Payload | CBOR | Unmarshal | 2108138 | 308410 | 6286 |
| Large Payload | JSON | Unmarshal | 4004613 | 526754 | 6984 |
| Small Struct | Sonic | Marshal | 3054 | 714 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 3937 | 288 | 2 |
| Small Struct | CBOR | Marshal | 11526 | 3218 | 2 |
| Small Struct | BEVE | Marshal | 14308 | 2593 | 3 |
| Small Struct | MessagePack | Marshal | 15276 | 4224 | 8 |
| Small Struct | JSON | Marshal | 23142 | 1936 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 590386 | 260 | 2 |
| Large Payload | Sonic | Marshal | 709007 | 211700 | 4 |
| Large Payload | BEVE | Marshal | 709758 | 205395 | 3 |
| Large Payload | CBOR | Marshal | 972728 | 189014 | 2 |
| Large Payload | MessagePack | Marshal | 1441579 | 526786 | 115 |
| Large Payload | JSON | Marshal | 2301548 | 222133 | 9 |
| Medium Payload | BEVE | Unmarshal | 140089 | 26108 | 59 |
| Medium Payload | Sonic | Unmarshal | 209747 | 49688 | 66 |
| Medium Payload | MessagePack | Unmarshal | 387466 | 46513 | 879 |
| Medium Payload | CBOR | Unmarshal | 400762 | 31208 | 643 |
| Medium Payload | JSON | Unmarshal | 1351999 | 63224 | 842 |
| Small Struct | Sonic | Unmarshal | 2352 | 478 | 4 |
| Small Struct | BEVE | Unmarshal | 9041 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 14150 | 1312 | 30 |
| Small Struct | CBOR | Unmarshal | 20700 | 1448 | 33 |
| Small Struct | JSON | Unmarshal | 61310 | 2312 | 44 |
| Medium Payload | BEVE ZeroCopy | Marshal | 69556 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 78514 | 18571 | 3 |
| Medium Payload | Sonic | Marshal | 78836 | 22652 | 4 |
| Medium Payload | CBOR | Marshal | 98199 | 19153 | 2 |
| Medium Payload | MessagePack | Marshal | 210489 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 245626 | 20785 | 9 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | MessagePack | Unmarshal | 687436 | 353821 | 6456 |
| Large Payload | BEVE | Unmarshal | 723007 | 237708 | 419 |
| Large Payload | Sonic | Unmarshal | 745892 | 389692 | 213 |
| Large Payload | CBOR | Unmarshal | 963837 | 304699 | 6217 |
| Large Payload | JSON | Unmarshal | 2583708 | 553678 | 7305 |
| Small Struct | BEVE | Marshal | 1529 | 928 | 3 |
| Small Struct | CBOR | Marshal | 2292 | 560 | 2 |
| Small Struct | Sonic | Marshal | 3653 | 1364 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 4100 | 289 | 2 |
| Small Struct | JSON | Marshal | 8513 | 3218 | 2 |
| Small Struct | MessagePack | Marshal | 9416 | 8321 | 9 |
| Large Payload | BEVE ZeroCopy | Marshal | 103445 | 286 | 2 |
| Large Payload | BEVE | Marshal | 377412 | 197284 | 3 |
| Large Payload | CBOR | Marshal | 610706 | 205677 | 2 |
| Large Payload | MessagePack | Marshal | 883193 | 526794 | 115 |
| Large Payload | Sonic | Marshal | 888032 | 230200 | 4 |
| Large Payload | JSON | Marshal | 1087003 | 230586 | 9 |
| Medium Payload | BEVE | Unmarshal | 79611 | 27292 | 59 |
| Medium Payload | Sonic | Unmarshal | 80381 | 42461 | 33 |
| Medium Payload | MessagePack | Unmarshal | 155165 | 34494 | 636 |
| Medium Payload | CBOR | Unmarshal | 226758 | 37912 | 773 |
| Medium Payload | JSON | Unmarshal | 597657 | 56760 | 740 |
| Small Struct | MessagePack | Unmarshal | 2972 | 544 | 14 |
| Small Struct | Sonic | Unmarshal | 3033 | 1563 | 6 |
| Small Struct | BEVE | Unmarshal | 4868 | 3000 | 4 |
| Small Struct | JSON | Unmarshal | 23521 | 1992 | 34 |
| Small Struct | CBOR | Unmarshal | 24826 | 5136 | 105 |
| Medium Payload | BEVE ZeroCopy | Marshal | 14525 | 134 | 2 |
| Medium Payload | CBOR | Marshal | 56121 | 18521 | 2 |
| Medium Payload | BEVE | Marshal | 56841 | 24721 | 3 |
| Medium Payload | JSON | Marshal | 86676 | 20791 | 9 |
| Medium Payload | Sonic | Marshal | 88624 | 25492 | 4 |
| Medium Payload | MessagePack | Marshal | 93329 | 65832 | 22 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1157606 | 225754 | 418 |
| Large Payload | MessagePack | Unmarshal | 1377340 | 375800 | 6898 |
| Large Payload | Sonic | Unmarshal | 1440202 | 542780 | 562 |
| Large Payload | CBOR | Unmarshal | 1997668 | 330188 | 6728 |
| Large Payload | JSON | Unmarshal | 3539630 | 530085 | 6960 |
| Small Struct | BEVE ZeroCopy | Marshal | 2409 | 288 | 2 |
| Small Struct | JSON | Marshal | 5374 | 656 | 2 |
| Small Struct | BEVE | Marshal | 8759 | 2080 | 3 |
| Small Struct | Sonic | Marshal | 10929 | 1995 | 3 |
| Small Struct | CBOR | Marshal | 20776 | 3217 | 2 |
| Small Struct | MessagePack | Marshal | 28196 | 4224 | 8 |
| Large Payload | BEVE ZeroCopy | Marshal | 229896 | 259 | 2 |
| Large Payload | Sonic | Marshal | 632861 | 217564 | 4 |
| Large Payload | BEVE | Marshal | 655694 | 180814 | 3 |
| Large Payload | CBOR | Marshal | 799748 | 198056 | 2 |
| Large Payload | MessagePack | Marshal | 1087105 | 526763 | 115 |
| Large Payload | JSON | Marshal | 1385034 | 215245 | 9 |
| Medium Payload | BEVE | Unmarshal | 78669 | 21819 | 59 |
| Medium Payload | Sonic | Unmarshal | 167361 | 48582 | 69 |
| Medium Payload | MessagePack | Unmarshal | 261510 | 30925 | 566 |
| Medium Payload | CBOR | Unmarshal | 409771 | 42504 | 877 |
| Medium Payload | JSON | Unmarshal | 554760 | 41816 | 529 |
| Small Struct | MessagePack | Unmarshal | 7332 | 1088 | 25 |
| Small Struct | Sonic | Unmarshal | 11219 | 3632 | 9 |
| Small Struct | BEVE | Unmarshal | 17098 | 3512 | 4 |
| Small Struct | CBOR | Unmarshal | 39188 | 5088 | 104 |
| Small Struct | JSON | Unmarshal | 52844 | 4264 | 70 |
| Medium Payload | BEVE ZeroCopy | Marshal | 23547 | 133 | 2 |
| Medium Payload | BEVE | Marshal | 52054 | 20625 | 3 |
| Medium Payload | Sonic | Marshal | 81818 | 25254 | 4 |
| Medium Payload | CBOR | Marshal | 108336 | 24675 | 2 |
| Medium Payload | MessagePack | Marshal | 128358 | 65829 | 22 |
| Medium Payload | JSON | Marshal | 159174 | 24910 | 9 |

