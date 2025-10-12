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
| Large Payload | BEVE | Unmarshal | 1509326 | 145088 | 419 |
| Large Payload | Sonic | Unmarshal | 2130819 | 363675 | 213 |
| Large Payload | MessagePack | Unmarshal | 2717719 | 365143 | 6687 |
| Large Payload | CBOR | Unmarshal | 2910866 | 320539 | 6527 |
| Large Payload | JSON | Unmarshal | 4909950 | 591036 | 7787 |
| Small Struct | BEVE | Marshal | 21258 | 928 | 3 |
| Small Struct | Sonic | Marshal | 39140 | 748 | 3 |
| Small Struct | JSON | Marshal | 39502 | 784 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 46595 | 288 | 2 |
| Small Struct | MessagePack | Marshal | 47480 | 4224 | 8 |
| Small Struct | CBOR | Marshal | 51026 | 2833 | 2 |
| Large Payload | BEVE | Marshal | 1399925 | 180938 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 1552977 | 189 | 2 |
| Large Payload | CBOR | Marshal | 1593554 | 205722 | 2 |
| Large Payload | MessagePack | Marshal | 1762625 | 526798 | 115 |
| Large Payload | JSON | Marshal | 2365524 | 198039 | 9 |
| Large Payload | Sonic | Marshal | 2609464 | 226685 | 4 |
| Medium Payload | BEVE | Unmarshal | 303258 | 14922 | 59 |
| Medium Payload | Sonic | Unmarshal | 419201 | 35944 | 33 |
| Medium Payload | MessagePack | Unmarshal | 526949 | 28188 | 512 |
| Medium Payload | CBOR | Unmarshal | 713173 | 29416 | 608 |
| Medium Payload | JSON | Unmarshal | 1454475 | 56120 | 726 |
| Small Struct | BEVE | Unmarshal | 25633 | 1720 | 4 |
| Small Struct | MessagePack | Unmarshal | 50284 | 680 | 17 |
| Small Struct | Sonic | Unmarshal | 69272 | 5071 | 6 |
| Small Struct | CBOR | Unmarshal | 101417 | 4040 | 87 |
| Small Struct | JSON | Unmarshal | 189330 | 2344 | 45 |
| Medium Payload | CBOR | Marshal | 223121 | 20567 | 2 |
| Medium Payload | BEVE | Marshal | 301754 | 20628 | 3 |
| Medium Payload | MessagePack | Marshal | 361799 | 33059 | 21 |
| Medium Payload | BEVE ZeroCopy | Marshal | 362144 | 132 | 2 |
| Medium Payload | JSON | Marshal | 429586 | 16692 | 9 |
| Medium Payload | Sonic | Marshal | 726890 | 33476 | 4 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1086434 | 153240 | 418 |
| Large Payload | Sonic | Unmarshal | 1507308 | 535613 | 557 |
| Large Payload | MessagePack | Unmarshal | 1691618 | 352347 | 6445 |
| Large Payload | CBOR | Unmarshal | 2056990 | 316732 | 6454 |
| Large Payload | JSON | Unmarshal | 3483057 | 503866 | 6548 |
| Small Struct | Sonic | Marshal | 3413 | 832 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 5394 | 288 | 2 |
| Small Struct | JSON | Marshal | 10944 | 784 | 2 |
| Small Struct | BEVE | Marshal | 14703 | 2593 | 3 |
| Small Struct | MessagePack | Marshal | 18567 | 4224 | 8 |
| Small Struct | CBOR | Marshal | 21250 | 3217 | 2 |
| Large Payload | Sonic | Marshal | 646981 | 203504 | 4 |
| Large Payload | BEVE ZeroCopy | Marshal | 736803 | 172 | 2 |
| Large Payload | BEVE | Marshal | 1010983 | 205312 | 3 |
| Large Payload | CBOR | Marshal | 1032310 | 197256 | 2 |
| Large Payload | MessagePack | Marshal | 1408516 | 526784 | 115 |
| Large Payload | JSON | Marshal | 2132942 | 230028 | 9 |
| Medium Payload | BEVE | Unmarshal | 140909 | 19996 | 59 |
| Medium Payload | Sonic | Unmarshal | 240517 | 66864 | 72 |
| Medium Payload | MessagePack | Unmarshal | 280298 | 32558 | 596 |
| Medium Payload | CBOR | Unmarshal | 342747 | 25672 | 525 |
| Medium Payload | JSON | Unmarshal | 1113194 | 50784 | 681 |
| Small Struct | BEVE | Unmarshal | 8347 | 1592 | 4 |
| Small Struct | MessagePack | Unmarshal | 24265 | 2720 | 57 |
| Small Struct | Sonic | Unmarshal | 27306 | 7781 | 10 |
| Small Struct | CBOR | Unmarshal | 49190 | 3976 | 85 |
| Small Struct | JSON | Unmarshal | 83663 | 3904 | 59 |
| Medium Payload | BEVE ZeroCopy | Marshal | 86833 | 134 | 2 |
| Medium Payload | Sonic | Marshal | 106955 | 28253 | 4 |
| Medium Payload | BEVE | Marshal | 123340 | 24726 | 3 |
| Medium Payload | CBOR | Marshal | 127000 | 20575 | 2 |
| Medium Payload | MessagePack | Marshal | 208411 | 65833 | 22 |
| Medium Payload | JSON | Marshal | 278335 | 24894 | 9 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 586426 | 162075 | 418 |
| Large Payload | Sonic | Unmarshal | 675200 | 416307 | 211 |
| Large Payload | MessagePack | Unmarshal | 799690 | 342713 | 6224 |
| Large Payload | CBOR | Unmarshal | 1044009 | 322635 | 6579 |
| Large Payload | JSON | Unmarshal | 2420804 | 541517 | 7059 |
| Small Struct | BEVE ZeroCopy | Marshal | 1731 | 288 | 2 |
| Small Struct | CBOR | Marshal | 2636 | 720 | 2 |
| Small Struct | JSON | Marshal | 5318 | 1040 | 2 |
| Small Struct | MessagePack | Marshal | 11222 | 4224 | 8 |
| Small Struct | BEVE | Marshal | 12438 | 2976 | 3 |
| Small Struct | Sonic | Marshal | 14212 | 2943 | 3 |
| Large Payload | BEVE ZeroCopy | Marshal | 388927 | 198 | 2 |
| Large Payload | BEVE | Marshal | 505363 | 180984 | 3 |
| Large Payload | CBOR | Marshal | 534546 | 205969 | 2 |
| Large Payload | MessagePack | Marshal | 805918 | 526788 | 115 |
| Large Payload | Sonic | Marshal | 853594 | 222457 | 4 |
| Large Payload | JSON | Marshal | 1017349 | 213814 | 9 |
| Medium Payload | BEVE | Unmarshal | 58884 | 14026 | 59 |
| Medium Payload | Sonic | Unmarshal | 75934 | 37189 | 31 |
| Medium Payload | MessagePack | Unmarshal | 157078 | 37358 | 694 |
| Medium Payload | CBOR | Unmarshal | 195945 | 30936 | 634 |
| Medium Payload | JSON | Unmarshal | 531247 | 46968 | 636 |
| Small Struct | BEVE | Unmarshal | 2470 | 472 | 4 |
| Small Struct | MessagePack | Unmarshal | 7683 | 1120 | 26 |
| Small Struct | Sonic | Unmarshal | 14623 | 5925 | 6 |
| Small Struct | CBOR | Unmarshal | 22299 | 3120 | 66 |
| Small Struct | JSON | Unmarshal | 74488 | 7816 | 110 |
| Medium Payload | BEVE ZeroCopy | Marshal | 29603 | 130 | 2 |
| Medium Payload | CBOR | Marshal | 63865 | 19167 | 2 |
| Medium Payload | BEVE | Marshal | 67834 | 20624 | 3 |
| Medium Payload | Sonic | Marshal | 83224 | 21268 | 4 |
| Medium Payload | MessagePack | Marshal | 120440 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 123896 | 24885 | 9 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE | Unmarshal | 1109777 | 145677 | 415 |
| Large Payload | MessagePack | Unmarshal | 1296905 | 329134 | 5943 |
| Large Payload | Sonic | Unmarshal | 1369755 | 507614 | 564 |
| Large Payload | CBOR | Unmarshal | 1725718 | 300235 | 6116 |
| Large Payload | JSON | Unmarshal | 3721883 | 541646 | 7144 |
| Small Struct | BEVE ZeroCopy | Marshal | 4970 | 288 | 2 |
| Small Struct | JSON | Marshal | 8661 | 912 | 2 |
| Small Struct | Sonic | Marshal | 9785 | 1963 | 3 |
| Small Struct | CBOR | Marshal | 10372 | 2192 | 2 |
| Small Struct | BEVE | Marshal | 10857 | 1440 | 3 |
| Small Struct | MessagePack | Marshal | 11598 | 4224 | 8 |
| Large Payload | BEVE ZeroCopy | Marshal | 343874 | 171 | 2 |
| Large Payload | Sonic | Marshal | 910587 | 218184 | 4 |
| Large Payload | CBOR | Marshal | 1000819 | 206175 | 2 |
| Large Payload | BEVE | Marshal | 1026692 | 189387 | 3 |
| Large Payload | MessagePack | Marshal | 1330588 | 526763 | 115 |
| Large Payload | JSON | Marshal | 1627239 | 224186 | 9 |
| Medium Payload | BEVE | Unmarshal | 79760 | 14489 | 59 |
| Medium Payload | MessagePack | Unmarshal | 225240 | 26459 | 478 |
| Medium Payload | Sonic | Unmarshal | 296792 | 48272 | 68 |
| Medium Payload | CBOR | Unmarshal | 358888 | 33432 | 687 |
| Medium Payload | JSON | Unmarshal | 626827 | 40848 | 520 |
| Small Struct | BEVE | Unmarshal | 1405 | 456 | 4 |
| Small Struct | Sonic | Unmarshal | 5559 | 2208 | 8 |
| Small Struct | CBOR | Unmarshal | 16124 | 3528 | 75 |
| Small Struct | MessagePack | Unmarshal | 17480 | 2528 | 55 |
| Small Struct | JSON | Unmarshal | 64080 | 7368 | 96 |
| Medium Payload | BEVE ZeroCopy | Marshal | 27401 | 130 | 2 |
| Medium Payload | CBOR | Marshal | 70243 | 18513 | 2 |
| Medium Payload | Sonic | Marshal | 83120 | 22269 | 4 |
| Medium Payload | BEVE | Marshal | 98551 | 19211 | 3 |
| Medium Payload | MessagePack | Marshal | 142991 | 65828 | 22 |
| Medium Payload | JSON | Marshal | 156258 | 22063 | 9 |

