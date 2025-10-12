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
| Small Struct | BEVE ZeroCopy | Marshal | 235.00 | 288 | 2 |
| Small Struct | JSON | Marshal | 583.10 | 464 | 2 |
| Small Struct | CBOR | Marshal | 607.60 | 848 | 2 |
| Small Struct | MessagePack | Marshal | 698.80 | 1152 | 6 |
| Small Struct | BEVE | Marshal | 1747 | 2977 | 3 |
| Small Struct | Sonic | Marshal | 2899 | 1469 | 3 |
| Small Struct | BEVE | Unmarshal | 897.00 | 1464 | 4 |
| Small Struct | CBOR | Unmarshal | 1588 | 952 | 23 |
| Small Struct | Sonic | Unmarshal | 1796 | 1722 | 6 |
| Small Struct | MessagePack | Unmarshal | 3881 | 3880 | 81 |
| Small Struct | JSON | Unmarshal | 5276 | 1448 | 32 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7324 | 128 | 2 |
| Medium Payload | BEVE | Marshal | 10676 | 16514 | 3 |
| Medium Payload | CBOR | Marshal | 20412 | 21847 | 2 |
| Medium Payload | MessagePack | Marshal | 32239 | 65833 | 22 |
| Medium Payload | Sonic | Marshal | 43515 | 22008 | 4 |
| Medium Payload | JSON | Marshal | 50740 | 27569 | 9 |
| Medium Payload | BEVE | Unmarshal | 18767 | 15850 | 59 |
| Medium Payload | Sonic | Unmarshal | 34280 | 35554 | 33 |
| Medium Payload | MessagePack | Unmarshal | 47895 | 32094 | 589 |
| Medium Payload | CBOR | Unmarshal | 56783 | 30536 | 621 |
| Medium Payload | JSON | Unmarshal | 205944 | 50776 | 666 |
| Large Payload | BEVE ZeroCopy | Marshal | 116291 | 478 | 2 |
| Large Payload | BEVE | Marshal | 172184 | 188927 | 3 |
| Large Payload | MessagePack | Marshal | 223009 | 526832 | 115 |
| Large Payload | CBOR | Marshal | 238545 | 189800 | 2 |
| Large Payload | JSON | Marshal | 444829 | 214018 | 9 |
| Large Payload | Sonic | Marshal | 461790 | 206869 | 4 |
| Large Payload | BEVE | Unmarshal | 160292 | 154482 | 418 |
| Large Payload | Sonic | Unmarshal | 313669 | 335231 | 209 |
| Large Payload | MessagePack | Unmarshal | 556732 | 347881 | 6342 |
| Large Payload | CBOR | Unmarshal | 665491 | 336462 | 6868 |
| Large Payload | JSON | Unmarshal | 1933466 | 545308 | 7076 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 690.10 | 288 | 2 |
| Small Struct | CBOR | Marshal | 1070 | 912 | 2 |
| Small Struct | Sonic | Marshal | 1239 | 1597 | 3 |
| Small Struct | BEVE | Marshal | 1689 | 1824 | 3 |
| Small Struct | MessagePack | Marshal | 2559 | 4224 | 8 |
| Small Struct | JSON | Marshal | 3203 | 1680 | 2 |
| Small Struct | BEVE | Unmarshal | 994.90 | 632 | 4 |
| Small Struct | Sonic | Unmarshal | 2701 | 4176 | 9 |
| Small Struct | MessagePack | Unmarshal | 4133 | 2904 | 63 |
| Small Struct | CBOR | Unmarshal | 6270 | 3176 | 68 |
| Small Struct | JSON | Unmarshal | 20429 | 4840 | 88 |
| Medium Payload | BEVE ZeroCopy | Marshal | 15343 | 141 | 2 |
| Medium Payload | Sonic | Marshal | 16225 | 22648 | 4 |
| Medium Payload | BEVE | Marshal | 17535 | 18618 | 3 |
| Medium Payload | CBOR | Marshal | 24873 | 24827 | 2 |
| Medium Payload | MessagePack | Marshal | 35175 | 65839 | 22 |
| Medium Payload | JSON | Marshal | 45781 | 22170 | 9 |
| Medium Payload | BEVE | Unmarshal | 19249 | 14075 | 59 |
| Medium Payload | Sonic | Unmarshal | 35668 | 52876 | 72 |
| Medium Payload | MessagePack | Unmarshal | 65407 | 43856 | 822 |
| Medium Payload | CBOR | Unmarshal | 90534 | 33256 | 680 |
| Medium Payload | JSON | Unmarshal | 201696 | 51624 | 675 |
| Large Payload | BEVE ZeroCopy | Marshal | 135646 | 479 | 2 |
| Large Payload | BEVE | Marshal | 174962 | 188849 | 3 |
| Large Payload | Sonic | Marshal | 175251 | 225544 | 4 |
| Large Payload | CBOR | Marshal | 219668 | 205884 | 2 |
| Large Payload | MessagePack | Marshal | 304718 | 526856 | 115 |
| Large Payload | JSON | Marshal | 439168 | 214026 | 9 |
| Large Payload | BEVE | Unmarshal | 187640 | 147471 | 417 |
| Large Payload | Sonic | Unmarshal | 350126 | 505819 | 560 |
| Large Payload | MessagePack | Unmarshal | 518042 | 315734 | 5682 |
| Large Payload | CBOR | Unmarshal | 727319 | 306457 | 6248 |
| Large Payload | JSON | Unmarshal | 2101470 | 521971 | 6819 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 882.80 | 288 | 2 |
| Small Struct | CBOR | Marshal | 1707 | 1937 | 2 |
| Small Struct | BEVE | Marshal | 2306 | 2977 | 3 |
| Small Struct | Sonic | Marshal | 2468 | 2080 | 3 |
| Small Struct | MessagePack | Marshal | 2498 | 4224 | 8 |
| Small Struct | JSON | Marshal | 4426 | 2834 | 2 |
| Small Struct | BEVE | Unmarshal | 674.80 | 376 | 4 |
| Small Struct | Sonic | Unmarshal | 3447 | 5619 | 6 |
| Small Struct | CBOR | Unmarshal | 3586 | 1896 | 42 |
| Small Struct | MessagePack | Unmarshal | 3862 | 3072 | 64 |
| Small Struct | JSON | Unmarshal | 16186 | 4424 | 75 |
| Medium Payload | BEVE ZeroCopy | Marshal | 10792 | 154 | 2 |
| Medium Payload | BEVE | Marshal | 14370 | 19218 | 3 |
| Medium Payload | CBOR | Marshal | 17649 | 19200 | 2 |
| Medium Payload | MessagePack | Marshal | 28870 | 65838 | 22 |
| Medium Payload | Sonic | Marshal | 34317 | 27785 | 4 |
| Medium Payload | JSON | Marshal | 46779 | 27597 | 9 |
| Medium Payload | BEVE | Unmarshal | 19305 | 16091 | 59 |
| Medium Payload | Sonic | Unmarshal | 27341 | 37882 | 33 |
| Medium Payload | MessagePack | Unmarshal | 57907 | 42849 | 807 |
| Medium Payload | CBOR | Unmarshal | 67301 | 34616 | 713 |
| Medium Payload | JSON | Unmarshal | 218915 | 63705 | 822 |
| Large Payload | BEVE ZeroCopy | Marshal | 108999 | 654 | 2 |
| Large Payload | BEVE | Marshal | 136475 | 172988 | 3 |
| Large Payload | CBOR | Marshal | 182265 | 197521 | 2 |
| Large Payload | MessagePack | Marshal | 255893 | 526869 | 115 |
| Large Payload | Sonic | Marshal | 293053 | 215502 | 4 |
| Large Payload | JSON | Marshal | 352228 | 197288 | 9 |
| Large Payload | BEVE | Unmarshal | 183143 | 160509 | 418 |
| Large Payload | Sonic | Unmarshal | 272489 | 374912 | 211 |
| Large Payload | MessagePack | Unmarshal | 482827 | 336708 | 6102 |
| Large Payload | CBOR | Unmarshal | 628763 | 317117 | 6472 |
| Large Payload | JSON | Unmarshal | 1964529 | 543780 | 7065 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 801.50 | 288 | 2 |
| Small Struct | Sonic | Marshal | 1729 | 1482 | 3 |
| Small Struct | CBOR | Marshal | 1875 | 1553 | 2 |
| Small Struct | BEVE | Marshal | 2793 | 1569 | 3 |
| Small Struct | MessagePack | Marshal | 3712 | 4224 | 8 |
| Small Struct | JSON | Marshal | 3877 | 1682 | 2 |
| Small Struct | BEVE | Unmarshal | 1529 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 2162 | 2235 | 8 |
| Small Struct | MessagePack | Unmarshal | 7225 | 4384 | 93 |
| Small Struct | CBOR | Unmarshal | 9912 | 4592 | 96 |
| Small Struct | JSON | Unmarshal | 18494 | 4304 | 71 |
| Medium Payload | BEVE ZeroCopy | Marshal | 14683 | 148 | 2 |
| Medium Payload | CBOR | Marshal | 21143 | 16486 | 2 |
| Medium Payload | Sonic | Marshal | 22909 | 25176 | 4 |
| Medium Payload | BEVE | Marshal | 23456 | 18586 | 3 |
| Medium Payload | MessagePack | Marshal | 42364 | 65828 | 22 |
| Medium Payload | JSON | Marshal | 57951 | 24933 | 9 |
| Medium Payload | BEVE | Unmarshal | 24186 | 14555 | 59 |
| Medium Payload | Sonic | Unmarshal | 42743 | 48050 | 69 |
| Medium Payload | MessagePack | Unmarshal | 55182 | 27228 | 488 |
| Medium Payload | CBOR | Unmarshal | 83283 | 29784 | 615 |
| Medium Payload | JSON | Unmarshal | 236778 | 53064 | 669 |
| Large Payload | BEVE ZeroCopy | Marshal | 156533 | 479 | 2 |
| Large Payload | Sonic | Marshal | 168193 | 227948 | 4 |
| Large Payload | CBOR | Marshal | 197150 | 182439 | 2 |
| Large Payload | BEVE | Marshal | 254542 | 197460 | 3 |
| Large Payload | MessagePack | Marshal | 305078 | 526770 | 115 |
| Large Payload | JSON | Marshal | 449019 | 197795 | 9 |
| Large Payload | BEVE | Unmarshal | 234576 | 157321 | 416 |
| Large Payload | Sonic | Unmarshal | 418314 | 513697 | 556 |
| Large Payload | MessagePack | Unmarshal | 625677 | 336833 | 6101 |
| Large Payload | CBOR | Unmarshal | 848976 | 315196 | 6435 |
| Large Payload | JSON | Unmarshal | 2433474 | 538446 | 7040 |

