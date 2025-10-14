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
| Large Payload | BEVE | Unmarshal | 1287437 | 280742 | 419 |
| Large Payload | Sonic | Unmarshal | 1695849 | 366114 | 209 |
| Large Payload | MessagePack | Unmarshal | 1895559 | 330754 | 5987 |
| Large Payload | CBOR | Unmarshal | 2144532 | 323643 | 6608 |
| Large Payload | JSON | Unmarshal | 3539083 | 528092 | 6855 |
| Small Struct | BEVE ZeroCopy | Marshal | 12599 | 288 | 2 |
| Small Struct | BEVE | Marshal | 26455 | 800 | 3 |
| Small Struct | MessagePack | Marshal | 28726 | 2176 | 7 |
| Small Struct | JSON | Marshal | 44671 | 1296 | 2 |
| Small Struct | CBOR | Marshal | 58951 | 2833 | 2 |
| Small Struct | Sonic | Marshal | 73936 | 1610 | 3 |
| Large Payload | BEVE | Marshal | 825877 | 180827 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 960816 | 164 | 2 |
| Large Payload | CBOR | Marshal | 1177940 | 197930 | 2 |
| Large Payload | MessagePack | Marshal | 1467060 | 526802 | 115 |
| Large Payload | JSON | Marshal | 1825695 | 205975 | 9 |
| Large Payload | Sonic | Marshal | 1994903 | 225568 | 4 |
| Medium Payload | BEVE | Unmarshal | 243423 | 31580 | 59 |
| Medium Payload | Sonic | Unmarshal | 317115 | 33008 | 33 |
| Medium Payload | MessagePack | Unmarshal | 394913 | 31548 | 576 |
| Medium Payload | CBOR | Unmarshal | 529274 | 30152 | 621 |
| Medium Payload | JSON | Unmarshal | 1701945 | 65848 | 853 |
| Small Struct | BEVE | Unmarshal | 21560 | 1336 | 4 |
| Small Struct | MessagePack | Unmarshal | 64586 | 3520 | 74 |
| Small Struct | Sonic | Unmarshal | 74886 | 4861 | 6 |
| Small Struct | CBOR | Unmarshal | 93849 | 3880 | 82 |
| Small Struct | JSON | Unmarshal | 145868 | 2312 | 44 |
| Medium Payload | BEVE | Marshal | 137895 | 20621 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 164848 | 130 | 2 |
| Medium Payload | MessagePack | Marshal | 173536 | 33060 | 21 |
| Medium Payload | CBOR | Marshal | 190371 | 20570 | 2 |
| Medium Payload | JSON | Marshal | 344609 | 20792 | 9 |
| Medium Payload | Sonic | Marshal | 371458 | 16820 | 4 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1267954 | 270096 | 418 |
| Large Payload | MessagePack | Unmarshal | 1596513 | 341913 | 6206 |
| Large Payload | Sonic | Unmarshal | 1654505 | 548416 | 571 |
| Large Payload | CBOR | Unmarshal | 1963504 | 325067 | 6628 |
| Large Payload | JSON | Unmarshal | 3821760 | 554260 | 7250 |
| Small Struct | BEVE ZeroCopy | Marshal | 1972 | 288 | 2 |
| Small Struct | Sonic | Marshal | 3105 | 373 | 3 |
| Small Struct | MessagePack | Marshal | 8017 | 2176 | 7 |
| Small Struct | CBOR | Marshal | 10629 | 2449 | 2 |
| Small Struct | JSON | Marshal | 11524 | 848 | 2 |
| Small Struct | BEVE | Marshal | 12120 | 2977 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 548740 | 180 | 2 |
| Large Payload | BEVE | Marshal | 655407 | 205261 | 3 |
| Large Payload | Sonic | Marshal | 746834 | 220714 | 4 |
| Large Payload | CBOR | Marshal | 1025010 | 205630 | 2 |
| Large Payload | MessagePack | Marshal | 1392365 | 526783 | 115 |
| Large Payload | JSON | Marshal | 2056029 | 221964 | 9 |
| Medium Payload | BEVE | Unmarshal | 139551 | 26140 | 59 |
| Medium Payload | Sonic | Unmarshal | 235074 | 63448 | 70 |
| Medium Payload | MessagePack | Unmarshal | 333216 | 40512 | 755 |
| Medium Payload | CBOR | Unmarshal | 425396 | 32008 | 662 |
| Medium Payload | JSON | Unmarshal | 1307143 | 57960 | 764 |
| Small Struct | BEVE | Unmarshal | 7177 | 1592 | 4 |
| Small Struct | MessagePack | Unmarshal | 13708 | 1224 | 28 |
| Small Struct | Sonic | Unmarshal | 14325 | 3924 | 9 |
| Small Struct | CBOR | Unmarshal | 20565 | 1192 | 28 |
| Small Struct | JSON | Unmarshal | 60110 | 2312 | 44 |
| Medium Payload | BEVE | Marshal | 75154 | 19209 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 77113 | 132 | 2 |
| Medium Payload | Sonic | Marshal | 118521 | 28201 | 4 |
| Medium Payload | CBOR | Marshal | 119427 | 20567 | 2 |
| Medium Payload | MessagePack | Marshal | 202743 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 245614 | 20794 | 9 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | Sonic | Unmarshal | 738316 | 373234 | 209 |
| Large Payload | BEVE | Unmarshal | 798942 | 273297 | 417 |
| Large Payload | MessagePack | Unmarshal | 966578 | 343960 | 6263 |
| Large Payload | CBOR | Unmarshal | 1183264 | 323723 | 6609 |
| Large Payload | JSON | Unmarshal | 2704958 | 590091 | 7510 |
| Small Struct | BEVE ZeroCopy | Marshal | 1662 | 288 | 2 |
| Small Struct | CBOR | Marshal | 3663 | 1040 | 2 |
| Small Struct | BEVE | Marshal | 5544 | 2080 | 3 |
| Small Struct | JSON | Marshal | 7592 | 1936 | 2 |
| Small Struct | Sonic | Marshal | 7997 | 2293 | 3 |
| Small Struct | MessagePack | Marshal | 15829 | 8321 | 9 |
| Large Payload | BEVE ZeroCopy | Marshal | 223496 | 170 | 2 |
| Large Payload | BEVE | Marshal | 377999 | 188792 | 3 |
| Large Payload | CBOR | Marshal | 516654 | 181442 | 2 |
| Large Payload | Sonic | Marshal | 861448 | 212863 | 4 |
| Large Payload | MessagePack | Marshal | 869775 | 526789 | 115 |
| Large Payload | JSON | Marshal | 1074291 | 230562 | 9 |
| Medium Payload | BEVE | Unmarshal | 82643 | 28893 | 59 |
| Medium Payload | Sonic | Unmarshal | 86319 | 39227 | 31 |
| Medium Payload | MessagePack | Unmarshal | 194434 | 44304 | 836 |
| Medium Payload | CBOR | Unmarshal | 232186 | 37720 | 769 |
| Medium Payload | JSON | Unmarshal | 495201 | 41048 | 568 |
| Small Struct | Sonic | Unmarshal | 2931 | 982 | 6 |
| Small Struct | BEVE | Unmarshal | 6859 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 9522 | 2080 | 45 |
| Small Struct | JSON | Unmarshal | 17592 | 1320 | 28 |
| Small Struct | CBOR | Unmarshal | 24429 | 4352 | 93 |
| Medium Payload | BEVE ZeroCopy | Marshal | 29547 | 131 | 2 |
| Medium Payload | BEVE | Marshal | 34894 | 16518 | 3 |
| Medium Payload | CBOR | Marshal | 57994 | 20567 | 2 |
| Medium Payload | Sonic | Marshal | 83328 | 21233 | 4 |
| Medium Payload | JSON | Marshal | 110551 | 19373 | 9 |
| Medium Payload | MessagePack | Marshal | 113048 | 65832 | 22 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1620658 | 278404 | 418 |
| Large Payload | MessagePack | Unmarshal | 1632283 | 360320 | 6583 |
| Large Payload | Sonic | Unmarshal | 1666194 | 533998 | 574 |
| Large Payload | CBOR | Unmarshal | 2173027 | 342540 | 6969 |
| Large Payload | JSON | Unmarshal | 4128212 | 560511 | 7171 |
| Small Struct | BEVE ZeroCopy | Marshal | 2884 | 288 | 2 |
| Small Struct | CBOR | Marshal | 16887 | 3217 | 2 |
| Small Struct | BEVE | Marshal | 18149 | 2977 | 3 |
| Small Struct | Sonic | Marshal | 18380 | 3294 | 3 |
| Small Struct | MessagePack | Marshal | 20077 | 4224 | 8 |
| Small Struct | JSON | Marshal | 28959 | 2833 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 232725 | 149 | 2 |
| Large Payload | BEVE | Marshal | 784094 | 197316 | 3 |
| Large Payload | Sonic | Marshal | 931686 | 225859 | 4 |
| Large Payload | CBOR | Marshal | 958274 | 198025 | 2 |
| Large Payload | MessagePack | Marshal | 1326334 | 526761 | 115 |
| Large Payload | JSON | Marshal | 1472518 | 207015 | 9 |
| Medium Payload | BEVE | Unmarshal | 125251 | 24187 | 59 |
| Medium Payload | CBOR | Unmarshal | 216893 | 20264 | 417 |
| Medium Payload | MessagePack | Unmarshal | 300136 | 40492 | 763 |
| Medium Payload | Sonic | Unmarshal | 306411 | 51560 | 68 |
| Medium Payload | JSON | Unmarshal | 833847 | 57152 | 748 |
| Small Struct | Sonic | Unmarshal | 6361 | 1319 | 7 |
| Small Struct | MessagePack | Unmarshal | 8955 | 1096 | 25 |
| Small Struct | BEVE | Unmarshal | 13691 | 3000 | 4 |
| Small Struct | CBOR | Unmarshal | 38075 | 4336 | 92 |
| Small Struct | JSON | Unmarshal | 91639 | 7208 | 91 |
| Medium Payload | BEVE ZeroCopy | Marshal | 33666 | 131 | 2 |
| Medium Payload | Sonic | Marshal | 66065 | 22278 | 4 |
| Medium Payload | BEVE | Marshal | 85211 | 18572 | 3 |
| Medium Payload | CBOR | Marshal | 113624 | 21860 | 2 |
| Medium Payload | JSON | Marshal | 140141 | 20799 | 9 |
| Medium Payload | MessagePack | Marshal | 170703 | 65828 | 22 |

