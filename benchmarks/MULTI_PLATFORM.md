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
| Large Payload | BEVE | Unmarshal | 1495027 | 156594 | 418 |
| Large Payload | Sonic | Unmarshal | 2162086 | 354661 | 213 |
| Large Payload | MessagePack | Unmarshal | 2665831 | 351540 | 6400 |
| Large Payload | CBOR | Unmarshal | 2750677 | 342347 | 6978 |
| Large Payload | JSON | Unmarshal | 4636869 | 574964 | 7519 |
| Small Struct | MessagePack | Marshal | 23282 | 4224 | 8 |
| Small Struct | BEVE ZeroCopy | Marshal | 32738 | 288 | 2 |
| Small Struct | BEVE | Marshal | 33446 | 2336 | 3 |
| Small Struct | CBOR | Marshal | 35227 | 2834 | 2 |
| Small Struct | JSON | Marshal | 50296 | 1296 | 2 |
| Small Struct | Sonic | Marshal | 92363 | 2915 | 3 |
| Large Payload | BEVE | Marshal | 1443112 | 197438 | 3 |
| Large Payload | CBOR | Marshal | 1539411 | 189521 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 1615624 | 163 | 2 |
| Large Payload | MessagePack | Marshal | 1718373 | 526797 | 115 |
| Large Payload | JSON | Marshal | 2454983 | 198145 | 9 |
| Large Payload | Sonic | Marshal | 2583096 | 201683 | 4 |
| Medium Payload | BEVE | Unmarshal | 282412 | 16090 | 59 |
| Medium Payload | Sonic | Unmarshal | 471887 | 35624 | 33 |
| Medium Payload | MessagePack | Unmarshal | 547202 | 30508 | 554 |
| Medium Payload | CBOR | Unmarshal | 788578 | 38424 | 791 |
| Medium Payload | JSON | Unmarshal | 1483106 | 48608 | 653 |
| Small Struct | BEVE | Unmarshal | 24046 | 664 | 4 |
| Small Struct | Sonic | Unmarshal | 32076 | 848 | 6 |
| Small Struct | MessagePack | Unmarshal | 76996 | 3256 | 70 |
| Small Struct | CBOR | Unmarshal | 144898 | 4808 | 103 |
| Small Struct | JSON | Unmarshal | 181838 | 2056 | 36 |
| Medium Payload | BEVE | Marshal | 219270 | 16519 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 224626 | 131 | 2 |
| Medium Payload | CBOR | Marshal | 266265 | 20586 | 2 |
| Medium Payload | MessagePack | Marshal | 501581 | 65831 | 22 |
| Medium Payload | Sonic | Marshal | 540832 | 21005 | 4 |
| Medium Payload | JSON | Marshal | 666444 | 24901 | 9 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1023440 | 148135 | 419 |
| Large Payload | Sonic | Unmarshal | 1514899 | 556596 | 584 |
| Large Payload | MessagePack | Unmarshal | 1715252 | 355146 | 6471 |
| Large Payload | CBOR | Unmarshal | 1919065 | 300459 | 6129 |
| Large Payload | JSON | Unmarshal | 3445124 | 502274 | 6621 |
| Small Struct | Sonic | Marshal | 5641 | 1475 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 7993 | 288 | 2 |
| Small Struct | CBOR | Marshal | 11862 | 2834 | 2 |
| Small Struct | MessagePack | Marshal | 13461 | 4224 | 8 |
| Small Struct | JSON | Marshal | 21024 | 1680 | 2 |
| Small Struct | BEVE | Marshal | 21682 | 2977 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 721776 | 189 | 2 |
| Large Payload | Sonic | Marshal | 752952 | 229396 | 4 |
| Large Payload | BEVE | Marshal | 972375 | 188865 | 3 |
| Large Payload | CBOR | Marshal | 987618 | 197520 | 2 |
| Large Payload | MessagePack | Marshal | 1326207 | 526781 | 115 |
| Large Payload | JSON | Marshal | 2101180 | 222010 | 9 |
| Medium Payload | BEVE | Unmarshal | 126078 | 15931 | 59 |
| Medium Payload | Sonic | Unmarshal | 251845 | 68008 | 80 |
| Medium Payload | MessagePack | Unmarshal | 324592 | 38848 | 723 |
| Medium Payload | CBOR | Unmarshal | 364745 | 26296 | 545 |
| Medium Payload | JSON | Unmarshal | 998595 | 45432 | 578 |
| Small Struct | MessagePack | Unmarshal | 7234 | 352 | 10 |
| Small Struct | BEVE | Unmarshal | 7593 | 1464 | 4 |
| Small Struct | Sonic | Unmarshal | 10940 | 2231 | 8 |
| Small Struct | CBOR | Unmarshal | 66266 | 4808 | 103 |
| Small Struct | JSON | Unmarshal | 134757 | 7272 | 93 |
| Medium Payload | Sonic | Marshal | 85504 | 19699 | 4 |
| Medium Payload | BEVE ZeroCopy | Marshal | 90591 | 132 | 2 |
| Medium Payload | BEVE | Marshal | 94707 | 18575 | 3 |
| Medium Payload | CBOR | Marshal | 120792 | 20576 | 2 |
| Medium Payload | MessagePack | Marshal | 139757 | 33060 | 21 |
| Medium Payload | JSON | Marshal | 212872 | 18735 | 9 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 599700 | 151624 | 419 |
| Large Payload | Sonic | Unmarshal | 684956 | 401543 | 211 |
| Large Payload | MessagePack | Unmarshal | 767654 | 352508 | 6432 |
| Large Payload | CBOR | Unmarshal | 1079307 | 321211 | 6545 |
| Large Payload | JSON | Unmarshal | 2461422 | 523380 | 6858 |
| Small Struct | CBOR | Marshal | 2750 | 624 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 3410 | 289 | 2 |
| Small Struct | BEVE | Marshal | 4659 | 1697 | 3 |
| Small Struct | MessagePack | Marshal | 5250 | 2176 | 7 |
| Small Struct | JSON | Marshal | 6392 | 2449 | 2 |
| Small Struct | Sonic | Marshal | 6968 | 2280 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 382165 | 224 | 2 |
| Large Payload | BEVE | Marshal | 564402 | 189018 | 3 |
| Large Payload | CBOR | Marshal | 614368 | 214149 | 2 |
| Large Payload | MessagePack | Marshal | 864755 | 526794 | 115 |
| Large Payload | Sonic | Marshal | 905517 | 229324 | 4 |
| Large Payload | JSON | Marshal | 1055974 | 205656 | 9 |
| Medium Payload | BEVE | Unmarshal | 62540 | 14762 | 59 |
| Medium Payload | Sonic | Unmarshal | 71868 | 29007 | 33 |
| Medium Payload | MessagePack | Unmarshal | 187567 | 41631 | 783 |
| Medium Payload | CBOR | Unmarshal | 209577 | 30744 | 634 |
| Medium Payload | JSON | Unmarshal | 650629 | 56801 | 752 |
| Small Struct | BEVE | Unmarshal | 2123 | 456 | 4 |
| Small Struct | CBOR | Unmarshal | 3012 | 280 | 9 |
| Small Struct | Sonic | Unmarshal | 12603 | 5543 | 6 |
| Small Struct | JSON | Unmarshal | 13035 | 872 | 21 |
| Small Struct | MessagePack | Unmarshal | 17288 | 3672 | 79 |
| Medium Payload | BEVE ZeroCopy | Marshal | 14312 | 132 | 2 |
| Medium Payload | CBOR | Marshal | 70468 | 20571 | 2 |
| Medium Payload | BEVE | Marshal | 78888 | 21910 | 3 |
| Medium Payload | MessagePack | Marshal | 87429 | 33060 | 21 |
| Medium Payload | Sonic | Marshal | 97889 | 21346 | 4 |
| Medium Payload | JSON | Marshal | 132661 | 24882 | 9 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1204057 | 153357 | 419 |
| Large Payload | Sonic | Unmarshal | 1531783 | 519396 | 557 |
| Large Payload | MessagePack | Unmarshal | 1661197 | 322873 | 5830 |
| Large Payload | CBOR | Unmarshal | 1909690 | 305035 | 6226 |
| Large Payload | JSON | Unmarshal | 3905786 | 523086 | 6802 |
| Small Struct | BEVE | Marshal | 5718 | 528 | 3 |
| Small Struct | CBOR | Marshal | 6698 | 624 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 7595 | 289 | 2 |
| Small Struct | MessagePack | Marshal | 8556 | 1152 | 6 |
| Small Struct | JSON | Marshal | 10397 | 1936 | 2 |
| Small Struct | Sonic | Marshal | 10695 | 3293 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 295413 | 163 | 2 |
| Large Payload | CBOR | Marshal | 1028476 | 206215 | 2 |
| Large Payload | Sonic | Marshal | 1056678 | 218242 | 4 |
| Large Payload | BEVE | Marshal | 1135576 | 180830 | 3 |
| Large Payload | MessagePack | Marshal | 1389102 | 526761 | 115 |
| Large Payload | JSON | Marshal | 1616141 | 223486 | 9 |
| Medium Payload | BEVE | Unmarshal | 89136 | 16474 | 59 |
| Medium Payload | Sonic | Unmarshal | 311279 | 52072 | 68 |
| Medium Payload | CBOR | Unmarshal | 337183 | 33976 | 702 |
| Medium Payload | MessagePack | Unmarshal | 367625 | 43885 | 823 |
| Medium Payload | JSON | Unmarshal | 777874 | 53224 | 710 |
| Small Struct | BEVE | Unmarshal | 5086 | 1208 | 4 |
| Small Struct | MessagePack | Unmarshal | 9813 | 1384 | 31 |
| Small Struct | CBOR | Unmarshal | 10792 | 1168 | 27 |
| Small Struct | Sonic | Unmarshal | 11333 | 2335 | 8 |
| Small Struct | JSON | Unmarshal | 85132 | 7592 | 103 |
| Medium Payload | BEVE ZeroCopy | Marshal | 40512 | 132 | 2 |
| Medium Payload | CBOR | Marshal | 88365 | 13643 | 2 |
| Medium Payload | Sonic | Marshal | 110240 | 22353 | 4 |
| Medium Payload | BEVE | Marshal | 110560 | 18579 | 3 |
| Medium Payload | MessagePack | Marshal | 154249 | 65828 | 22 |
| Medium Payload | JSON | Marshal | 198780 | 27583 | 9 |

