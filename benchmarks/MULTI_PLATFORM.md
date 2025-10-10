# Multi-Platform Benchmark Results

| CPU | OS | Artifacts |
|-----|----|-----------|
| Apple M1 (Virtual) | Darwin | [Markdown](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.md) · [JSON](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.json) · [PNG](benchmarks/benchmark-darwin-apple-m1-virtual/benchmark.png) |
| AMD EPYC 7763 64-Core Processor | Linux | [Markdown](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.md) · [JSON](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.json) · [PNG](benchmarks/benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png) |
| Neoverse-N2 | Linux | [Markdown](benchmarks/benchmark-linux-neoverse-n2/benchmark.md) · [JSON](benchmarks/benchmark-linux-neoverse-n2/benchmark.json) · [PNG](benchmarks/benchmark-linux-neoverse-n2/benchmark.png) |
| unknown | MINGW64_NT-10.0-26100 | [Markdown](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.md) · [JSON](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.json) · [PNG](benchmarks/benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png) |

## Apple M1 (Virtual) — Darwin

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 959.60 | 145 | 1 |
| Small Struct | BEVE | Marshal | 1150 | 1938 | 2 |
| Small Struct | Sonic | Marshal | 1353 | 750 | 3 |
| Small Struct | CBOR | Marshal | 1964 | 2833 | 2 |
| Small Struct | MessagePack | Marshal | 1975 | 4224 | 8 |
| Small Struct | JSON | Marshal | 2346 | 1552 | 2 |
| Small Struct | BEVE | Unmarshal | 980.90 | 952 | 4 |
| Small Struct | Sonic | Unmarshal | 3010 | 4513 | 6 |
| Small Struct | CBOR | Unmarshal | 3012 | 1480 | 34 |
| Small Struct | MessagePack | Unmarshal | 3746 | 3681 | 79 |
| Small Struct | JSON | Unmarshal | 16303 | 4488 | 77 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8051 | 64 | 1 |
| Medium Payload | BEVE | Marshal | 15370 | 22037 | 2 |
| Medium Payload | CBOR | Marshal | 17309 | 20566 | 2 |
| Medium Payload | MessagePack | Marshal | 22609 | 65834 | 22 |
| Medium Payload | JSON | Marshal | 33148 | 20783 | 9 |
| Medium Payload | Sonic | Marshal | 40110 | 20777 | 4 |
| Medium Payload | BEVE | Unmarshal | 13665 | 15883 | 59 |
| Medium Payload | MessagePack | Unmarshal | 30333 | 25669 | 459 |
| Medium Payload | Sonic | Unmarshal | 31495 | 36924 | 33 |
| Medium Payload | CBOR | Unmarshal | 54437 | 31657 | 655 |
| Medium Payload | JSON | Unmarshal | 195909 | 53801 | 678 |
| Large Payload | BEVE ZeroCopy | Marshal | 74613 | 239 | 1 |
| Large Payload | BEVE | Marshal | 103353 | 205762 | 2 |
| Large Payload | CBOR | Marshal | 138972 | 180925 | 2 |
| Large Payload | MessagePack | Marshal | 192296 | 526830 | 115 |
| Large Payload | JSON | Marshal | 301634 | 205302 | 9 |
| Large Payload | Sonic | Marshal | 443589 | 222454 | 4 |
| Large Payload | BEVE | Unmarshal | 138660 | 155268 | 418 |
| Large Payload | Sonic | Unmarshal | 267707 | 343230 | 213 |
| Large Payload | MessagePack | Unmarshal | 369024 | 372469 | 6831 |
| Large Payload | CBOR | Unmarshal | 461905 | 311113 | 6328 |
| Large Payload | JSON | Unmarshal | 1745345 | 568917 | 7495 |

