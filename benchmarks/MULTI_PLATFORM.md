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
| Small Struct | BEVE ZeroCopy | Marshal | 317.30 | 144 | 1 |
| Small Struct | BEVE | Marshal | 789.60 | 1683 | 2 |
| Small Struct | CBOR | Marshal | 929.30 | 1424 | 2 |
| Small Struct | Sonic | Marshal | 2492 | 1568 | 3 |
| Small Struct | MessagePack | Marshal | 2496 | 8321 | 9 |
| Small Struct | JSON | Marshal | 3680 | 2834 | 2 |
| Small Struct | BEVE | Unmarshal | 923.90 | 1464 | 4 |
| Small Struct | MessagePack | Unmarshal | 1726 | 1704 | 38 |
| Small Struct | Sonic | Unmarshal | 1788 | 2394 | 6 |
| Small Struct | CBOR | Unmarshal | 3189 | 2736 | 58 |
| Small Struct | JSON | Unmarshal | 9980 | 3912 | 59 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6664 | 77 | 1 |
| Medium Payload | BEVE | Marshal | 10072 | 22017 | 2 |
| Medium Payload | CBOR | Marshal | 14220 | 20566 | 2 |
| Medium Payload | MessagePack | Marshal | 21120 | 65834 | 22 |
| Medium Payload | JSON | Marshal | 33321 | 24868 | 9 |
| Medium Payload | Sonic | Marshal | 46737 | 27638 | 4 |
| Medium Payload | BEVE | Unmarshal | 12885 | 16362 | 59 |
| Medium Payload | Sonic | Unmarshal | 23736 | 31326 | 33 |
| Medium Payload | MessagePack | Unmarshal | 36134 | 38013 | 709 |
| Medium Payload | CBOR | Unmarshal | 51950 | 39817 | 814 |
| Medium Payload | JSON | Unmarshal | 172582 | 56608 | 753 |
| Large Payload | BEVE ZeroCopy | Marshal | 63827 | 239 | 1 |
| Large Payload | BEVE | Marshal | 109533 | 209527 | 2 |
| Large Payload | CBOR | Marshal | 151339 | 197837 | 2 |
| Large Payload | MessagePack | Marshal | 210318 | 526817 | 115 |
| Large Payload | JSON | Marshal | 324549 | 221864 | 9 |
| Large Payload | Sonic | Marshal | 394615 | 222363 | 4 |
| Large Payload | BEVE | Unmarshal | 125146 | 140163 | 417 |
| Large Payload | Sonic | Unmarshal | 300511 | 345192 | 213 |
| Large Payload | MessagePack | Unmarshal | 358890 | 357603 | 6536 |
| Large Payload | CBOR | Unmarshal | 499209 | 335193 | 6829 |
| Large Payload | JSON | Unmarshal | 1715485 | 543946 | 7292 |

## AMD EPYC 7763 64-Core Processor — Linux

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 1249 | 145 | 1 |
| Small Struct | Sonic | Marshal | 1746 | 2507 | 3 |
| Small Struct | BEVE | Marshal | 1869 | 2454 | 2 |
| Small Struct | CBOR | Marshal | 2098 | 2193 | 2 |
| Small Struct | MessagePack | Marshal | 3979 | 8322 | 9 |
| Small Struct | JSON | Marshal | 4646 | 2449 | 2 |
| Small Struct | BEVE | Unmarshal | 751.00 | 472 | 4 |
| Small Struct | Sonic | Unmarshal | 1853 | 2373 | 8 |
| Small Struct | CBOR | Unmarshal | 5016 | 2760 | 59 |
| Small Struct | MessagePack | Unmarshal | 6027 | 4792 | 102 |
| Small Struct | JSON | Unmarshal | 10356 | 2344 | 45 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8693 | 64 | 1 |
| Medium Payload | Sonic | Marshal | 14689 | 19586 | 4 |
| Medium Payload | BEVE | Marshal | 16555 | 20821 | 2 |
| Medium Payload | CBOR | Marshal | 21699 | 20689 | 2 |
| Medium Payload | MessagePack | Marshal | 34432 | 65838 | 22 |
| Medium Payload | JSON | Marshal | 54447 | 27702 | 9 |
| Medium Payload | BEVE | Unmarshal | 18744 | 14779 | 59 |
| Medium Payload | Sonic | Unmarshal | 33387 | 50327 | 67 |
| Medium Payload | MessagePack | Unmarshal | 58726 | 39760 | 741 |
| Medium Payload | CBOR | Unmarshal | 68504 | 31848 | 653 |
| Medium Payload | JSON | Unmarshal | 195983 | 46168 | 621 |
| Large Payload | BEVE ZeroCopy | Marshal | 113570 | 327 | 1 |
| Large Payload | BEVE | Marshal | 159345 | 198199 | 2 |
| Large Payload | Sonic | Marshal | 168522 | 218546 | 4 |
| Large Payload | CBOR | Marshal | 217543 | 206583 | 2 |
| Large Payload | MessagePack | Marshal | 338844 | 526858 | 115 |
| Large Payload | JSON | Marshal | 452169 | 215425 | 9 |
| Large Payload | BEVE | Unmarshal | 185939 | 152719 | 417 |
| Large Payload | Sonic | Unmarshal | 374939 | 557450 | 597 |
| Large Payload | MessagePack | Unmarshal | 558893 | 363536 | 6644 |
| Large Payload | CBOR | Unmarshal | 704326 | 323354 | 6582 |
| Large Payload | JSON | Unmarshal | 2209866 | 503946 | 6719 |

