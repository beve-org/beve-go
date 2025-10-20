# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 78379 | 39 | 0 |
| Large Payload | BEVE | Marshal | 115834 | 188500 | 1 |
| Large Payload | Sonic | Marshal | 153967 | 215667 | 3 |
| Large Payload | CBOR | Marshal | 209001 | 196761 | 1 |
| Large Payload | MessagePack | Marshal | 307159 | 526778 | 115 |
| Large Payload | JSON | Marshal | 432164 | 213282 | 8 |
| Large Payload | BEVE | Unmarshal | 247052 | 286052 | 418 |
| Large Payload | Sonic | Unmarshal | 353334 | 537555 | 574 |
| Large Payload | MessagePack | Unmarshal | 582977 | 366635 | 6708 |
| Large Payload | CBOR | Unmarshal | 702219 | 311482 | 6356 |
| Large Payload | JSON | Unmarshal | 2132895 | 514857 | 6602 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7410 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11376 | 18438 | 1 |
| Medium Payload | Sonic | Marshal | 15104 | 20799 | 3 |
| Medium Payload | CBOR | Marshal | 23759 | 21780 | 1 |
| Medium Payload | JSON | Marshal | 36843 | 18663 | 8 |
| Medium Payload | MessagePack | Marshal | 36925 | 65783 | 22 |
| Medium Payload | BEVE | Unmarshal | 23986 | 26814 | 59 |
| Medium Payload | Sonic | Unmarshal | 39274 | 55864 | 74 |
| Medium Payload | CBOR | Unmarshal | 58679 | 24904 | 516 |
| Medium Payload | MessagePack | Unmarshal | 67725 | 42322 | 789 |
| Medium Payload | JSON | Unmarshal | 199111 | 46416 | 611 |
| Small Struct | Sonic | Marshal | 386 | 293 | 2 |
| Small Struct | BEVE | Marshal | 597 | 512 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 875 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1363 | 1280 | 1 |
| Small Struct | MessagePack | Marshal | 4568 | 8201 | 9 |
| Small Struct | JSON | Marshal | 5438 | 3073 | 1 |
| Small Struct | CBOR | Unmarshal | 1108 | 232 | 7 |
| Small Struct | BEVE | Unmarshal | 1114 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 4113 | 7404 | 10 |
| Small Struct | MessagePack | Unmarshal | 4516 | 3272 | 70 |
| Small Struct | JSON | Unmarshal | 28272 | 8008 | 116 |
