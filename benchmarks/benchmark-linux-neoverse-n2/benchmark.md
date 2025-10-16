# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 71066 | 233 | 2 |
| Large Payload | BEVE | Marshal | 108504 | 189513 | 3 |
| Large Payload | CBOR | Marshal | 185376 | 190723 | 2 |
| Large Payload | MessagePack | Marshal | 262168 | 526856 | 115 |
| Large Payload | Sonic | Marshal | 316531 | 237760 | 4 |
| Large Payload | JSON | Marshal | 370891 | 207762 | 9 |
| Large Payload | BEVE | Unmarshal | 212162 | 253411 | 417 |
| Large Payload | Sonic | Unmarshal | 289974 | 401676 | 213 |
| Large Payload | MessagePack | Unmarshal | 513373 | 345724 | 6297 |
| Large Payload | CBOR | Unmarshal | 630014 | 304954 | 6219 |
| Large Payload | JSON | Unmarshal | 1996612 | 539795 | 7102 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7422 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 12700 | 24736 | 3 |
| Medium Payload | CBOR | Marshal | 18392 | 19174 | 2 |
| Medium Payload | MessagePack | Marshal | 21799 | 33063 | 21 |
| Medium Payload | Sonic | Marshal | 29131 | 22238 | 4 |
| Medium Payload | JSON | Marshal | 43254 | 24888 | 9 |
| Medium Payload | BEVE | Unmarshal | 22343 | 28286 | 59 |
| Medium Payload | Sonic | Unmarshal | 25152 | 32890 | 33 |
| Medium Payload | MessagePack | Unmarshal | 55410 | 39792 | 744 |
| Medium Payload | CBOR | Unmarshal | 62123 | 30232 | 628 |
| Medium Payload | JSON | Unmarshal | 182810 | 50680 | 650 |
| Small Struct | BEVE ZeroCopy | Marshal | 432 | 288 | 2 |
| Small Struct | BEVE | Marshal | 765 | 928 | 3 |
| Small Struct | CBOR | Marshal | 1428 | 1424 | 2 |
| Small Struct | JSON | Marshal | 2405 | 1425 | 2 |
| Small Struct | Sonic | Marshal | 3834 | 2947 | 3 |
| Small Struct | MessagePack | Marshal | 3862 | 8321 | 9 |
| Small Struct | BEVE | Unmarshal | 778 | 568 | 4 |
| Small Struct | CBOR | Unmarshal | 1298 | 320 | 10 |
| Small Struct | Sonic | Unmarshal | 3167 | 5282 | 6 |
| Small Struct | MessagePack | Unmarshal | 4868 | 4024 | 86 |
| Small Struct | JSON | Unmarshal | 11014 | 3688 | 52 |
