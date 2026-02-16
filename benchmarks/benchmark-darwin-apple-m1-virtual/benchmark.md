# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 70802 | 39 | 0 |
| Large Payload | BEVE | Marshal | 131150 | 196676 | 1 |
| Large Payload | CBOR | Marshal | 226068 | 196735 | 1 |
| Large Payload | MessagePack | Marshal | 382558 | 526753 | 115 |
| Large Payload | JSON | Marshal | 476869 | 213305 | 8 |
| Large Payload | Sonic | Marshal | 541102 | 214375 | 3 |
| Large Payload | BEVE | Unmarshal | 299733 | 267951 | 418 |
| Large Payload | Sonic | Unmarshal | 439712 | 343437 | 211 |
| Large Payload | MessagePack | Unmarshal | 701914 | 373513 | 6854 |
| Large Payload | CBOR | Unmarshal | 860729 | 309290 | 6304 |
| Large Payload | JSON | Unmarshal | 2628970 | 537211 | 7045 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7579 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 11676 | 18436 | 1 |
| Medium Payload | CBOR | Marshal | 25424 | 19085 | 1 |
| Medium Payload | MessagePack | Marshal | 34275 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 42587 | 18665 | 8 |
| Medium Payload | Sonic | Marshal | 46864 | 19312 | 3 |
| Medium Payload | BEVE | Unmarshal | 25684 | 27228 | 59 |
| Medium Payload | Sonic | Unmarshal | 41078 | 42983 | 33 |
| Medium Payload | MessagePack | Unmarshal | 66654 | 39902 | 747 |
| Medium Payload | CBOR | Unmarshal | 70560 | 33480 | 689 |
| Medium Payload | JSON | Unmarshal | 226595 | 55944 | 721 |
| Small Struct | BEVE ZeroCopy | Marshal | 470 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1716 | 1792 | 1 |
| Small Struct | BEVE | Marshal | 2145 | 1792 | 1 |
| Small Struct | MessagePack | Marshal | 2733 | 4104 | 8 |
| Small Struct | JSON | Marshal | 3538 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 4397 | 1857 | 2 |
| Small Struct | BEVE | Unmarshal | 1824 | 1848 | 4 |
| Small Struct | Sonic | Unmarshal | 2447 | 2365 | 6 |
| Small Struct | MessagePack | Unmarshal | 4410 | 2752 | 58 |
| Small Struct | CBOR | Unmarshal | 4483 | 2472 | 54 |
| Small Struct | JSON | Unmarshal | 29135 | 7912 | 113 |
