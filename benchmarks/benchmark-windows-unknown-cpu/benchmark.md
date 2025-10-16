# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80755 | 105 | 0 |
| Large Payload | BEVE | Marshal | 109194 | 188476 | 1 |
| Large Payload | Sonic | Marshal | 161169 | 214958 | 3 |
| Large Payload | CBOR | Marshal | 224326 | 196711 | 1 |
| Large Payload | MessagePack | Marshal | 285229 | 526710 | 115 |
| Large Payload | JSON | Marshal | 489334 | 221487 | 8 |
| Large Payload | BEVE | Unmarshal | 336760 | 261160 | 418 |
| Large Payload | Sonic | Unmarshal | 439969 | 570793 | 592 |
| Large Payload | MessagePack | Unmarshal | 675057 | 349445 | 6354 |
| Large Payload | CBOR | Unmarshal | 890510 | 340266 | 6932 |
| Large Payload | JSON | Unmarshal | 2653688 | 553643 | 7233 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9270 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12875 | 18436 | 1 |
| Medium Payload | Sonic | Marshal | 19274 | 27598 | 3 |
| Medium Payload | CBOR | Marshal | 22554 | 19091 | 1 |
| Medium Payload | MessagePack | Marshal | 35442 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 44110 | 19305 | 8 |
| Medium Payload | BEVE | Unmarshal | 26252 | 24699 | 59 |
| Medium Payload | Sonic | Unmarshal | 53501 | 64787 | 77 |
| Medium Payload | MessagePack | Unmarshal | 72985 | 38397 | 713 |
| Medium Payload | CBOR | Unmarshal | 79751 | 30136 | 621 |
| Medium Payload | JSON | Unmarshal | 231608 | 48984 | 624 |
| Small Struct | BEVE ZeroCopy | Marshal | 590 | 1 | 0 |
| Small Struct | BEVE | Marshal | 1180 | 1280 | 1 |
| Small Struct | CBOR | Marshal | 1193 | 1024 | 1 |
| Small Struct | Sonic | Marshal | 1824 | 2390 | 2 |
| Small Struct | MessagePack | Marshal | 3312 | 4104 | 8 |
| Small Struct | JSON | Marshal | 5385 | 2689 | 1 |
| Small Struct | BEVE | Unmarshal | 1937 | 2616 | 4 |
| Small Struct | MessagePack | Unmarshal | 2011 | 688 | 17 |
| Small Struct | Sonic | Unmarshal | 2138 | 2203 | 8 |
| Small Struct | CBOR | Unmarshal | 6620 | 2664 | 56 |
| Small Struct | JSON | Unmarshal | 7168 | 1352 | 29 |
