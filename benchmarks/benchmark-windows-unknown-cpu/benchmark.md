# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 75913 | 52 | 0 |
| Large Payload | BEVE | Marshal | 122060 | 188466 | 1 |
| Large Payload | Sonic | Marshal | 179733 | 215875 | 3 |
| Large Payload | CBOR | Marshal | 250045 | 204881 | 1 |
| Large Payload | MessagePack | Marshal | 329642 | 526713 | 115 |
| Large Payload | JSON | Marshal | 504479 | 221486 | 8 |
| Large Payload | BEVE | Unmarshal | 290221 | 270726 | 418 |
| Large Payload | Sonic | Unmarshal | 442862 | 538877 | 579 |
| Large Payload | MessagePack | Unmarshal | 685485 | 339974 | 6163 |
| Large Payload | CBOR | Unmarshal | 945670 | 328842 | 6704 |
| Large Payload | JSON | Unmarshal | 2572831 | 511178 | 6743 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6745 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 13336 | 18440 | 1 |
| Medium Payload | Sonic | Marshal | 19745 | 24932 | 3 |
| Medium Payload | CBOR | Marshal | 25626 | 19084 | 1 |
| Medium Payload | MessagePack | Marshal | 37536 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 51180 | 24803 | 8 |
| Medium Payload | BEVE | Unmarshal | 29634 | 27867 | 59 |
| Medium Payload | Sonic | Unmarshal | 42417 | 46533 | 68 |
| Medium Payload | CBOR | Unmarshal | 77424 | 27144 | 560 |
| Medium Payload | MessagePack | Unmarshal | 85462 | 41214 | 775 |
| Medium Payload | JSON | Unmarshal | 295302 | 63752 | 845 |
| Small Struct | BEVE | Marshal | 384 | 208 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 829 | 0 | 0 |
| Small Struct | CBOR | Marshal | 2145 | 2048 | 1 |
| Small Struct | Sonic | Marshal | 2260 | 2760 | 2 |
| Small Struct | MessagePack | Marshal | 3178 | 4104 | 8 |
| Small Struct | JSON | Marshal | 5690 | 2688 | 1 |
| Small Struct | Sonic | Unmarshal | 1700 | 1248 | 7 |
| Small Struct | BEVE | Unmarshal | 2353 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 6344 | 4032 | 86 |
| Small Struct | CBOR | Unmarshal | 8876 | 3880 | 82 |
| Small Struct | JSON | Unmarshal | 30280 | 7752 | 108 |
