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
| Large Payload | BEVE | Unmarshal | 1111419 | 152481 | 418 |
| Large Payload | Sonic | Unmarshal | 1759607 | 346360 | 211 |
| Large Payload | MessagePack | Unmarshal | 2063338 | 347429 | 6333 |
| Large Payload | CBOR | Unmarshal | 2238993 | 335930 | 6848 |
| Large Payload | JSON | Unmarshal | 3578360 | 496957 | 6545 |
| Small Struct | CBOR | Marshal | 520.40 | 432 | 2 |
| Small Struct | BEVE | Marshal | 9788 | 496 | 3 |
| Small Struct | MessagePack | Marshal | 14240 | 1152 | 6 |
| Small Struct | BEVE ZeroCopy | Marshal | 30885 | 288 | 2 |
| Small Struct | Sonic | Marshal | 37264 | 1731 | 3 |
| Small Struct | JSON | Marshal | 46846 | 1936 | 2 |
| Large Payload | BEVE | Marshal | 1000146 | 197420 | 3 |
| Large Payload | CBOR | Marshal | 1011417 | 189393 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 1163388 | 180 | 2 |
| Large Payload | MessagePack | Marshal | 1432618 | 526800 | 115 |
| Large Payload | JSON | Marshal | 1893788 | 214445 | 9 |
| Large Payload | Sonic | Marshal | 2090595 | 234651 | 4 |
| Medium Payload | BEVE | Unmarshal | 141852 | 16106 | 59 |
| Medium Payload | Sonic | Unmarshal | 224419 | 33121 | 33 |
| Medium Payload | MessagePack | Unmarshal | 361742 | 35661 | 656 |
| Medium Payload | CBOR | Unmarshal | 414926 | 33160 | 679 |
| Medium Payload | JSON | Unmarshal | 1226807 | 56056 | 750 |
| Small Struct | BEVE | Unmarshal | 4777 | 552 | 4 |
| Small Struct | MessagePack | Unmarshal | 33550 | 2080 | 45 |
| Small Struct | Sonic | Unmarshal | 33786 | 2753 | 6 |
| Small Struct | CBOR | Unmarshal | 46226 | 2472 | 54 |
| Small Struct | JSON | Unmarshal | 85233 | 1256 | 26 |
| Medium Payload | BEVE | Marshal | 171407 | 16523 | 3 |
| Medium Payload | CBOR | Marshal | 186146 | 21855 | 2 |
| Medium Payload | BEVE ZeroCopy | Marshal | 236037 | 132 | 2 |
| Medium Payload | MessagePack | Marshal | 282811 | 65831 | 22 |
| Medium Payload | JSON | Marshal | 321932 | 22081 | 9 |
| Medium Payload | Sonic | Marshal | 349749 | 19578 | 4 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1107924 | 158537 | 418 |
| Large Payload | Sonic | Unmarshal | 1568976 | 555572 | 571 |
| Large Payload | MessagePack | Unmarshal | 1764572 | 362509 | 6624 |
| Large Payload | CBOR | Unmarshal | 2078447 | 334157 | 6797 |
| Large Payload | JSON | Unmarshal | 3793162 | 573027 | 7532 |
| Small Struct | CBOR | Marshal | 2750 | 384 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 5446 | 288 | 2 |
| Small Struct | BEVE | Marshal | 7348 | 1184 | 3 |
| Small Struct | Sonic | Marshal | 11454 | 2948 | 3 |
| Small Struct | MessagePack | Marshal | 13993 | 4224 | 8 |
| Small Struct | JSON | Marshal | 40892 | 2833 | 2 |
| Large Payload | Sonic | Marshal | 728315 | 228516 | 4 |
| Large Payload | BEVE ZeroCopy | Marshal | 753666 | 171 | 2 |
| Large Payload | CBOR | Marshal | 1078700 | 197345 | 2 |
| Large Payload | BEVE | Marshal | 1088013 | 205224 | 3 |
| Large Payload | MessagePack | Marshal | 1416918 | 526784 | 115 |
| Large Payload | JSON | Marshal | 2196941 | 221940 | 9 |
| Medium Payload | BEVE | Unmarshal | 136253 | 18315 | 59 |
| Medium Payload | Sonic | Unmarshal | 265124 | 70810 | 80 |
| Medium Payload | MessagePack | Unmarshal | 270599 | 28221 | 508 |
| Medium Payload | CBOR | Unmarshal | 378793 | 28392 | 584 |
| Medium Payload | JSON | Unmarshal | 1265257 | 59448 | 771 |
| Small Struct | CBOR | Unmarshal | 5307 | 232 | 7 |
| Small Struct | BEVE | Unmarshal | 9435 | 2104 | 4 |
| Small Struct | Sonic | Unmarshal | 9450 | 2094 | 8 |
| Small Struct | MessagePack | Unmarshal | 33814 | 4352 | 92 |
| Small Struct | JSON | Unmarshal | 71476 | 3720 | 53 |
| Medium Payload | BEVE ZeroCopy | Marshal | 82565 | 132 | 2 |
| Medium Payload | Sonic | Marshal | 102105 | 25507 | 4 |
| Medium Payload | BEVE | Marshal | 102266 | 16522 | 3 |
| Medium Payload | CBOR | Marshal | 127044 | 20570 | 2 |
| Medium Payload | MessagePack | Marshal | 204490 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 288898 | 24894 | 9 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 598554 | 151608 | 418 |
| Large Payload | Sonic | Unmarshal | 656135 | 405046 | 213 |
| Large Payload | MessagePack | Unmarshal | 803322 | 364094 | 6661 |
| Large Payload | CBOR | Unmarshal | 1105018 | 323323 | 6600 |
| Large Payload | JSON | Unmarshal | 2499003 | 512891 | 6746 |
| Small Struct | MessagePack | Marshal | 5458 | 2176 | 7 |
| Small Struct | BEVE ZeroCopy | Marshal | 6288 | 289 | 2 |
| Small Struct | CBOR | Marshal | 8641 | 2833 | 2 |
| Small Struct | Sonic | Marshal | 9595 | 2278 | 3 |
| Small Struct | BEVE | Marshal | 10509 | 2336 | 3 |
| Small Struct | JSON | Marshal | 12542 | 1552 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 290295 | 171 | 2 |
| Large Payload | CBOR | Marshal | 531310 | 189620 | 2 |
| Large Payload | BEVE | Marshal | 610748 | 197573 | 3 |
| Large Payload | MessagePack | Marshal | 868233 | 526795 | 115 |
| Large Payload | Sonic | Marshal | 870737 | 219473 | 4 |
| Large Payload | JSON | Marshal | 1080446 | 230600 | 9 |
| Medium Payload | BEVE | Unmarshal | 76298 | 18539 | 59 |
| Medium Payload | Sonic | Unmarshal | 102842 | 44867 | 33 |
| Medium Payload | MessagePack | Unmarshal | 167966 | 36302 | 673 |
| Medium Payload | CBOR | Unmarshal | 198931 | 29336 | 601 |
| Medium Payload | JSON | Unmarshal | 638076 | 55208 | 711 |
| Small Struct | BEVE | Unmarshal | 3897 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 12741 | 5157 | 6 |
| Small Struct | CBOR | Unmarshal | 14847 | 1928 | 43 |
| Small Struct | MessagePack | Unmarshal | 16533 | 3928 | 83 |
| Small Struct | JSON | Unmarshal | 20786 | 1320 | 28 |
| Medium Payload | BEVE ZeroCopy | Marshal | 42861 | 131 | 2 |
| Medium Payload | CBOR | Marshal | 62502 | 18527 | 2 |
| Medium Payload | BEVE | Marshal | 66518 | 20619 | 3 |
| Medium Payload | Sonic | Marshal | 103861 | 22624 | 4 |
| Medium Payload | MessagePack | Marshal | 114724 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 132704 | 24882 | 9 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1200121 | 148688 | 419 |
| Large Payload | Sonic | Unmarshal | 1303769 | 572355 | 598 |
| Large Payload | MessagePack | Unmarshal | 1370908 | 358723 | 6549 |
| Large Payload | CBOR | Unmarshal | 1697437 | 324699 | 6607 |
| Large Payload | JSON | Unmarshal | 3902100 | 550503 | 7287 |
| Small Struct | CBOR | Marshal | 1378 | 288 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 3642 | 288 | 2 |
| Small Struct | Sonic | Marshal | 10869 | 2866 | 3 |
| Small Struct | JSON | Marshal | 12949 | 1680 | 2 |
| Small Struct | MessagePack | Marshal | 13849 | 4224 | 8 |
| Small Struct | BEVE | Marshal | 14914 | 2977 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 334424 | 163 | 2 |
| Large Payload | Sonic | Marshal | 896497 | 208615 | 4 |
| Large Payload | CBOR | Marshal | 973353 | 197807 | 2 |
| Large Payload | BEVE | Marshal | 1148393 | 197467 | 3 |
| Large Payload | MessagePack | Marshal | 1319481 | 526764 | 115 |
| Large Payload | JSON | Marshal | 1593294 | 223275 | 9 |
| Medium Payload | BEVE | Unmarshal | 81893 | 15562 | 59 |
| Medium Payload | MessagePack | Unmarshal | 296700 | 32828 | 604 |
| Medium Payload | CBOR | Unmarshal | 319882 | 30216 | 624 |
| Medium Payload | Sonic | Unmarshal | 351515 | 44983 | 67 |
| Medium Payload | JSON | Unmarshal | 955189 | 67128 | 872 |
| Small Struct | BEVE | Unmarshal | 6719 | 1016 | 4 |
| Small Struct | CBOR | Unmarshal | 9007 | 760 | 19 |
| Small Struct | Sonic | Unmarshal | 13304 | 4157 | 9 |
| Small Struct | MessagePack | Unmarshal | 19250 | 3112 | 65 |
| Small Struct | JSON | Unmarshal | 26428 | 1992 | 34 |
| Medium Payload | BEVE ZeroCopy | Marshal | 29259 | 128 | 2 |
| Medium Payload | CBOR | Marshal | 65555 | 16467 | 2 |
| Medium Payload | Sonic | Marshal | 72812 | 25129 | 4 |
| Medium Payload | MessagePack | Marshal | 94732 | 33059 | 21 |
| Medium Payload | BEVE | Marshal | 105638 | 20625 | 3 |
| Medium Payload | JSON | Marshal | 144380 | 22067 | 9 |

