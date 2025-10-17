# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 89323 | 52 | 0 |
| Large Payload | BEVE | Marshal | 144248 | 188440 | 1 |
| Large Payload | Sonic | Marshal | 190627 | 215080 | 3 |
| Large Payload | CBOR | Marshal | 216046 | 188544 | 1 |
| Large Payload | MessagePack | Marshal | 284193 | 526708 | 115 |
| Large Payload | JSON | Marshal | 515311 | 221485 | 8 |
| Large Payload | BEVE | Unmarshal | 283555 | 275654 | 419 |
| Large Payload | Sonic | Unmarshal | 435129 | 569977 | 579 |
| Large Payload | MessagePack | Unmarshal | 685814 | 370504 | 6806 |
| Large Payload | CBOR | Unmarshal | 895950 | 335370 | 6837 |
| Large Payload | JSON | Unmarshal | 2575251 | 518549 | 6851 |
| Medium Payload | BEVE ZeroCopy | Marshal | 10941 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 15791 | 16389 | 1 |
| Medium Payload | Sonic | Marshal | 20771 | 20755 | 3 |
| Medium Payload | CBOR | Marshal | 25664 | 19083 | 1 |
| Medium Payload | MessagePack | Marshal | 43125 | 65773 | 22 |
| Medium Payload | JSON | Marshal | 70663 | 27502 | 8 |
| Medium Payload | BEVE | Unmarshal | 40159 | 31068 | 59 |
| Medium Payload | Sonic | Unmarshal | 55616 | 56002 | 70 |
| Medium Payload | MessagePack | Unmarshal | 93810 | 42063 | 788 |
| Medium Payload | CBOR | Unmarshal | 120723 | 38584 | 792 |
| Medium Payload | JSON | Unmarshal | 315639 | 57488 | 737 |
| Small Struct | BEVE ZeroCopy | Marshal | 1194 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1418 | 1408 | 1 |
| Small Struct | CBOR | Marshal | 1882 | 1280 | 1 |
| Small Struct | Sonic | Marshal | 2845 | 2113 | 2 |
| Small Struct | JSON | Marshal | 3185 | 896 | 1 |
| Small Struct | MessagePack | Marshal | 6602 | 8200 | 9 |
| Small Struct | CBOR | Unmarshal | 2744 | 616 | 16 |
| Small Struct | BEVE | Unmarshal | 2898 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 6642 | 3448 | 72 |
| Small Struct | Sonic | Unmarshal | 7098 | 7398 | 10 |
| Small Struct | JSON | Unmarshal | 11520 | 2120 | 38 |
