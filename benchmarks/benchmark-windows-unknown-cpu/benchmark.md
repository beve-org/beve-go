# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 95292 | 312 | 2 |
| Large Payload | BEVE | Marshal | 133364 | 189161 | 3 |
| Large Payload | Sonic | Marshal | 157748 | 211427 | 4 |
| Large Payload | CBOR | Marshal | 225646 | 206068 | 2 |
| Large Payload | MessagePack | Marshal | 283847 | 526765 | 115 |
| Large Payload | JSON | Marshal | 488262 | 215613 | 9 |
| Large Payload | BEVE | Unmarshal | 288667 | 274596 | 417 |
| Large Payload | Sonic | Unmarshal | 428119 | 518759 | 564 |
| Large Payload | MessagePack | Unmarshal | 671523 | 341200 | 6203 |
| Large Payload | CBOR | Unmarshal | 854170 | 313977 | 6385 |
| Large Payload | JSON | Unmarshal | 2329751 | 482395 | 6348 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9786 | 138 | 2 |
| Medium Payload | BEVE | Marshal | 13658 | 16533 | 3 |
| Medium Payload | Sonic | Marshal | 17139 | 20930 | 4 |
| Medium Payload | CBOR | Marshal | 26189 | 21879 | 2 |
| Medium Payload | MessagePack | Marshal | 37111 | 65828 | 22 |
| Medium Payload | JSON | Marshal | 48011 | 22090 | 9 |
| Medium Payload | BEVE | Unmarshal | 28069 | 30426 | 59 |
| Medium Payload | Sonic | Unmarshal | 53153 | 63846 | 77 |
| Medium Payload | MessagePack | Unmarshal | 60819 | 27579 | 497 |
| Medium Payload | CBOR | Unmarshal | 83275 | 27816 | 575 |
| Medium Payload | JSON | Unmarshal | 285654 | 60104 | 802 |
| Small Struct | BEVE ZeroCopy | Marshal | 658 | 289 | 2 |
| Small Struct | Sonic | Marshal | 1698 | 2023 | 3 |
| Small Struct | CBOR | Marshal | 2352 | 2449 | 2 |
| Small Struct | BEVE | Marshal | 2579 | 1696 | 3 |
| Small Struct | MessagePack | Marshal | 5469 | 8320 | 9 |
| Small Struct | JSON | Marshal | 6670 | 3218 | 2 |
| Small Struct | BEVE | Unmarshal | 1643 | 2104 | 4 |
| Small Struct | CBOR | Unmarshal | 2492 | 808 | 20 |
| Small Struct | Sonic | Unmarshal | 5896 | 7750 | 10 |
| Small Struct | MessagePack | Unmarshal | 6598 | 4288 | 90 |
| Small Struct | JSON | Unmarshal | 26364 | 7272 | 93 |
