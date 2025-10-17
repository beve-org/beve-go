# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80999 | 39 | 0 |
| Large Payload | BEVE | Marshal | 116168 | 196653 | 1 |
| Large Payload | Sonic | Marshal | 162062 | 231913 | 3 |
| Large Payload | CBOR | Marshal | 226086 | 213177 | 1 |
| Large Payload | MessagePack | Marshal | 300496 | 526776 | 115 |
| Large Payload | JSON | Marshal | 468873 | 229695 | 8 |
| Large Payload | BEVE | Unmarshal | 226601 | 262621 | 416 |
| Large Payload | Sonic | Unmarshal | 360256 | 564395 | 599 |
| Large Payload | MessagePack | Unmarshal | 579151 | 366204 | 6700 |
| Large Payload | CBOR | Unmarshal | 654545 | 291434 | 5934 |
| Large Payload | JSON | Unmarshal | 2330647 | 559323 | 7210 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8085 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 10724 | 16387 | 1 |
| Medium Payload | CBOR | Marshal | 17251 | 16401 | 1 |
| Medium Payload | Sonic | Marshal | 17618 | 25039 | 3 |
| Medium Payload | MessagePack | Marshal | 34612 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 58737 | 32999 | 8 |
| Medium Payload | BEVE | Unmarshal | 23270 | 27134 | 59 |
| Medium Payload | Sonic | Unmarshal | 33482 | 49329 | 62 |
| Medium Payload | MessagePack | Unmarshal | 47330 | 29070 | 527 |
| Medium Payload | CBOR | Unmarshal | 76883 | 37752 | 774 |
| Medium Payload | JSON | Unmarshal | 232738 | 56184 | 728 |
| Small Struct | BEVE ZeroCopy | Marshal | 885 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1001 | 1792 | 1 |
| Small Struct | CBOR | Marshal | 1131 | 1024 | 1 |
| Small Struct | Sonic | Marshal | 1595 | 2356 | 2 |
| Small Struct | MessagePack | Marshal | 4043 | 8202 | 9 |
| Small Struct | JSON | Marshal | 4574 | 2688 | 1 |
| Small Struct | Sonic | Unmarshal | 969 | 842 | 6 |
| Small Struct | BEVE | Unmarshal | 988 | 1016 | 4 |
| Small Struct | MessagePack | Unmarshal | 4341 | 3192 | 68 |
| Small Struct | CBOR | Unmarshal | 8585 | 5192 | 107 |
| Small Struct | JSON | Unmarshal | 13225 | 3816 | 56 |
