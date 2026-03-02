# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 57468 | 26 | 0 |
| Large Payload | BEVE | Marshal | 78232 | 196648 | 1 |
| Large Payload | CBOR | Marshal | 150500 | 204938 | 1 |
| Large Payload | MessagePack | Marshal | 190111 | 526752 | 115 |
| Large Payload | JSON | Marshal | 324948 | 213305 | 8 |
| Large Payload | Sonic | Marshal | 394055 | 205369 | 3 |
| Large Payload | BEVE | Unmarshal | 195912 | 282577 | 418 |
| Large Payload | Sonic | Unmarshal | 270279 | 335497 | 211 |
| Large Payload | MessagePack | Unmarshal | 387957 | 345701 | 6284 |
| Large Payload | CBOR | Unmarshal | 457082 | 293977 | 6001 |
| Large Payload | JSON | Unmarshal | 1727354 | 536299 | 6982 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5700 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9269 | 24581 | 1 |
| Medium Payload | CBOR | Marshal | 13126 | 18444 | 1 |
| Medium Payload | MessagePack | Marshal | 21213 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 35020 | 24810 | 8 |
| Medium Payload | Sonic | Marshal | 38679 | 21952 | 3 |
| Medium Payload | BEVE | Unmarshal | 15727 | 27548 | 59 |
| Medium Payload | Sonic | Unmarshal | 24454 | 30299 | 33 |
| Medium Payload | MessagePack | Unmarshal | 35007 | 32941 | 607 |
| Medium Payload | CBOR | Unmarshal | 50892 | 34520 | 706 |
| Medium Payload | JSON | Unmarshal | 157930 | 49048 | 642 |
| Small Struct | BEVE ZeroCopy | Marshal | 495 | 0 | 0 |
| Small Struct | BEVE | Marshal | 688 | 1792 | 1 |
| Small Struct | JSON | Marshal | 993 | 704 | 1 |
| Small Struct | CBOR | Marshal | 1485 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 2829 | 8201 | 9 |
| Small Struct | Sonic | Marshal | 3731 | 2107 | 2 |
| Small Struct | BEVE | Unmarshal | 717 | 1080 | 4 |
| Small Struct | CBOR | Unmarshal | 2776 | 2088 | 46 |
| Small Struct | Sonic | Unmarshal | 3302 | 5653 | 6 |
| Small Struct | MessagePack | Unmarshal | 3357 | 4000 | 85 |
| Small Struct | JSON | Unmarshal | 13003 | 4296 | 71 |
