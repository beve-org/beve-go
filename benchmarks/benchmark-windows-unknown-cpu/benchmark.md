# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80728 | 65 | 0 |
| Large Payload | BEVE | Marshal | 126362 | 196697 | 1 |
| Large Payload | Sonic | Marshal | 153641 | 206392 | 3 |
| Large Payload | CBOR | Marshal | 250815 | 196689 | 1 |
| Large Payload | MessagePack | Marshal | 278505 | 526707 | 115 |
| Large Payload | JSON | Marshal | 505246 | 229652 | 8 |
| Large Payload | BEVE | Unmarshal | 273828 | 272358 | 418 |
| Large Payload | Sonic | Unmarshal | 411071 | 522803 | 573 |
| Large Payload | MessagePack | Unmarshal | 866248 | 345388 | 6289 |
| Large Payload | CBOR | Unmarshal | 1043174 | 330953 | 6747 |
| Large Payload | JSON | Unmarshal | 2658986 | 543830 | 7061 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7529 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 12152 | 18436 | 1 |
| Medium Payload | Sonic | Marshal | 17274 | 20800 | 3 |
| Medium Payload | CBOR | Marshal | 26734 | 24588 | 1 |
| Medium Payload | MessagePack | Marshal | 36430 | 65773 | 22 |
| Medium Payload | JSON | Marshal | 45465 | 18661 | 8 |
| Medium Payload | BEVE | Unmarshal | 27708 | 27355 | 59 |
| Medium Payload | Sonic | Unmarshal | 46368 | 56929 | 76 |
| Medium Payload | MessagePack | Unmarshal | 68207 | 34668 | 638 |
| Medium Payload | CBOR | Unmarshal | 90187 | 34328 | 705 |
| Medium Payload | JSON | Unmarshal | 280817 | 59808 | 769 |
| Small Struct | BEVE ZeroCopy | Marshal | 572 | 0 | 0 |
| Small Struct | CBOR | Marshal | 927 | 768 | 1 |
| Small Struct | Sonic | Marshal | 1317 | 1843 | 2 |
| Small Struct | BEVE | Marshal | 2201 | 2689 | 1 |
| Small Struct | JSON | Marshal | 2713 | 1280 | 1 |
| Small Struct | MessagePack | Marshal | 4389 | 8200 | 9 |
| Small Struct | BEVE | Unmarshal | 797 | 408 | 4 |
| Small Struct | Sonic | Unmarshal | 5788 | 7344 | 10 |
| Small Struct | MessagePack | Unmarshal | 7425 | 5216 | 107 |
| Small Struct | CBOR | Unmarshal | 8602 | 3944 | 84 |
| Small Struct | JSON | Unmarshal | 21275 | 4480 | 77 |
