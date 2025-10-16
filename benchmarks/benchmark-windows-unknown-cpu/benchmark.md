# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77285 | 286 | 2 |
| Large Payload | BEVE | Marshal | 124312 | 205570 | 3 |
| Large Payload | Sonic | Marshal | 177101 | 227788 | 4 |
| Large Payload | CBOR | Marshal | 219136 | 198191 | 2 |
| Large Payload | MessagePack | Marshal | 298892 | 526762 | 115 |
| Large Payload | JSON | Marshal | 484994 | 214877 | 9 |
| Large Payload | BEVE | Unmarshal | 288817 | 278503 | 417 |
| Large Payload | Sonic | Unmarshal | 438398 | 571011 | 591 |
| Large Payload | MessagePack | Unmarshal | 690593 | 353442 | 6437 |
| Large Payload | CBOR | Unmarshal | 902936 | 333306 | 6791 |
| Large Payload | JSON | Unmarshal | 2638360 | 554204 | 7247 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8297 | 138 | 2 |
| Medium Payload | BEVE | Marshal | 14190 | 19211 | 3 |
| Medium Payload | CBOR | Marshal | 19719 | 16462 | 2 |
| Medium Payload | Sonic | Marshal | 21659 | 27893 | 4 |
| Medium Payload | MessagePack | Marshal | 35317 | 65829 | 22 |
| Medium Payload | JSON | Marshal | 51943 | 22110 | 9 |
| Medium Payload | BEVE | Unmarshal | 32836 | 30556 | 59 |
| Medium Payload | Sonic | Unmarshal | 50116 | 63655 | 77 |
| Medium Payload | MessagePack | Unmarshal | 77487 | 38756 | 726 |
| Medium Payload | CBOR | Unmarshal | 89579 | 32408 | 665 |
| Medium Payload | JSON | Unmarshal | 210942 | 39288 | 536 |
| Small Struct | BEVE ZeroCopy | Marshal | 443 | 290 | 2 |
| Small Struct | CBOR | Marshal | 680 | 528 | 2 |
| Small Struct | BEVE | Marshal | 2726 | 3362 | 3 |
| Small Struct | Sonic | Marshal | 2858 | 2919 | 3 |
| Small Struct | MessagePack | Marshal | 5075 | 8320 | 9 |
| Small Struct | JSON | Marshal | 5136 | 2193 | 2 |
| Small Struct | BEVE | Unmarshal | 2140 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 3599 | 3888 | 9 |
| Small Struct | MessagePack | Unmarshal | 5647 | 3264 | 70 |
| Small Struct | CBOR | Unmarshal | 7664 | 3208 | 69 |
| Small Struct | JSON | Unmarshal | 8079 | 1416 | 31 |
