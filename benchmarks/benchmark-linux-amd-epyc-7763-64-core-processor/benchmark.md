# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77185 | 259 | 2 |
| Large Payload | BEVE | Marshal | 117470 | 188923 | 3 |
| Large Payload | Sonic | Marshal | 164417 | 208735 | 4 |
| Large Payload | CBOR | Marshal | 221458 | 197692 | 2 |
| Large Payload | MessagePack | Marshal | 342692 | 526840 | 115 |
| Large Payload | JSON | Marshal | 463372 | 222090 | 9 |
| Large Payload | BEVE | Unmarshal | 242656 | 265759 | 416 |
| Large Payload | Sonic | Unmarshal | 380457 | 548636 | 586 |
| Large Payload | MessagePack | Unmarshal | 578953 | 342599 | 6237 |
| Large Payload | CBOR | Unmarshal | 803013 | 345114 | 7020 |
| Large Payload | JSON | Unmarshal | 2335302 | 526722 | 6949 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8589 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 12949 | 19258 | 3 |
| Medium Payload | Sonic | Marshal | 14363 | 19258 | 4 |
| Medium Payload | CBOR | Marshal | 19737 | 18555 | 2 |
| Medium Payload | MessagePack | Marshal | 26670 | 33063 | 21 |
| Medium Payload | JSON | Marshal | 48927 | 24968 | 9 |
| Medium Payload | BEVE | Unmarshal | 25124 | 30687 | 59 |
| Medium Payload | Sonic | Unmarshal | 32533 | 45515 | 65 |
| Medium Payload | MessagePack | Unmarshal | 50676 | 29839 | 539 |
| Medium Payload | CBOR | Unmarshal | 83033 | 38456 | 786 |
| Medium Payload | JSON | Unmarshal | 231274 | 54520 | 726 |
| Small Struct | BEVE ZeroCopy | Marshal | 689 | 290 | 2 |
| Small Struct | BEVE | Marshal | 1661 | 2979 | 3 |
| Small Struct | MessagePack | Marshal | 1664 | 2176 | 7 |
| Small Struct | Sonic | Marshal | 1712 | 2293 | 3 |
| Small Struct | CBOR | Marshal | 3203 | 3218 | 2 |
| Small Struct | JSON | Marshal | 4217 | 2193 | 2 |
| Small Struct | BEVE | Unmarshal | 1052 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 2640 | 3814 | 9 |
| Small Struct | MessagePack | Unmarshal | 5234 | 3840 | 80 |
| Small Struct | CBOR | Unmarshal | 6072 | 3144 | 67 |
| Small Struct | JSON | Unmarshal | 6516 | 1352 | 29 |
