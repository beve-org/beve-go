# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 90601 | 65 | 0 |
| Large Payload | BEVE | Marshal | 148157 | 196645 | 1 |
| Large Payload | Sonic | Marshal | 209774 | 215506 | 3 |
| Large Payload | CBOR | Marshal | 236731 | 213096 | 1 |
| Large Payload | MessagePack | Marshal | 372540 | 526722 | 115 |
| Large Payload | JSON | Marshal | 560296 | 205103 | 8 |
| Large Payload | BEVE | Unmarshal | 297703 | 289221 | 417 |
| Large Payload | Sonic | Unmarshal | 458167 | 600743 | 607 |
| Large Payload | MessagePack | Unmarshal | 655818 | 336015 | 6092 |
| Large Payload | CBOR | Unmarshal | 962420 | 319754 | 6500 |
| Large Payload | JSON | Unmarshal | 2638110 | 534659 | 7050 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8701 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 13334 | 20483 | 1 |
| Medium Payload | Sonic | Marshal | 18665 | 16859 | 3 |
| Medium Payload | MessagePack | Marshal | 29245 | 33002 | 21 |
| Medium Payload | CBOR | Marshal | 30756 | 19083 | 1 |
| Medium Payload | JSON | Marshal | 49039 | 20707 | 8 |
| Medium Payload | BEVE | Unmarshal | 32096 | 29244 | 59 |
| Medium Payload | Sonic | Unmarshal | 59179 | 57512 | 75 |
| Medium Payload | MessagePack | Unmarshal | 103616 | 39712 | 742 |
| Medium Payload | CBOR | Unmarshal | 141919 | 38664 | 792 |
| Medium Payload | JSON | Unmarshal | 297999 | 55888 | 740 |
| Small Struct | CBOR | Marshal | 528 | 352 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 905 | 0 | 0 |
| Small Struct | Sonic | Marshal | 2029 | 2740 | 2 |
| Small Struct | JSON | Marshal | 2458 | 1152 | 1 |
| Small Struct | BEVE | Marshal | 2843 | 2688 | 1 |
| Small Struct | MessagePack | Marshal | 3234 | 4104 | 8 |
| Small Struct | MessagePack | Unmarshal | 1918 | 832 | 20 |
| Small Struct | BEVE | Unmarshal | 2044 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 5859 | 7333 | 10 |
| Small Struct | CBOR | Unmarshal | 7346 | 3272 | 71 |
| Small Struct | JSON | Unmarshal | 28245 | 7528 | 101 |
