# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 97985 | 65 | 0 |
| Large Payload | BEVE | Marshal | 141504 | 188454 | 1 |
| Large Payload | Sonic | Marshal | 204103 | 215597 | 3 |
| Large Payload | CBOR | Marshal | 229168 | 204911 | 1 |
| Large Payload | MessagePack | Marshal | 399793 | 526726 | 115 |
| Large Payload | JSON | Marshal | 638086 | 221492 | 8 |
| Large Payload | BEVE | Unmarshal | 365064 | 267144 | 416 |
| Large Payload | Sonic | Unmarshal | 441218 | 548233 | 593 |
| Large Payload | MessagePack | Unmarshal | 657160 | 331843 | 6006 |
| Large Payload | CBOR | Unmarshal | 892015 | 329609 | 6714 |
| Large Payload | JSON | Unmarshal | 2520067 | 514628 | 6701 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7250 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 14142 | 20481 | 1 |
| Medium Payload | Sonic | Marshal | 16711 | 20807 | 3 |
| Medium Payload | MessagePack | Marshal | 27516 | 33002 | 21 |
| Medium Payload | CBOR | Marshal | 29240 | 20496 | 1 |
| Medium Payload | JSON | Marshal | 49211 | 21997 | 8 |
| Medium Payload | BEVE | Unmarshal | 27849 | 26843 | 59 |
| Medium Payload | Sonic | Unmarshal | 59420 | 66686 | 78 |
| Medium Payload | CBOR | Unmarshal | 88874 | 28168 | 580 |
| Medium Payload | MessagePack | Unmarshal | 89023 | 37726 | 703 |
| Medium Payload | JSON | Unmarshal | 281243 | 55960 | 721 |
| Small Struct | BEVE ZeroCopy | Marshal | 716 | 0 | 0 |
| Small Struct | JSON | Marshal | 894 | 320 | 1 |
| Small Struct | CBOR | Marshal | 894 | 768 | 1 |
| Small Struct | MessagePack | Marshal | 1141 | 1032 | 6 |
| Small Struct | BEVE | Marshal | 1893 | 2304 | 1 |
| Small Struct | Sonic | Marshal | 2224 | 2725 | 2 |
| Small Struct | BEVE | Unmarshal | 1831 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 3104 | 3531 | 9 |
| Small Struct | MessagePack | Unmarshal | 6012 | 3584 | 76 |
| Small Struct | CBOR | Unmarshal | 7330 | 3168 | 68 |
| Small Struct | JSON | Unmarshal | 23364 | 4648 | 82 |
