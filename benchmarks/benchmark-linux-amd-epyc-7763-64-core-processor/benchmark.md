# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77609 | 207 | 2 |
| Large Payload | BEVE | Marshal | 118398 | 180597 | 3 |
| Large Payload | Sonic | Marshal | 158166 | 200681 | 4 |
| Large Payload | CBOR | Marshal | 211628 | 197431 | 2 |
| Large Payload | MessagePack | Marshal | 310283 | 526835 | 115 |
| Large Payload | JSON | Marshal | 430380 | 205283 | 9 |
| Large Payload | BEVE | Unmarshal | 232720 | 274721 | 418 |
| Large Payload | Sonic | Unmarshal | 359296 | 540289 | 570 |
| Large Payload | MessagePack | Unmarshal | 581570 | 357387 | 6517 |
| Large Payload | CBOR | Unmarshal | 739225 | 319945 | 6509 |
| Large Payload | JSON | Unmarshal | 2098406 | 472506 | 6205 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8538 | 134 | 2 |
| Medium Payload | Sonic | Marshal | 13571 | 17084 | 4 |
| Medium Payload | BEVE | Marshal | 14108 | 21921 | 3 |
| Medium Payload | CBOR | Marshal | 27806 | 24698 | 2 |
| Medium Payload | MessagePack | Marshal | 41098 | 65839 | 22 |
| Medium Payload | JSON | Marshal | 41662 | 19469 | 9 |
| Medium Payload | BEVE | Unmarshal | 23030 | 25118 | 59 |
| Medium Payload | Sonic | Unmarshal | 46696 | 73373 | 81 |
| Medium Payload | CBOR | Unmarshal | 69314 | 30312 | 624 |
| Medium Payload | MessagePack | Unmarshal | 70467 | 47506 | 900 |
| Medium Payload | JSON | Unmarshal | 226490 | 52008 | 716 |
| Small Struct | BEVE ZeroCopy | Marshal | 796 | 290 | 2 |
| Small Struct | Sonic | Marshal | 1416 | 2023 | 3 |
| Small Struct | CBOR | Marshal | 1584 | 1552 | 2 |
| Small Struct | JSON | Marshal | 1698 | 848 | 2 |
| Small Struct | BEVE | Marshal | 1774 | 3362 | 3 |
| Small Struct | MessagePack | Marshal | 2888 | 4225 | 8 |
| Small Struct | BEVE | Unmarshal | 929 | 1016 | 4 |
| Small Struct | MessagePack | Unmarshal | 3944 | 2528 | 55 |
| Small Struct | Sonic | Unmarshal | 4279 | 7420 | 10 |
| Small Struct | CBOR | Unmarshal | 8179 | 4392 | 94 |
| Small Struct | JSON | Unmarshal | 14702 | 4008 | 62 |
