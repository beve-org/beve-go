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
| Large Payload | BEVE | Unmarshal | 1061839 | 151600 | 419 |
| Large Payload | Sonic | Unmarshal | 1707611 | 366316 | 213 |
| Large Payload | MessagePack | Unmarshal | 2029736 | 345716 | 6291 |
| Large Payload | CBOR | Unmarshal | 2203998 | 319515 | 6510 |
| Large Payload | JSON | Unmarshal | 3535936 | 539092 | 6972 |
| Small Struct | CBOR | Marshal | 365.50 | 288 | 2 |
| Small Struct | BEVE | Marshal | 814.40 | 1696 | 3 |
| Small Struct | Sonic | Marshal | 846.20 | 382 | 3 |
| Small Struct | MessagePack | Marshal | 19027 | 8321 | 9 |
| Small Struct | BEVE ZeroCopy | Marshal | 21717 | 288 | 2 |
| Small Struct | JSON | Marshal | 41366 | 1296 | 2 |
| Large Payload | BEVE | Marshal | 858904 | 189042 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 889719 | 233 | 2 |
| Large Payload | CBOR | Marshal | 1277962 | 189823 | 2 |
| Large Payload | MessagePack | Marshal | 1408925 | 526802 | 115 |
| Large Payload | JSON | Marshal | 1845362 | 213728 | 9 |
| Large Payload | Sonic | Marshal | 2089784 | 216949 | 4 |
| Medium Payload | BEVE | Unmarshal | 140941 | 15418 | 59 |
| Medium Payload | Sonic | Unmarshal | 263512 | 35374 | 33 |
| Medium Payload | MessagePack | Unmarshal | 508849 | 37982 | 710 |
| Medium Payload | CBOR | Unmarshal | 555993 | 30552 | 625 |
| Medium Payload | JSON | Unmarshal | 1075877 | 54616 | 691 |
| Small Struct | BEVE | Unmarshal | 1614 | 1592 | 4 |
| Small Struct | Sonic | Unmarshal | 22522 | 2481 | 6 |
| Small Struct | MessagePack | Unmarshal | 44084 | 784 | 19 |
| Small Struct | CBOR | Unmarshal | 78265 | 2792 | 60 |
| Small Struct | JSON | Unmarshal | 371677 | 7816 | 110 |
| Medium Payload | BEVE | Marshal | 122710 | 20619 | 3 |
| Medium Payload | BEVE ZeroCopy | Marshal | 126351 | 134 | 2 |
| Medium Payload | CBOR | Marshal | 165435 | 24690 | 2 |
| Medium Payload | MessagePack | Marshal | 273059 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 297798 | 18749 | 9 |
| Medium Payload | Sonic | Marshal | 490775 | 25262 | 4 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1045426 | 152922 | 418 |
| Large Payload | Sonic | Unmarshal | 1676737 | 495222 | 548 |
| Large Payload | MessagePack | Unmarshal | 1677074 | 350141 | 6380 |
| Large Payload | CBOR | Unmarshal | 1964356 | 324491 | 6599 |
| Large Payload | JSON | Unmarshal | 3944807 | 562586 | 7348 |
| Small Struct | BEVE ZeroCopy | Marshal | 7820 | 289 | 2 |
| Small Struct | BEVE | Marshal | 9188 | 2977 | 3 |
| Small Struct | CBOR | Marshal | 10075 | 1936 | 2 |
| Small Struct | Sonic | Marshal | 11076 | 2956 | 3 |
| Small Struct | MessagePack | Marshal | 26600 | 8321 | 9 |
| Small Struct | JSON | Marshal | 33421 | 2833 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 536215 | 286 | 2 |
| Large Payload | BEVE | Marshal | 665717 | 188772 | 3 |
| Large Payload | Sonic | Marshal | 769631 | 221328 | 4 |
| Large Payload | CBOR | Marshal | 1013375 | 198100 | 2 |
| Large Payload | MessagePack | Marshal | 1443824 | 526782 | 115 |
| Large Payload | JSON | Marshal | 2039922 | 205904 | 9 |
| Medium Payload | BEVE | Unmarshal | 113443 | 15723 | 59 |
| Medium Payload | Sonic | Unmarshal | 216689 | 55827 | 68 |
| Medium Payload | MessagePack | Unmarshal | 324616 | 36783 | 684 |
| Medium Payload | CBOR | Unmarshal | 398394 | 31544 | 647 |
| Medium Payload | JSON | Unmarshal | 1096385 | 50776 | 644 |
| Small Struct | BEVE | Unmarshal | 9336 | 1848 | 4 |
| Small Struct | MessagePack | Unmarshal | 15246 | 2056 | 44 |
| Small Struct | Sonic | Unmarshal | 22308 | 7025 | 10 |
| Small Struct | CBOR | Unmarshal | 46973 | 4640 | 98 |
| Small Struct | JSON | Unmarshal | 151803 | 7792 | 109 |
| Medium Payload | BEVE ZeroCopy | Marshal | 58998 | 136 | 2 |
| Medium Payload | BEVE | Marshal | 85008 | 20622 | 3 |
| Medium Payload | CBOR | Marshal | 95062 | 14412 | 2 |
| Medium Payload | Sonic | Marshal | 106629 | 25475 | 4 |
| Medium Payload | MessagePack | Marshal | 205933 | 65833 | 22 |
| Medium Payload | JSON | Marshal | 278643 | 24888 | 9 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 542801 | 147401 | 419 |
| Large Payload | Sonic | Unmarshal | 693451 | 375292 | 211 |
| Large Payload | MessagePack | Unmarshal | 717536 | 372065 | 6820 |
| Large Payload | CBOR | Unmarshal | 1051033 | 318909 | 6495 |
| Large Payload | JSON | Unmarshal | 2452704 | 527085 | 6880 |
| Small Struct | BEVE ZeroCopy | Marshal | 1389 | 288 | 2 |
| Small Struct | MessagePack | Marshal | 2968 | 1152 | 6 |
| Small Struct | BEVE | Marshal | 4820 | 2977 | 3 |
| Small Struct | JSON | Marshal | 7358 | 1936 | 2 |
| Small Struct | CBOR | Marshal | 8574 | 2833 | 2 |
| Small Struct | Sonic | Marshal | 10722 | 2970 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 121527 | 286 | 2 |
| Large Payload | BEVE | Marshal | 423928 | 205823 | 3 |
| Large Payload | CBOR | Marshal | 555848 | 198369 | 2 |
| Large Payload | Sonic | Marshal | 854680 | 220380 | 4 |
| Large Payload | MessagePack | Marshal | 886904 | 526796 | 115 |
| Large Payload | JSON | Marshal | 983670 | 189621 | 9 |
| Medium Payload | BEVE | Unmarshal | 55899 | 17274 | 59 |
| Medium Payload | Sonic | Unmarshal | 81973 | 31833 | 33 |
| Medium Payload | MessagePack | Unmarshal | 160000 | 33726 | 621 |
| Medium Payload | CBOR | Unmarshal | 220440 | 31784 | 657 |
| Medium Payload | JSON | Unmarshal | 652995 | 66872 | 884 |
| Small Struct | BEVE | Unmarshal | 1796 | 888 | 4 |
| Small Struct | Sonic | Unmarshal | 7614 | 2744 | 6 |
| Small Struct | CBOR | Unmarshal | 9795 | 1288 | 30 |
| Small Struct | MessagePack | Unmarshal | 14596 | 4000 | 85 |
| Small Struct | JSON | Unmarshal | 83706 | 7880 | 112 |
| Medium Payload | BEVE ZeroCopy | Marshal | 23581 | 133 | 2 |
| Medium Payload | BEVE | Marshal | 38880 | 21903 | 3 |
| Medium Payload | CBOR | Marshal | 66850 | 21855 | 2 |
| Medium Payload | Sonic | Marshal | 78764 | 19168 | 4 |
| Medium Payload | MessagePack | Marshal | 83895 | 33061 | 21 |
| Medium Payload | JSON | Marshal | 120075 | 24874 | 9 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1015118 | 155151 | 416 |
| Large Payload | MessagePack | Unmarshal | 1194394 | 339206 | 6154 |
| Large Payload | Sonic | Unmarshal | 1576098 | 538430 | 574 |
| Large Payload | CBOR | Unmarshal | 1681541 | 327611 | 6695 |
| Large Payload | JSON | Unmarshal | 3708514 | 528093 | 6854 |
| Small Struct | BEVE ZeroCopy | Marshal | 2848 | 288 | 2 |
| Small Struct | Sonic | Marshal | 4694 | 521 | 3 |
| Small Struct | BEVE | Marshal | 10933 | 2336 | 3 |
| Small Struct | CBOR | Marshal | 13165 | 2833 | 2 |
| Small Struct | MessagePack | Marshal | 14331 | 4224 | 8 |
| Small Struct | JSON | Marshal | 17944 | 2834 | 2 |
| Large Payload | BEVE ZeroCopy | Marshal | 276709 | 207 | 2 |
| Large Payload | Sonic | Marshal | 707293 | 209777 | 4 |
| Large Payload | CBOR | Marshal | 807108 | 198084 | 2 |
| Large Payload | BEVE | Marshal | 842610 | 188978 | 3 |
| Large Payload | MessagePack | Marshal | 1073250 | 526759 | 115 |
| Large Payload | JSON | Marshal | 1498046 | 215611 | 9 |
| Medium Payload | BEVE | Unmarshal | 200945 | 17466 | 59 |
| Medium Payload | MessagePack | Unmarshal | 344865 | 38635 | 719 |
| Medium Payload | CBOR | Unmarshal | 357647 | 27672 | 566 |
| Medium Payload | Sonic | Unmarshal | 391507 | 58509 | 76 |
| Medium Payload | JSON | Unmarshal | 727769 | 47992 | 641 |
| Small Struct | BEVE | Unmarshal | 3843 | 424 | 4 |
| Small Struct | MessagePack | Unmarshal | 16371 | 3904 | 82 |
| Small Struct | Sonic | Unmarshal | 17067 | 7771 | 10 |
| Small Struct | JSON | Unmarshal | 19822 | 1992 | 34 |
| Small Struct | CBOR | Unmarshal | 20423 | 3104 | 66 |
| Medium Payload | BEVE ZeroCopy | Marshal | 27809 | 133 | 2 |
| Medium Payload | Sonic | Marshal | 69580 | 27789 | 4 |
| Medium Payload | CBOR | Marshal | 76650 | 19160 | 2 |
| Medium Payload | BEVE | Marshal | 116709 | 20624 | 3 |
| Medium Payload | JSON | Marshal | 147694 | 19375 | 9 |
| Medium Payload | MessagePack | Marshal | 255013 | 65831 | 22 |

