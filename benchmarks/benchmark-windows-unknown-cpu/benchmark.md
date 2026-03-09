# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 85110 | 65 | 0 |
| Large Payload | BEVE | Marshal | 146092 | 213017 | 1 |
| Large Payload | Sonic | Marshal | 163439 | 215257 | 3 |
| Large Payload | CBOR | Marshal | 222016 | 196735 | 1 |
| Large Payload | MessagePack | Marshal | 277698 | 526706 | 115 |
| Large Payload | JSON | Marshal | 481165 | 205072 | 8 |
| Large Payload | BEVE | Unmarshal | 293610 | 260773 | 419 |
| Large Payload | Sonic | Unmarshal | 467186 | 564519 | 582 |
| Large Payload | MessagePack | Unmarshal | 789234 | 344952 | 6258 |
| Large Payload | CBOR | Unmarshal | 893410 | 313609 | 6394 |
| Large Payload | JSON | Unmarshal | 2812263 | 522420 | 6897 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6655 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 12510 | 18433 | 1 |
| Medium Payload | Sonic | Marshal | 15916 | 19354 | 3 |
| Medium Payload | MessagePack | Marshal | 24483 | 33002 | 21 |
| Medium Payload | CBOR | Marshal | 28202 | 20491 | 1 |
| Medium Payload | JSON | Marshal | 43876 | 20713 | 8 |
| Medium Payload | BEVE | Unmarshal | 33168 | 32476 | 59 |
| Medium Payload | Sonic | Unmarshal | 50506 | 50619 | 61 |
| Medium Payload | MessagePack | Unmarshal | 76115 | 33788 | 620 |
| Medium Payload | CBOR | Unmarshal | 108067 | 34168 | 700 |
| Medium Payload | JSON | Unmarshal | 321598 | 62904 | 787 |
| Small Struct | BEVE ZeroCopy | Marshal | 349 | 0 | 0 |
| Small Struct | BEVE | Marshal | 811 | 768 | 1 |
| Small Struct | Sonic | Marshal | 1223 | 1445 | 2 |
| Small Struct | CBOR | Marshal | 2328 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 2758 | 4104 | 8 |
| Small Struct | JSON | Marshal | 6624 | 3072 | 1 |
| Small Struct | MessagePack | Unmarshal | 2261 | 1024 | 24 |
| Small Struct | BEVE | Unmarshal | 2655 | 3512 | 4 |
| Small Struct | Sonic | Unmarshal | 4381 | 4389 | 9 |
| Small Struct | CBOR | Unmarshal | 10274 | 4808 | 103 |
| Small Struct | JSON | Unmarshal | 11277 | 2312 | 44 |
