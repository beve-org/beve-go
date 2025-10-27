# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 71165 | 52 | 0 |
| Large Payload | BEVE | Marshal | 138360 | 196745 | 1 |
| Large Payload | Sonic | Marshal | 183250 | 207670 | 3 |
| Large Payload | CBOR | Marshal | 228454 | 188576 | 1 |
| Large Payload | MessagePack | Marshal | 322249 | 526732 | 115 |
| Large Payload | JSON | Marshal | 462852 | 205109 | 8 |
| Large Payload | BEVE | Unmarshal | 299448 | 278129 | 417 |
| Large Payload | Sonic | Unmarshal | 499996 | 548648 | 581 |
| Large Payload | MessagePack | Unmarshal | 720540 | 367501 | 6723 |
| Large Payload | CBOR | Unmarshal | 879795 | 333545 | 6804 |
| Large Payload | JSON | Unmarshal | 2339904 | 538786 | 7065 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7314 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 15135 | 21773 | 1 |
| Medium Payload | Sonic | Marshal | 20151 | 22272 | 3 |
| Medium Payload | CBOR | Marshal | 23479 | 19088 | 1 |
| Medium Payload | JSON | Marshal | 42818 | 19306 | 8 |
| Medium Payload | MessagePack | Marshal | 44351 | 65776 | 22 |
| Medium Payload | BEVE | Unmarshal | 31964 | 28957 | 59 |
| Medium Payload | Sonic | Unmarshal | 60086 | 69592 | 80 |
| Medium Payload | MessagePack | Unmarshal | 70344 | 35247 | 657 |
| Medium Payload | CBOR | Unmarshal | 90610 | 36856 | 756 |
| Medium Payload | JSON | Unmarshal | 239200 | 57416 | 767 |
| Small Struct | Sonic | Marshal | 374 | 268 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 407 | 0 | 0 |
| Small Struct | BEVE | Marshal | 990 | 1152 | 1 |
| Small Struct | CBOR | Marshal | 1132 | 896 | 1 |
| Small Struct | MessagePack | Marshal | 1205 | 1032 | 6 |
| Small Struct | JSON | Marshal | 3607 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 1725 | 2104 | 4 |
| Small Struct | CBOR | Unmarshal | 3976 | 1544 | 35 |
| Small Struct | Sonic | Unmarshal | 4253 | 4715 | 9 |
| Small Struct | MessagePack | Unmarshal | 5312 | 3616 | 77 |
| Small Struct | JSON | Unmarshal | 27772 | 8072 | 118 |
