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
| Small Struct | BEVE ZeroCopy | Marshal | 667.30 | 144 | 1 |
| Small Struct | CBOR | Marshal | 844.40 | 784 | 2 |
| Small Struct | BEVE | Marshal | 1201 | 1939 | 2 |
| Small Struct | Sonic | Marshal | 1494 | 778 | 3 |
| Small Struct | JSON | Marshal | 3344 | 1936 | 2 |
| Small Struct | MessagePack | Marshal | 4248 | 8321 | 9 |
| Small Struct | BEVE | Unmarshal | 668.80 | 488 | 4 |
| Small Struct | CBOR | Unmarshal | 3186 | 1576 | 36 |
| Small Struct | Sonic | Unmarshal | 4377 | 5921 | 6 |
| Small Struct | MessagePack | Unmarshal | 5191 | 4824 | 103 |
| Small Struct | JSON | Unmarshal | 6744 | 1992 | 34 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8479 | 64 | 1 |
| Medium Payload | BEVE | Marshal | 14999 | 18692 | 2 |
| Medium Payload | CBOR | Marshal | 28570 | 20581 | 2 |
| Medium Payload | MessagePack | Marshal | 40686 | 65835 | 22 |
| Medium Payload | JSON | Marshal | 42365 | 24893 | 9 |
| Medium Payload | Sonic | Marshal | 62377 | 25059 | 4 |
| Medium Payload | BEVE | Unmarshal | 21358 | 18251 | 59 |
| Medium Payload | Sonic | Unmarshal | 45396 | 46214 | 33 |
| Medium Payload | MessagePack | Unmarshal | 52371 | 35070 | 650 |
| Medium Payload | CBOR | Unmarshal | 71720 | 27401 | 568 |
| Medium Payload | JSON | Unmarshal | 278359 | 62953 | 813 |
| Large Payload | BEVE ZeroCopy | Marshal | 85451 | 327 | 1 |
| Large Payload | BEVE | Marshal | 161864 | 210668 | 2 |
| Large Payload | CBOR | Marshal | 192946 | 197666 | 2 |
| Large Payload | MessagePack | Marshal | 270390 | 526829 | 115 |
| Large Payload | JSON | Marshal | 417002 | 214541 | 9 |
| Large Payload | Sonic | Marshal | 520393 | 223254 | 4 |
| Large Payload | BEVE | Unmarshal | 179574 | 155219 | 418 |
| Large Payload | Sonic | Unmarshal | 310664 | 375717 | 211 |
| Large Payload | MessagePack | Unmarshal | 452869 | 348262 | 6335 |
| Large Payload | CBOR | Unmarshal | 578937 | 316409 | 6449 |
| Large Payload | JSON | Unmarshal | 1974525 | 502285 | 6693 |

## AMD EPYC 7763 64-Core Processor — Linux

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | Sonic | Marshal | 408.70 | 476 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 864.60 | 145 | 1 |
| Small Struct | MessagePack | Marshal | 994.60 | 640 | 5 |
| Small Struct | CBOR | Marshal | 1741 | 1681 | 2 |
| Small Struct | BEVE | Marshal | 1756 | 2193 | 2 |
| Small Struct | JSON | Marshal | 4629 | 2450 | 2 |
| Small Struct | BEVE | Unmarshal | 903.80 | 696 | 4 |
| Small Struct | MessagePack | Unmarshal | 3134 | 2104 | 46 |
| Small Struct | Sonic | Unmarshal | 4301 | 7772 | 10 |
| Small Struct | CBOR | Unmarshal | 6194 | 3184 | 68 |
| Small Struct | JSON | Unmarshal | 18159 | 4392 | 74 |
| Medium Payload | BEVE ZeroCopy | Marshal | 10875 | 77 | 1 |
| Medium Payload | BEVE | Marshal | 14030 | 16756 | 2 |
| Medium Payload | Sonic | Marshal | 14337 | 19431 | 4 |
| Medium Payload | CBOR | Marshal | 23498 | 22010 | 2 |
| Medium Payload | MessagePack | Marshal | 27395 | 33063 | 21 |
| Medium Payload | JSON | Marshal | 38800 | 18829 | 9 |
| Medium Payload | BEVE | Unmarshal | 19877 | 15899 | 59 |
| Medium Payload | Sonic | Unmarshal | 43298 | 68048 | 76 |
| Medium Payload | MessagePack | Unmarshal | 64283 | 44345 | 839 |
| Medium Payload | CBOR | Unmarshal | 67749 | 28712 | 592 |
| Medium Payload | JSON | Unmarshal | 242836 | 59225 | 795 |
| Large Payload | BEVE ZeroCopy | Marshal | 106700 | 327 | 1 |
| Large Payload | BEVE | Marshal | 150781 | 190006 | 2 |
| Large Payload | Sonic | Marshal | 172221 | 227527 | 4 |
| Large Payload | CBOR | Marshal | 217791 | 205711 | 2 |
| Large Payload | MessagePack | Marshal | 317683 | 526856 | 115 |
| Large Payload | JSON | Marshal | 459430 | 215426 | 9 |
| Large Payload | BEVE | Unmarshal | 186289 | 152735 | 419 |
| Large Payload | Sonic | Unmarshal | 362018 | 533813 | 572 |
| Large Payload | MessagePack | Unmarshal | 553750 | 357106 | 6519 |
| Large Payload | CBOR | Unmarshal | 703785 | 296393 | 6040 |
| Large Payload | JSON | Unmarshal | 2255799 | 517378 | 6783 |

