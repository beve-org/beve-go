# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 84593 | 65 | 0 |
| Large Payload | BEVE | Marshal | 134487 | 196670 | 1 |
| Large Payload | Sonic | Marshal | 175012 | 223157 | 3 |
| Large Payload | CBOR | Marshal | 226971 | 196762 | 1 |
| Large Payload | MessagePack | Marshal | 309258 | 526707 | 115 |
| Large Payload | JSON | Marshal | 535655 | 221485 | 8 |
| Large Payload | BEVE | Unmarshal | 286325 | 265027 | 418 |
| Large Payload | Sonic | Unmarshal | 462112 | 546771 | 582 |
| Large Payload | MessagePack | Unmarshal | 707748 | 344454 | 6265 |
| Large Payload | CBOR | Unmarshal | 961880 | 337626 | 6873 |
| Large Payload | JSON | Unmarshal | 2617392 | 480314 | 6366 |
| Medium Payload | BEVE ZeroCopy | Marshal | 11067 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 18927 | 18437 | 1 |
| Medium Payload | Sonic | Marshal | 24306 | 20903 | 3 |
| Medium Payload | CBOR | Marshal | 31345 | 21770 | 1 |
| Medium Payload | MessagePack | Marshal | 33803 | 33002 | 21 |
| Medium Payload | JSON | Marshal | 71328 | 20710 | 8 |
| Medium Payload | BEVE | Unmarshal | 28568 | 24987 | 59 |
| Medium Payload | Sonic | Unmarshal | 57177 | 58603 | 76 |
| Medium Payload | MessagePack | Unmarshal | 82612 | 37549 | 695 |
| Medium Payload | CBOR | Unmarshal | 106524 | 35032 | 718 |
| Medium Payload | JSON | Unmarshal | 300499 | 57040 | 777 |
| Small Struct | BEVE ZeroCopy | Marshal | 395 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1260 | 1324 | 2 |
| Small Struct | MessagePack | Marshal | 1352 | 1032 | 6 |
| Small Struct | BEVE | Marshal | 2555 | 2689 | 1 |
| Small Struct | CBOR | Marshal | 2875 | 2688 | 1 |
| Small Struct | JSON | Marshal | 5688 | 2304 | 1 |
| Small Struct | Sonic | Unmarshal | 902 | 570 | 5 |
| Small Struct | BEVE | Unmarshal | 2575 | 3512 | 4 |
| Small Struct | MessagePack | Unmarshal | 6960 | 3488 | 73 |
| Small Struct | CBOR | Unmarshal | 11239 | 4232 | 89 |
| Small Struct | JSON | Unmarshal | 31793 | 7624 | 104 |
