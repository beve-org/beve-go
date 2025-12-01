# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 84811 | 79 | 0 |
| Large Payload | BEVE | Marshal | 132631 | 196658 | 1 |
| Large Payload | Sonic | Marshal | 185208 | 206865 | 3 |
| Large Payload | CBOR | Marshal | 270602 | 196702 | 1 |
| Large Payload | MessagePack | Marshal | 328325 | 526715 | 115 |
| Large Payload | JSON | Marshal | 510288 | 196911 | 8 |
| Large Payload | BEVE | Unmarshal | 309053 | 273669 | 417 |
| Large Payload | Sonic | Unmarshal | 444163 | 521097 | 574 |
| Large Payload | CBOR | Unmarshal | 858286 | 315274 | 6433 |
| Large Payload | MessagePack | Unmarshal | 917411 | 339678 | 6168 |
| Large Payload | JSON | Unmarshal | 2346912 | 469371 | 6206 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7769 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 13539 | 20487 | 1 |
| Medium Payload | Sonic | Marshal | 17777 | 20780 | 3 |
| Medium Payload | CBOR | Marshal | 23907 | 18440 | 1 |
| Medium Payload | MessagePack | Marshal | 39635 | 65773 | 22 |
| Medium Payload | JSON | Marshal | 49245 | 20710 | 8 |
| Medium Payload | BEVE | Unmarshal | 32433 | 25243 | 59 |
| Medium Payload | Sonic | Unmarshal | 50649 | 55350 | 72 |
| Medium Payload | MessagePack | Unmarshal | 79824 | 38702 | 721 |
| Medium Payload | CBOR | Unmarshal | 94927 | 34472 | 706 |
| Medium Payload | JSON | Unmarshal | 314052 | 61816 | 833 |
| Small Struct | BEVE ZeroCopy | Marshal | 447 | 0 | 0 |
| Small Struct | CBOR | Marshal | 806 | 576 | 1 |
| Small Struct | BEVE | Marshal | 1512 | 2048 | 1 |
| Small Struct | Sonic | Marshal | 1810 | 2127 | 2 |
| Small Struct | MessagePack | Marshal | 4837 | 8200 | 9 |
| Small Struct | JSON | Marshal | 6490 | 3072 | 1 |
| Small Struct | BEVE | Unmarshal | 1377 | 1016 | 4 |
| Small Struct | Sonic | Unmarshal | 2024 | 1968 | 8 |
| Small Struct | CBOR | Unmarshal | 3407 | 1160 | 27 |
| Small Struct | JSON | Unmarshal | 3752 | 552 | 15 |
| Small Struct | MessagePack | Unmarshal | 7534 | 4600 | 96 |