## AMD EPYC 7763 64-Core Processor — Linux

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE | Marshal | 978.50 | 723 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 1110 | 144 | 1 |
| Small Struct | Sonic | Marshal | 1311 | 1781 | 3 |
| Small Struct | JSON | Marshal | 1351 | 624 | 2 |
| Small Struct | MessagePack | Marshal | 1783 | 2176 | 7 |
| Small Struct | CBOR | Marshal | 3037 | 2834 | 2 |
| Small Struct | BEVE | Unmarshal | 1035 | 824 | 4 |
| Small Struct | CBOR | Unmarshal | 1637 | 424 | 12 |
| Small Struct | Sonic | Unmarshal | 2697 | 4155 | 9 |
| Small Struct | MessagePack | Unmarshal | 4792 | 3232 | 69 |
| Small Struct | JSON | Unmarshal | 25707 | 7520 | 101 |
| Medium Payload | BEVE ZeroCopy | Marshal | 10054 | 64 | 1 |
| Medium Payload | Sonic | Marshal | 19102 | 25552 | 4 |
| Medium Payload | BEVE | Marshal | 19159 | 20833 | 2 |
| Medium Payload | CBOR | Marshal | 24223 | 20715 | 2 |
| Medium Payload | MessagePack | Marshal | 36655 | 65838 | 22 |
| Medium Payload | JSON | Marshal | 43644 | 18895 | 9 |
| Medium Payload | BEVE | Unmarshal | 22152 | 18331 | 59 |
| Medium Payload | Sonic | Unmarshal | 45679 | 72722 | 82 |
| Medium Payload | MessagePack | Unmarshal | 62242 | 42336 | 793 |
| Medium Payload | CBOR | Unmarshal | 77440 | 33960 | 702 |
| Medium Payload | JSON | Unmarshal | 164862 | 36200 | 467 |
| Large Payload | BEVE ZeroCopy | Marshal | 105672 | 239 | 1 |
| Large Payload | BEVE | Marshal | 158649 | 205344 | 2 |
| Large Payload | Sonic | Marshal | 175274 | 233686 | 4 |
| Large Payload | CBOR | Marshal | 189982 | 173624 | 2 |
| Large Payload | MessagePack | Marshal | 319108 | 526858 | 115 |
| Large Payload | JSON | Marshal | 450542 | 214726 | 9 |
| Large Payload | BEVE | Unmarshal | 190944 | 155504 | 418 |
| Large Payload | Sonic | Unmarshal | 386098 | 573169 | 596 |
| Large Payload | MessagePack | Unmarshal | 569663 | 373926 | 6851 |
| Large Payload | CBOR | Unmarshal | 772319 | 334891 | 6818 |
| Large Payload | JSON | Unmarshal | 2222517 | 514081 | 6670 |

## Neoverse-N2 — Linux

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 269.30 | 144 | 1 |
| Small Struct | CBOR | Marshal | 849.50 | 720 | 2 |
| Small Struct | BEVE | Marshal | 1348 | 1938 | 2 |
| Small Struct | MessagePack | Marshal | 3585 | 8321 | 9 |
| Small Struct | Sonic | Marshal | 3744 | 2933 | 3 |
| Small Struct | JSON | Marshal | 4607 | 2835 | 2 |
| Small Struct | BEVE | Unmarshal | 1357 | 1720 | 4 |
| Small Struct | MessagePack | Unmarshal | 2051 | 1224 | 28 |
| Small Struct | Sonic | Unmarshal | 3077 | 4723 | 6 |
| Small Struct | CBOR | Unmarshal | 3203 | 1640 | 37 |
| Small Struct | JSON | Unmarshal | 9043 | 2280 | 43 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8778 | 77 | 1 |
| Medium Payload | BEVE | Marshal | 11141 | 16655 | 2 |
| Medium Payload | CBOR | Marshal | 17343 | 18533 | 2 |
| Medium Payload | Sonic | Marshal | 31983 | 24896 | 4 |
| Medium Payload | MessagePack | Marshal | 33278 | 65837 | 22 |
| Medium Payload | JSON | Marshal | 44960 | 24907 | 9 |
| Medium Payload | BEVE | Unmarshal | 18805 | 15516 | 59 |
| Medium Payload | Sonic | Unmarshal | 29037 | 38138 | 33 |
| Medium Payload | CBOR | Unmarshal | 46079 | 19417 | 405 |
| Medium Payload | MessagePack | Unmarshal | 56412 | 40144 | 751 |
| Medium Payload | JSON | Unmarshal | 230184 | 67993 | 863 |
| Large Payload | BEVE ZeroCopy | Marshal | 88017 | 502 | 1 |
| Large Payload | BEVE | Marshal | 145356 | 215178 | 2 |
| Large Payload | CBOR | Marshal | 195551 | 207116 | 2 |
| Large Payload | MessagePack | Marshal | 264721 | 526869 | 115 |
| Large Payload | Sonic | Marshal | 303024 | 216930 | 4 |
| Large Payload | JSON | Marshal | 413455 | 222923 | 9 |
| Large Payload | BEVE | Unmarshal | 176591 | 148795 | 418 |
| Large Payload | Sonic | Unmarshal | 285423 | 384871 | 211 |
| Large Payload | MessagePack | Unmarshal | 489644 | 344438 | 6263 |
| Large Payload | CBOR | Unmarshal | 613241 | 309627 | 6312 |
| Large Payload | JSON | Unmarshal | 1906618 | 509315 | 6772 |