## Neoverse-N2 — Linux

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 501.30 | 144 | 1 |
| Small Struct | CBOR | Marshal | 851.30 | 720 | 2 |
| Small Struct | BEVE | Marshal | 1461 | 2196 | 2 |
| Small Struct | MessagePack | Marshal | 1573 | 2176 | 7 |
| Small Struct | JSON | Marshal | 2466 | 1424 | 2 |
| Small Struct | Sonic | Marshal | 3899 | 2919 | 3 |
| Small Struct | BEVE | Unmarshal | 1178 | 1208 | 4 |
| Small Struct | MessagePack | Unmarshal | 1343 | 536 | 14 |
| Small Struct | Sonic | Unmarshal | 2512 | 4001 | 6 |
| Small Struct | JSON | Unmarshal | 4141 | 872 | 21 |
| Small Struct | CBOR | Unmarshal | 5175 | 3080 | 65 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9702 | 84 | 1 |
| Medium Payload | BEVE | Marshal | 14732 | 24971 | 2 |
| Medium Payload | CBOR | Marshal | 18410 | 20556 | 2 |
| Medium Payload | MessagePack | Marshal | 22291 | 33063 | 21 |
| Medium Payload | Sonic | Marshal | 29575 | 22357 | 4 |
| Medium Payload | JSON | Marshal | 42782 | 24948 | 9 |
| Medium Payload | BEVE | Unmarshal | 18835 | 16283 | 59 |
| Medium Payload | Sonic | Unmarshal | 31227 | 43751 | 33 |
| Medium Payload | MessagePack | Unmarshal | 50793 | 35888 | 669 |
| Medium Payload | CBOR | Unmarshal | 63905 | 32873 | 672 |
| Medium Payload | JSON | Unmarshal | 209783 | 60504 | 779 |
| Large Payload | BEVE ZeroCopy | Marshal | 88334 | 502 | 1 |
| Large Payload | BEVE | Marshal | 138543 | 214042 | 2 |
| Large Payload | CBOR | Marshal | 187118 | 206248 | 2 |
| Large Payload | MessagePack | Marshal | 257169 | 526869 | 115 |
| Large Payload | Sonic | Marshal | 316575 | 231592 | 4 |
| Large Payload | JSON | Marshal | 368129 | 197637 | 9 |
| Large Payload | BEVE | Unmarshal | 170903 | 146474 | 417 |
| Large Payload | Sonic | Unmarshal | 278261 | 370544 | 213 |
| Large Payload | MessagePack | Unmarshal | 459965 | 322301 | 5821 |
| Large Payload | CBOR | Unmarshal | 629189 | 322766 | 6579 |
| Large Payload | JSON | Unmarshal | 1925657 | 522101 | 6870 |

## unknown — MINGW64_NT-10.0-26100

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 675.10 | 144 | 1 |
| Small Struct | Sonic | Marshal | 1086 | 956 | 3 |
| Small Struct | MessagePack | Marshal | 1287 | 640 | 5 |
| Small Struct | JSON | Marshal | 2207 | 1040 | 2 |
| Small Struct | BEVE | Marshal | 3507 | 2836 | 2 |
| Small Struct | CBOR | Marshal | 3805 | 3216 | 2 |
| Small Struct | BEVE | Unmarshal | 1165 | 888 | 4 |
| Small Struct | Sonic | Unmarshal | 2194 | 2256 | 8 |
| Small Struct | MessagePack | Unmarshal | 3277 | 1792 | 40 |
| Small Struct | CBOR | Unmarshal | 8089 | 3464 | 73 |
| Small Struct | JSON | Unmarshal | 15296 | 3848 | 57 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9011 | 70 | 1 |
| Medium Payload | Sonic | Marshal | 19390 | 20890 | 4 |
| Medium Payload | BEVE | Marshal | 19599 | 18689 | 2 |
| Medium Payload | CBOR | Marshal | 27369 | 24697 | 2 |
| Medium Payload | MessagePack | Marshal | 41227 | 65828 | 22 |
| Medium Payload | JSON | Marshal | 47905 | 20783 | 9 |
| Medium Payload | BEVE | Unmarshal | 24969 | 15402 | 59 |
| Medium Payload | Sonic | Unmarshal | 46292 | 56582 | 74 |
| Medium Payload | MessagePack | Unmarshal | 56365 | 27819 | 503 |
| Medium Payload | CBOR | Unmarshal | 60904 | 19256 | 398 |
| Medium Payload | JSON | Unmarshal | 255517 | 54209 | 705 |
| Large Payload | BEVE ZeroCopy | Marshal | 122724 | 327 | 1 |
| Large Payload | Sonic | Marshal | 158788 | 226885 | 4 |
| Large Payload | BEVE | Marshal | 185811 | 194903 | 2 |
| Large Payload | CBOR | Marshal | 214255 | 191155 | 2 |
| Large Payload | MessagePack | Marshal | 295473 | 526772 | 115 |
| Large Payload | JSON | Marshal | 460655 | 215405 | 9 |
| Large Payload | BEVE | Unmarshal | 220426 | 151317 | 418 |
| Large Payload | Sonic | Unmarshal | 443693 | 579965 | 605 |
| Large Payload | MessagePack | Unmarshal | 671259 | 356016 | 6500 |
| Large Payload | CBOR | Unmarshal | 872683 | 316572 | 6457 |
| Large Payload | JSON | Unmarshal | 2546476 | 535160 | 6903 |

