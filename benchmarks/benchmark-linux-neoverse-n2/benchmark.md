# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69840 | 105 | 0 |
| Large Payload | BEVE | Marshal | 105457 | 188563 | 1 |
| Large Payload | CBOR | Marshal | 201772 | 205002 | 1 |
| Large Payload | MessagePack | Marshal | 283715 | 526807 | 115 |
| Large Payload | Sonic | Marshal | 306640 | 217683 | 3 |
| Large Payload | JSON | Marshal | 384776 | 213579 | 8 |
| Large Payload | BEVE | Unmarshal | 221084 | 271915 | 419 |
| Large Payload | Sonic | Unmarshal | 285224 | 391976 | 213 |
| Large Payload | MessagePack | Unmarshal | 523331 | 358220 | 6544 |
| Large Payload | CBOR | Unmarshal | 669311 | 333148 | 6795 |
| Large Payload | JSON | Unmarshal | 1945492 | 521317 | 6867 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7465 | 10 | 0 |
| Medium Payload | BEVE | Marshal | 10431 | 20491 | 1 |
| Medium Payload | CBOR | Marshal | 18730 | 20499 | 1 |
| Medium Payload | Sonic | Marshal | 31516 | 25004 | 3 |
| Medium Payload | MessagePack | Marshal | 31902 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 37549 | 21988 | 8 |
| Medium Payload | BEVE | Unmarshal | 22215 | 27902 | 59 |
| Medium Payload | Sonic | Unmarshal | 29040 | 39618 | 33 |
| Medium Payload | MessagePack | Unmarshal | 51051 | 35623 | 660 |
| Medium Payload | CBOR | Unmarshal | 67096 | 34216 | 702 |
| Medium Payload | JSON | Unmarshal | 173272 | 47384 | 615 |
| Small Struct | BEVE ZeroCopy | Marshal | 609 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 703 | 520 | 5 |
| Small Struct | BEVE | Marshal | 834 | 1408 | 1 |
| Small Struct | CBOR | Marshal | 2315 | 2688 | 1 |
| Small Struct | Sonic | Marshal | 2432 | 1878 | 2 |
| Small Struct | JSON | Marshal | 2572 | 1536 | 1 |
| Small Struct | BEVE | Unmarshal | 802 | 728 | 4 |
| Small Struct | MessagePack | Unmarshal | 1358 | 496 | 13 |
| Small Struct | CBOR | Unmarshal | 2828 | 1352 | 31 |
| Small Struct | Sonic | Unmarshal | 2881 | 4979 | 6 |
| Small Struct | JSON | Unmarshal | 26022 | 8072 | 118 |
