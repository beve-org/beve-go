# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 68742 | 65 | 0 |
| Large Payload | BEVE | Marshal | 104449 | 180330 | 1 |
| Large Payload | CBOR | Marshal | 193288 | 205082 | 1 |
| Large Payload | MessagePack | Marshal | 267855 | 526802 | 115 |
| Large Payload | Sonic | Marshal | 291499 | 208193 | 3 |
| Large Payload | JSON | Marshal | 356177 | 197086 | 8 |
| Large Payload | BEVE | Unmarshal | 213795 | 260037 | 418 |
| Large Payload | Sonic | Unmarshal | 275107 | 371881 | 211 |
| Large Payload | MessagePack | Unmarshal | 476142 | 316658 | 5697 |
| Large Payload | CBOR | Unmarshal | 682102 | 338427 | 6897 |
| Large Payload | JSON | Unmarshal | 1899315 | 513100 | 6730 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7016 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 9411 | 18437 | 1 |
| Medium Payload | CBOR | Marshal | 17151 | 18447 | 1 |
| Medium Payload | Sonic | Marshal | 29284 | 22138 | 3 |
| Medium Payload | MessagePack | Marshal | 31701 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 38383 | 21988 | 8 |
| Medium Payload | BEVE | Unmarshal | 23144 | 31391 | 59 |
| Medium Payload | Sonic | Unmarshal | 32483 | 46216 | 33 |
| Medium Payload | MessagePack | Unmarshal | 55776 | 39776 | 745 |
| Medium Payload | CBOR | Unmarshal | 69386 | 34904 | 718 |
| Medium Payload | JSON | Unmarshal | 208148 | 59496 | 776 |
| Small Struct | BEVE ZeroCopy | Marshal | 662 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1307 | 2689 | 1 |
| Small Struct | CBOR | Marshal | 1772 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 2493 | 4104 | 8 |
| Small Struct | JSON | Marshal | 2733 | 1536 | 1 |
| Small Struct | Sonic | Marshal | 3524 | 2802 | 2 |
| Small Struct | BEVE | Unmarshal | 1287 | 2104 | 4 |
| Small Struct | Sonic | Unmarshal | 3283 | 5724 | 6 |
| Small Struct | MessagePack | Unmarshal | 3850 | 3136 | 66 |
| Small Struct | CBOR | Unmarshal | 5978 | 3560 | 76 |
| Small Struct | JSON | Unmarshal | 9699 | 2408 | 47 |