## Neoverse-N2 — Linux

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 308.00 | 145 | 1 |
| Small Struct | BEVE | Marshal | 1092 | 1554 | 2 |
| Small Struct | JSON | Marshal | 1612 | 912 | 2 |
| Small Struct | CBOR | Marshal | 2454 | 2836 | 2 |
| Small Struct | Sonic | Marshal | 2504 | 2008 | 3 |
| Small Struct | MessagePack | Marshal | 3536 | 8322 | 9 |
| Small Struct | BEVE | Unmarshal | 1441 | 1848 | 4 |
| Small Struct | Sonic | Unmarshal | 2163 | 3006 | 6 |
| Small Struct | MessagePack | Unmarshal | 4544 | 3873 | 81 |
| Small Struct | CBOR | Unmarshal | 5278 | 3176 | 68 |
| Small Struct | JSON | Unmarshal | 26069 | 8072 | 118 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9098 | 90 | 1 |
| Medium Payload | BEVE | Marshal | 12669 | 20786 | 2 |
| Medium Payload | CBOR | Marshal | 17051 | 18547 | 2 |
| Medium Payload | MessagePack | Marshal | 21473 | 33063 | 21 |
| Medium Payload | Sonic | Marshal | 30460 | 22423 | 4 |
| Medium Payload | JSON | Marshal | 42702 | 24907 | 9 |
| Medium Payload | BEVE | Unmarshal | 18329 | 15482 | 59 |
| Medium Payload | Sonic | Unmarshal | 30028 | 42409 | 33 |
| Medium Payload | MessagePack | Unmarshal | 48474 | 33439 | 615 |
| Medium Payload | CBOR | Unmarshal | 53348 | 25129 | 519 |
| Medium Payload | JSON | Unmarshal | 184791 | 50841 | 667 |
| Large Payload | BEVE ZeroCopy | Marshal | 87656 | 415 | 1 |
| Large Payload | BEVE | Marshal | 129619 | 195894 | 2 |
| Large Payload | CBOR | Marshal | 176851 | 189848 | 2 |
| Large Payload | MessagePack | Marshal | 259553 | 526870 | 115 |
| Large Payload | Sonic | Marshal | 304087 | 225305 | 4 |
| Large Payload | JSON | Marshal | 385227 | 214902 | 9 |
| Large Payload | BEVE | Unmarshal | 178696 | 161340 | 417 |
| Large Payload | Sonic | Unmarshal | 276357 | 375461 | 211 |
| Large Payload | MessagePack | Unmarshal | 488135 | 350662 | 6397 |
| Large Payload | CBOR | Unmarshal | 639737 | 336203 | 6858 |
| Large Payload | JSON | Unmarshal | 1941359 | 530542 | 6958 |

## unknown — MINGW64_NT-10.0-26100

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 478.10 | 144 | 1 |
| Small Struct | BEVE | Marshal | 881.00 | 721 | 2 |
| Small Struct | JSON | Marshal | 1727 | 848 | 2 |
| Small Struct | Sonic | Marshal | 2482 | 2520 | 3 |
| Small Struct | CBOR | Marshal | 2667 | 2450 | 2 |
| Small Struct | MessagePack | Marshal | 3170 | 4224 | 8 |
| Small Struct | BEVE | Unmarshal | 1088 | 696 | 4 |
| Small Struct | MessagePack | Unmarshal | 1335 | 352 | 10 |
| Small Struct | Sonic | Unmarshal | 2193 | 2245 | 8 |
| Small Struct | JSON | Unmarshal | 4764 | 840 | 20 |
| Small Struct | CBOR | Unmarshal | 7749 | 3592 | 77 |
| Medium Payload | BEVE ZeroCopy | Marshal | 14406 | 84 | 1 |
| Medium Payload | Sonic | Marshal | 21851 | 25477 | 4 |
| Medium Payload | BEVE | Marshal | 25571 | 20739 | 2 |
| Medium Payload | CBOR | Marshal | 25582 | 20590 | 2 |
| Medium Payload | MessagePack | Marshal | 42031 | 65830 | 22 |
| Medium Payload | JSON | Marshal | 60227 | 24946 | 9 |
| Medium Payload | BEVE | Unmarshal | 26455 | 17914 | 59 |
| Medium Payload | Sonic | Unmarshal | 44491 | 50652 | 72 |
| Medium Payload | MessagePack | Unmarshal | 79115 | 43390 | 815 |
| Medium Payload | CBOR | Unmarshal | 80908 | 31145 | 639 |
| Medium Payload | JSON | Unmarshal | 308603 | 70137 | 916 |
| Large Payload | BEVE ZeroCopy | Marshal | 106546 | 415 | 1 |
| Large Payload | Sonic | Marshal | 165644 | 229020 | 4 |
| Large Payload | BEVE | Marshal | 178251 | 193418 | 2 |
| Large Payload | CBOR | Marshal | 209470 | 189931 | 2 |
| Large Payload | MessagePack | Marshal | 289965 | 526773 | 115 |
| Large Payload | JSON | Marshal | 500922 | 232492 | 9 |
| Large Payload | BEVE | Unmarshal | 233447 | 156954 | 419 |
| Large Payload | Sonic | Unmarshal | 396422 | 497901 | 546 |
| Large Payload | MessagePack | Unmarshal | 639667 | 345475 | 6301 |
| Large Payload | CBOR | Unmarshal | 815452 | 325324 | 6632 |
| Large Payload | JSON | Unmarshal | 2468981 | 524534 | 6866 |