## unknown — MINGW64_NT-10.0-26100

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 426.60 | 144 | 1 |
| Small Struct | MessagePack | Marshal | 1337 | 1152 | 6 |
| Small Struct | Sonic | Marshal | 1778 | 2009 | 3 |
| Small Struct | CBOR | Marshal | 2079 | 1680 | 2 |
| Small Struct | BEVE | Marshal | 3014 | 2839 | 2 |
| Small Struct | JSON | Marshal | 3525 | 1680 | 2 |
| Small Struct | BEVE | Unmarshal | 1846 | 1208 | 4 |
| Small Struct | CBOR | Unmarshal | 3112 | 1096 | 26 |
| Small Struct | Sonic | Unmarshal | 4025 | 4432 | 9 |
| Small Struct | JSON | Unmarshal | 6760 | 1320 | 28 |
| Small Struct | MessagePack | Unmarshal | 7969 | 4736 | 100 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9786 | 84 | 1 |
| Medium Payload | Sonic | Marshal | 19612 | 22656 | 4 |
| Medium Payload | BEVE | Marshal | 20904 | 18775 | 2 |
| Medium Payload | CBOR | Marshal | 26655 | 21872 | 2 |
| Medium Payload | MessagePack | Marshal | 41159 | 65829 | 22 |
| Medium Payload | JSON | Marshal | 59995 | 24932 | 9 |
| Medium Payload | BEVE | Unmarshal | 28955 | 16602 | 59 |
| Medium Payload | Sonic | Unmarshal | 52229 | 61489 | 76 |
| Medium Payload | MessagePack | Unmarshal | 72514 | 38492 | 720 |
| Medium Payload | CBOR | Unmarshal | 89190 | 32057 | 659 |
| Medium Payload | JSON | Unmarshal | 280703 | 58489 | 779 |
| Large Payload | BEVE ZeroCopy | Marshal | 114938 | 415 | 1 |
| Large Payload | Sonic | Marshal | 178765 | 211501 | 4 |
| Large Payload | BEVE | Marshal | 191904 | 213308 | 2 |
| Large Payload | CBOR | Marshal | 216042 | 198118 | 2 |
| Large Payload | MessagePack | Marshal | 297419 | 526772 | 115 |
| Large Payload | JSON | Marshal | 469720 | 206871 | 9 |
| Large Payload | BEVE | Unmarshal | 240761 | 163814 | 419 |
| Large Payload | Sonic | Unmarshal | 426381 | 544938 | 589 |
| Large Payload | MessagePack | Unmarshal | 682259 | 354464 | 6468 |
| Large Payload | CBOR | Unmarshal | 799036 | 292780 | 5969 |
| Large Payload | JSON | Unmarshal | 2533670 | 529686 | 6874 |

