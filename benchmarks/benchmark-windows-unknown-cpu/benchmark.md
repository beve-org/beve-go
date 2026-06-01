# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 81541 | 52 | 0 |
| Large Payload | BEVE | Marshal | 137097 | 204849 | 1 |
| Large Payload | Sonic | Marshal | 181417 | 215368 | 3 |
| Large Payload | CBOR | Marshal | 239107 | 196738 | 1 |
| Large Payload | MessagePack | Marshal | 320516 | 526709 | 115 |
| Large Payload | JSON | Marshal | 519127 | 213265 | 8 |
| Large Payload | BEVE | Unmarshal | 332891 | 270405 | 418 |
| Large Payload | Sonic | Unmarshal | 485672 | 571315 | 598 |
| Large Payload | CBOR | Unmarshal | 884924 | 332107 | 6777 |
| Large Payload | MessagePack | Unmarshal | 1349481 | 341768 | 6222 |
| Large Payload | JSON | Unmarshal | 3127519 | 567332 | 7382 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7308 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 15784 | 21764 | 1 |
| Medium Payload | Sonic | Marshal | 22045 | 24899 | 3 |
| Medium Payload | CBOR | Marshal | 29306 | 24589 | 1 |
| Medium Payload | MessagePack | Marshal | 40119 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 62284 | 24809 | 8 |
| Medium Payload | BEVE | Unmarshal | 35329 | 33499 | 59 |
| Medium Payload | Sonic | Unmarshal | 55203 | 60604 | 76 |
| Medium Payload | MessagePack | Unmarshal | 82421 | 36853 | 685 |
| Medium Payload | CBOR | Unmarshal | 109366 | 36104 | 738 |
| Medium Payload | JSON | Unmarshal | 290364 | 56392 | 709 |
| Small Struct | BEVE ZeroCopy | Marshal | 964 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1072 | 1182 | 2 |
| Small Struct | BEVE | Marshal | 2122 | 1792 | 1 |
| Small Struct | JSON | Marshal | 2326 | 1024 | 1 |
| Small Struct | CBOR | Marshal | 3057 | 3073 | 1 |
| Small Struct | MessagePack | Marshal | 4680 | 8200 | 9 |
| Small Struct | Sonic | Unmarshal | 799 | 389 | 3 |
| Small Struct | BEVE | Unmarshal | 2556 | 3000 | 4 |
| Small Struct | CBOR | Unmarshal | 6407 | 2416 | 52 |
| Small Struct | MessagePack | Unmarshal | 7874 | 4344 | 92 |
| Small Struct | JSON | Unmarshal | 8642 | 1352 | 29 |
