# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 78447 | 52 | 0 |
| Large Payload | BEVE | Marshal | 119411 | 196667 | 1 |
| Large Payload | Sonic | Marshal | 155764 | 215471 | 3 |
| Large Payload | CBOR | Marshal | 203856 | 188563 | 1 |
| Large Payload | MessagePack | Marshal | 305516 | 526777 | 115 |
| Large Payload | JSON | Marshal | 448803 | 221477 | 8 |
| Large Payload | BEVE | Unmarshal | 238845 | 272736 | 417 |
| Large Payload | Sonic | Unmarshal | 360864 | 563202 | 585 |
| Large Payload | MessagePack | Unmarshal | 560819 | 343079 | 6229 |
| Large Payload | CBOR | Unmarshal | 738996 | 336330 | 6831 |
| Large Payload | JSON | Unmarshal | 2299713 | 535121 | 7057 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7761 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 11162 | 18436 | 1 |
| Medium Payload | Sonic | Marshal | 19250 | 27752 | 3 |
| Medium Payload | CBOR | Marshal | 22000 | 20499 | 1 |
| Medium Payload | MessagePack | Marshal | 38389 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 39414 | 19303 | 8 |
| Medium Payload | BEVE | Unmarshal | 25209 | 29567 | 59 |
| Medium Payload | Sonic | Unmarshal | 31872 | 44176 | 64 |
| Medium Payload | MessagePack | Unmarshal | 60095 | 35376 | 655 |
| Medium Payload | CBOR | Unmarshal | 70539 | 32216 | 662 |
| Medium Payload | JSON | Unmarshal | 196788 | 46648 | 611 |
| Small Struct | Sonic | Marshal | 440 | 357 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 604 | 0 | 0 |
| Small Struct | BEVE | Marshal | 639 | 1024 | 1 |
| Small Struct | CBOR | Marshal | 2619 | 2688 | 1 |
| Small Struct | MessagePack | Marshal | 4192 | 8201 | 9 |
| Small Struct | JSON | Marshal | 5553 | 3073 | 1 |
| Small Struct | BEVE | Unmarshal | 906 | 888 | 4 |
| Small Struct | MessagePack | Unmarshal | 1265 | 400 | 11 |
| Small Struct | Sonic | Unmarshal | 1294 | 1328 | 7 |
| Small Struct | CBOR | Unmarshal | 8573 | 5168 | 106 |
| Small Struct | JSON | Unmarshal | 10829 | 2400 | 47 |
