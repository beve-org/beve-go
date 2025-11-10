# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 83697 | 39 | 0 |
| Large Payload | BEVE | Marshal | 112405 | 180281 | 1 |
| Large Payload | Sonic | Marshal | 159209 | 223598 | 3 |
| Large Payload | CBOR | Marshal | 211557 | 196759 | 1 |
| Large Payload | MessagePack | Marshal | 318938 | 526778 | 115 |
| Large Payload | JSON | Marshal | 440614 | 213282 | 8 |
| Large Payload | BEVE | Unmarshal | 230236 | 255613 | 417 |
| Large Payload | Sonic | Unmarshal | 353729 | 544221 | 576 |
| Large Payload | MessagePack | Unmarshal | 601444 | 374286 | 6862 |
| Large Payload | CBOR | Unmarshal | 708513 | 311002 | 6340 |
| Large Payload | JSON | Unmarshal | 2261082 | 530411 | 6889 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8499 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 12500 | 20486 | 1 |
| Medium Payload | Sonic | Marshal | 19722 | 27813 | 3 |
| Medium Payload | CBOR | Marshal | 19758 | 18449 | 1 |
| Medium Payload | MessagePack | Marshal | 27202 | 33007 | 21 |
| Medium Payload | JSON | Marshal | 44116 | 21988 | 8 |
| Medium Payload | BEVE | Unmarshal | 26085 | 32159 | 59 |
| Medium Payload | Sonic | Unmarshal | 41552 | 64208 | 78 |
| Medium Payload | MessagePack | Unmarshal | 52497 | 33743 | 625 |
| Medium Payload | CBOR | Unmarshal | 62947 | 27640 | 570 |
| Medium Payload | JSON | Unmarshal | 273799 | 67736 | 891 |
| Small Struct | BEVE ZeroCopy | Marshal | 832 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1223 | 2304 | 1 |
| Small Struct | Sonic | Marshal | 1872 | 2789 | 2 |
| Small Struct | CBOR | Marshal | 2237 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 2431 | 4104 | 8 |
| Small Struct | JSON | Marshal | 3626 | 2048 | 1 |
| Small Struct | BEVE | Unmarshal | 1239 | 1592 | 4 |
| Small Struct | Sonic | Unmarshal | 2861 | 4406 | 9 |
| Small Struct | CBOR | Unmarshal | 5113 | 2816 | 61 |
| Small Struct | MessagePack | Unmarshal | 5319 | 4064 | 87 |
| Small Struct | JSON | Unmarshal | 26815 | 7816 | 110 |
