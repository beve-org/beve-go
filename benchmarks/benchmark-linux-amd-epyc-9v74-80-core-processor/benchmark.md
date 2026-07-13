# AMD EPYC 9V74 80-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 71940 | 26 | 0 |
| Large Payload | BEVE | Marshal | 125315 | 204863 | 1 |
| Large Payload | Sonic | Marshal | 154926 | 208296 | 3 |
| Large Payload | CBOR | Marshal | 209786 | 196837 | 1 |
| Large Payload | MessagePack | Marshal | 329248 | 526780 | 115 |
| Large Payload | JSON | Marshal | 412190 | 213283 | 8 |
| Large Payload | BEVE | Unmarshal | 239095 | 257727 | 418 |
| Large Payload | Sonic | Unmarshal | 422277 | 550418 | 588 |
| Large Payload | MessagePack | Unmarshal | 597249 | 379677 | 6983 |
| Large Payload | CBOR | Unmarshal | 850954 | 327771 | 6684 |
| Large Payload | JSON | Unmarshal | 1922236 | 476242 | 6279 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8031 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 12183 | 20486 | 1 |
| Medium Payload | Sonic | Marshal | 17750 | 25075 | 3 |
| Medium Payload | CBOR | Marshal | 19505 | 18452 | 1 |
| Medium Payload | MessagePack | Marshal | 39295 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 50276 | 27494 | 8 |
| Medium Payload | BEVE | Unmarshal | 25493 | 31583 | 59 |
| Medium Payload | Sonic | Unmarshal | 43699 | 60542 | 75 |
| Medium Payload | MessagePack | Unmarshal | 56005 | 37024 | 690 |
| Medium Payload | CBOR | Unmarshal | 87368 | 35928 | 740 |
| Medium Payload | JSON | Unmarshal | 201071 | 47864 | 660 |
| Small Struct | BEVE ZeroCopy | Marshal | 506 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1646 | 2048 | 1 |
| Small Struct | Sonic | Marshal | 2007 | 2747 | 2 |
| Small Struct | CBOR | Marshal | 2108 | 2305 | 1 |
| Small Struct | MessagePack | Marshal | 4069 | 8201 | 9 |
| Small Struct | JSON | Marshal | 4530 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 688 | 504 | 4 |
| Small Struct | Sonic | Unmarshal | 1323 | 1258 | 7 |
| Small Struct | MessagePack | Unmarshal | 5873 | 4705 | 99 |
| Small Struct | CBOR | Unmarshal | 9125 | 4232 | 89 |
| Small Struct | JSON | Unmarshal | 11635 | 3656 | 51 |
