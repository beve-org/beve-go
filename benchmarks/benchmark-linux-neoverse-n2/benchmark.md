# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69453 | 65 | 0 |
| Large Payload | BEVE | Marshal | 104905 | 180422 | 1 |
| Large Payload | CBOR | Marshal | 185008 | 188872 | 1 |
| Large Payload | MessagePack | Marshal | 286921 | 526805 | 115 |
| Large Payload | Sonic | Marshal | 316286 | 225050 | 3 |
| Large Payload | JSON | Marshal | 393474 | 213448 | 8 |
| Large Payload | BEVE | Unmarshal | 227286 | 268074 | 419 |
| Large Payload | Sonic | Unmarshal | 291417 | 401893 | 211 |
| Large Payload | MessagePack | Unmarshal | 523465 | 354685 | 6462 |
| Large Payload | CBOR | Unmarshal | 642610 | 308298 | 6293 |
| Large Payload | JSON | Unmarshal | 1947444 | 521092 | 6858 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7897 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9829 | 18444 | 1 |
| Medium Payload | CBOR | Marshal | 17892 | 18447 | 1 |
| Medium Payload | Sonic | Marshal | 29739 | 22119 | 3 |
| Medium Payload | MessagePack | Marshal | 32602 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 38781 | 21991 | 8 |
| Medium Payload | BEVE | Unmarshal | 22457 | 27870 | 59 |
| Medium Payload | Sonic | Unmarshal | 34573 | 52413 | 33 |
| Medium Payload | MessagePack | Unmarshal | 55994 | 39232 | 732 |
| Medium Payload | CBOR | Unmarshal | 62056 | 29624 | 608 |
| Medium Payload | JSON | Unmarshal | 209726 | 58568 | 780 |
| Small Struct | CBOR | Marshal | 625 | 416 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 704 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1033 | 1793 | 1 |
| Small Struct | Sonic | Marshal | 2155 | 1459 | 2 |
| Small Struct | MessagePack | Marshal | 2570 | 4104 | 8 |
| Small Struct | JSON | Marshal | 4374 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 898 | 728 | 4 |
| Small Struct | Sonic | Unmarshal | 2860 | 4455 | 6 |
| Small Struct | MessagePack | Unmarshal | 2988 | 2112 | 46 |
| Small Struct | CBOR | Unmarshal | 5941 | 3528 | 75 |
| Small Struct | JSON | Unmarshal | 8455 | 2216 | 41 |
