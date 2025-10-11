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
| Small Struct | CBOR | Marshal | 383.10 | 400 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 396.30 | 144 | 1 |
| Small Struct | JSON | Marshal | 644.60 | 496 | 2 |
| Small Struct | MessagePack | Marshal | 971.00 | 1152 | 6 |
| Small Struct | Sonic | Marshal | 1272 | 750 | 3 |
| Small Struct | BEVE | Marshal | 1661 | 2194 | 2 |
| Small Struct | BEVE | Unmarshal | 1218 | 1592 | 4 |
| Small Struct | Sonic | Unmarshal | 1926 | 2796 | 6 |
| Small Struct | MessagePack | Unmarshal | 4926 | 3848 | 80 |
| Small Struct | CBOR | Unmarshal | 5880 | 4296 | 91 |
| Small Struct | JSON | Unmarshal | 14171 | 4360 | 73 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6077 | 64 | 1 |
| Medium Payload | BEVE | Marshal | 10485 | 18675 | 2 |
| Medium Payload | CBOR | Marshal | 18021 | 19184 | 2 |
| Medium Payload | JSON | Marshal | 27116 | 16688 | 9 |
| Medium Payload | MessagePack | Marshal | 29459 | 65834 | 22 |
| Medium Payload | Sonic | Marshal | 42748 | 20929 | 4 |
| Medium Payload | BEVE | Unmarshal | 14207 | 15274 | 59 |
| Medium Payload | Sonic | Unmarshal | 40033 | 39847 | 33 |
| Medium Payload | MessagePack | Unmarshal | 40666 | 34109 | 631 |
| Medium Payload | CBOR | Unmarshal | 61315 | 36169 | 744 |
| Medium Payload | JSON | Unmarshal | 144854 | 34321 | 453 |
| Large Payload | BEVE ZeroCopy | Marshal | 69311 | 239 | 1 |
| Large Payload | BEVE | Marshal | 126147 | 207068 | 2 |
| Large Payload | CBOR | Marshal | 153871 | 197668 | 2 |
| Large Payload | MessagePack | Marshal | 276701 | 526828 | 115 |
| Large Payload | JSON | Marshal | 379991 | 222036 | 9 |
| Large Payload | Sonic | Marshal | 458200 | 223393 | 4 |
| Large Payload | BEVE | Unmarshal | 183431 | 152373 | 418 |
| Large Payload | Sonic | Unmarshal | 293750 | 345400 | 213 |
| Large Payload | MessagePack | Unmarshal | 375728 | 339606 | 6155 |
| Large Payload | CBOR | Unmarshal | 483499 | 313660 | 6391 |
| Large Payload | JSON | Unmarshal | 2084870 | 550659 | 7227 |

## AMD EPYC 7763 64-Core Processor — Linux

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 768.90 | 144 | 1 |
| Small Struct | MessagePack | Marshal | 846.00 | 640 | 5 |
| Small Struct | BEVE | Marshal | 1046 | 1299 | 2 |
| Small Struct | CBOR | Marshal | 1339 | 1297 | 2 |
| Small Struct | Sonic | Marshal | 2338 | 3347 | 3 |
| Small Struct | JSON | Marshal | 5798 | 3219 | 2 |
| Small Struct | BEVE | Unmarshal | 624.20 | 312 | 3 |
| Small Struct | CBOR | Unmarshal | 2605 | 1096 | 26 |
| Small Struct | Sonic | Unmarshal | 2874 | 4443 | 9 |
| Small Struct | MessagePack | Unmarshal | 3249 | 2272 | 49 |
| Small Struct | JSON | Unmarshal | 12577 | 3648 | 51 |
| Medium Payload | BEVE ZeroCopy | Marshal | 11070 | 80 | 1 |
| Medium Payload | BEVE | Marshal | 15215 | 19395 | 2 |
| Medium Payload | Sonic | Marshal | 17452 | 25315 | 4 |
| Medium Payload | CBOR | Marshal | 22915 | 21991 | 2 |
| Medium Payload | MessagePack | Marshal | 35212 | 65839 | 22 |
| Medium Payload | JSON | Marshal | 52075 | 25092 | 9 |
| Medium Payload | BEVE | Unmarshal | 17732 | 13178 | 59 |
| Medium Payload | Sonic | Unmarshal | 34612 | 51788 | 71 |
| Medium Payload | MessagePack | Unmarshal | 47686 | 29566 | 534 |
| Medium Payload | CBOR | Unmarshal | 55730 | 23816 | 490 |
| Medium Payload | JSON | Unmarshal | 201433 | 49097 | 642 |
| Large Payload | BEVE ZeroCopy | Marshal | 107121 | 415 | 1 |
| Large Payload | BEVE | Marshal | 161555 | 210379 | 2 |
| Large Payload | Sonic | Marshal | 165432 | 219335 | 4 |
| Large Payload | CBOR | Marshal | 201063 | 190015 | 2 |
| Large Payload | MessagePack | Marshal | 310737 | 526856 | 115 |
| Large Payload | JSON | Marshal | 445689 | 214199 | 9 |
| Large Payload | BEVE | Unmarshal | 183882 | 151327 | 418 |
| Large Payload | Sonic | Unmarshal | 347244 | 511633 | 552 |
| Large Payload | MessagePack | Unmarshal | 574247 | 358176 | 6556 |
| Large Payload | CBOR | Unmarshal | 682497 | 313769 | 6403 |
| Large Payload | JSON | Unmarshal | 2139315 | 504386 | 6617 |

