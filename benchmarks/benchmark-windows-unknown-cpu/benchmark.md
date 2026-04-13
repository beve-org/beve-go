# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79355 | 52 | 0 |
| Large Payload | BEVE | Marshal | 114660 | 188451 | 1 |
| Large Payload | Sonic | Marshal | 164434 | 222477 | 3 |
| Large Payload | CBOR | Marshal | 217638 | 196680 | 1 |
| Large Payload | MessagePack | Marshal | 289679 | 526708 | 115 |
| Large Payload | JSON | Marshal | 465961 | 205072 | 8 |
| Large Payload | BEVE | Unmarshal | 267139 | 259622 | 417 |
| Large Payload | Sonic | Unmarshal | 443989 | 566547 | 594 |
| Large Payload | MessagePack | Unmarshal | 659275 | 340195 | 6171 |
| Large Payload | CBOR | Unmarshal | 813208 | 298378 | 6094 |
| Large Payload | JSON | Unmarshal | 2438606 | 495290 | 6515 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8299 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 11860 | 18433 | 1 |
| Medium Payload | Sonic | Marshal | 15530 | 20731 | 3 |
| Medium Payload | CBOR | Marshal | 22279 | 20488 | 1 |
| Medium Payload | MessagePack | Marshal | 34100 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 50888 | 24805 | 8 |
| Medium Payload | BEVE | Unmarshal | 29011 | 29659 | 59 |
| Medium Payload | Sonic | Unmarshal | 44381 | 50860 | 70 |
| Medium Payload | MessagePack | Unmarshal | 76817 | 43070 | 812 |
| Medium Payload | CBOR | Unmarshal | 98658 | 39032 | 804 |
| Medium Payload | JSON | Unmarshal | 247073 | 53496 | 685 |
| Small Struct | BEVE | Marshal | 437 | 384 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 679 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 1164 | 1032 | 6 |
| Small Struct | JSON | Marshal | 1495 | 576 | 1 |
| Small Struct | Sonic | Marshal | 1509 | 1438 | 2 |
| Small Struct | CBOR | Marshal | 1912 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 762 | 376 | 4 |
| Small Struct | MessagePack | Unmarshal | 4530 | 2752 | 58 |
| Small Struct | Sonic | Unmarshal | 5507 | 7360 | 10 |
| Small Struct | CBOR | Unmarshal | 7339 | 3432 | 72 |
| Small Struct | JSON | Unmarshal | 11473 | 2344 | 45 |
