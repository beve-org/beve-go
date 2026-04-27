# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 73409 | 65 | 0 |
| Large Payload | BEVE | Marshal | 107715 | 180258 | 1 |
| Large Payload | Sonic | Marshal | 162572 | 216034 | 3 |
| Large Payload | CBOR | Marshal | 226576 | 213096 | 1 |
| Large Payload | MessagePack | Marshal | 294291 | 526713 | 115 |
| Large Payload | JSON | Marshal | 518157 | 229679 | 8 |
| Large Payload | BEVE | Unmarshal | 312989 | 273127 | 419 |
| Large Payload | Sonic | Unmarshal | 548609 | 535949 | 576 |
| Large Payload | MessagePack | Unmarshal | 706576 | 359239 | 6566 |
| Large Payload | CBOR | Unmarshal | 865411 | 325113 | 6647 |
| Large Payload | JSON | Unmarshal | 2549983 | 527468 | 6893 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7594 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 11763 | 18436 | 1 |
| Medium Payload | Sonic | Marshal | 20103 | 27611 | 3 |
| Medium Payload | CBOR | Marshal | 23690 | 19085 | 1 |
| Medium Payload | MessagePack | Marshal | 34062 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 57549 | 27497 | 8 |
| Medium Payload | BEVE | Unmarshal | 27337 | 27803 | 58 |
| Medium Payload | Sonic | Unmarshal | 57948 | 59347 | 75 |
| Medium Payload | CBOR | Unmarshal | 67862 | 24872 | 515 |
| Medium Payload | MessagePack | Unmarshal | 81382 | 43166 | 810 |
| Medium Payload | JSON | Unmarshal | 296118 | 69016 | 881 |
| Small Struct | BEVE ZeroCopy | Marshal | 758 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1698 | 2113 | 2 |
| Small Struct | BEVE | Marshal | 1934 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 2654 | 4104 | 8 |
| Small Struct | CBOR | Marshal | 2820 | 2688 | 1 |
| Small Struct | JSON | Marshal | 5237 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 757 | 504 | 4 |
| Small Struct | Sonic | Unmarshal | 4152 | 4389 | 9 |
| Small Struct | MessagePack | Unmarshal | 5273 | 3456 | 72 |
| Small Struct | CBOR | Unmarshal | 5297 | 2408 | 52 |
| Small Struct | JSON | Unmarshal | 13227 | 3664 | 51 |