## Neoverse-N2 — Linux

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 1120 | 144 | 1 |
| Small Struct | BEVE | Marshal | 1414 | 2195 | 2 |
| Small Struct | CBOR | Marshal | 2108 | 2450 | 2 |
| Small Struct | MessagePack | Marshal | 2280 | 4224 | 8 |
| Small Struct | Sonic | Marshal | 2946 | 2293 | 3 |
| Small Struct | JSON | Marshal | 3269 | 1936 | 2 |
| Small Struct | BEVE | Unmarshal | 759.50 | 520 | 4 |
| Small Struct | Sonic | Unmarshal | 1701 | 2124 | 6 |
| Small Struct | MessagePack | Unmarshal | 3180 | 2464 | 53 |
| Small Struct | CBOR | Unmarshal | 7095 | 4616 | 97 |
| Small Struct | JSON | Unmarshal | 7219 | 2024 | 35 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9866 | 90 | 1 |
| Medium Payload | BEVE | Marshal | 13915 | 22060 | 2 |
| Medium Payload | CBOR | Marshal | 21521 | 24666 | 2 |
| Medium Payload | Sonic | Marshal | 28265 | 20948 | 4 |
| Medium Payload | MessagePack | Marshal | 29876 | 65838 | 22 |
| Medium Payload | JSON | Marshal | 38161 | 20797 | 9 |
| Medium Payload | BEVE | Unmarshal | 19860 | 18427 | 59 |
| Medium Payload | Sonic | Unmarshal | 31596 | 46449 | 33 |
| Medium Payload | MessagePack | Unmarshal | 56357 | 42719 | 806 |
| Medium Payload | CBOR | Unmarshal | 69365 | 36793 | 753 |
| Medium Payload | JSON | Unmarshal | 193158 | 56705 | 708 |
| Large Payload | BEVE ZeroCopy | Marshal | 86211 | 415 | 1 |
| Large Payload | BEVE | Marshal | 146459 | 221448 | 2 |
| Large Payload | CBOR | Marshal | 184332 | 197523 | 2 |
| Large Payload | MessagePack | Marshal | 255065 | 526868 | 115 |
| Large Payload | Sonic | Marshal | 293127 | 218020 | 4 |
| Large Payload | JSON | Marshal | 401198 | 221871 | 9 |
| Large Payload | BEVE | Unmarshal | 172829 | 154924 | 418 |
| Large Payload | Sonic | Unmarshal | 291424 | 422907 | 211 |
| Large Payload | MessagePack | Unmarshal | 480653 | 342612 | 6214 |
| Large Payload | CBOR | Unmarshal | 592769 | 296795 | 6056 |
| Large Payload | JSON | Unmarshal | 1895059 | 519695 | 6722 |

## unknown — MINGW64_NT-10.0-26100

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE | Marshal | 603.40 | 464 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 1239 | 146 | 1 |
| Small Struct | Sonic | Marshal | 1428 | 1383 | 3 |
| Small Struct | JSON | Marshal | 1467 | 624 | 2 |
| Small Struct | CBOR | Marshal | 3047 | 2835 | 2 |
| Small Struct | MessagePack | Marshal | 3384 | 4224 | 8 |
| Small Struct | BEVE | Unmarshal | 975.70 | 552 | 4 |
| Small Struct | Sonic | Unmarshal | 3976 | 4464 | 9 |
| Small Struct | CBOR | Unmarshal | 4306 | 904 | 22 |
| Small Struct | MessagePack | Unmarshal | 6047 | 4448 | 95 |
| Small Struct | JSON | Unmarshal | 18104 | 4424 | 75 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9017 | 87 | 1 |
| Medium Payload | BEVE | Marshal | 20493 | 20732 | 2 |
| Medium Payload | Sonic | Marshal | 22152 | 25122 | 4 |
| Medium Payload | CBOR | Marshal | 25240 | 21914 | 2 |
| Medium Payload | JSON | Marshal | 45950 | 22091 | 9 |
| Medium Payload | MessagePack | Marshal | 48404 | 65830 | 22 |
| Medium Payload | BEVE | Unmarshal | 24123 | 15051 | 59 |
| Medium Payload | Sonic | Unmarshal | 54620 | 62369 | 78 |
| Medium Payload | MessagePack | Unmarshal | 77076 | 42721 | 803 |
| Medium Payload | CBOR | Unmarshal | 81262 | 36024 | 736 |
| Medium Payload | JSON | Unmarshal | 226890 | 56152 | 731 |
| Large Payload | BEVE ZeroCopy | Marshal | 92488 | 328 | 1 |
| Large Payload | BEVE | Marshal | 160135 | 197040 | 2 |
| Large Payload | Sonic | Marshal | 200631 | 221054 | 4 |
| Large Payload | CBOR | Marshal | 235864 | 199932 | 2 |
| Large Payload | MessagePack | Marshal | 333297 | 526799 | 115 |
| Large Payload | JSON | Marshal | 449797 | 214543 | 9 |
| Large Payload | BEVE | Unmarshal | 227714 | 149372 | 417 |
| Large Payload | Sonic | Unmarshal | 477255 | 570307 | 591 |
| Large Payload | MessagePack | Unmarshal | 590107 | 319825 | 5761 |
| Large Payload | CBOR | Unmarshal | 742246 | 311724 | 6350 |
| Large Payload | JSON | Unmarshal | 2134729 | 508812 | 6755 |

