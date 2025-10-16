# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66715 | 207 | 2 |
| Large Payload | BEVE | Marshal | 143824 | 197397 | 3 |
| Large Payload | CBOR | Marshal | 159078 | 181494 | 2 |
| Large Payload | MessagePack | Marshal | 424008 | 526811 | 115 |
| Large Payload | JSON | Marshal | 464768 | 213842 | 9 |
| Large Payload | Sonic | Marshal | 592874 | 215573 | 4 |
| Large Payload | Sonic | Unmarshal | 320071 | 359067 | 211 |
| Large Payload | BEVE | Unmarshal | 364278 | 278032 | 419 |
| Large Payload | MessagePack | Unmarshal | 419910 | 335345 | 6078 |
| Large Payload | CBOR | Unmarshal | 513777 | 334923 | 6826 |
| Large Payload | JSON | Unmarshal | 1997878 | 527410 | 6844 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8969 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 13984 | 24725 | 3 |
| Medium Payload | CBOR | Marshal | 21668 | 16473 | 2 |
| Medium Payload | MessagePack | Marshal | 27570 | 33061 | 21 |
| Medium Payload | JSON | Marshal | 46737 | 22071 | 9 |
| Medium Payload | Sonic | Marshal | 64538 | 24925 | 4 |
| Medium Payload | BEVE | Unmarshal | 32792 | 34846 | 59 |
| Medium Payload | Sonic | Unmarshal | 44900 | 41289 | 33 |
| Medium Payload | MessagePack | Unmarshal | 60997 | 36430 | 669 |
| Medium Payload | CBOR | Unmarshal | 71066 | 30024 | 618 |
| Medium Payload | JSON | Unmarshal | 263932 | 60712 | 814 |
| Small Struct | BEVE ZeroCopy | Marshal | 294 | 289 | 2 |
| Small Struct | BEVE | Marshal | 1702 | 2978 | 3 |
| Small Struct | MessagePack | Marshal | 2512 | 4224 | 8 |
| Small Struct | CBOR | Marshal | 2518 | 2449 | 2 |
| Small Struct | JSON | Marshal | 2804 | 1936 | 2 |
| Small Struct | Sonic | Marshal | 4928 | 2250 | 3 |
| Small Struct | BEVE | Unmarshal | 2631 | 3384 | 4 |
| Small Struct | CBOR | Unmarshal | 2967 | 1576 | 36 |
| Small Struct | Sonic | Unmarshal | 4234 | 5491 | 6 |
| Small Struct | MessagePack | Unmarshal | 4870 | 3168 | 67 |
| Small Struct | JSON | Unmarshal | 9066 | 2096 | 37 |
