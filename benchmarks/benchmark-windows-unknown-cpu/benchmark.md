# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79742 | 79 | 0 |
| Large Payload | BEVE | Marshal | 118504 | 196671 | 1 |
| Large Payload | Sonic | Marshal | 153797 | 206275 | 3 |
| Large Payload | CBOR | Marshal | 213431 | 188488 | 1 |
| Large Payload | MessagePack | Marshal | 276772 | 526707 | 115 |
| Large Payload | JSON | Marshal | 495578 | 221486 | 8 |
| Large Payload | BEVE | Unmarshal | 271131 | 264708 | 417 |
| Large Payload | Sonic | Unmarshal | 440542 | 569494 | 588 |
| Large Payload | MessagePack | Unmarshal | 648244 | 339109 | 6166 |
| Large Payload | CBOR | Unmarshal | 788637 | 290890 | 5920 |
| Large Payload | JSON | Unmarshal | 2665704 | 564156 | 7371 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9953 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 13540 | 20482 | 1 |
| Medium Payload | Sonic | Marshal | 16308 | 19416 | 3 |
| Medium Payload | CBOR | Marshal | 21112 | 18440 | 1 |
| Medium Payload | MessagePack | Marshal | 33133 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 51198 | 24809 | 8 |
| Medium Payload | BEVE | Unmarshal | 28563 | 28699 | 59 |
| Medium Payload | Sonic | Unmarshal | 53810 | 66584 | 79 |
| Medium Payload | MessagePack | Unmarshal | 68509 | 35788 | 661 |
| Medium Payload | CBOR | Unmarshal | 94625 | 37560 | 769 |
| Medium Payload | JSON | Unmarshal | 246877 | 51960 | 691 |
| Small Struct | BEVE ZeroCopy | Marshal | 255 | 0 | 0 |
| Small Struct | Sonic | Marshal | 626 | 684 | 2 |
| Small Struct | CBOR | Marshal | 1423 | 1152 | 1 |
| Small Struct | MessagePack | Marshal | 1436 | 1032 | 6 |
| Small Struct | BEVE | Marshal | 1470 | 1536 | 1 |
| Small Struct | JSON | Marshal | 3600 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 1419 | 1592 | 4 |
| Small Struct | Sonic | Unmarshal | 3333 | 3888 | 9 |
| Small Struct | MessagePack | Unmarshal | 3434 | 1696 | 38 |
| Small Struct | CBOR | Unmarshal | 5400 | 2312 | 51 |
| Small Struct | JSON | Unmarshal | 29237 | 7560 | 102 |
