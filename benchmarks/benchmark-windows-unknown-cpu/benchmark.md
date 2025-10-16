# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79066 | 207 | 2 |
| Large Payload | BEVE | Marshal | 112627 | 189320 | 3 |
| Large Payload | Sonic | Marshal | 160020 | 227663 | 4 |
| Large Payload | CBOR | Marshal | 250731 | 206012 | 2 |
| Large Payload | MessagePack | Marshal | 290034 | 526762 | 115 |
| Large Payload | JSON | Marshal | 485097 | 215195 | 9 |
| Large Payload | BEVE | Unmarshal | 277790 | 276772 | 418 |
| Large Payload | Sonic | Unmarshal | 420960 | 525164 | 570 |
| Large Payload | MessagePack | Unmarshal | 707512 | 353463 | 6448 |
| Large Payload | CBOR | Unmarshal | 902030 | 332330 | 6786 |
| Large Payload | JSON | Unmarshal | 2562383 | 531213 | 7000 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8140 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 13301 | 18580 | 3 |
| Medium Payload | Sonic | Marshal | 16466 | 20880 | 4 |
| Medium Payload | CBOR | Marshal | 21894 | 18529 | 2 |
| Medium Payload | MessagePack | Marshal | 25391 | 33058 | 21 |
| Medium Payload | JSON | Marshal | 46029 | 20777 | 9 |
| Medium Payload | BEVE | Unmarshal | 28699 | 28763 | 59 |
| Medium Payload | Sonic | Unmarshal | 37217 | 42330 | 61 |
| Medium Payload | MessagePack | Unmarshal | 69611 | 36076 | 665 |
| Medium Payload | CBOR | Unmarshal | 98442 | 38520 | 797 |
| Medium Payload | JSON | Unmarshal | 261119 | 56440 | 730 |
| Small Struct | BEVE ZeroCopy | Marshal | 992 | 289 | 2 |
| Small Struct | JSON | Marshal | 1455 | 656 | 2 |
| Small Struct | Sonic | Marshal | 1721 | 2093 | 3 |
| Small Struct | MessagePack | Marshal | 2120 | 2176 | 7 |
| Small Struct | BEVE | Marshal | 2174 | 2978 | 3 |
| Small Struct | CBOR | Marshal | 3236 | 3217 | 2 |
| Small Struct | BEVE | Unmarshal | 1922 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 3255 | 1784 | 40 |
| Small Struct | JSON | Unmarshal | 5060 | 872 | 21 |
| Small Struct | Sonic | Unmarshal | 5611 | 7825 | 10 |
| Small Struct | CBOR | Unmarshal | 6640 | 3176 | 68 |
