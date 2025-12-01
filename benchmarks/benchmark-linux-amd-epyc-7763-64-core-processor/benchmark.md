# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 74510 | 26 | 0 |
| Large Payload | BEVE | Marshal | 119625 | 204847 | 1 |
| Large Payload | Sonic | Marshal | 145530 | 207392 | 3 |
| Large Payload | CBOR | Marshal | 207405 | 196760 | 1 |
| Large Payload | MessagePack | Marshal | 316124 | 526778 | 115 |
| Large Payload | JSON | Marshal | 468513 | 229695 | 8 |
| Large Payload | BEVE | Unmarshal | 241544 | 280770 | 418 |
| Large Payload | Sonic | Unmarshal | 355693 | 557281 | 592 |
| Large Payload | MessagePack | Unmarshal | 577953 | 361832 | 6610 |
| Large Payload | CBOR | Unmarshal | 701890 | 314809 | 6410 |
| Large Payload | JSON | Unmarshal | 2498352 | 605786 | 7876 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7208 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 17468 | 24582 | 1 |
| Medium Payload | CBOR | Marshal | 18803 | 16398 | 1 |
| Medium Payload | Sonic | Marshal | 19537 | 27735 | 3 |
| Medium Payload | MessagePack | Marshal | 35606 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 36678 | 18666 | 8 |
| Medium Payload | BEVE | Unmarshal | 22947 | 25854 | 59 |
| Medium Payload | Sonic | Unmarshal | 35115 | 51431 | 70 |
| Medium Payload | MessagePack | Unmarshal | 55467 | 35504 | 655 |
| Medium Payload | CBOR | Unmarshal | 74853 | 31960 | 661 |
| Medium Payload | JSON | Unmarshal | 223817 | 55048 | 719 |
| Small Struct | BEVE ZeroCopy | Marshal | 665 | 0 | 0 |
| Small Struct | JSON | Marshal | 1026 | 416 | 1 |
| Small Struct | BEVE | Marshal | 1137 | 1536 | 1 |
| Small Struct | Sonic | Marshal | 1576 | 2349 | 2 |
| Small Struct | CBOR | Marshal | 2413 | 2689 | 1 |
| Small Struct | MessagePack | Marshal | 2629 | 4104 | 8 |
| Small Struct | CBOR | Unmarshal | 1397 | 368 | 11 |
| Small Struct | BEVE | Unmarshal | 1692 | 2616 | 4 |
| Small Struct | Sonic | Unmarshal | 2455 | 3793 | 9 |
| Small Struct | MessagePack | Unmarshal | 3670 | 2496 | 54 |
| Small Struct | JSON | Unmarshal | 4836 | 904 | 22 |
