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
| Small Struct | BEVE ZeroCopy | Marshal | 880.30 | 288 | 2 |
| Small Struct | BEVE | Marshal | 1384 | 2336 | 3 |
| Small Struct | MessagePack | Marshal | 1595 | 4224 | 8 |
| Small Struct | CBOR | Marshal | 1891 | 2450 | 2 |
| Small Struct | JSON | Marshal | 5201 | 2833 | 2 |
| Small Struct | Sonic | Marshal | 5520 | 2877 | 3 |
| Small Struct | BEVE | Unmarshal | 1016 | 1336 | 4 |
| Small Struct | MessagePack | Unmarshal | 2232 | 2080 | 45 |
| Small Struct | Sonic | Unmarshal | 2357 | 3336 | 6 |
| Small Struct | CBOR | Unmarshal | 5010 | 4232 | 89 |
| Small Struct | JSON | Unmarshal | 23302 | 7976 | 115 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9532 | 128 | 2 |
| Medium Payload | CBOR | Marshal | 14912 | 21860 | 2 |
| Medium Payload | BEVE | Marshal | 15294 | 24721 | 3 |
| Medium Payload | MessagePack | Marshal | 15977 | 33061 | 21 |
| Medium Payload | JSON | Marshal | 28298 | 19376 | 9 |
| Medium Payload | Sonic | Marshal | 40278 | 22169 | 4 |
| Medium Payload | BEVE | Unmarshal | 13909 | 16698 | 59 |
| Medium Payload | Sonic | Unmarshal | 32982 | 44661 | 33 |
| Medium Payload | MessagePack | Unmarshal | 38700 | 37677 | 700 |
| Medium Payload | CBOR | Unmarshal | 58705 | 30329 | 623 |
| Medium Payload | JSON | Unmarshal | 175262 | 48696 | 626 |
| Large Payload | BEVE ZeroCopy | Marshal | 108428 | 391 | 2 |
| Large Payload | BEVE | Marshal | 128672 | 180643 | 3 |
| Large Payload | CBOR | Marshal | 150791 | 198012 | 2 |
| Large Payload | MessagePack | Marshal | 198922 | 526830 | 115 |
| Large Payload | JSON | Marshal | 346325 | 222210 | 9 |
| Large Payload | Sonic | Marshal | 387054 | 205474 | 4 |
| Large Payload | BEVE | Unmarshal | 124741 | 157667 | 418 |
| Large Payload | MessagePack | Unmarshal | 353150 | 352135 | 6433 |
| Large Payload | Sonic | Unmarshal | 371282 | 370644 | 213 |
| Large Payload | CBOR | Unmarshal | 453277 | 315914 | 6436 |
| Large Payload | JSON | Unmarshal | 1533701 | 474989 | 6364 |

## AMD EPYC 7763 64-Core Processor — Linux

