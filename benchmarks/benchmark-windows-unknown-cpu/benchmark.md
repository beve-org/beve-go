# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79814 | 207 | 2 |
| Large Payload | BEVE | Marshal | 114721 | 188815 | 3 |
| Large Payload | Sonic | Marshal | 157543 | 218854 | 4 |
| Large Payload | CBOR | Marshal | 219275 | 190256 | 2 |
| Large Payload | MessagePack | Marshal | 286686 | 526764 | 115 |
| Large Payload | JSON | Marshal | 474808 | 207210 | 9 |
| Large Payload | BEVE | Unmarshal | 276165 | 273666 | 419 |
| Large Payload | Sonic | Unmarshal | 435857 | 553267 | 588 |
| Large Payload | MessagePack | Unmarshal | 642308 | 325868 | 5874 |
| Large Payload | CBOR | Unmarshal | 865966 | 311129 | 6352 |
| Large Payload | JSON | Unmarshal | 2570991 | 532956 | 6979 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8551 | 138 | 2 |
| Medium Payload | BEVE | Marshal | 12709 | 16526 | 3 |
| Medium Payload | Sonic | Marshal | 17685 | 19513 | 4 |
| Medium Payload | CBOR | Marshal | 21719 | 18521 | 2 |
| Medium Payload | MessagePack | Marshal | 35039 | 65828 | 22 |
| Medium Payload | JSON | Marshal | 50538 | 19395 | 9 |
| Medium Payload | BEVE | Unmarshal | 27974 | 26715 | 58 |
| Medium Payload | Sonic | Unmarshal | 52807 | 66130 | 79 |
| Medium Payload | MessagePack | Unmarshal | 67813 | 35308 | 652 |
| Medium Payload | CBOR | Unmarshal | 85839 | 32248 | 663 |
| Medium Payload | JSON | Unmarshal | 262988 | 55464 | 727 |
| Small Struct | BEVE | Marshal | 721 | 672 | 3 |
| Small Struct | Sonic | Marshal | 779 | 792 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 783 | 290 | 2 |
| Small Struct | CBOR | Marshal | 2062 | 1937 | 2 |
| Small Struct | JSON | Marshal | 2278 | 1040 | 2 |
| Small Struct | MessagePack | Marshal | 3528 | 4224 | 8 |
| Small Struct | BEVE | Unmarshal | 1636 | 1592 | 4 |
| Small Struct | CBOR | Unmarshal | 2633 | 904 | 22 |
| Small Struct | Sonic | Unmarshal | 3587 | 4443 | 9 |
| Small Struct | MessagePack | Unmarshal | 5498 | 3224 | 69 |
| Small Struct | JSON | Unmarshal | 32679 | 7944 | 114 |
