# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79472 | 26 | 0 |
| Large Payload | BEVE | Marshal | 115986 | 196679 | 1 |
| Large Payload | Sonic | Marshal | 148134 | 207444 | 3 |
| Large Payload | CBOR | Marshal | 210415 | 196760 | 1 |
| Large Payload | MessagePack | Marshal | 315860 | 526778 | 115 |
| Large Payload | JSON | Marshal | 446388 | 213309 | 8 |
| Large Payload | BEVE | Unmarshal | 242248 | 276289 | 418 |
| Large Payload | Sonic | Unmarshal | 353979 | 529823 | 569 |
| Large Payload | MessagePack | Unmarshal | 571194 | 350023 | 6358 |
| Large Payload | CBOR | Unmarshal | 696293 | 308234 | 6275 |
| Large Payload | JSON | Unmarshal | 2262034 | 534921 | 6941 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8018 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 13116 | 21766 | 1 |
| Medium Payload | Sonic | Marshal | 13370 | 18921 | 3 |
| Medium Payload | CBOR | Marshal | 22413 | 20509 | 1 |
| Medium Payload | MessagePack | Marshal | 35739 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 36039 | 18663 | 8 |
| Medium Payload | BEVE | Unmarshal | 23433 | 26718 | 59 |
| Medium Payload | Sonic | Unmarshal | 44341 | 63178 | 76 |
| Medium Payload | MessagePack | Unmarshal | 55708 | 36848 | 678 |
| Medium Payload | CBOR | Unmarshal | 68357 | 31208 | 644 |
| Medium Payload | JSON | Unmarshal | 225235 | 56904 | 731 |
| Small Struct | BEVE ZeroCopy | Marshal | 631 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1082 | 1460 | 2 |
| Small Struct | MessagePack | Marshal | 1134 | 1032 | 6 |
| Small Struct | BEVE | Marshal | 1475 | 2688 | 1 |
| Small Struct | JSON | Marshal | 2629 | 1280 | 1 |
| Small Struct | CBOR | Marshal | 2838 | 3073 | 1 |
| Small Struct | BEVE | Unmarshal | 1650 | 2616 | 4 |
| Small Struct | CBOR | Unmarshal | 1693 | 520 | 14 |
| Small Struct | Sonic | Unmarshal | 3259 | 4683 | 9 |
| Small Struct | MessagePack | Unmarshal | 4327 | 3200 | 68 |
| Small Struct | JSON | Unmarshal | 4847 | 904 | 22 |
