# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 72598 | 207 | 2 |
| Large Payload | BEVE | Marshal | 135988 | 189282 | 3 |
| Large Payload | CBOR | Marshal | 167110 | 205504 | 2 |
| Large Payload | MessagePack | Marshal | 211850 | 526814 | 115 |
| Large Payload | JSON | Marshal | 397453 | 222142 | 9 |
| Large Payload | Sonic | Marshal | 481848 | 223918 | 4 |
| Large Payload | BEVE | Unmarshal | 223113 | 265647 | 417 |
| Large Payload | Sonic | Unmarshal | 280161 | 333240 | 213 |
| Large Payload | MessagePack | Unmarshal | 451547 | 363607 | 6668 |
| Large Payload | CBOR | Unmarshal | 804733 | 329529 | 6720 |
| Large Payload | JSON | Unmarshal | 1755515 | 532761 | 6911 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6693 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 11829 | 18573 | 3 |
| Medium Payload | CBOR | Marshal | 29751 | 24685 | 2 |
| Medium Payload | JSON | Marshal | 30404 | 18722 | 9 |
| Medium Payload | MessagePack | Marshal | 34135 | 65834 | 22 |
| Medium Payload | Sonic | Marshal | 55581 | 25015 | 4 |
| Medium Payload | BEVE | Unmarshal | 29202 | 27548 | 58 |
| Medium Payload | Sonic | Unmarshal | 38357 | 43693 | 33 |
| Medium Payload | CBOR | Unmarshal | 55917 | 26520 | 550 |
| Medium Payload | MessagePack | Unmarshal | 65327 | 41519 | 775 |
| Medium Payload | JSON | Unmarshal | 237781 | 58840 | 754 |
| Small Struct | JSON | Marshal | 699 | 496 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 820 | 288 | 2 |
| Small Struct | CBOR | Marshal | 1038 | 1296 | 2 |
| Small Struct | BEVE | Marshal | 1039 | 2081 | 3 |
| Small Struct | MessagePack | Marshal | 3146 | 8321 | 9 |
| Small Struct | Sonic | Marshal | 5275 | 2905 | 3 |
| Small Struct | CBOR | Unmarshal | 1061 | 432 | 12 |
| Small Struct | BEVE | Unmarshal | 1344 | 1592 | 4 |
| Small Struct | Sonic | Unmarshal | 2303 | 3290 | 6 |
| Small Struct | MessagePack | Unmarshal | 2309 | 2328 | 51 |
| Small Struct | JSON | Unmarshal | 5517 | 1384 | 30 |
