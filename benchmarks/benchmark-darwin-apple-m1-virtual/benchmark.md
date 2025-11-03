# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 59441 | 52 | 0 |
| Large Payload | BEVE | Marshal | 82819 | 196649 | 1 |
| Large Payload | CBOR | Marshal | 157190 | 196727 | 1 |
| Large Payload | MessagePack | Marshal | 210799 | 526753 | 115 |
| Large Payload | JSON | Marshal | 356992 | 213280 | 8 |
| Large Payload | Sonic | Marshal | 501643 | 222360 | 3 |
| Large Payload | BEVE | Unmarshal | 170109 | 270449 | 419 |
| Large Payload | Sonic | Unmarshal | 324353 | 335975 | 209 |
| Large Payload | MessagePack | Unmarshal | 507426 | 343716 | 6254 |
| Large Payload | CBOR | Unmarshal | 599769 | 313530 | 6398 |
| Large Payload | JSON | Unmarshal | 2008185 | 521634 | 6868 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5853 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9365 | 18436 | 1 |
| Medium Payload | CBOR | Marshal | 15377 | 21774 | 1 |
| Medium Payload | MessagePack | Marshal | 26835 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 46590 | 27495 | 8 |
| Medium Payload | Sonic | Marshal | 51996 | 24917 | 3 |
| Medium Payload | BEVE | Unmarshal | 20913 | 30685 | 59 |
| Medium Payload | Sonic | Unmarshal | 35089 | 46335 | 33 |
| Medium Payload | MessagePack | Unmarshal | 35442 | 31469 | 574 |
| Medium Payload | CBOR | Unmarshal | 55024 | 36744 | 758 |
| Medium Payload | JSON | Unmarshal | 184269 | 42304 | 560 |
| Small Struct | BEVE | Marshal | 225 | 160 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 562 | 0 | 0 |
| Small Struct | CBOR | Marshal | 697 | 576 | 1 |
| Small Struct | JSON | Marshal | 1124 | 512 | 1 |
| Small Struct | MessagePack | Marshal | 2510 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 6568 | 3124 | 2 |
| Small Struct | MessagePack | Unmarshal | 1435 | 600 | 15 |
| Small Struct | BEVE | Unmarshal | 1549 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 2127 | 2796 | 6 |
| Small Struct | CBOR | Unmarshal | 4156 | 2128 | 47 |
| Small Struct | JSON | Unmarshal | 22251 | 7720 | 107 |
