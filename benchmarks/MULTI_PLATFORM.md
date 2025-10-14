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

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1382391 | 282356 | 418 |
| Large Payload | Sonic | Unmarshal | 1785293 | 540912 | 564 |
| Large Payload | MessagePack | Unmarshal | 1960935 | 358303 | 6521 |
| Large Payload | CBOR | Unmarshal | 2137647 | 306299 | 6241 |
| Large Payload | JSON | Unmarshal | 4116227 | 566794 | 7444 |
| Small Struct | BEVE ZeroCopy | Marshal | 1962 | 289 | 2 |
| Small Struct | Sonic | Marshal | 2357 | 641 | 3 |
| Small Struct | JSON | Marshal | 5727 | 560 | 2 |
| Small Struct | BEVE | Marshal | 10888 | 2080 | 3 |
| Small Struct | CBOR | Marshal | 14369 | 2833 | 2 |
| Small Struct | MessagePack | Marshal | 15690 | 4224 | 8 |
| Large Payload | BEVE ZeroCopy | Marshal | 587562 | 259 | 2 |
| Large Payload | BEVE | Marshal | 604587 | 188772 | 3 |
| Large Payload | Sonic | Marshal | 708687 | 220765 | 4 |
| Large Payload | CBOR | Marshal | 1022516 | 205716 | 2 |
| Large Payload | MessagePack | Marshal | 1444610 | 526783 | 115 |
| Large Payload | JSON | Marshal | 2388444 | 221871 | 9 |
| Medium Payload | BEVE | Unmarshal | 153848 | 28445 | 59 |
| Medium Payload | Sonic | Unmarshal | 205493 | 47602 | 71 |
| Medium Payload | MessagePack | Unmarshal | 351165 | 42608 | 803 |
| Medium Payload | CBOR | Unmarshal | 458104 | 34752 | 716 |
| Medium Payload | JSON | Unmarshal | 1287432 | 55576 | 759 |
| Small Struct | Sonic | Unmarshal | 9428 | 2370 | 8 |
| Small Struct | BEVE | Unmarshal | 13115 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 13703 | 1664 | 37 |
| Small Struct | CBOR | Unmarshal | 17350 | 1256 | 29 |
| Small Struct | JSON | Unmarshal | 158065 | 7368 | 96 |
| Medium Payload | BEVE ZeroCopy | Marshal | 65699 | 136 | 2 |
| Medium Payload | BEVE | Marshal | 68894 | 18570 | 3 |
| Medium Payload | Sonic | Marshal | 100977 | 22646 | 4 |
| Medium Payload | CBOR | Marshal | 131632 | 20570 | 2 |
| Medium Payload | JSON | Marshal | 227312 | 18743 | 9 |
| Medium Payload | MessagePack | Marshal | 235382 | 65832 | 22 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | MessagePack | Unmarshal | 708161 | 355902 | 6494 |
| Large Payload | BEVE | Unmarshal | 758134 | 285493 | 417 |
| Large Payload | Sonic | Unmarshal | 762498 | 372324 | 209 |
| Large Payload | CBOR | Unmarshal | 1046717 | 324556 | 6608 |
| Large Payload | JSON | Unmarshal | 2647401 | 573884 | 7450 |
| Small Struct | Sonic | Marshal | 1110 | 420 | 3 |
| Small Struct | BEVE | Marshal | 1622 | 2594 | 3 |
| Small Struct | CBOR | Marshal | 2384 | 1296 | 2 |
| Small Struct | MessagePack | Marshal | 3255 | 1152 | 6 |
| Small Struct | BEVE ZeroCopy | Marshal | 4307 | 289 | 2 |
| Small Struct | JSON | Marshal | 8403 | 1168 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 199601 | 259 | 2 |
| Large Payload | BEVE | Marshal | 376586 | 188907 | 3 |
| Large Payload | CBOR | Marshal | 553098 | 197952 | 2 |
| Large Payload | MessagePack | Marshal | 842551 | 526798 | 115 |
| Large Payload | Sonic | Marshal | 935884 | 230658 | 4 |
| Large Payload | JSON | Marshal | 1054072 | 205745 | 9 |
| Medium Payload | Sonic | Unmarshal | 87100 | 42709 | 33 |
| Medium Payload | BEVE | Unmarshal | 87449 | 31933 | 59 |
| Medium Payload | MessagePack | Unmarshal | 145992 | 31822 | 583 |
| Medium Payload | CBOR | Unmarshal | 224384 | 35928 | 738 |
| Medium Payload | JSON | Unmarshal | 572304 | 48889 | 652 |
| Small Struct | BEVE | Unmarshal | 1444 | 600 | 4 |
| Small Struct | Sonic | Unmarshal | 8452 | 3075 | 6 |
| Small Struct | CBOR | Unmarshal | 9707 | 4360 | 93 |
| Small Struct | MessagePack | Unmarshal | 19796 | 5153 | 105 |
| Small Struct | JSON | Unmarshal | 66356 | 7304 | 94 |
| Medium Payload | BEVE ZeroCopy | Marshal | 13513 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 50021 | 24721 | 3 |
| Medium Payload | CBOR | Marshal | 76166 | 24661 | 2 |
| Medium Payload | JSON | Marshal | 97949 | 19376 | 9 |
| Medium Payload | MessagePack | Marshal | 107752 | 65832 | 22 |
| Medium Payload | Sonic | Marshal | 124409 | 28402 | 4 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1334329 | 262017 | 418 |
| Large Payload | Sonic | Unmarshal | 1412887 | 523827 | 568 |
| Large Payload | MessagePack | Unmarshal | 1479432 | 342544 | 6230 |
| Large Payload | CBOR | Unmarshal | 1848957 | 308635 | 6285 |
| Large Payload | JSON | Unmarshal | 3851570 | 541838 | 7137 |
| Small Struct | BEVE ZeroCopy | Marshal | 3743 | 289 | 2 |
| Small Struct | JSON | Marshal | 10236 | 1296 | 2 |
| Small Struct | BEVE | Marshal | 10292 | 1568 | 3 |
| Small Struct | CBOR | Marshal | 12309 | 3217 | 2 |
| Small Struct | Sonic | Marshal | 12889 | 2539 | 3 |
| Small Struct | MessagePack | Marshal | 16894 | 4224 | 8 |
| Large Payload | BEVE ZeroCopy | Marshal | 228394 | 259 | 2 |
| Large Payload | Sonic | Marshal | 577903 | 209714 | 4 |
| Large Payload | CBOR | Marshal | 772214 | 181553 | 2 |
| Large Payload | BEVE | Marshal | 985823 | 205073 | 3 |
| Large Payload | MessagePack | Marshal | 1130270 | 526759 | 115 |
| Large Payload | JSON | Marshal | 1433619 | 215191 | 9 |
| Medium Payload | BEVE | Unmarshal | 90366 | 29212 | 59 |
| Medium Payload | MessagePack | Unmarshal | 215159 | 30619 | 555 |
| Medium Payload | Sonic | Unmarshal | 229713 | 47756 | 68 |
| Medium Payload | CBOR | Unmarshal | 308817 | 33432 | 682 |
| Medium Payload | JSON | Unmarshal | 751920 | 57961 | 789 |
| Small Struct | BEVE | Unmarshal | 13315 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 13987 | 2432 | 52 |
| Small Struct | JSON | Unmarshal | 20999 | 2184 | 40 |
| Small Struct | CBOR | Unmarshal | 23640 | 3944 | 84 |
| Small Struct | Sonic | Unmarshal | 27687 | 4173 | 9 |
| Medium Payload | BEVE ZeroCopy | Marshal | 24631 | 136 | 2 |
| Medium Payload | Sonic | Marshal | 81018 | 22272 | 4 |
| Medium Payload | CBOR | Marshal | 81969 | 18517 | 2 |
| Medium Payload | BEVE | Marshal | 104659 | 21908 | 3 |
| Medium Payload | MessagePack | Marshal | 151562 | 65828 | 22 |
| Medium Payload | JSON | Marshal | 169458 | 22074 | 9 |