![Benchmark Chart](benchmark-linux-amd-epyc-7763-64-core-processor/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | CBOR | Marshal | 1318 | 1297 | 2 |
| Small Struct | Sonic | Marshal | 1763 | 2565 | 3 |
| Small Struct | BEVE | Marshal | 1807 | 2081 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 1917 | 289 | 2 |
| Small Struct | MessagePack | Marshal | 2578 | 4224 | 8 |
| Small Struct | JSON | Marshal | 3331 | 1681 | 2 |
| Small Struct | BEVE | Unmarshal | 743.90 | 392 | 4 |
| Small Struct | Sonic | Unmarshal | 4411 | 7388 | 10 |
| Small Struct | MessagePack | Unmarshal | 5410 | 4224 | 88 |
| Small Struct | CBOR | Unmarshal | 6109 | 2824 | 61 |
| Small Struct | JSON | Unmarshal | 17693 | 4392 | 74 |
| Medium Payload | BEVE ZeroCopy | Marshal | 14319 | 151 | 2 |
| Medium Payload | BEVE | Marshal | 15218 | 16542 | 3 |
| Medium Payload | Sonic | Marshal | 17431 | 23028 | 4 |
| Medium Payload | CBOR | Marshal | 21874 | 20729 | 2 |
| Medium Payload | MessagePack | Marshal | 25622 | 33064 | 21 |
| Medium Payload | JSON | Marshal | 42411 | 20838 | 9 |
| Medium Payload | BEVE | Unmarshal | 20970 | 16523 | 59 |
| Medium Payload | Sonic | Unmarshal | 40645 | 62351 | 76 |
| Medium Payload | MessagePack | Unmarshal | 49334 | 30766 | 561 |
| Medium Payload | CBOR | Unmarshal | 59883 | 24872 | 516 |
| Medium Payload | JSON | Unmarshal | 205905 | 48992 | 675 |
| Large Payload | BEVE ZeroCopy | Marshal | 125441 | 391 | 2 |
| Large Payload | Sonic | Marshal | 167323 | 225796 | 4 |
| Large Payload | BEVE | Marshal | 181638 | 188673 | 3 |
| Large Payload | CBOR | Marshal | 212422 | 189316 | 2 |
| Large Payload | MessagePack | Marshal | 309912 | 526850 | 115 |
| Large Payload | JSON | Marshal | 443306 | 205832 | 9 |
| Large Payload | BEVE | Unmarshal | 197693 | 157680 | 417 |
| Large Payload | Sonic | Unmarshal | 406126 | 594532 | 606 |
| Large Payload | MessagePack | Unmarshal | 565820 | 359763 | 6577 |
| Large Payload | CBOR | Unmarshal | 714024 | 315451 | 6427 |
| Large Payload | JSON | Unmarshal | 2295641 | 561173 | 7230 |

## Neoverse-N2 — Linux

![Benchmark Chart](benchmark-linux-neoverse-n2/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 1158 | 288 | 2 |
| Small Struct | BEVE | Marshal | 1923 | 2337 | 3 |
| Small Struct | Sonic | Marshal | 1938 | 1326 | 3 |
| Small Struct | MessagePack | Marshal | 2321 | 4225 | 8 |
| Small Struct | CBOR | Marshal | 2503 | 2836 | 2 |
| Small Struct | JSON | Marshal | 3673 | 2451 | 2 |
| Small Struct | BEVE | Unmarshal | 734.70 | 408 | 4 |
| Small Struct | Sonic | Unmarshal | 1154 | 1001 | 6 |
| Small Struct | MessagePack | Unmarshal | 1280 | 448 | 12 |
| Small Struct | CBOR | Unmarshal | 2213 | 952 | 23 |
| Small Struct | JSON | Unmarshal | 16269 | 4424 | 75 |
| Medium Payload | BEVE ZeroCopy | Marshal | 13042 | 154 | 2 |
| Medium Payload | BEVE | Marshal | 14705 | 19223 | 3 |
| Medium Payload | CBOR | Marshal | 19592 | 21877 | 2 |
| Medium Payload | Sonic | Marshal | 27333 | 19645 | 4 |
| Medium Payload | MessagePack | Marshal | 29383 | 65837 | 22 |
| Medium Payload | JSON | Marshal | 39766 | 22092 | 9 |
| Medium Payload | BEVE | Unmarshal | 18967 | 15082 | 59 |
| Medium Payload | Sonic | Unmarshal | 31756 | 42876 | 33 |
| Medium Payload | MessagePack | Unmarshal | 56946 | 41328 | 778 |
| Medium Payload | CBOR | Unmarshal | 63408 | 31864 | 656 |
| Medium Payload | JSON | Unmarshal | 177920 | 46889 | 634 |
| Large Payload | BEVE ZeroCopy | Marshal | 119767 | 479 | 2 |
| Large Payload | BEVE | Marshal | 155892 | 197393 | 3 |
| Large Payload | CBOR | Marshal | 177949 | 189845 | 2 |
| Large Payload | MessagePack | Marshal | 265994 | 526865 | 115 |
| Large Payload | Sonic | Marshal | 306557 | 225487 | 4 |
| Large Payload | JSON | Marshal | 362929 | 206008 | 9 |
| Large Payload | BEVE | Unmarshal | 176479 | 148346 | 419 |
| Large Payload | Sonic | Unmarshal | 307440 | 422361 | 211 |
| Large Payload | MessagePack | Unmarshal | 504928 | 364647 | 6661 |
| Large Payload | CBOR | Unmarshal | 649016 | 334107 | 6809 |
| Large Payload | JSON | Unmarshal | 1898479 | 507867 | 6701 |

## unknown — MINGW64_NT-10.0-26100

![Benchmark Chart](benchmark-mingw64-nt-10-0-26100-unknown/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | CBOR | Marshal | 673.60 | 320 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 698.40 | 288 | 2 |
| Small Struct | BEVE | Marshal | 911.30 | 768 | 3 |
| Small Struct | Sonic | Marshal | 1830 | 1823 | 3 |
| Small Struct | JSON | Marshal | 2152 | 1040 | 2 |
| Small Struct | MessagePack | Marshal | 5902 | 8321 | 9 |
| Small Struct | BEVE | Unmarshal | 1793 | 1464 | 4 |
| Small Struct | CBOR | Unmarshal | 1998 | 464 | 13 |
| Small Struct | Sonic | Unmarshal | 4249 | 4155 | 9 |
| Small Struct | MessagePack | Unmarshal | 5773 | 3968 | 84 |
| Small Struct | JSON | Unmarshal | 29883 | 7880 | 112 |
| Medium Payload | BEVE ZeroCopy | Marshal | 12249 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 23125 | 18597 | 3 |
| Medium Payload | CBOR | Marshal | 24596 | 20579 | 2 |
| Medium Payload | Sonic | Marshal | 24803 | 25420 | 4 |
| Medium Payload | MessagePack | Marshal | 44217 | 65832 | 22 |
| Medium Payload | JSON | Marshal | 46649 | 20812 | 9 |
| Medium Payload | BEVE | Unmarshal | 27184 | 18395 | 58 |
| Medium Payload | Sonic | Unmarshal | 58582 | 66209 | 78 |
| Medium Payload | MessagePack | Unmarshal | 62506 | 34895 | 641 |
| Medium Payload | CBOR | Unmarshal | 75643 | 33097 | 683 |
| Medium Payload | JSON | Unmarshal | 171941 | 40521 | 524 |
| Large Payload | BEVE ZeroCopy | Marshal | 127796 | 392 | 2 |
| Large Payload | Sonic | Marshal | 173683 | 211188 | 4 |
| Large Payload | BEVE | Marshal | 198948 | 205411 | 3 |
| Large Payload | CBOR | Marshal | 230193 | 206000 | 2 |
| Large Payload | MessagePack | Marshal | 341397 | 526802 | 115 |
| Large Payload | JSON | Marshal | 456431 | 214368 | 9 |
| Large Payload | BEVE | Unmarshal | 238187 | 155086 | 417 |
| Large Payload | Sonic | Unmarshal | 480840 | 567763 | 593 |
| Large Payload | MessagePack | Unmarshal | 641923 | 361235 | 6597 |
| Large Payload | CBOR | Unmarshal | 726354 | 299243 | 6108 |
| Large Payload | JSON | Unmarshal | 2038417 | 505549 | 6642 |

