# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 61910 | 26 | 0 |
| Large Payload | BEVE | Marshal | 103161 | 180264 | 1 |
| Large Payload | CBOR | Marshal | 193478 | 196757 | 1 |
| Large Payload | MessagePack | Marshal | 254142 | 526756 | 115 |
| Large Payload | JSON | Marshal | 384952 | 213280 | 8 |
| Large Payload | Sonic | Marshal | 505113 | 222260 | 3 |
| Large Payload | BEVE | Unmarshal | 228955 | 258989 | 418 |
| Large Payload | Sonic | Unmarshal | 258873 | 329771 | 211 |
| Large Payload | CBOR | Unmarshal | 512699 | 321961 | 6582 |
| Large Payload | MessagePack | Unmarshal | 539229 | 346980 | 6302 |
| Large Payload | JSON | Unmarshal | 1899005 | 526041 | 6863 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5229 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 15617 | 20484 | 1 |
| Medium Payload | CBOR | Marshal | 17646 | 18454 | 1 |
| Medium Payload | MessagePack | Marshal | 27989 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 34278 | 18658 | 8 |
| Medium Payload | Sonic | Marshal | 50007 | 19321 | 3 |
| Medium Payload | BEVE | Unmarshal | 21096 | 27516 | 59 |
| Medium Payload | Sonic | Unmarshal | 36454 | 32520 | 33 |
| Medium Payload | MessagePack | Unmarshal | 48839 | 32989 | 606 |
| Medium Payload | CBOR | Unmarshal | 70429 | 33960 | 695 |
| Medium Payload | JSON | Unmarshal | 209360 | 49304 | 645 |
| Small Struct | BEVE ZeroCopy | Marshal | 434 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 1079 | 1032 | 6 |
| Small Struct | Sonic | Marshal | 1148 | 464 | 2 |
| Small Struct | BEVE | Marshal | 1190 | 2048 | 1 |
| Small Struct | JSON | Marshal | 1331 | 704 | 1 |
| Small Struct | CBOR | Marshal | 1991 | 1536 | 1 |
| Small Struct | BEVE | Unmarshal | 2610 | 1592 | 4 |
| Small Struct | CBOR | Unmarshal | 4735 | 2312 | 51 |
| Small Struct | MessagePack | Unmarshal | 5763 | 3904 | 82 |
| Small Struct | Sonic | Unmarshal | 6375 | 5397 | 6 |
| Small Struct | JSON | Unmarshal | 16389 | 2280 | 43 |
